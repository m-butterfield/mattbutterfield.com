let chipsContainer;

export function initTagInput() {
  chipsContainer = document.getElementById("tag-chips");
  const input = document.getElementById("tag-input");
  const hidden = document.getElementById("tags");
  const addButton = document.getElementById("add-tag-button");

  if (hidden.value) {
    hidden.value.split(",").map(t => t.trim()).filter(t => t).forEach(tag => addChip(tag));
  }

  addButton.addEventListener("click", () => handleAdd());
  input.addEventListener("keydown", e => {
    if (e.key === "Enter") {
      e.preventDefault();
      handleAdd();
    }
  });

  function handleAdd() {
    const value = input.value.trim();
    if (!value || hasChip(value)) return;
    addChip(value);
    input.value = "";
    input.focus();
    syncHidden();
  }

  function addChip(text) {
    const chip = document.createElement("span");
    chip.className = "tag-chip";
    chip.dataset.tag = text.toLowerCase();

    const textSpan = document.createElement("span");
    textSpan.textContent = text;
    chip.appendChild(textSpan);

    const removeBtn = document.createElement("button");
    removeBtn.type = "button";
    removeBtn.className = "tag-chip-remove";
    removeBtn.textContent = "\u00D7";
    removeBtn.addEventListener("click", () => {
      chip.remove();
      syncHidden();
    });
    chip.appendChild(removeBtn);

    chipsContainer.appendChild(chip);
  }

  function hasChip(text) {
    const key = text.toLowerCase();
    return Array.from(chipsContainer.children).some(chip => chip.dataset.tag === key);
  }

  function syncHidden() {
    hidden.value = Array.from(chipsContainer.querySelectorAll(".tag-chip"))
      .map(chip => chip.firstChild.textContent)
      .join(", ");
  }
}

export function getTags() {
  return Array.from(chipsContainer.querySelectorAll(".tag-chip"))
    .map(chip => chip.firstChild.textContent);
}
