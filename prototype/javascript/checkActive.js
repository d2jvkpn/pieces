let n = 0;
let maxN = 5;
let interval = 2000; // 2s

let checkActive = setInterval(
  () => {
    n++;
    if (n > maxN) {
      console.log("!!! clearInterval");
      clearInterval(checkActive);
      return;
    }
    console.log(">>>", n, new Date());
  }
, interval); // 2s

