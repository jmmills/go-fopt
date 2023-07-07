.PHONY: bin clean test dep coverage

dep:
	go mod tidy

bin: fopt

test: cover.out

coverage: cover.html

clean:
	rm -vf fopt cover.out cover.html

fopt:
	go build 

cover.out:
	go test -coverprofile cover.out ./...

cover.html: cover.out
	go tool cover -html=cover.out -o cover.html

