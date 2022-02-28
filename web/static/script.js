window.addEventListener("DOMContentLoaded", (_) => {
    const ws = new WebSocket("ws://" + window.location.host + "/ws");
    const thread = document.getElementById("chat-thread");

    ws.addEventListener("message", function (e) {
        const data = JSON.parse(e.data);
        const msg = document.createElement("li");
        msg.innerHTML = `<strong>${data.username}</strong>: ${data.text}`;

        thread.appendChild(msg);
        thread.scrollTop = thread.scrollHeight;
    });

    const form = document.getElementById("chat-form");
    form.addEventListener("submit", function (event) {
        event.preventDefault();
        const username = document.getElementById("input-username");
        const text = document.getElementById("input-text");
        ws.send(
            JSON.stringify({
                username: username.value,
                text: text.value,
            })
        );
        text.value = "";
    });
});
