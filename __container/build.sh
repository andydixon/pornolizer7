#!/bin/bash
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
mv Pornolizer7 ~/go/src/
cd ~/go/src/Pornolizer7
go get
go build
cd ..
mv Pornolizer7 $DIR
cd $DIR
cd docker
mv ../Pornolizer7/Pornolizer7 .
docker build -t andydixon:pornolize-beta .
rm -f Pornolizer7
