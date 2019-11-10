TEST?=$$(go list ./...)
GOFMT_FILES?=$$(find . -name '*.go' |grep -v vendor)
PKG_NAME=treasuredata
VETARGS?=

default: test vet

clean:
	rm -Rf $(CURDIR)/bin/*

build: clean vet
	go install

test: vet
	GO111MODULE=on TF_ACC= go test $(TEST) $(TESTARGS) -timeout=30s -parallel=4

testacc: vet
	GO111MODULE=on TF_ACC=1 go test $(TEST) -v $(TESTARGS) -timeout 120m

vet: fmt
	@echo "go vet $(VETARGS) ./..."
	@GO111MODULE=on go vet $(VETARGS) ./... ; if [ $$? -eq 1 ]; then \
		echo ""; \
		echo "Vet found suspicious constructs. Please check the reported constructs"; \
		echo "and fix them if necessary before submitting the code for review."; \
		exit 1; \
	fi

fmt:
	@echo "==> Fixing source code with gofmt..."
	gofmt -s -w ./$(PKG_NAME)

.PHONY: default test vet fmt
