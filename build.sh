#!/bin/sh
cd assets
npm run-script build
cd ..
go-bindata -o assets.go assets/dist/...
go run *.go
