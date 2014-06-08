/* globals process, __dirname */
var gulp = require('gulp');
var spawn = require('child_process').spawn;
var es = require('event-stream');

var upstream = function(proc) {
  proc.stderr.pipe(process.stderr);
  proc.stdout.pipe(process.stdout);  

  return es.merge(
    proc.stdout,
    proc.stderr
  );
};

gulp.task('exec:gobindata:dev', function(){
  var proc = spawn('go-bindata', ['-o=bin/assets_dev.go', '-pkg=res', '-tags=debug', '-debug=true', '-prefix='+__dirname+'/src', 'src/...']);  
  return upstream(proc);
});

gulp.task('exec:gobindata:prod', function(){
  var proc = spawn('go-bindata', ['-o=bin/assets_prod.go', '-pkg=res', '-tags=release', '-debug=false', '-prefix='+__dirname+'/src', 'src/...']);  
  return upstream(proc);
});

gulp.task('build:prod', ['exec:gobindata:prod']);
gulp.task('build:dev', ['exec:gobindata:dev']);

gulp.task('default', ['build:prod', 'build:dev']);