let peerConnection = new RTCPeerConnection({"iceServers": [{"urls": "stun:stun.l.google.com:19302"}]}),
    ws = new WebSocket(window.location.protocol === "https:" ? "wss://" : "ws://" + window.location.host + '/video/connections' + window.location.search);

navigator.mediaDevices.getUserMedia({video: true, audio: true}).then(stream => {
  let element = document.getElementById('local_video');
  element.srcObject = stream;
  element.play().then(() => {
    stream.getTracks().forEach(track => peerConnection.addTrack(track, stream));
    peerConnection.onnegotiationneeded = () => {
      peerConnection.createOffer().then(offer => {
        return peerConnection.setLocalDescription(offer);
      }).then(() => {
        ws.send(JSON.stringify(peerConnection.localDescription));
      });
    }
  });
});

peerConnection.ontrack = evt => {
  let element = document.getElementById('remote_video');
  if (element.srcObject === evt.streams[0]) return;
  element.srcObject = evt.streams[0];
  element.play();
};

peerConnection.onicecandidate = evt => {
  if (evt.candidate) ws.send(JSON.stringify({type: 'candidate', ice: evt.candidate}));
};

ws.onmessage = (evt) => {
  const message = JSON.parse(evt.data);
  switch (message.type) {
    case 'offer': {
      peerConnection.setRemoteDescription(message).then(() => {
        return peerConnection.createAnswer()
      }).then(answer => {
        return peerConnection.setLocalDescription(answer)
      }).then(() => {
        ws.send(JSON.stringify(peerConnection.localDescription));
      });
      break;
    }
    case 'answer': {
      peerConnection.setRemoteDescription(message);
      break;
    }
    case 'candidate': {
      peerConnection.addIceCandidate(new RTCIceCandidate(message.ice));
      break;
    }
  }
};
