package config

import (
    "crypto/tls"
    "fmt"
    "net"
    "net/url"
    "os"
    "runtime"
    "strings"
    "sync"

    "github.com/pkg/errors"
    "github.com/rs/zerolog/log"
    "github.com/xmlking/configor"

    configPB "github.com/xmlking/micro-starter-kit/shared/proto/config"
    uTLS "github.com/xmlking/micro-starter-kit/shared/util/tls"
)

var (
    Configor   *configor.Configor
    cfg        configPB.Configuration
    configLock = new(sync.RWMutex)

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

    Configor = configor.New(&configor.Config{UsePkger: true, ErrorOnUnmatchedKeys: true})
    log.Info().Msgf("loading configuration from file: %s", configPath)
    if err := Configor.Load(&cfg, configPath); err != nil {
        if strings.Contains(err.Error(), "no such file") {
            log.Panic().Err(err).Msgf("missing config file at %s", configPath)
        } else {
            log.Fatal().Err(err).Send()
        }
    }
}

/**
  Helper Functions
*/
func IsProduction() bool {
    return Configor.GetEnvironment() == "production"
}

func GetBuildInfo() string {
    return fmt.Sprintf(versionMsg, Version, BuildDate, runtime.Version(), runtime.Compiler, runtime.GOOS, runtime.GOARCH,
        GitCommit, GitBranch, GitState, GitSummary)
}

func GetConfig() configPB.Configuration { // FIXME: return a deep copy?
    configLock.RLock()
    defer configLock.RUnlock()
    return cfg
}

func CreateServerCerts() (tlsConfig *tls.Config, err error) {
    configLock.RLock()
    defer configLock.RUnlock()
    tlsConf := cfg.Features.Tls
    return uTLS.GetTLSConfig(tlsConf.CertFile, tlsConf.KeyFile, tlsConf.CaFile, tlsConf.Servername)
}

func GetListener(endpoint string) (lis net.Listener, err error) {
    configLock.RLock()
    defer configLock.RUnlock()

    u, err := url.Parse(endpoint)
    if err != nil {
        return nil, err
    }

    switch u.Scheme {
    case "unix":
        return net.Listen("unix", u.Path)
    case "tcp", "dns", "kubernetes":
        tlsConf := cfg.Features.Tls
        if tlsConf.Enabled {
            if tlsConfig, err := uTLS.GetTLSConfig(tlsConf.CertFile, tlsConf.KeyFile, tlsConf.CaFile, tlsConf.Servername); err != nil {
                return nil, err
            } else {
                return tls.Listen("tcp", fmt.Sprintf("0:%s", u.Port()), tlsConfig)
            }
        } else {
            return net.Listen("tcp", fmt.Sprintf("0:%s", u.Port()))
        }
    default:
        return nil, errors.New(fmt.Sprintf("unknown scheme: %s in endpoint: %s", u.Scheme, endpoint))
    }
}
