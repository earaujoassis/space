#!/usr/bin/env bash

if [ -f configs/jwks/development_key.private.pem ];
then
    echo "> Development keys for JWKS already exists"
else
    openssl genrsa -out configs/jwks/development_key.private.pem 2048
    openssl rsa -in configs/jwks/development_key.private.pem -pubout -out configs/jwks/development_key.public.pem
    echo "> Development keys for JWKS generated"
fi
