# Logger

This logger basically configure **zerolog** so that you can log via `github.com/rs/zerolog/log`

## Usage

Import `shared/logger` package. It will be self-initialized. 

```golang
import  "github.com/xmlking/micro-starter-kit/shared/logger"
```

Your can set **Logger** config via Environment Variables: `CONFIGOR_LOG_LEVEL` , `CONFIGOR_LOG_FORMAT`

## Test
```
CONFIGOR_LOG_LEVEL=info CONFIGOR_LOG_FORMAT=json go test github.com/xmlking/micro-starter-kit/shared/logger  -count=1
```
