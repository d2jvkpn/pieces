let href = new URL(window.location.href);
let search = new URLSearchParams(href.search);
let elem = search.get("token");
