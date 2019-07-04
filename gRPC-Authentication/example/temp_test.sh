#!/usr/bin/env bash

CUST_NAME="$1"
PSWD="$2"
SERVER_CN="$3"
EXPIRY_DAYS="$4"
COUNTRY="$5"
STATE="$6"
LOCATION="$7"
ORG="$8"
OU="$9"

mkdir "$CUST_NAME"_server

# Step 1: Generate Certificate Authority + Trust Certificate (ca.crt)
openssl genrsa -passout pass:"$PSWD" -des3 -out "$CUST_NAME"_server/"$CUST_NAME"_ca.key 4096
openssl req -passin pass:1111 -new -x509 -days "$EXPIRY_DAYS" -key "$CUST_NAME"_server/"$CUST_NAME"_ca.key -out "$CUST_NAME"_server/"$CUST_NAME"_ca.crt -subj "/CN=${SERVER_CN}/C=${COUNTRY}/ST=${STATE}/L=${LOCATION}/O=${ORG}/OU=${OU}"

# Step 2: Generate the Server Private Key (server.key)
openssl genrsa -passout pass:"$PSWD" -des3 -out "$CUST_NAME"_server/"$CUST_NAME"_server.key 4096

# Step 3: Get a certificate signing request from the CA (server.csr)
openssl req -passin pass:"$PSWD" -new -key "$CUST_NAME"_server/"$CUST_NAME"_server.key -out "$CUST_NAME"_server/"$CUST_NAME"_server.csr -subj "/CN=${SERVER_CN}/C=${COUNTRY}/ST=${STATE}/L=${LOCATION}/O=${ORG}/OU=${OU}"

# Step 4: Sign the certificate with the CA we created (it's called self signing) - server.crt
openssl x509 -req -passin pass:"$PSWD" -days "$EXPIRY_DAYS" -in "$CUST_NAME"_server/"$CUST_NAME"_server.csr -CA "$CUST_NAME"_server/"$CUST_NAME"_ca.crt -CAkey "$CUST_NAME"_server/"$CUST_NAME"_ca.key -set_serial 01 -out "$CUST_NAME"_server/"$CUST_NAME"_server.crt

# Step 5: Convert the server certificate to .pem format (server.pem) - usable by gRPC
openssl pkcs8 -topk8 -nocrypt -passin pass:"$PSWD" -in "$CUST_NAME"_server/"$CUST_NAME"_server.key -out "$CUST_NAME"_server/"$CUST_NAME"_server.pem

