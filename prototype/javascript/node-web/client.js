/// import packages
const ws = require('ws');
const yargs = require('yargs');

// dirty hack
process.env['NODE_TLS_REJECT_UNAUTHORIZED'] = 0;

function newDate() {
  let now = new Date();
  let offset = now.getTimezoneOffset();

  if (offset === 0) { return now.toISOString() };

  now = new Date(now.setMinutes(now.getMinutes() - now.getTimezoneOffset()));

  function padH0 (value, len=2) { return value.toString().padStart(len, '0')}

  function offsetString(offset) {
    if (offset === 0) { return "Z" }

    let hour = padH0(Math.floor(Math.abs(offset) / 60));
    let minute = padH0(Math.abs(offset) % 60);
    return `${(offset < 0) ? "+" : "-"}${hour}:${minute}`;
  }

  return now.toISOString().slice(0, -1) + offsetString(offset);
}

var log = console.log;
console.log = function(...args){
  log.apply(console, [newDate()].concat(args));
};

const argv = yargs(process.argv.slice(2))
  .option("addr", {
      description: "ws address",
      // alias: "p",
      type: "string", default: 'ws://127.0.0.1:8080/ws/open/talk',
  })
  .help()
  .alias("help", "h")
  .argv;

/// variables
// console.log(argv);
let pingSecs = 5;
let checkSecs = 30;
let wsc = null;

function connect() {
  wsc = new ws.WebSocket(argv.addr);

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
    wsc.send(JSON.stringify({kind: "hello", msg: msg, name: "Evol"}));

//    ping = setInterval(function() {
//      wsc.send(JSON.stringify({kind: "ping", msg: Date.now().toString(), id: newId()}));
//    }, pingSecs*1000);

    ping = setInterval(function() {
      console.log(`--> ping: ~`);
      wsc.ping();
    }, pingSecs*1000);
  });

  wsc.on("message", function (event) {
    console.log('<-- message:', event.toString());
  });

  wsc.on("pong", function (event) {
    console.log(`<-- pong: ${event.toString() || "~"}`);
  });

  wsc.on("close", function (code, reason) {
    clearInterval(ping); // must
    console.log(`<== close: ${code}, ${reason}`);
  });

  wsc.on("error", function (error) {
    console.log(`!!! error: ${error}`);
  });

  function checkWsIsAlive(addr) {
    let interval = setInterval(function() {
      // https://developer.mozilla.org/en-US/docs/Web/API/WebSocket/readyState
      // 0:CONNECTING, 1:OPEN, 2:CLOSING, 3:CLOSED
      console.log(`~~~ wsc.readyState: ${wsc.readyState}`);

      if (wsc.readyState == 3) {
        console.log(`!!! try to reconnect to ${argv.addr}`);
        connect();
        clearInterval(interval);
        return;
      };
    }, checkSecs*1000);
  }

  checkWsIsAlive(argv.addr);
}

connect();
