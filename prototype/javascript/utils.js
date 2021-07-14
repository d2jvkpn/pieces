export function toDatetime(at=null) {
  if (!at) at = new Date();
  function padH0 (value, len=2) { return value.toString().padStart(len, '0')}

  function timezoneOffset(offset) {
    if (offset === 0) { return "Z" }

    let hour = padH0(Math.floor(Math.abs(offset) / 60));
    let minute = padH0(Math.abs(offset) % 60);
    return `${(offset < 0) ? "+" : "-"}${hour}:${minute}`;
  }

  at.date = `${at.getFullYear()}-${padH0(at.getMonth() + 1)}-${padH0(at.getDate())}`;
  at.time = `${padH0(at.getHours())}:${padH0(at.getMinutes())}:${padH0(at.getSeconds())}`;
  at.ms = padH0(at.getMilliseconds(), 3);
  at.tz = timezoneOffset(at.getTimezoneOffset());

  at.datetime = `${at.date}T${at.time}`;
  at.rfc3339 = at.datetime + `${at.tz}`;
  at.rfc3339ms = at.datetime + `.${at.ms}${at.tz}`;

  return at
}

export function addDate(at, years, months, days) {
    if (!at) at = new Date();
    at.setDate(at.getDate() + days);
    at.setMonth(at.getMonth() + months);
    at.setFullYear(at.getFullYear() + years);
    at = toDatetime(at);
    return at;
}

export function randStr(length) {
  let result   = [];
  let chars    = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789';
  let charsLen = chars.length;

  for ( let i = 0; i < length; i++ ) {
    result.push(chars.charAt(Math.floor(Math.random() * charsLen)));
  }

  return result.join('');
}

export function validUsername(name) {
  return /^[\p{Script=Han}a-zA-Z0-9-_\.]+$/u.test(name);
}

export function validUserTel(tel) {
  return /^1[3456789]\d{9}$/.test(tel);
}
