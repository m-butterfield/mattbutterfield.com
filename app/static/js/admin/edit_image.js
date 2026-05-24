import { disableForm } from "./upload_base.js";

document.querySelector("#save-button").addEventListener("click", async function(e) {
  e.preventDefault();

  const createdDate = document.querySelector("#created-date").value;
  if (createdDate === "") {
    alert("Please provide a created date.");
    return;
  }

  disableForm();

  const tagsValue = document.querySelector("#tags").value;
  const tags = tagsValue ? tagsValue.split(",").map(t => t.trim()).filter(t => t) : [];

  fetch("/admin/update_image", {
    method: "POST",
    headers: new Headers({"Content-Type": "application/json"}),
    body: JSON.stringify({
      imageID: document.querySelector("#image-id").value,
      caption: document.querySelector("#caption").value,
      location: document.querySelector("#location").value,
      createdDate: createdDate,
      camera: document.querySelector("#camera").value,
      lens: document.querySelector("#lens").value,
      film: document.querySelector("#film").value,
      tags: tags,
    }),
  }).then(resp => {
    if (resp.ok) {
      return resp.json();
    }
    throw new Error("update failed");
  }).then(data => {
    if (data.redirect) {
      window.location.href = data.redirect;
    }
  }).catch(err => {
    alert("error saving image");
    console.log(err);
  });
});
