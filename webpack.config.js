const path = require("path");

module.exports = {
  entry: "./app/static/ts/index.ts",
  devtool: "inline-source-map",
  module: {
    rules: [
      {
        test: /\.tsx?$/,
        use: "ts-loader",
        exclude: /node_modules/,
      },
      {
        test: /\.csd$/,
        use: "raw-loader",
        include: path.resolve(__dirname, "app", "static", "csound"),
      },
    ],
  },
  resolve: {
    extensions: [".tsx", ".ts", ".js"],
  },
  output: {
    filename: "bundle.js",
    path: path.resolve(__dirname, "app", "static", "js", "dist"),
  },
};
