#!/bin/bash

set -e

echo "Generating RSA keys..."

mkdir -p auth-service/keys
mkdir -p protected-service/keys

echo "Generating private key..."

openssl genpkey \
  -algorithm RSA \
  -out auth-service/keys/private.pem \
  -pkeyopt rsa_keygen_bits:2048


echo "Generating public key..."

openssl rsa \
  -pubout \
  -in auth-service/keys/private.pem \
  -out auth-service/keys/public.pem


echo "Copying public key to protected-service..."

cp auth-service/keys/public.pem \
   protected-service/keys/public.pem


echo "Keys generated successfully!"