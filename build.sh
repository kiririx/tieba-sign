#!/bin/sh

# git clone https://github.com/kiririx/tieba-sign.git
git pull
go -d -v ./...
go build
mv tieba-sign ~/runner/tieba-sign
cd ~/runner/tieba-sign
./tieba-sign