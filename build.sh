#!/bin/sh

# git clone https://github.com/kiririx/tieba-sign.git
git pull
go get -d -v ./...
go build
mv tieba-sign ~/runner/tieba-sign
cd ~/runner/tieba-sign
./tieba-sign