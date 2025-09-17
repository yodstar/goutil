@set GOOS=linux
: @set GOARCH=386
: @set AR=i686-w64-mingw32-ar
: @set CC=i686-w64-mingw32-gcc
: @set CXX=i686-w64-mingw32-g++
: @set GOARCH=amd64
: @set AR=x86_64-w64-mingw32-ar
: @set CC=x86_64-w64-mingw32-gcc
: @set CXX=x86_64-w64-mingw32-g++
: @set CGO_ENABLED=1
@set GOARCH=arm64

@set TARGET=gohttpd

: @del build\%TARGET%
@del %TARGET%

@go fmt %TARGET%

@go build -o %TARGET% -ldflags "-w -s" %TARGET%
