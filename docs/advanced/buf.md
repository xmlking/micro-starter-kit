# Buf

[Buf](https://buf.build/) is quality check tool for __Protobuf__ files

### Info

```bash
# To list all files Buf is configured to use:
buf ls-files
# To see your currently configured lint or breaking checkers:
buf check ls-lint-checkers
buf check ls-breaking-checkers
```

### Build

```bash
# check
buf image build -o /dev/null
buf image build -o image.bin
```

### Lint

```bash
buf check lint
# We can also output errors in a format you can then copy into your buf.yaml file
buf check lint --error-format=config-ignore-yaml
# Run breaking change detection
buf check breaking --against-input image.bin
```