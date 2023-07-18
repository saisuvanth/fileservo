GO_CMD=go
GO_BUILD=$(GO_CMD) build
GO_CLEAN=$(GO_CMD) clean


server-dev:
	export APP_ENV=dev
	$(GO_CMD) run main.go
	
	
