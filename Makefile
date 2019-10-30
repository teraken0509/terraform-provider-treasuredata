TEST?=$$(go list ./... | grep -v '/vendor/')
VETARGS?=
GOFMT_FILES?=$$(find . -name '*.go' | grep -v vendor)

default: test vet

tools:
	go get -u github.com/kardianos/govendor

clean:
	rm -Rf $(CURDIR)/bin/*

build: clean vet
	govendor build -o $(CURDIR)/bin/terraform-provider-treasuredata $(CURDIR)/builtin/bins/provider-treasuredata/main.go

test: vet
	TF_ACC= go test $(TEST) $(TESTARGS) -timeout=30s -parallel=4

testacc: vet
	TF_ACC=1 go test $(TEST) -v $(TESTARGS) -timeout 120m

vet: fmt
	@for target in "." "./builtin/..."; do \
		echo "go vet $(VETARGS) $$target"; \
		go vet $(VETARGS) $$target ; if [ $$? -eq 1 ]; then \
			echo ""; \
			echo "Vet found suspicious constructs. Please check the reported constructs"; \
			echo "and fix them if necessary before submitting the code for review."; \
			exit 1; \
		fi \
	done

fmt:
	gofmt -w $(GOFMT_FILES)


.PHONY: default test vet fmt