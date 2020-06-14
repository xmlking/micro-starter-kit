package config

import (
    "crypto/tls"
    "fmt"
    "net"
    "os"
    "runtime"
    "strings"
    "sync"

    "github.com/pkg/errors"
    "github.com/rs/zerolog/log"
    "github.com/xmlking/configor"
    "google.golang.org/grpc/resolver"

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

func IsProduction() bool {
    return Configor.GetEnvironment() == "production"
}

func IsSecure() bool {
    configLock.RLock()
    defer configLock.RUnlock()
    return cfg.Features.Tls.Enabled
}

func GetListener(endpoint string) (lis net.Listener, err error) {
    configLock.RLock()
    defer configLock.RUnlock()

    target := ParseTarget(endpoint)

    switch target.Scheme {
    case "unix":
        return net.Listen("unix", target.Endpoint)
    case "tcp", "dns", "kubernetes":
        var port string
        if _, port, err = net.SplitHostPort(target.Endpoint); err == nil {
            if port == "" {
                port = "0"
            }
        } else {
            return nil, errors.New(fmt.Sprintf("unable to parse host and port from endpoint: %s", endpoint))
        }

        tlsConf := cfg.Features.Tls
        if tlsConf.Enabled {
            if tlsConfig, err := uTLS.GetTLSConfig(tlsConf.CertFile, tlsConf.KeyFile, tlsConf.CaFile, tlsConf.Servername); err != nil {
                return nil, err
            } else {
                return tls.Listen("tcp", fmt.Sprintf("0:%s", port), tlsConfig)
            }
        } else {
            return net.Listen("tcp", fmt.Sprintf("0:%s", port))
        }
    default:
        return nil, errors.New(fmt.Sprintf("unknown scheme: %s in endpoint: %s", target.Scheme, endpoint))
    }
}

//*** Copied from https://github.com/grpc/grpc-go/blob/master/internal/grpcutil/target.go ***/
// split2 returns the values from strings.SplitN(s, sep, 2).
// If sep is not found, it returns ("", "", false) instead.
func split2(s, sep string) (string, string, bool) {
    spl := strings.SplitN(s, sep, 2)
    if len(spl) < 2 {
        return "", "", false
    }
    return spl[0], spl[1], true
}

// ParseTarget splits target into a resolver.Target struct containing scheme,
// authority and endpoint.
//
// If target is not a valid scheme://authority/endpoint, it returns {Endpoint:
// target}.
func ParseTarget(target string) (ret resolver.Target) {
    var ok bool
    ret.Scheme, ret.Endpoint, ok = split2(target, "://")
    if !ok {
        return resolver.Target{Endpoint: target}
    }
    ret.Authority, ret.Endpoint, ok = split2(ret.Endpoint, "/")
    if !ok {
        return resolver.Target{Endpoint: target}
    }
    return ret
}
