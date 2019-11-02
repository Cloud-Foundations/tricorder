all: test

format:
	gofmt -s -w .

format-imports:
	goimports -w .


test:
	@find * -name '*_test.go' |\
	sed -e 's@^@github.com/Cloud-Foundations/tricorder/@' -e 's@/[^/]*$$@@' |\
	sort -u | xargs go test
