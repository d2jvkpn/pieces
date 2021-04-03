const fs = require('fs');
const { Console } = require('console');

const loggers = [];
const logFunctions = [];
const logColorFunctions = [];

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

console.log = function (msg) {
	logFunctions.forEach(func => { func(msg); });
}

console.logColor = function (color, msg) {
	logColorFunctions.forEach(func => { func(color, msg); });
}

const AllAttributesOff = '\x1b[0m';
const BoldOn = '\x1b[1m';
const Black = '\x1b[30m';
const Red = '\x1b[31m';
const Green = '\x1b[32m';
const Yellow = '\x1b[33m';
const Blue = '\x1b[34m';
const Magenta = '\x1b[35m';
const Cyan = '\x1b[36m';
const White = '\x1b[37m';


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
	now.rfc3339 =  `${now.datetime}${now.tz}`;
	now.rfc3339ms = `${now.datetime}.${now.ms}${now.tz}`;

	return now
}

function RegisterFileLogger(path) {
	if (path == null) path = './log';
	if (!fs.existsSync(path)) fs.mkdirSync(path);
	
	let now = newDatetime();

	var output = fs.createWriteStream(`${path}/${now.datetime.replace(/:/g, '-')}.log`);
	var fileLogger = new Console(output);
	logFunctions.push(function (msg) {
		fileLogger.log(`${now.rfc3339ms} ${msg}`);
	});

	logColorFunctions.push(function (color, msg) {
		fileLogger.log(`${now.rfc3339ms} ${msg}`);
	});
	loggers.push(fileLogger);
}

function RegisterConsoleLogger() {
	var consoleLogger = new Console(process.stdout, process.stderr)
	logFunctions.push(function (msg) {
		consoleLogger.log(`${timeToString()} ${msg}`);
	});
	
	let now = newDatetime();

	logColorFunctions.push(function (color, msg) {
		consoleLogger.log(`${BoldOn}${color}${now.rfc3339ms} ${msg}${AllAttributesOff}`);
	});
	loggers.push(consoleLogger);
}


function trace(msg) {
	console.logColor(Green, `[TRACE] ${msg}`);
}

function debug(msg) {
	console.logColor(Blue, `[DEBUG] ${msg}`);
}

function info(msg) {
	console.logColor(White, `[INFO] ${msg}`);
}

function warn(msg) {
	console.logColor(Yellow, `[WARN] ${msg}`);
}

function error(msg) {
	console.logColor(Red, `[ERROR] ${msg}`);
	console.trace();
}

function fatal(msg) {
	console.logColor(Red, `[FATAL] ${msg}`);
	console.trace();
	process.exit(1);
}

module.exports = {
	//Functions
	RegisterFileLogger,
	RegisterConsoleLogger,
	newDatetime, toDatetime,
	trace, debug, info, warn, error, fatal,

	//Variables
	AllAttributesOff,
	BoldOn,
	Black,
	Red,
	Green,
	Yellow,
	Blue,
	Magenta,
	Cyan,
	White
}
