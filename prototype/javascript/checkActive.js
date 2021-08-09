////
let n = 0;
let maxN = 5;
let interval = 2000;

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
, interval);

////
let activity = {
  interval: 20*1000,
  interval2: 30*1000,
  delta: 10*1000,
  last: new Date(),
  n: 0,
  m: 2,
};

updateActivity = () => {
  activity.last = new Date();
  activity.n+=1;
}

let cycle = setInterval(
  () => {
    let dura = new Date() - activity.last;
    let d1 = activity.interval2 + activity.delta;
    let d2 = activity.m*activity.interval2 + activity.delta;

    if (activity.n === 0) {
      if (dura > d1) console.log(`INFO: not initialized!`);
      return;
    }

    if (dura > d1) {
       console.log(`WARN: no activity for ${dura/1000}s, lastActivity=${activity.last}`);
    }

    if (dura > d2) {
      console.log(`ERROR: no activity for ${dura/1000}s, lastActivity=${activity.last}`);
      clearInterval(cycle);
      return;
    }
  },
activity.interval);
