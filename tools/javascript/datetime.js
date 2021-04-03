function toDatetime(now) {
	function padH0 (value, len=2) { return value.toString().padStart(len, '0')}

	function timezoneOffset(offset) {
		if (offset === 0) { return "Z" }

		let hour = padH0(Math.floor(Math.abs(offset) / 60));
		let minute = padH0(Math.abs(offset) % 60);
		return `${(offset < 0) ? "+" : "-"}${hour}:${minute}`;
	}

	now.date = `${now.getFullYear()}-${padH0(now.getMonth() + 1)}-${padH0(now.getDate())}`;
	now.time = `${padH0(now.getHours())}:${padH0(now.getMinutes())}:${padH0(now.getSeconds())}`;
	now.ms = padH0(now.getMilliseconds(), 3);
	now.tz = timezoneOffset(now.getTimezoneOffset());

    now.datetime = `${now.date}T${now.time}`;
	now.rfc3339 = now.datetime + `${now.tz}`;
	now.rfc3339ms = now.datetime + `.${now.ms}${now.tz}`;

	return now
}

function newDatetime() {
	let now = new Date();
	return toDatetime(now);
}
