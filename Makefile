all: test

test: testdata/testdata.zip
	go test -v -race

bench: bin/benchstat
	go test -v -bench=. -cpu=1,4 -count=3 -benchtime=100000x | tee new.txt
	bin/benchstat old.txt new.txt

check: bin/gofumports bin/go-consistent bin/golangci-lint
	bin/gofumports -w -l -local github.com/AlekSi/applehealth .
	bin/go-consistent -exclude=tools -pedantic ./...
	bin/golangci-lint run --config=.golangci-required.yml
	bin/golangci-lint run --new

bin/golangci-lint:
	go build -modfile=tools/go.mod -o bin/golangci-lint github.com/golangci/golangci-lint/cmd/golangci-lint

bin/go-consistent:
	go build -modfile=tools/go.mod -o bin/go-consistent github.com/quasilyte/go-consistent

bin/reviewdog:
	go build -modfile=tools/go.mod -o bin/reviewdog github.com/reviewdog/reviewdog/cmd/reviewdog

bin/benchstat:
	go build -modfile=tools/go.mod -o bin/benchstat golang.org/x/perf/cmd/benchstat

bin/gofumports:
	go build -modfile=tools/go.mod -o bin/gofumports mvdan.cc/gofumpt/gofumports

testdata/testdata.zip: testdata/testdata.xml
	zip testdata/testdata.zip testdata/testdata.xml
