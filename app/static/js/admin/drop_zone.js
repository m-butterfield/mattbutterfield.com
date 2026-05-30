export function initDropZone() {
  const dropZone = document.getElementById("drop-zone");
  const fileInput = document.getElementById("image-file");

  ["dragenter", "dragover", "dragleave", "drop"].forEach(event => {
    dropZone.addEventListener(event, e => e.preventDefault());
  });

  dropZone.addEventListener("dragenter", () => dropZone.classList.add("drag-over"));
  dropZone.addEventListener("dragover", () => dropZone.classList.add("drag-over"));
  dropZone.addEventListener("dragleave", () => dropZone.classList.remove("drag-over"));
  dropZone.addEventListener("drop", e => {
    dropZone.classList.remove("drag-over");
    const file = e.dataTransfer.files[0];
    if (file && file.type === "image/jpeg") {
      fileInput.files = e.dataTransfer.files;
      dropZone.querySelector("p").textContent = file.name;
    }
  });

  dropZone.addEventListener("click", () => fileInput.click());

  fileInput.addEventListener("change", () => {
    if (fileInput.files[0]) {
      dropZone.querySelector("p").textContent = fileInput.files[0].name;
    }
  });
}
