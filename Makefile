.PHONY : test build clean format docker-build

build:
	go build github.com/telkomdev/go-libreoffice/cmd/golo

docker-build:
	docker build -t golo .

test:
	go test ./...

test-verbose:
	go test -v ./...

format:
	find . -name "*.go" -not -path "./vendor/*" -not -path ".git/*" | xargs gofmt -s -d -w

clean:
	rm golo *.txt
