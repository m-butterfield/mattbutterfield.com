document.body.addEventListener("play", event => {
  Array.from(document.body.getElementsByTagName("audio")).forEach(element => {
    if (element !== event.target) {
      element.pause();
    }
  });
}, true);
document.body.addEventListener("ended", event => {
  const nextID = event.target.dataset.next;
  if (nextID) {
    document.getElementById(nextID).play();
  }
}, true);
