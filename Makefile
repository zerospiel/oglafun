.PHONY: .test
.test:
	go test ./... -count=1

.PHONY: test
test: .test

.PHONY: .bench
.bench:
	go test ./... -bench=$(run)

.PHONY: bench
bench: .bench