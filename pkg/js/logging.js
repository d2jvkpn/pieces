const fs = require('fs');
const { Console } = require('console');

const {datetime, datetimeStr, toDatetime} = require('./time.js')

const loggers = [];
const logFunctions = [];
const logColorFunctions = [];

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


function dateTimeAsFilename() {
	let now = datetime();
	return `${now.date}T${now.time.replace(/:/g, '-')}`;
}

function RegisterFileLogger(path) {
	if (path == null) path = './';
	if (!fs.existsSync(path)) fs.mkdirSync(path);

	var output = fs.createWriteStream(`${path}/${dateTimeAsFilename()}.log`);
	var fileLogger = new Console(output);
	logFunctions.push(function (msg) {
		fileLogger.log(`${datetimeStr()} ${msg}`);
	});

	logColorFunctions.push(function (color, msg) {
		fileLogger.log(`${datetimeStr()} ${msg}`);
	});
	loggers.push(fileLogger);
}

function RegisterConsoleLogger() {
	var consoleLogger = new Console(process.stdout, process.stderr)
	logFunctions.push(function (msg) {
		consoleLogger.log(`${datetimeStr()} ${msg}`);
	});

	logColorFunctions.push(function (color, msg) {
		consoleLogger.log(`${BoldOn}${color}${datetimeStr()} ${msg}${AllAttributesOff}`);
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
	datetime, datetimeStr, toDatetime,
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
