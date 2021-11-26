// import packages
const url = require('url');
const http = require('http');

const express = require('express');
const ws = require('ws');
const yargs = require('yargs');

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
  .option("port", {
      description: "http listening port",
      // alias: "p",
      type: "integer", default: 8080,
  })
  .help()
  .alias("help", "h")
  .argv;

/// variables
const apiPath = "/api/time";
const wsPath = '/ws/talk';

const app = express();
const server = http.Server(app);
const wss = new ws.Server({ noServer: true });

/*
// https://github.com/websockets/ws
const server = createServer({
  cert: readFileSync('/path/to/cert.pem'),
  key: readFileSync('/path/to/key.pem')
});
const wss = new WebSocketServer({ server });
*/


/// http handler
app.get(apiPath, function(req, res){
  console.log(`>>> http request ${apiPath}: "${req.connection.remoteAddress}"`);

  res.json({"time": new Date()});
});

/// websocket server
wss.on("connection", function (conn) {
  let ip = conn._socket.remoteAddress;
  let port = conn._socket.remotePort;
  let clientId = `${ip}:${port}`
  console.log(`>>> ws connection ${wsPath}: "${clientId}"`);

  function sendData(data) {
    console.log(`--> ${clientId} message: ${JSON.stringify(data)}`);
    conn.send(JSON.stringify(data));
  }

  ///
  conn.on("message", function (event) {
    let data = {};

    try {
      data = JSON.parse(event);
    } catch(err) {
      let rec = {data: event.toString(), error: err.message};

      console.log(`!!! ${clientId} message parse failed: ${JSON.stringify(rec)}`);
      // conn.close(1008, "cannot parse to json");
      return;
    }

    console.log(`<-- ${clientId} message: ${JSON.stringify(data)}`);
    if (data.kind == "hello") {
      sendData({kind: "hello", msg: "NICE TO MEET YOU"});
    } else if (data.kind == "ping") {
      let dalay = 1000;

      setTimeout(function() {
        sendData({kind: "pong", msg: Date.now().toString(), delay: dalay, id: data.id});
      }, 1000);
    } else {
      console.log(`!!! ${clientId} unknown kind: ${data.kind}`);
    }
  });

  conn.on("close", function (code, reason) {
    console.log(`client ${clientId} onclose: ${code} - ${reason}`);
  })

  conn.on("error", function (error) {
    console.log(`client ${clientId} onerror: ${error}`);
  })

  setTimeout(function() {
    sendData({kind: "goodbye", msg: "SEE YOU NEXT TIME"});
    console.log(`==> ${clientId} close connection`);
    console.log(`<<< ${clientId} session end`);
    conn.close(4001, "TIMEOUT");
  }, 30*1000);
});

///
server.on('upgrade', function upgrade(request, socket, head) {
  let req = url.parse(request.url, true);

  if (req.pathname === wsPath) {
    if (req.query) {
      console.log("==> id:", req.query["id"]);
    }

    wss.handleUpgrade(request, socket, head, function done(conn) {
      wss.emit("connection", conn, request);
    });
  } else {
    console.log(`unkown upgrade url path: ${pathname}`);
    socket.destroy();
  }
});

/// run
server.listen(argv.port, function () {
  console.log(`>>> HTTP webserver listening on "*:${argv.port}"`);
  console.log(`    wsPath: ${wsPath}, apiPath: ${apiPath}`);
});
