goreleaser --snapshot --skip-publish --rm-dist
git tag v0.1.0
git push origin v0.1.0
goreleaser --rm-dist 