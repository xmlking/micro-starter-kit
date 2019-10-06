# CI/CD

## Events

- Push
  - push.yaml
- PR
  - pr.yaml
- Release
  - publish.yaml
- Deployment
  - e2e.yaml
  - deploy.yaml

## FAQ

- How do I install private modules?

```yaml
- name: Configure git for private modules
  env:
    TOKEN: ${{ secrets.PERSONAL_ACCESS_TOKEN }}
  run: git config --global url."https://YOUR_GITHUB_USERNAME:${TOKEN}@github.com".insteadOf "https://github.com"
```

- Deployment
  GitHub has first class support for [deployment workflows](https://developer.github.com/v3/repos/deployments/)
  Trigger a deployment via the API (we have internal tooling to ease this)

```bash
curl POST -H "Authorization: token $GITHUB_TOKEN" \
          -H "Accept: application/vnd.github.ant-man-preview+json"  \
          -H "Content-Type: application/json" \
          https://api.github.com/repos/org/repo/deployments \
          --data '{"ref": "master", "environment": "production"}'
```

## Ref

- [GitHub Actions for Go](https://github.com/mvdan/github-actions-golang)
- https://github.com/fsouza/fake-gcs-server/blob/master/.github/workflows/main.yml
- https://github.com/mvdan/sh/blob/master/.github/workflows/test.yml
- https://github.com/kjk/siser/blob/master/.github/workflows/go.yml
- [kind, k8s](https://github.com/olegchorny/bookinfo-productpage/blob/master/.github/workflows/pr.yml)
  - [kind](https://github.com/engineerd/setup-kind)
  - [kind workflow](https://github.com/kubevault/operator/blob/master/.github/workflows/go.yml)
- [fuzz testing](https://fuzzit.dev/2019/10/02/how-to-fuzz-go-code-with-go-fuzz-continuously/)
- [status](https://github.com/deliverybot/status)
  - [deliverybot](https://deliverybot.dev/)
- [kustomize](https://github.com/imranismail/setup-kustomize)
