const path = require('path')
const loaders = require('./webpack/loaders')
const plugins = require('./webpack/plugins')
const env = process.env.NODE_ENV

module.exports = {
  entry: './src/js/index.js',
  module: {
    rules: [
      loaders.JSLoader,
      loaders.CSSLoader,
    ]
  },
  plugins: [
    plugins.MiniCssExtractPlugin,
    plugins.ManifestPlugin
  ],
  output: {
    filename: env == 'prod' ? '[name].[contenthash].js' : '[name].js',
    path: path.resolve(__dirname, 'public')
  },
}