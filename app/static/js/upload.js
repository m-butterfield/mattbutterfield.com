const songName = document.querySelector("#song-name"),
  description = document.querySelector("#description"),
  uploadButton = document.querySelector("#upload-button");

uploadButton.addEventListener("click", submit);

async function submit(e) {
  e.preventDefault();
  const audioFile = document.querySelector("#audio-file").files[0];
  if (!audioFile || audioFile.type !== "audio/wav") {
    alert("Please provide a .wav file.");
    return;
  }
  const imageFile = document.querySelector("#image-file").files[0];
  if (!imageFile || imageFile.type !== "image/jpeg") {
    alert("Please provide a 640x640 .jpg file.");
    return;
  }
  if (songName.value.trim() === "") {
    alert("Please provide a song name.");
    return;
  }
  songName.disabled = true;
  description.disabled = true;
  uploadButton.disabled = true;

  const now = Date.now(),
    audioFileName = `${audioFile.name}?${now}`,
    imageFileName = `${imageFile.name}?${now}`;

  await Promise.all([
    uploadFile(audioFile, audioFileName),
    uploadFile(imageFile, imageFileName),
  ]);

  saveSong(audioFileName, imageFileName);
}

function uploadFile(fileObj, fileName) {
  return fetch("/admin/signed_upload_url", {
    method: "POST",
    body: JSON.stringify({
      fileName: fileName,
      contentType: fileObj.type,
    }),
  }).then(r => r.json()).then(data => {
    return upload(data.url, fileObj);
  }).catch(err => console.log(err));
}

function upload(url, file) {
  return fetch(url, {
    method: "PUT",
    headers: new Headers({
      "Content-Type": file.type,
    }),
    body: file,
  }).catch(err => console.log(err));
}

function saveSong(audioFileName, imageFileName) {
  fetch("/admin/save_song", {
    method: "POST",
    body: JSON.stringify({
      audioFileName: audioFileName,
      imageFileName: imageFileName,
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
    uploadButton.disabled = false;
  });
}
