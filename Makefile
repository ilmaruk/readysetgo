.PHONY: dependency test

profilePath := ./.build/profile.out
coveragePath := ./.build/coverage.txt

dependency:
	go mod vendor

test:
	@echo "" > $(coveragePath)
	@for d in $(shell go list ./... | grep -v vendor); do \
		go test -mod=mod -race -v -coverprofile=$(profilePath) -covermode=atomic $$d || exit 1; \
		[ -f $(profilePath) ] && cat $(profilePath) >> $(coveragePath) && rm $(profilePath); \
	done
