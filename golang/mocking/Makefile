mockery-mock-generate: ## Generate mocks
	docker run --rm -v "$(PWD):/app" -w /app -t vektra/mockery --all --dir /app --case underscore --output ./mockery_mocks

gomock-mock-generate:
	docker run -v $(PWD):/app -w /app ekofr/gomock:go-1.17 mockgen -source=interface.go -destination=gomock_mocks/store.go -package=gomock
