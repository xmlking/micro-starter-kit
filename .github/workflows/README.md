# CI/CD

## FAQ

- How do I install private modules?

```yaml
- name: Configure git for private modules
  env:
    TOKEN: ${{ secrets.PERSONAL_ACCESS_TOKEN }}
  run: git config --global url."https://YOUR_GITHUB_USERNAME:${TOKEN}@github.com".insteadOf "https://github.com"
```

## Ref

- [GitHub Actions for Go](https://github.com/mvdan/github-actions-golang)
- https://github.com/fsouza/fake-gcs-server/blob/master/.github/workflows/main.yml
- https://github.com/mvdan/sh/blob/master/.github/workflows/test.yml
- https://github.com/kjk/siser/blob/master/.github/workflows/go.yml
