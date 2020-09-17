const path = require("path");

module.exports = {
	mode: "development",
	entry: "./index.js",
	output: {
		filename: "app.js",
		path: path.resolve(__dirname, "./dist/"),
		publicPath: "./dist/",
	},
	module: {
		rules: [{
			test: /\.css$/i,
			use: ["style-loader", "css-loader"],
		}, {
			test: /wrapped\.wasm$/i,
			type: "javascript/auto",
			use: ["./go_wasm_loader.js", "file-loader"],
		}],
	},
};
