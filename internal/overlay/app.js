async function fetchStats() {
    const res = await fetch("http://localhost:8080/stats");
    const data = res.json();

    document.getElementById("lines").innerText = data.lines;
    document.getElementById("errors").innerText = data.errors;
    document.getElementById("keys").innerText = data.keys;
}

setInterval(fetchStats, 500);