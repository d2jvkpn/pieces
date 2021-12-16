let args = process.argv;

let host = args.length > 2 ? args[2] : 'http://127.0.0.1:8080';
host = host.startsWith("http") ? host : "http://" + host;

const http = host.startsWith("https:") ? require('https') : require('http');

const pathname = "/api/open/time";

http.get(`${host}${pathname}`, (res) => {
  // console.log('statusCode:', res.statusCode);
  // console.log('headers:', res.headers);
  if (res.statusCode != 200) throw new Error(`statusCode: ${res.statusCode}`);

  res.on('data', (d) => {
    process.stdout.write(d);
  });
}).on('error', (e) => {
  console.error(e);
});
