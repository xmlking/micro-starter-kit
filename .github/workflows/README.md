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
  export GITHUB_TOKEN=16650cad4e8a4332284255f7caf76134eb965b45
  curl POST -H "Authorization: token $GITHUB_TOKEN" \
            -H "Accept: application/vnd.github.ant-man-preview+json"  \
            -H "Content-Type: application/json" \
            https://api.github.com/repos/xmlking/micro-starter-kit/deployments \
            --data '{"ref": "develop", "environment": "e2e", "payload":   "payload": { "what": "deployment for e2e testing"}}'
  ```

## TODO

- [A Github Action to automatically bump and tag master, on merge, with the latest SemVer formatted version.](https://github.com/anothrNick/github-tag-action)

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
- [Deploy to Cloud Run](https://github.com/Preetam/contrast/blob/master/.github/workflows/push.yml)
- GitHub Actions Blogs
  - [radu](https://radu-matei.com/)
