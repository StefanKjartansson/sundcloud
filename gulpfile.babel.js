import del from 'del';
import eslint from 'gulp-eslint';
import glp from 'gulp-load-plugins'
import gulp from 'gulp';
import gutil from 'gulp-util';
import webpack from 'webpack';
import webpackConfig from './webpack.config.js';
import {FRONTEND_DIR, PROJECT, DIST_PATH} from './config';

const plugins = glp();


gulp.task('clean', (cb) => {
  del(`${DIST_PATH}/*`, cb);
});

// The development server (the recommended option for development)
gulp.task('watch', ['webpack-dev-server']);

// Production build
gulp.task('build', ['clean', 'webpack:build']);

gulp.task('webpack:build', (callback) => {
  // modify some webpack config options
  let myConfig = Object.create(webpackConfig);
  myConfig.plugins = myConfig.plugins.concat(
    new webpack.DefinePlugin({
      'process.env': {
        // This has effect on the react lib size
        'NODE_ENV': JSON.stringify('production')
      }
    }),
    new webpack.optimize.DedupePlugin(),
    new webpack.optimize.UglifyJsPlugin()
  );

  // run webpack
  webpack(myConfig, (err, stats) => {
    if (err) throw new gutil.PluginError('webpack:build', err);
    gutil.log('[webpack:build]', stats.toString({
      colors: true
    }));
    callback();
  });
});

// modify some webpack config options
let myDevConfig = Object.create(webpackConfig);
myDevConfig.devtool = 'sourcemap';
myDevConfig.debug = true;

// create a single instance of the compiler to allow caching
let devCompiler = webpack(myDevConfig);

gulp.task('webpack:build-dev', (callback) => {
  // run webpack
  devCompiler.run(function(err, stats) {
    if (err) throw new gutil.PluginError('webpack:build-dev', err);
    gutil.log('[webpack:build-dev]', stats.toString({
      colors: true
    }));
    callback();
  });
});

gulp.task('lint', () => {
  return gulp.src([`${FRONTEND_DIR}/**/*.js`])
    // eslint() attaches the lint output to the eslint property
    // of the file object so it can be used by other modules.
    .pipe(eslint())
    // eslint.format() outputs the lint results to the console.
    // Alternatively use eslint.formatEach() (see Docs).
    .pipe(eslint.format())
    // To have the process exit with an error code (1) on
    // lint error, return the stream and pipe to failOnError last.
    .pipe(eslint.failOnError());
});
