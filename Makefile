.PHONY: run
run:
	go run ./cmd/grin-api

.PHONY: run-webapp
run-webapp:
	yarn --cwd app/grin web