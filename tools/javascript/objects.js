class Error {
	constructor (ok, errmsg, data=null) {
		this.ok = ok;          // Boolean
		this.errmsg = errmsg;  // String
		this.data = data;      // Object
	}
}

class ResponseData {
	constructor (requestId, code, message, errmsg, data=null) {
		this.requestId = requestId; // String, request id
		this.code = code;           // Number, response code, 0 for ok
		this.message = message;     // String, message show to user
		this.errmsg = errmsg;       // String, program error message
		this.data = data;           // Object, response data
	}
	
	toJson() { return JSON.stringify(this) }
}
