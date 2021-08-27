const express = require('express');
const http = require('http');
const ws = require('ws');
const url = require('url');

///
var httpPort = 9000;
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


///
app.get(apiPath, function(req, res){
  console.log(`>>> http request ${apiPath}: "${req.connection.remoteAddress}"`);

  res.json({"time": new Date()});
});

///
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

      console.error(`<-- ${clientId} message parse failed: ${JSON.stringify(rec)}`);
      // conn.close(1008, "cannot parse to json");
      return;
    }

    console.log(`<-- ${clientId} message: ${JSON.stringify(data)}`);
    if (data.kind == "hello") {
      sendData({kind: "hello", msg: "NICE TO MEET YOU"});
    } else if (data.kind == "ping") {
      let dalay = 1000;

      setTimeout(function() {
        sendData({kind: "pong", msg: Date.now(), delay: dalay, id: data.id});
      }, 1000);
    } else {
      console.error(`!!! ${clientId} unknown kind: ${data.kind}`);
    }
  });

  setTimeout(function() {
    sendData({kind: "goodbye", msg: "SEE YOU NEXT TIME"});
    console.warn(`==> ${clientId} close connection`);
    conn.close(4001, "TIMEOUT");
  }, 30*1000);

});

///
server.on('upgrade', function upgrade(request, socket, head) {
  const { pathname } = url.parse(request.url);

  if (pathname === wsPath) {
    wss.handleUpgrade(request, socket, head, function done(conn) {
      wss.emit("connection", conn, request);
    });
  } else {
    console.warn(`unkown upgrade url path: ${pathname}`)
    socket.destroy();
  }
});

///
server.listen(httpPort, function () {
  console.log(`>>> HTTP webserver listening on "*:${httpPort}"`);
  console.log(`    wsPath: ${wsPath}, apiPath: ${apiPath}`);
});
