@echo off
IF NOT EXIST releases (
    mkdir release
)
go build -ldflags "-s -w" main.go
move /Y main.exe release
