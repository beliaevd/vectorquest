var gulp = require('gulp');
const concat = require('gulp-concat');
var merge = require('merge-stream');

var sass = require('gulp-sass')(require('sass'));


//styles all 
gulp.task('sass', function(){
var admin = gulp.src('static/admin/scss/portal.scss')
        .pipe(sass({outputStyle: 'compressed'}).on('error', sass.logError))
        .pipe(gulp.dest('static/admin/css/'));

var user = gulp.src('static/user/scss/app.scss')
    .pipe(sass({outputStyle: 'compressed'}).on('error', sass.logError))
    .pipe(gulp.dest('static/user/css/'));

return merge(admin, user)
});


//js public task
gulp.task('js', function (){
   return gulp.src(['static/user/js/**/*js'])
       .pipe(concat('App.js'))
       .pipe(gulp.dest('static/build/js'));
});