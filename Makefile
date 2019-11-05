TEST?=$$(go list ./...)
VETARGS?=

default: test vet

clean:
	rm -Rf $(CURDIR)/bin/*

build: clean vet
	GO111MODULE=on go build -o $(CURDIR)/bin/terraform-provider-treasuredata $(CURDIR)/builtin/bins/provider-treasuredata/main.go

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
	gofmt -w .


.PHONY: default test vet fmt