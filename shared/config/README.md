# Config


## Usage
customize **Configor** at runtime with Environment Variables 

```bash
export CONFIGOR_FILE_PATH=/config/config.yaml
export CONFIGOR_DEBUG_MODE=true
export CONFIGOR_VERBOSE_MODE=true
export CONFIGOR_USE_PKGER=true
export CONFIGOR_ENV=prod

export CONFIGOR_ENV_PREFIX=APP
export APP_FEATURES_TLS_ENABLED=true

# for example 
MICRO_SERVER_ADDRESS=0.0.0.0:9999 CONFIGOR_ENV_PREFIX=APP APP_FEATURES_TLS_ENABLED=true ./build/greeter-service
```
