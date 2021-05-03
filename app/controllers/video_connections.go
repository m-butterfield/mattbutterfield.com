package controllers

import (
	"cloud.google.com/go/pubsub"
	"context"
	"fmt"
	"log"
	"net/http"
	"nhooyr.io/websocket"
	"sort"
	"strings"
)

func VideoConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := websocket.Accept(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer closeWS(ws)
	userID := strings.ToLower(r.URL.Query().Get("userID"))
	peerID := strings.ToLower(r.URL.Query().Get("peerID"))

	peers := []string{userID, peerID}
	sort.Strings(peers)
	topicName := fmt.Sprintf("video-%s-%s", peers[0], peers[1])
	topic := pubSub.Topic(topicName)
	topic.EnableMessageOrdering = true

	ctx := context.Background()
	exists, err := topic.Exists(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if !exists {
		log.Printf("Topic %s doesn't exist - creating it", topicName)
		_, err = pubSub.CreateTopic(ctx, topicName)
		if err != nil {
			log.Fatal(err)
		}
	}

	cctx, cancelFunc := context.WithCancel(ctx)
	go wsLoop(ctx, cancelFunc, ws, topic, userID)
	pubSubLoop(cctx, ctx, ws, topic, userID)
}

func wsLoop(ctx context.Context, cancelFunc context.CancelFunc, ws *websocket.Conn, topic *pubsub.Topic, userID string) {
	log.Printf("Starting wsLoop for %s...", userID)
	orderingKey := fmt.Sprintf("%s-%s", userID, topic.ID())
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

func pubSubLoop(cctx, ctx context.Context, ws *websocket.Conn, topic *pubsub.Topic, userID string) {
	log.Printf("Starting pubSubLoop for %s...", userID)
	subscriptionName := fmt.Sprintf("%s-%s", userID, topic.ID())
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
		if m.Attributes["sender"] == userID {
			log.Println("skipping message from self")
			return
		}
		log.Printf("Received message to pubSub: ")
		if err := ws.Write(ctx, websocket.MessageText, m.Data); err != nil {
			log.Printf("Error writing message to %s: %s", userID, err)
			return
		}
	}); err != nil {
		log.Printf("Error setting up subscription Receive: %s", err)
	}
	log.Printf("Shutting down pubSubLoop for %s...", userID)
}

func closeWS(ws *websocket.Conn) {
	// can check if already closed here
	if err := ws.Close(websocket.StatusNormalClosure, ""); err != nil {
		log.Printf("Error closing: %s", err)
	}
}
