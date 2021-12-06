import { disableForm, uploadFile, saveUpload } from "./upload_base.js";

document.querySelector("#upload-button").addEventListener("click", async function(e) {
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
  const createdDate = document.querySelector("#created-date").value;
  if (createdDate === "") {
    alert("Please provide a created date.");
    return;
  }
  const songName = document.querySelector("#song-name").value;
  if (songName.trim() === "") {
    alert("Please provide a song name.");
    return;
  }

  disableForm();

  const now = Date.now(),
    audioFileName = `${audioFile.name}?${now}`,
    imageFileName = `${imageFile.name}?${now}`;

  await Promise.all([
    uploadFile(audioFile, audioFileName),
    uploadFile(imageFile, imageFileName),
  ]).catch(err => {
    alert("error uploading song");
    console.log(err);
  });

  saveUpload("save_song", {
    audioFileName: audioFileName,
    imageFileName: imageFileName,
    createdDate: createdDate,
    songName: songName,
    description: document.querySelector("#description").value,
  }).catch(err => {
    alert("error saving song");
    console.log(err);
  });
});
