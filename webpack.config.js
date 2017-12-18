// Based on https://github.com/preboot/angular-webpack/blob/master/webpack.config.js

// Helper: root() is defined at the bottom
var path = require('path');
var webpack = require('webpack');
var CopyWebpackPlugin = require('copy-webpack-plugin');

var ENV = process.env.ENV;
var isDeploy = ENV === 'deploy';
var isProd = (ENV === 'prod') || isDeploy;
var isDev = (ENV === 'dev') || (!isProd);
var envText = isDeploy ? 'deploy' : isProd ? 'prod' : 'dev'

console.debug('Resolved build environment: '  + envText);

// Webpack Plugins
var CommonsChunkPlugin = webpack.optimize.CommonsChunkPlugin;
var autoprefixer = require('autoprefixer');
var HtmlWebpackPlugin = require('html-webpack-plugin');

module.exports = () => {
  var config = {};

  if (isDev) {
    config.devtool = 'eval-source-map';
  } else {
    config.devtool = 'source-map';
  }

  config.entry = {
    app: root('websnake-frontend', 'main.ts'),
    polyfill: root('websnake-frontend', 'polyfill.ts'),
    vendor: root('websnake-frontend', 'vendor.ts'),
  };

  config.output = {
    path: root('dist', envText),
    filename: '[name].bundle.js',
  };

  config.resolve = {
    extensions: ['.ts', '.js', '.json', '.css', '.scss', '.html'],
    alias: [{
      alias: 'jquery',
      name: 'jquery/src/jquery',
    }],
    modules: [
      root('.'),
      'node_modules',
    ],
  };

  var atlConfigFile = root('websnake-frontend', 'tsconfig.json');
  config.module = {
    rules: [
      {test: require.resolve("jquery"), loaders: ["expose-loader?$", "expose-loader?jQuery"] },
      {test: /\.ts$/, loader: 'awesome-typescript-loader?configFileName=' + atlConfigFile},
      {test: /\.(gif|png|jpg|svg|woff|woff2|ttf|eot)$/, loader: 'file-loader'},
      {test: /\.json$/, loader: 'json-loader'},
      {test: /\.(css|scss|sass)$/, loaders: ['to-string-loader', 'css-loader', 'sass-loader']},
      {test: /\.html$/, loader: 'html-loader'}
    ]
  };

  config.plugins = [
    new CopyWebpackPlugin([{
      from: root('websnake-frontend', 'static'),
      flatten: true,
    }]),
    new CommonsChunkPlugin({
      name: ['vendor', 'polyfill'],
    }),
    new webpack.ProvidePlugin({
      jQuery: 'jquery',
      $: 'jquery',
      jquery: 'jquery',
    }),
    new webpack.LoaderOptionsPlugin({
      options: {
        postcss: [
          require('autoprefixer')(),
        ]
      }
    }),
    new webpack.DefinePlugin({
      'process.env': {
        ENV: JSON.stringify(envText),
      }
    }),
    new webpack.NoEmitOnErrorsPlugin(),
    new HtmlWebpackPlugin({
      template: root('websnake-frontend', 'index.html'),
      chunksSortMode: 'dependency',
    }),
  ];

  if (isProd) {
    config.plugins.push(
      new webpack.optimize.UglifyJsPlugin({sourceMap: true, mangle: { keep_fnames: true }})
    );
  }

  config.devServer = {
    contentBase: root('websnake-frontend'),
    historyApiFallback: true,
    quiet: false,
    stats: 'normal', // none (or false), errors-only, minimal, normal (or true) and verbose
    port: process.env.WEBPACK_DEV_SERVER_PORT || 8888,
  };

  return config
}


// Helper functions
function root(args) {
  args = Array.prototype.slice.call(arguments, 0);
  return path.join.apply(path, [__dirname].concat(args));
}
