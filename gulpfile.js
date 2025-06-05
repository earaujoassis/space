var gulp       = require('gulp')
var sass       = require('gulp-sass')(require('sass'))
var rename     = require('gulp-rename')
var uglify     = require('gulp-uglify')
var browserify = require('browserify')
var babelify   = require('babelify')
var source     = require('vinyl-source-stream')
var buffer     = require('vinyl-buffer')

var satellites = ['callisto', 'io']
var environment = process.env.NODE_ENV

gulp.task('styles', function () {
    return gulp
        .src([
            './web/core/styles/core.scss',
            './web/core/styles/errors.scss',
            './web/amalthea/styles/amalthea.scss',
            './web/callisto/styles/callisto.scss',
            './web/ganymede/styles/ganymede.scss',
            './web/io/styles/io.scss',
        ])
        .pipe(
            sass({
                includePaths: [
                    'node_modules/foundation-sites/scss',
                    'node_modules/normalize.scss',
                    'web/core/styles',
                    'web'
                ],
                outputStyle: environment == 'production' ? 'compressed' : 'expanded',
                quietDeps: true,
                silenceDeprecations: ['legacy-js-api', 'import', 'global-builtin', 'mixed-decls']
            }).on('error', sass.logError)
        )
        .pipe(gulp.dest('./web/public/css/'))
})

Array.prototype.forEach.call(satellites, function(satellite) {
    if (environment == 'production') {
        gulp.task(satellite, function() {
            return browserify('./web/' + satellite + '/index.jsx')
                .transform(babelify, {presets: ['@babel/preset-env', '@babel/react']})
                .bundle()
                .pipe(source('index.jsx'))
                .pipe(buffer())
                .pipe(rename(satellite + '.min.js'))
                .pipe(uglify())
                .pipe(gulp.dest('./web/public/js/'))
        })
    } else {
        gulp.task(satellite, function() {
            return browserify('./web/' + satellite + '/index.jsx')
                .transform(babelify, {presets: ['@babel/preset-env', '@babel/react']})
                .bundle()
                .pipe(source('index.jsx'))
                .pipe(buffer())
                .pipe(rename(satellite + '.min.js'))
                .pipe(gulp.dest('./web/public/js/'))
        })
    }
})

gulp.task('scripts', gulp.series(satellites))

gulp.task('default', gulp.series('scripts', 'styles'))
