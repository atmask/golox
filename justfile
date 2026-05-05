build:
    go build -o golox ./cmd/golox/

run *ARGS:
    go run ./cmd/golox/ {{ARGS}}

test:
    go test ./...

clean:
    rm -f golox
