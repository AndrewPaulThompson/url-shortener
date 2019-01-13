const MiniCssExtractPlugin = require('mini-css-extract-plugin')
const ManifestPlugin = require('webpack-manifest-plugin')
const env = process.env.NODE_ENV

module.exports = {
  MiniCssExtractPlugin: new MiniCssExtractPlugin({
      filename: env == 'prod' ? '[name].[contenthash].css' : '[name].css',
  }),
  MiniCssExtractLoader: MiniCssExtractPlugin.loader,
  ManifestPlugin: new ManifestPlugin()
}