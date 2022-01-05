.PHONY: run
run:
	go run ./cmd/grin-api

.PHONY: cert
cert:
	openssl genrsa -out cert/id_rsa 4096
	openssl rsa -in cert/id_rsa -pubout -out cert/id_rsa.pub

.PHONY: run-auth
run-auth:
	go run ./cmd/grin-auth