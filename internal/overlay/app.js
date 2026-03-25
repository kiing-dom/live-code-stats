const ws  = new WebSocket("ws://localhost:8080/ws")

ws.onopen = () => console.log("Connected!")
ws.onclose = () => console.log("Disconnected!");


ws.onmessage = (event) => {
    const data = JSON.parse(event.data);

    updateStat("lines", data.lines);
    updateStat("errors", data.errors);
    updateStat("keys", data.keystrokes);
    updateStat("file", formatFile(data.file_name));
}

function updateStat(id, value) {
    const el = document.getElementById(id);
    if (el.innerText === String(value)) return;
    el.innerText = value;
    el.classList.remove("flash");
    el.style.animation = "none";
    el.style.animation = "";
    el.classList.add("flash");
}


function formatFile(path) {
    if (!path) return "-";

    return path.split(/[/\\]/).pop()
}