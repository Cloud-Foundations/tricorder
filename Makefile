all: go-test

go-test:
	@find * -name '*_test.go' |\
	sed -e 's@^@github.com/Cloud-Foundations/tricorder/@' -e 's@/[^/]*$$@@' |\
	sort -u | xargs go test
