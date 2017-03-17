#!/bin/bash -el

go build -o ./bin/bb
chmod 744 ./bin/bb
codesign -s $CERT ./bin/bb
