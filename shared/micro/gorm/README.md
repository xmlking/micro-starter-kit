# gormlog

[GORM](https://github.com/jinzhu/gorm) logger implementation using __go-micro__ [meta logger](https://github.com/micro/go-micro/tree/master/logger).

## Usage

```go
var debug bool // shows if we have debug enabled in our app

db, err := gorm.Open("postgres", dsn)
if err != nil {
    panic(err)
}

if debug {
    // By default, gorm logs only errors. If we set LogMode to true,
    // then all queries will be logged.
    // WARNING: if you explicitly set this to false, then even
    // errors won't be logged.
    db.LogMode(true)
}

log := logger.NewLogger()

db.SetLogger(gormlog.NewGormLogger(log, gormlog.WithLevel(logger.DebugLevel)))
```
