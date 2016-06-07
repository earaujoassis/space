var gulp       = require('gulp');
var stylus     = require('gulp-stylus');

var browserify = require('browserify');
var babelify   = require('babelify');
var source     = require('vinyl-source-stream');
var buffer     = require('vinyl-buffer');
var rename     = require('gulp-rename');
var uglify     = require('gulp-uglify');

gulp.task('styles', function () {
  gulp.src('./jupiter/assets/styles/base.styl')
    .pipe(stylus())
    .pipe(gulp.dest('./jupiter/public/css/'));
});

gulp.task('ganymede', function() {
    return browserify('ganymede/index.jsx')
        .transform(babelify, {presets: ['es2015', 'react']})
        .bundle()
        .pipe(source('index.jsx'))
        .pipe(buffer())
        .pipe(rename('ganymede.min.js'))
        //.pipe(uglify())
        .pipe(gulp.dest('jupiter/public/js/'));
});

gulp.task('scripts', ['ganymede']);
gulp.task('default', ['scripts', 'styles']);
