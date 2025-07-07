#!/bin/bash
RUN_NAME=hertz_service
mkdir -p output/bin output/conf output/docs
cp script/* output 2>/dev/null
cp conf/* output/conf
cp docs/* output/docs
chmod +x output/bootstrap.sh
go build -o output/bin/${RUN_NAME}