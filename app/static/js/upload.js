document.querySelector("#upload-button").addEventListener("click", submit);

function submit(e) {
  e.preventDefault();
  const audioFile = document.querySelector("#audio-file").files[0];
  const fileName = `${audioFile.name}?${Date.now()}`;

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
  const songName = document.querySelector("#song-name").value;
  const description = document.querySelector("#description").value;
  fetch("/admin/save_song", {
    method: "POST",
    body: JSON.stringify({
      fileName: fileName,
      songName: songName,
      description: description,
    }),
  }).then(resp => {
    if (resp.status === 201) {
      alert("success!");
    }
  }).catch(err => console.log(err));
}
