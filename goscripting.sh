#!/bin/bash 
go get github.com/erning/gorun
mv ~/go/bin/gorun /usr/local/bin/
echo ':golang:E::go::/usr/local/bin/gorun:OC' | tee /proc/sys/fs/binfmt_misc/register