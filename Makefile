lint:
	@go vet ./...
	@staticcheck ./...

tag:
	@git tag $(version)
	@git push origin $(version)

install:
	@gh extension install chelnak/gh-environments

upgrade:
	@gh extension upgrade chelnak/gh-environments