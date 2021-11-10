/// import packages
const ws = require('ws');

/// variables
let addr = 'ws://127.0.0.1:9000/ws/talk';
let pingSec = 5;
let wsc = null;

function connect() {
  wsc = new ws.WebSocket();

  function checkAlive(addr) {
    let interval = setInterval(function() {
      // https://developer.mozilla.org/en-US/docs/Web/API/WebSocket/readyState
      // 0:CONNECTING, 1:OPEN, 2:CLOSING, 3:CLOSED
      if (wsc.readyState === 3) {
        console.log(`try to reconnect to ${addr}`);
        clearInterval(interval);
        connect();
      };
    }, pingSec * 1000);
  }

  checkAlive();
}

connect();

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
  }, pingSec*1000);
});

wsc.on("message", function (event) {
  console.log('<-- message: %s', event.toString());
});

wsc.on("close", function (code, reason) {
  clearInterval(ping); // must
  console.log(`<== close: ${code}, ${reason}`);
});
