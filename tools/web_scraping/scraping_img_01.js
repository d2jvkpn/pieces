// let elems = document.getElementsByClassName("div-with-img");

// let data = Array.prototype.map.call(elems, e => {
//    return e.firstChild.getAttribute("src");
// });

// element_type=img, class_name=target
var elems = document.querySelectorAll(`img[target]`);

let data = Array.prototype.map.call(elems, e => e.getAttribute("src"));

const link = document.createElement("a");
link.href = `data:text/json,${encodeURIComponent(JSON.stringify(data))}`;

let ts = new Date().getTime()
link.download = `links.${ts}.json`;
link.click();

// jq -r .[] links.xxxx.json | xargs -i wget -c {}
// jq -r .[] links.xxxx.json | xargs -n 1 -P 8 -i wget -c {}
