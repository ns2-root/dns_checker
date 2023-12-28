#!/bin/bash

yum install tar -y

sudo yum install golang -y

yum install gunzip -y

cd golang/

sudo tar -zxvf go1.21.5.linux-amd64.tar.gz -C /usr/local/

echo 'export PATH=$PATH:/usr/local/go/bin' >> "$HOME/.profile"

echo 'export PATH=$PATH:/usr/local/go/bin' >> "$HOME/.bashrc"

cd 

rm -rf dns_checker/run.sh

rm -rf dns_checker/golang

cd

exec bash