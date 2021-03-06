.DEFAULT_GOAL := run

.PHONY: run
run:
	@echo Running f1
	@go run cmd/main.go

.PHONY: plot
plot:
	@echo Test plotting
	@mkdir -p plot
	@python script/main.py test/group.json plot
