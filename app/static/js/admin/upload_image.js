import { disableForm, uploadFile, saveUpload } from "./upload_base.js";

document.querySelector("#upload-button").addEventListener("click", async function(e) {
  e.preventDefault();
  const imageFile = document.querySelector("#image-file").files[0];
  if (!imageFile || imageFile.type !== "image/jpeg") {
    alert("Please provide a .jpg file.");
    return;
  }
  const createdDate = document.querySelector("#created-date").value;
  if (createdDate === "") {
    alert("Please provide a created date.");
    return;
  }
  const imageType = document.querySelector("#image-type").value;

  disableForm();

  const imageFileName = `${imageFile.name}?${Date.now()}`;

  await uploadFile(imageFile, imageFileName).catch(err => {
    alert("error uploading image");
    console.log(err);
  });

  saveUpload("save_image", {
    imageFileName: imageFileName,
    createdDate: createdDate,
    caption: document.querySelector("#caption").value,
    location: document.querySelector("#location").value,
    imageType: imageType,
  }).catch(err => {
    alert("error saving image");
    console.log(err);
  });
});
