#!/bin/bash

export GOROOT="/usr/local/go"
export PATH=$GOROOT/bin:$PATH

pushd /tmp/
yum -y install tar gzip
curl https://dl.google.com/go/go1.17.7.linux-amd64.tar.gz -o go1.17.7.linux-amd64.tar.gz
rm -rf $GOROOT && tar -C /usr/local -xzf ./go1.17.7.linux-amd64.tar.gz
rm -rf go1.17.7.linux-amd64.tar.gz
popd
