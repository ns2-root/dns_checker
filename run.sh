#!/bin/bash

cd golang/

yum install tar -y

yum install gunzip -y

sudo tar -zxvf go1.21.5.linux-amd64.tar.gz -C /usr/local/

echo 'export PATH=$PATH:/usr/local/go/bin' >> "$HOME/.profile"

source "$HOME/.profile"

cd ~

rm -rf dns_checker/run.sh

#rm -rf dns_checker/golang/

exec bash