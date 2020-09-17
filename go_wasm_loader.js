"use strict";

const fs = require("fs");

const runtime = (wasmPath) => new Promise(async (resolve) => {
	window.__go_startup_callback = () => {
		const exports = new Proxy(window.__go_exports, {
			get(target, prop) {
				if (typeof target[prop] === "function") {
					return (...args) => {
						let raisedErr;
						const resp = target[prop].bind({
							__go_error_callback(err) {
								raisedErr = err;
							}
						})(...args);

						if (raisedErr) {
							throw raisedErr;
						}
						return resp;
					};
				}
				return target[prop];
			}
		});

		delete window.__go_exports;
		delete window.__go_startup_callback;
		resolve(exports);
	};

	const go = new Go();
	const module = await WebAssembly.instantiateStreaming(fetch(wasmPath), go.importObject);
	go.run(module.instance);
});

module.exports = async (wasmPath) => {
	return `
		const require = undefined;
		${await fs.promises.readFile("/usr/local/opt/go/libexec/misc/wasm/wasm_exec.js", "UTF-8")}
		export default (${runtime.toString()})(${wasmPath.replace(/^export default (.+?);/, "$1")});
	`;
};
