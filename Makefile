TEST?=$$(go list ./... | grep -v 'vendor')
GOFMT_FILES?=$$(find . -name '*.go' | grep -v vendor)

.PHONY: generate test tools lint depscheck fmt fmtcheck vet testacc

generate:
	go generate ./...

test:
	go test -count=1 ./...

tools:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.62.2

lint:
	@echo "==> Checking source code against linters..."
	golangci-lint run ./...

depscheck:
	@echo "==> Checking source code with 'git diff'..."
	@git diff --check || exit 1
	@echo "==> Checking source code with go mod tidy..."
	@go mod tidy
	@git diff --exit-code -- go.mod go.sum || \
		(echo; echo "Unexpected difference in go.mod/go.sum files. Run 'go mod tidy' command or revert any go.mod/go.sum changes and commit."; exit 1)
	@echo "==> Checking source code with go mod vendor..."
	@go mod vendor
	@git diff --exit-code -- vendor || \
		(echo; echo "Unexpected difference in vendor/ directory. Run 'go mod vendor' command or revert any go.mod/go.sum/vendor changes and commit."; exit 1)

fmt:
	@echo "==> Fixing source code with gofmt..."
	gofmt -s -w $(GOFMT_FILES)

vet:
	@echo "go vet ."
	@go vet $$(go list ./... | grep -v vendor/) ; if [ $$? -eq 1 ]; then \
		echo ""; \
		echo "Vet found suspicious constructs. Please check the reported constructs"; \
		echo "and fix them if necessary before submitting the code for review."; \
		exit 1; \
	fi

testacc: fmtcheck
	go test -count=1 -parallel=1 -v ./...
