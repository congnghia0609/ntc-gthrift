#!/bin/bash
# Author:       nghiatc
# Email:        congnghia0609@gmail.com

#source /etc/profile

echo "Install library dependencies..."

go get -u github.com/tools/godep
go get -u github.com/apache/thrift/lib/go/thrift

echo "Install dependencies complete..."
