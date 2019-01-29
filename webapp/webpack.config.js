const HtmlWebPackPlugin = require('html-webpack-plugin');
const path = require('path');
module.exports = {
  devServer: {
    contentBase: './dist',
    historyApiFallback: true,
    headers: { 'Access-Control-Allow-Origin': '*' }
  },
  entry: {
    javascript: './src/index.js',
    html: './src/index.html'
  },
  module: {
    rules: [
      {
        test: /\.css$/,
        use: ['style-loader', 'css-loader']
      },
      {
        test: /\.js$/,
        exclude: /node_modules/,
        use: {
          loader: 'babel-loader'
        }
      },
      {
        test: /\.html$/,
        use: [
          {
            loader: 'html-loader'
          }
        ]
      },
      {
        test: /\.(png|woff|woff2|eot|ttf|svg)$/,
        loader: 'url-loader?limit=100000'
      }
    ]
  },
  plugins: [
    new HtmlWebPackPlugin({
      template: './src/index.html',
      filename: './index.html'
    })
  ]
};
