package config

import (
    "crypto/tls"
    "fmt"
    "os"
    "runtime"
    "strings"
    "sync"

    "github.com/rs/zerolog/log"
    "github.com/xmlking/configor"

    configPB "github.com/xmlking/micro-starter-kit/shared/proto/config"
    uTLS "github.com/xmlking/micro-starter-kit/shared/util/tls"
)

var (
    Configor   *configor.Configor
    cfg        configPB.ServiceConfiguration
    configLock = new(sync.RWMutex)
    _appName   string

    // Version is populated by govvv in compile-time.
    Version = "untouched"
    // BuildDate is populated by govvv.
    BuildDate string
    // GitCommit is populated by govvv.
    GitCommit string
    // GitBranch is populated by govvv.
    GitBranch string
    // GitState is populated by govvv.
    GitState string
    // GitSummary is populated by govvv.
    GitSummary string
)

// VersionMsg is the message that is shown after process started.
const versionMsg = `
version     : %s
build date  : %s
go version  : %s
go compiler : %s
platform    : %s/%s
git commit  : %s
git branch  : %s
git state   : %s
git summary : %s
`

func init() {
    configPath, exists := os.LookupEnv("CONFIGOR_FILE_PATH")
    if !exists {
        configPath = "/config/config.yaml"
    }

    Configor = configor.New(&configor.Config{UsePkger: true})
    log.Info().Msgf("loading configuration from file: %s", configPath)
    if err := Configor.Load(&cfg, configPath); err != nil {
        if strings.Contains(err.Error(), "no such file") {
            log.Panic().Err(err).Msgf("missing config file at %s", configPath)
        } else {
            log.Fatal().Err(err).Msg("")
        }
    }
}

/**
  Helper Functions
*/

func GetAppName() string {
    configLock.RLock()
    defer configLock.RUnlock()
    return _appName
}

func SetAppName(appName string) {
    configLock.Lock()
    defer configLock.Unlock()
    _appName = appName
}

func IsProduction() bool {
    return Configor.GetEnvironment() == "production"
}

func GetBuildInfo() string {
    return fmt.Sprintf(versionMsg, Version, BuildDate, runtime.Version(), runtime.Compiler, runtime.GOOS, runtime.GOARCH,
        GitCommit, GitBranch, GitState, GitSummary)
}

func GetServiceConfig() configPB.ServiceConfiguration {
    configLock.RLock()
    defer configLock.RUnlock()
    return cfg
}

func CreateServerCerts() (tlsConf *tls.Config, err error) {
    configLock.RLock()
    defer configLock.RUnlock()
    // ff := cfg.Features["mtls"]
    // tlsConf, err = util.GetTLSConfig(ff.certfile, ff.keyfile, ff.cafile, ff.servername)
    tlsConf, err = uTLS.GetTLSConfig("", "", "", "")

    if err != nil {
        tlsConf, err = uTLS.GetSelfSignedTLSConfig("*")
    }
    return
}

//func GetFeatureFlags() FeatureFlags {
//	configLock.RLock()
//	defer configLock.RUnlock()
//	return &featureFlags{features: cfg.Features}
//}
