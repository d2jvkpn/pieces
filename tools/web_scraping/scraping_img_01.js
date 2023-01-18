// var elems = document.getElementsByClassName("div-with-img");

// var data = Array.prototype.map.call(elems, e => {
//    return e.firstChild.getAttribute("src");
// });

// element_type=img, class_name=target
var elems = document.querySelectorAll(`img[target]`);
var links = Array.prototype.map.call(elems, e => e.getAttribute("src"));

var url = new URL(document.URL);

var at = new Date(); // .getTime()
var day = `${at.getFullYear()}-${at.getMonth()+1}-${at.getDate()}`;
var clock = `${at.getHours()}-${at.getMinutes()}-${at.getMinutes()}`;

var data = {
  url: document.URL,
  title: document.title,
  time: at.toISOString(),
  links: links,
}

var link = document.createElement("a");
link.href = `data:text/json,${encodeURIComponent(JSON.stringify(data))}`;
link.download = `scrap_${url.hostname}_${day}T${clock}.json`;
link.click();

// jq -r .links[] scrap_*.json | xargs -n 1 -P 8 -i wget -c {}
