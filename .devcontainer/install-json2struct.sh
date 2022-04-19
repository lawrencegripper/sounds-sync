#! /bin/bash

mkdir -p /tmp/json2struct
cd /tmp/json2struct
curl -fsSLO https://github.com/marhaupe/json2struct/releases/download/v1.3.0/json2struct_1.3.0_linux_64-bit.tar.gz
tar xzvf json2struct*.tar.gz -C /usr/local/bin json2struct
rm -rf /tmp/json2struct