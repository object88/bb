import express from 'express';
import fs from 'fs';
import path from 'path';
import webpack from 'webpack';

import Handlebars from 'handlebars';
import WebpackDevMiddleware from 'webpack-dev-middleware';
import WebpackHotMiddleware from 'webpack-hot-middleware';

import httpProxy from 'http-proxy';

const APP_PORT = 3000;
const GRAPHQL_PORT = 8081;

const templatePath = path.join(__dirname, 'public', 'index.handlebar');
const source = fs.readFileSync(templatePath, 'utf8');
const template = Handlebars.compile(source);

const html = template({apiKey: process.env.API_KEY, clientId: process.env.CLIENT_ID});

var apiProxy = httpProxy.createProxyServer();

// Serve the Relay app
var compiler = webpack({
  entry: path.resolve(__dirname, 'js', 'app.js'),
  module: {
    rules: [
      {
        exclude: /node_modules/,
        use: [{
          loader: 'babel-loader',
        }],
        //query: {plugins: [path.resolve(__dirname, 'build', 'babelRelayPlugin')]},
        // resolveLoader: {
        //   root: path.join(__dirname, 'node_modules')
        // },
        test: /\.js$/
      }
    ]
  },
  output: {filename: 'app.js', path: '/'}
});
const middleware = WebpackDevMiddleware(compiler, {
  contentBase: '/public/',
  publicPath: '/js/',
  stats: {
    colors: true,
    hash: false,
    timings: true,
    chunks: false,
    chunkModules: false,
    modules: false
  }
});
const app = express();
app.use(middleware);
app.use(WebpackHotMiddleware(compiler));
// Proxy api requests
app.use("/graphql", function(req, res) {
  console.log(`Proxying request for '${req.baseUrl}'`)
  req.url = req.baseUrl; // Janky hack...
  apiProxy.web(req, res, {
    target: { host: "localhost", port: GRAPHQL_PORT }
  });
});
app.get('/', function response(req, res) {
  console.log('Request for index.html...');
  res.write(html);
  res.end();
});
app.get('/css/*', function response(req, res) {
  console.log(`Received request '${req.url}'`);
  res.sendFile(path.join(__dirname, 'public', req.url));
});

app.listen(APP_PORT, 'localhost', function onStart(err) {
  if (err) {
    console.log(err);
  }
  console.info(`Listening on port ${APP_PORT}. Open up http://0.0.0.0:${APP_PORT}/ in your browser.`);
});
