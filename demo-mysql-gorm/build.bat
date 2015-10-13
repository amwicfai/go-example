set GOOS=linux
set GOARCH=amd64
set CGO_ENABLED=0
go build

set GOOS=windows
set GOARCH=amd64
set CGO_ENABLED=0
go build