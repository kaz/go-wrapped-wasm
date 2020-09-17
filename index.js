"use strict";

import wasmbin from "./wrapped.wasm";

import "simplemde/dist/simplemde.min.css";
import SimpleMDE from "simplemde";

(async () => {
	const wasm = await wasmbin;

	new SimpleMDE({
		element: document.querySelector("textarea"),
		previewRender: wasm.md2html,
	});

	try {
		wasm.errortest();
	} catch (e) {
		console.info(e);
	}

	try {
		wasm.panictest();
	} catch (e) {
		console.info(e);
	}
})();
