function padH0 (value, len=2) { return value.toString().padStart(len, '0')}

  function timezoneOffset(offset) {
    if (offset === 0) { return "Z" }

    let hour = padH0(Math.floor(Math.abs(offset) / 60));
    let minute = padH0(Math.abs(offset) % 60);
    return `${(offset < 0) ? "+" : "-"}${hour}:${minute}`;
  }

  let now = new Date();
  
  now.date = `${now.getFullYear()}-${padH0(now.getMonth() + 1)}-${padH0(now.getDate())}`;
  now.time = `${padH0(now.getHours())}:${padH0(now.getMinutes())}:${padH0(now.getSeconds())}`;
  now.ms = padH0(now.getMilliseconds(), 3);
  now.tz = timezoneOffset(now.getTimezoneOffset());

  now.rfc3339 = `${now.date}T${now.time}${now.tz}`;
  now.rfc3339ms = `${now.date}T${now.time}.${now.ms}${now.tz}`;

  return now
}
