export function uploadFile(fileObj, fileName) {
  return fetch("/admin/signed_upload_url", {
    method: "POST",
    body: JSON.stringify({
      fileName: fileName,
      contentType: fileObj.type,
    }),
  }).then(r => r.json()).then(data => {
    return upload(data.url, fileObj);
  });
}

function upload(url, file) {
  return fetch(url, {
    method: "PUT",
    headers: new Headers({
      "Content-Type": file.type,
    }),
    body: file,
  });
}

export function saveUpload(action, body) {
  return fetch(`/admin/${action}`, {
    method: "POST",
    body: JSON.stringify(body),
  }).then(resp => {
    if (resp.status === 201) {
      alert("success!");
    }
  }).finally(enableForm);
}

export function disableForm() {
  document.querySelectorAll(".upload-form-element").forEach(e => e.disabled = true);
}

function enableForm() {
  document.querySelectorAll(".upload-form-element").forEach(e => e.disabled = false);
}
