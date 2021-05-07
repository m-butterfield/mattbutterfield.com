package controllers

import (
	"cloud.google.com/go/pubsub"
	"context"
	"fmt"
	"log"
	"net/http"
	"nhooyr.io/websocket"
	"strings"
)

func MultiVideoConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := websocket.Accept(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer closeWS(ws)
	meetingID := strings.ToLower(r.URL.Query().Get("meetingID"))
	userID := strings.ToLower(r.URL.Query().Get("userID"))

	topicInName := fmt.Sprintf("reflector-input-%s", meetingID)
	topicOutName := fmt.Sprintf("reflector-output-%s", meetingID)
	topicIn := pubSub.Topic(topicInName)
	topicIn.EnableMessageOrdering = true

	ctx := context.Background()
	exists, err := topicIn.Exists(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if !exists {
		log.Printf("Topic %s doesn't exist - creating it", topicInName)
		_, err = pubSub.CreateTopic(ctx, topicInName)
		if err != nil {
			log.Fatal(err)
		}
	}

	topicOut := pubSub.Topic(topicOutName)
	topicOut.EnableMessageOrdering = true

	exists, err = topicOut.Exists(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if !exists {
		log.Printf("Topic %s doesn't exist - creating it", topicOutName)
		_, err = pubSub.CreateTopic(ctx, topicOutName)
		if err != nil {
			log.Fatal(err)
		}
	}

	cctx, cancelFunc := context.WithCancel(ctx)
	go multiWsLoop(ctx, cancelFunc, ws, topicIn, userID)
	multiPubSubLoop(cctx, ctx, ws, topicOut, userID)
}

func multiWsLoop(ctx context.Context, cancelFunc context.CancelFunc, ws *websocket.Conn, topic *pubsub.Topic, userID string) {
	log.Printf("Starting wsLoop for %s...", userID)
	orderingKey := fmt.Sprintf("%s-%s", userID, topic.ID())
	msg := &pubsub.Message{
		Data:        []byte("setup"),
		Attributes:  map[string]string{"sender": userID},
		OrderingKey: orderingKey,
	}
	if _, err := topic.Publish(ctx, msg).Get(ctx); err != nil {
		log.Printf("Could not publish message: %s", err)
		return
	}
	for {
		if _, message, err := ws.Read(ctx); err != nil {
			// could check for 'close' here and tell peer we have closed
			log.Printf("Error reading message %s", err)
			break
		} else {
			log.Printf("Received message to websocket: ")
			msg := &pubsub.Message{
				Data:        message,
				Attributes:  map[string]string{"sender": userID},
				OrderingKey: orderingKey,
			}
			if _, err = topic.Publish(ctx, msg).Get(ctx); err != nil {
				log.Printf("Could not publish message: %s", err)
				return
			}
		}
	}
	cancelFunc()
	log.Printf("Shutting down wsLoop for %s...", userID)
}

func multiPubSubLoop(cctx, ctx context.Context, ws *websocket.Conn, topic *pubsub.Topic, userID string) {
	log.Printf("Starting pubSubLoop for %s...", userID)
	subscriptionName := fmt.Sprintf("client-%s-%s", userID, topic.ID())
	sub := pubSub.Subscription(subscriptionName)
	if exists, err := sub.Exists(ctx); err != nil {
		log.Printf("Error checking if sub exists: %s", err)
		return
	} else if !exists {
		log.Printf("Creating subscription: %s", subscriptionName)
		if _, err = pubSub.CreateSubscription(
			context.Background(),
			subscriptionName,
			pubsub.SubscriptionConfig{
				Topic:                 topic,
				EnableMessageOrdering: true,
			},
		); err != nil {
			log.Printf("Error creating subscription: %s", err)
			return
		}
	}
	if err := sub.Receive(cctx, func(c context.Context, m *pubsub.Message) {
		m.Ack()
		if m.Attributes["sender"] != userID && m.Attributes["messageType"] != "newPeer" {
			return
		}
		if m.Attributes["sender"] == userID && m.Attributes["messageType"] == "newPeer" {
			return
		}
		log.Println("Received message to pubSub...")
		if err := ws.Write(ctx, websocket.MessageText, m.Data); err != nil {
			log.Printf("Error writing message to %s: %s", userID, err)
			return
		}
	}); err != nil {
		log.Printf("Error setting up subscription Receive: %s", err)
	}
	log.Printf("Shutting down pubSubLoop for %s...", userID)
}
