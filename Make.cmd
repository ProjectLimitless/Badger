echo off;
echo "Limitless.Badger Windows Compile Script - temporarily changes GOPATH, it is reverted once the script ends."
SET GOPATH=%cd%\vendor
go build -o bin/badger.exe src/main.go
