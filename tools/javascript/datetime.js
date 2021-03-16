function datetime() {
  let now = new Date();

  function padStart (value, len=2) { return value.toString().padStart(len, '0')}

  function timezoneOffset(offset) {
    offset = Math.abs(offset);
    if (offset === 0) { return "Z" }
    let hour = padStart(Math.floor(offset / 60));
    let minute = padStart(offset % 60);
    return `${(offset > 0) ? "+" : "-"}${hour}:${minute}`;
  }

  let year = now.getFullYear();
  let month = padStart(now.getMonth() + 1);
  let date = padStart(now.getDate());
  let hour = padStart(now.getHours());
  let minute = padStart(now.getMinutes());
  let second = padStart(now.getSeconds());
  let ms = padStart(now.getMilliseconds(), 3);
  let tz = timezoneOffset(now.getTimezoneOffset());

  now.rfc3339 = `${year}-${month}-${date}T${hour}:${minute}:${second}${tz}`;
  now.rfc3339ms = `${year}-${month}-${date}T${hour}:${minute}:${second}.${ms}${tz}`;

  return now
}
