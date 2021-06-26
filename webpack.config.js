/* eslint-disable quote-props */
const path = require('path');
const TerserPlugin = require('terser-webpack-plugin');

module.exports = {
  entry: {
    'europa': './web/europa/index.js'
  },
  output: {
    path: path.resolve(__dirname, './web/public/js/'),
    publicPath: '/public/',
    filename: '[name].min.js'
  },
  resolve: {
    alias: {
      '@core': path.resolve(__dirname, 'web/core'),
      '@europa': path.resolve(__dirname, 'web/europa')
    }
  },
  devtool: 'cheap-module-source-map',
  plugins: [],
  optimization: {
    minimizer: [new TerserPlugin({ extractComments: false })]
  },
  module: {
    rules: [
      {
        test: /\.(js|jsx)$/,
        exclude: /node_modules/,
        use: {
          loader: 'babel-loader'
        }
      }
    ]
  }
};
