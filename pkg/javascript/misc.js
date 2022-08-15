var crypto = require("crypto");
export function randStr(length) {
//  let result   = [];
//  let chars    = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789';
//  let charsLen = chars.length;

//  for ( let i = 0; i < length; i++ ) {
//    result.push(chars.charAt(Math.floor(Math.random() * charsLen)));
//  }

//  return result.join('');
  return crypto.randomBytes(Math.round(length/2)).toString('hex').slice(0, length);
}
