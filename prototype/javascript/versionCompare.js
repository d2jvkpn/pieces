// if a > b; then return > 0
// if a < b; then return < 0
// if a == b; then return == 0
function versionCompare(a, b) {
	function version2Slice(v) {
		return v.split(".").map(x => parseInt(x)).
			filter(function (value) {return !Number.isNaN(value)});
	}

	var s1 = version2Slice(a), s2 = version2Slice(b);
	// console.log(s1, s2);
	// console.log(parseInt("12a"));

	for (let i = 0; i < Math.min(s1.length, s2.length); i++) {
		if (s1[i] > s2[i]) { return i+1;}
		if (s1[i] < s2[i]) { return -i-1;}
	}

	return 0;
}

//// test
// console.log(versionCompare("2.8.2", "2.8.10"));
// console.log(versionCompare("2.8.12", "2.8.10"));
// console.log(versionCompare("2.8.12", "2.8.10a"));