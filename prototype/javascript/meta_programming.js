var m = {
  "a": [{".x": 1, ".y": 2}],
  "b": [
    {"[0].x": 99},
  ],
};

var data = {
  a: {x: 0},
  b: [{x: 0}],
};

console.log(">>> m:", JSON.stringify(m))
console.log(">>> data:", JSON.stringify(data));

Object.keys(m).forEach(f => {
  m[f].forEach(v => {
    Object.keys(v).forEach(k => {
      let code = `data.${f}${k} = ${v[k]}`;
      console.log(`~~~ ${code}`);
      eval(code);
    });
  });
});

console.log(">>> data:", JSON.stringify(data));
