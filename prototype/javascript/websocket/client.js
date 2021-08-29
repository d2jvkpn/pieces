/// import packages
const ws = require('ws');

/// variables
const wsc = new ws.WebSocket('ws://127.0.0.1:9000/ws/talk');

let ping = null;

let id = 0;
function newId() {
  id += 1;
  return id;
}

// websocket client
wsc.on("open", function () {
  let msg = "My name is Evol";
  console.log(`--> message: hello, ${msg}`);
  wsc.send(JSON.stringify({kind: "hello", msg: msg}));

  ping = setInterval(function() {
    wsc.send(JSON.stringify({kind: "ping", msg: Date.now(), id: newId()}));
  }, 5*1000);
});

wsc.on("message", function (event) {
  console.log('<-- message: %s', event.toString());
});

wsc.on("close", function (code, reason) {
  clearInterval(ping); // must
  console.log(`<== close: ${code}, ${reason}`);
});
