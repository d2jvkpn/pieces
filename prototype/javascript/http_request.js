let args = process.argv;

let host = args.length > 2 ? args[2] : 'https://www.random.org';
const https = host.startsWith("https:") ? require('https') : require('http');

const pathname = "/integers/?num=1&min=1&max=10&col=1&base=10&format=plain&rnd=new";

https.get(`${host}${pathname}`, (res) => {
  // console.log('statusCode:', res.statusCode);
  // console.log('headers:', res.headers);
  if (res.statusCode != 200) throw new Error(`statusCode: ${res.statusCode}`);

  res.on('data', (d) => {
    process.stdout.write(d);
  });

}).on('error', (e) => {
  console.error(e);
});
