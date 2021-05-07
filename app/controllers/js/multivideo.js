let firstOffer = true,
  currentPeer = new URLSearchParams(window.location.search).get("userID"),
  iceCandidates = [],
  streamsPeers = {},
  peerConnection = new RTCPeerConnection({"iceServers": [{"urls": "stun:stun.l.google.com:19302"}]}),
  ws = new WebSocket((window.location.protocol === "https:" ? "wss://" : "ws://") + window.location.host + "/multivideo/connections" + window.location.search);

let processOffer = offer => {
  peerConnection.setRemoteDescription(offer).then(() => {
    if (iceCandidates.length) {
      iceCandidates.forEach(candidate => {
        peerConnection.addIceCandidate(candidate);
      });
      iceCandidates = [];
    }
    return peerConnection.createAnswer();
  }).then(answer => {
    return peerConnection.setLocalDescription(answer);
  }).then(() => {
    ws.send(JSON.stringify(peerConnection.localDescription));
  });
};

ws.onmessage = evt => {
  const message = JSON.parse(evt.data);
  switch (message.type) {
  case "offer": {
    if (firstOffer) {
      navigator.mediaDevices.getUserMedia({video: true, audio: true}).then(stream => {
        let element = document.getElementById("local-video");
        element.srcObject = stream;
        element.play().then(() => {
          stream.getTracks().forEach(track => {
            peerConnection.addTrack(track, stream);
          });
          firstOffer = false;
          processOffer(message);
        });
      });
    } else {
      processOffer(message);
    }
    break;
  }
  case "candidate": {
    if (!peerConnection.remoteDescription) {
      iceCandidates.push(new RTCIceCandidate(message.ice));
    } else {
      peerConnection.addIceCandidate(new RTCIceCandidate(message.ice));
    }
    break;
  }
  case "newPeer": {
    streamsPeers[message.peerStreamID] = message.peerName;
    break;
  }
  case "peerList": {
    streamsPeers = message.peerList;
    break;
  }
  }
};

peerConnection.ontrack = evt => {
  if (evt.track.kind === "audio") return;
  let peerName = streamsPeers[evt.streams[0].id];
  if (!peerName || peerName === currentPeer) return;
  const elementID = peerName + "-video";
  let element = document.getElementById(elementID);
  if (!element) {
    let div = document.createElement("div");
    div.innerHTML = "<p>" + peerName + ":</p><video id=\"" + elementID + "\" autoPlay controls playsInline style=\"width: 650px; margin-bottom: 50px;\"></video>";
    document.getElementById("videos-div").appendChild(div);
    element = document.getElementById(elementID);
  } else if (element.srcObject === evt.streams[0]) {
    return;
  }
  element.srcObject = evt.streams[0];
  element.play();
};

peerConnection.onicecandidate = evt => {
  if (evt.candidate) ws.send(JSON.stringify({type: "candidate", ice: evt.candidate}));
};
