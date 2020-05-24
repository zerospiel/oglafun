.PHONY: .test
.test:
	go test ./... -count=1 -run=$(run)

.PHONY: test
test: .test

.PHONY: .bench
.bench:
	go test -benchmem ./... -bench=^\($(run)\)$$

.PHONY: bench
bench: .bench

.PHONY: goimports
goimports:
	$(info #Running goimports...)
	find . -name "*.go" | grep -vE "^(\.\/pkg\/)|vendor|_mock.go" | xargs -n1 goimports -w

.PHONY: .run
.run:
	go run .

.PHONY: run
run: .run