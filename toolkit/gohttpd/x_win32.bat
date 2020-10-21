set GOOS=windows
: set GOARCH=386
: set CGO_ENABLED=1

set TARGET=gohttpd

go fmt %TARGET%

go build -o %TARGET%.exe -ldflags "-H windowsgui -w -s" %TARGET%

@pause