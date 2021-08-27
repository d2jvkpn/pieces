function range(start, end, step) {
  step = step * (end - start) > 0 ? step : -step;
  let r = Math.abs(end - start);
  let s = Math.abs(step);
  if (r < s || step == 0) return [start];

  let l = Math.floor((r + s) / s);
  // console.log(`start=${start}, end=${end}, step=${step}, l=${l}`)

  return Array(l).fill().map((_, idx) => start + idx*step);
}

// range(18, 9, 2);
// range(10, 18, 2);
// range(10, 18, 3);
// range(18, 10, 3);
// range(10, 10, 2);
