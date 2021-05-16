GO ?= go
TEST_RUN = ""
.PHONY: test

test:
	$(GO) test -v ./... -run $(TEST_RUN)

clean-cache:
	$(GO) clean -testcache