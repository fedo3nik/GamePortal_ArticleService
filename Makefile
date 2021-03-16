go_lint:
	docker run --rm -v ${PWD}:/app -w /app/ golangci/golangci-lint:v1.36-alpine golangci-lint run -v --timeout=5m

postgres_run:
	docker run --name postgres -p 5432:5432 -e POSTGRES_PASSWORD=article -d postgres:9.6.20
