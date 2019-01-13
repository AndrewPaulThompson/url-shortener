const plugins = require('./plugins')

const JSLoader = {
  test: /\.js$/,
  exclude: /node_modules/,
  use: {
    loader: 'babel-loader',
    options: {
      presets: ['@babel/preset-env']
    }
  }
}

const CSSLoader = {
  test: /\.pcss$/,
  use: [
    plugins.MiniCssExtractLoader,
    {
      loader: 'css-loader',
      options: {importLoaders: 1},
    },
    {
      loader: 'postcss-loader',
      options: {
        config: {
          path: __dirname + '/postcss.config.js'
        }
      },
    },
    {
      loader: 'sass-loader'
    }
  ]
}

module.exports = {
  JSLoader: JSLoader,
  CSSLoader: CSSLoader
}