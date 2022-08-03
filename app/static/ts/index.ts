import { playSong } from "./csound";

const element = document.getElementById("csound-div");
const button = document.createElement("button");
button.innerHTML = "click me";
button.onclick = playSong;
element.appendChild(button);
