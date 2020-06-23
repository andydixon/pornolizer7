#!/bin/bash
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
cd $DIR
cd ..
mkdir -p ~/go/src/Pornolizer7
cp -r * ~/go/src/Pornolizer7
cd ~/go/src/Pornolizer7
go get
go build
cd $DIR
cd docker
mv ~/go/src/Pornolizer7/Pornolizer7 ./
docker build -t andydixon:pornolizer7 . --no-cache
#rm -f Pornolizer7
