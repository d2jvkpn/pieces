function datetime() {
  function padH0 (value, len=2) { return value.toString().padStart(len, '0')}

  function timezoneOffset(offset) {
    if (offset === 0) { return "Z" }

    let hour = padH0(Math.floor(Math.abs(offset) / 60));
    let minute = padH0(Math.abs(offset) % 60);
    return `${(offset < 0) ? "+" : "-"}${hour}:${minute}`;
  }

  let now = new Date();

  let date = `${now.getFullYear()}-${padH0(now.getMonth() + 1)}-${padH0(now.getDate())}`;
  let clock = `${padH0(now.getHours())}:${padH0(now.getMinutes())}:${padH0(now.getSeconds())}`;
  let ms = padH0(now.getMilliseconds(), 3);
  let tz = timezoneOffset(now.getTimezoneOffset());

  now.tz = tz;
  now.rfc3339 = `${date}T${clock}${tz}`;
  now.rfc3339ms = `${date}T${clock}.${ms}${tz}`;

  return now
}
