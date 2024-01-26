.PHONY: test
test:
	@echo "Running tests..."
	@go test -v ./...


check: test
	@echo "Running nilaway..."
	nilaway ./...
