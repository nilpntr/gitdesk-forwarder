.PHONY: build-dev
build-dev:
	mkdir -p bin
	go build -tags dev -ldflags "-s -w -X 'github.com/nilpntr/gitdesk-forwarder/cmd.version=dev'" -o bin/gitdesk-forwarder github.com/nilpntr/gitdesk-forwarder

.PHONY: run
run: build-dev
	@./bin/gitdesk-forwarder

.PHONY: version
version: build-dev
	@./bin/gitdesk-forwarder version