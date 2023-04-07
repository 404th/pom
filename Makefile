bin:
	go-bindata -pkg biny -o biny/biny.go music/

build:
	go build -o pom github.com/404th/helloworld

PHONY: genbin