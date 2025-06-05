var gulp       = require('gulp')
var sass       = require('gulp-sass')(require('sass'))

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

gulp.task('default', gulp.series('styles'))
