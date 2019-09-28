#!/bin/bash

## Create TLS keypairs
openssl req -newkey rsa:2048 -nodes -keyout tls.key -x509 -days 365 -out tls.cert
cp tls.key ../simpleServer/
cp tls.cert ../simpleServer/

cp tls.key ../bindShell/
cp tls.cert ../bindShell/

rm tls.key tls.cert