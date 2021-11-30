document.querySelector("#upload-button").addEventListener("click", submit);

const songName = document.querySelector("#song-name"),
  description = document.querySelector("#description");

function submit(e) {
  e.preventDefault();
  const audioFile = document.querySelector("#audio-file").files[0];
  if (!audioFile || audioFile.type !== "audio/wav") {
    alert("Please provide a wav file.");
    return;
  }
  const fileName = `${audioFile.name}?${Date.now()}`;
  if (songName.value.trim() === "") {
    alert("Please provide a song name.");
    return;
  }
  songName.disabled = true;
  description.disabled = true;

  fetch("/admin/signed_upload_url", {
    method: "POST",
    body: JSON.stringify({
      fileName: fileName,
      contentType: audioFile.type,
    }),
  }).then(r => r.json()).then(data => {
    upload(data.url, audioFile, fileName);
  }).catch(err => console.log(err));
}

function upload(url, file, fileName) {
  fetch(url, {
    method: "PUT",
    headers: new Headers({
      "Content-Type": file.type,
    }),
    body: file,
  }).then(resp => {
    if (resp.status === 200) {
      saveSong(fileName);
    }
  }).catch(err => console.log(err));
}

function saveSong(fileName) {
  fetch("/admin/save_song", {
    method: "POST",
    body: JSON.stringify({
      fileName: fileName,
      songName: songName.value,
      description: description.value,
    }),
  }).then(resp => {
    if (resp.status === 201) {
      alert("success!");
    }
  }).catch(err => {
    alert("error saving song");
    console.log(err);
  }).finally(() => {
    songName.disabled = false;
    description.disabled = false;
  });
}
