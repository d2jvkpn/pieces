function convert(ts) {
  function padH0 (value, len=2) { return value.toString().padStart(len, '0')}

  function timezoneOffset(offset) {
    if (offset === 0) { return "Z" }

    let hour = padH0(Math.floor(Math.abs(offset) / 60));
    let minute = padH0(Math.abs(offset) % 60);
    return `${(offset < 0) ? "+" : "-"}${hour}:${minute}`;
  }

  ts.date = `${ts.getFullYear()}-${padH0(ts.getMonth() + 1)}-${padH0(ts.getDate())}`;
  ts.time = `${padH0(ts.getHours())}:${padH0(ts.getMinutes())}:${padH0(ts.getSeconds())}`;
  ts.ms = padH0(ts.getMilliseconds(), 3);
  ts.tz = timezoneOffset(ts.getTimezoneOffset());

  ts.datetime = `${ts.date}T${ts.time}`;
  ts.rfc3339 = ts.datetime + `${ts.tz}`;
  ts.rfc3339ms = ts.datetime + `.${ts.ms}${ts.tz}`;

  return ts;
}

function now() {
  let ts = new Date();
  return convert(ts);
}

function nowStr() {
  let ts = now();
  return ts.rfc3339ms;
}

// timestamp as file basename
function basename() {
  let ts = now();
  return `${ts.date}T${ts.time.replace(/:/g, '-')}`;
}

module.exports = {
  convert, now, nowStr, basename,
}

/*
const logts = require('log-timestamp');

logts(function() {
  return "[" + timestamp.nowStr() + "] %s";
});
*/
