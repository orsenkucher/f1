.DEFAULT_GOAL := run

.PHONY: run
run:
	@echo Running f1
	@go run cmd/main.go
