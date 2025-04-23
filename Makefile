TEST?=$$(go list . | grep -v 'vendor')

all, help:
	@awk 'BEGIN {FS = ":.*##"; printf "\nMakefile help:\n  make \033[36m<target>\033[0m\n"} /^[0-9a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

default: testacc

testacc: ## Run acceptance tests using the docker-compose file
	docker compose -f ./docker/docker-compose.yaml up -d
	go test $(TEST) -v $(TESTARGS) -timeout 10m
	docker compose -f ./docker/docker-compose.yaml down
	git checkout HEAD -- ./docker/conf/AdGuardHome.yaml
