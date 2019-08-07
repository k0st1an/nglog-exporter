build:
	go install ./

linux:
	GOOS=linux GOARCH=amd64 go build -o nglog-exporter ./

macos:
	GOOS=darwin GOARCH=amd64 go build -o nglog-exporter ./

.PHONY: build linux macos
