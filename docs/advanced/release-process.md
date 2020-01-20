# Release process

- Draft a release
- Upload release assets
- Publish a release

Flow:

1. Manually tag with `make release/draft VERSION=v0.1.1`
2. `Release Management` Pipeline create draft release, upload assets, build and publish docker images.
3. On release success event, `Deploy` Pipeline get triggered and applications get deployed to `QA` k8s cluster.
4. On deployment success event, `E2E` Pipeline get triggered and E2E tests will be executed

Before beginning, you will need your Github personal access token.

```bash
make release/draft VERSION=v0.1.1
```

At this point, you should inspect the release in the Github web UI. If it looks reasonable, proceed:

```bash
export GITHUB_TOKEN=$YOUR_TOKEN
make release/publish
```
