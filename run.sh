#! /bin/bash

ifconfig lo 127.0.0.1

echo "127.0.0.1    kms.ap-southeast-1.amazonaws.com" >> /etc/hosts

/usr/local/bin/socat TCP4-LISTEN:443,bind=127.0.0.1,fork VSOCK-CONNECT:3:8000 &

/usr/local/bin/socat VSOCK-LISTEN:9090,fork,reuseaddr TCP:localhost:9000 & /myapp

