// Buffer.alloc(number);

let str1 = "hello, world, 你好!";
let buff1 = Buffer.from(str1);
let eStr1 = buff1.toString('base64');

let buff2 = Buffer.from(eStr1, 'base64');
let str2 = buff2.toString();
