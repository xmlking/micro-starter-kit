# GitFlow

## Branches

- Production branch: master
- Develop branch: develop
- Feature prefix: feature/
- Release prefix: release/
- Hotfix prefix: hotfix/

## Setup

```bash
# Start using git-flow by initializing it inside an existing git repository
git flow init [-d] # The -d flag will accept all defaults.
# Start a new feature
git flow feature start grpc
# Finish up a feature
git flow feature finish grpc
# Publish a feature
git flow feature publish grpc
# Get a feature published by another user.
git flow feature pull origin grpc
## Make a release
# Start a release
git flow release start '0.1.0'
# It's wise to publish the release branch after creating it to allow release commits by other developers
git flow release publish '0.1.0'
## Finish up a release,
# Merges the release branch back into 'master'
# Tags the release with its name
# Back-merges the release into 'develop'
# Removes the release branch
git flow release finish '0.1.0'
# Don't forget to push your tags with
git push origin --tags
```

## Reference

- <https://nvie.com/posts/a-successful-git-branching-model/>
- <https://www.atlassian.com/git/tutorials/comparing-workflows/gitflow-workflow>
- <https://danielkummer.github.io/git-flow-cheatsheet/>
