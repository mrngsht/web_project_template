#!/bin/sh

errcheck ./... 
go build -race -o /tmp/back main.go

