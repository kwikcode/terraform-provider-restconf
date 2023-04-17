git tag v0.1.0
git push origin v0.1.0
goreleaser --snapshot --skip-publish --rm-dist
goreleaser --rm-dist 