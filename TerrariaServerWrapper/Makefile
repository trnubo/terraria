snapshot:
	git tag v0.0.0
	goreleaser --snapshot --clean
	git tag --delete v0.0.0

_ci_build: snapshot
