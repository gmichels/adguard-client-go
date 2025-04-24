all, help:
	@awk 'BEGIN {FS = ":.*##"; printf "\nMakefile help:\n  make \033[36m<target>\033[0m\n"} /^[0-9a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

default: testacc

testacc: ## Run acceptance tests using the docker-compose file
	@echo "==============Running install tests"
	rm -f ./docker/conf/AdGuardHome.yaml
	docker compose -f ./docker/docker-compose.yaml up -d
	go test -failfast -coverprofile=coverage_install.out -v -run=TestInstall -timeout 10m
	docker compose -f ./docker/docker-compose.yaml down
	git checkout HEAD -- ./docker
	@echo "==============Running functional tests"
	docker compose -f ./docker/docker-compose.yaml up -d
	go test -failfast -coverprofile=coverage_rest.out -v -skip=TestInstall -timeout 10m
	docker compose -f ./docker/docker-compose.yaml down
	git checkout HEAD -- ./docker
	@echo "Generating coverage report"
	grep install.go coverage_install.out > coverage_install.tmp.out
	grep -v install.go coverage_rest.out > coverage.out
	cat coverage_install.tmp.out >> coverage.out
