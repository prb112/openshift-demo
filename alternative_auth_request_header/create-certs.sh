#!/usr/bin/env bash

# ----------------------------------------------------------------------------
# (C) Copyright IBM Corp. 2022
#
# SPDX-License-Identifier: Apache-2.0
# ----------------------------------------------------------------------------

# HN is the hostname
HN='*.'"${1}"

mkdir -p tmp && cd tmp

# The Root CA
# Instead of using -passin pass:change-password -passout pass:change-password
# opted for no password with nodes
DAYS="10960"
BITS="4096"
openssl req -new -newkey rsa:${BITS} -x509 -nodes -sha512 -keyout ca.key -out ca.crt -days ${DAYS} -subj '/C=US/O=IBM Corporation/OU=Power Systems/OU=OCP'

# Create the mTLS certificate
echo "mtls - certs are being generated and signed"
openssl req -newkey rsa:${BITS} -days ${DAYS} -nodes -sha512 -keyout client.key -out client-req.pem -subj '/C=US/O=IBM Corporation/OU=Power Systems/OU=OCP/CN='$HN -config ../server.conf -reqexts SAN
openssl x509 -req -in client-req.pem -days ${DAYS} -sha512 -CA ca.crt -CAkey ca.key -days 10960 -CAcreateserial -out client.crt -extfile ../server.conf -extensions SAN

# Verify the subject
openssl x509 -in client.crt -text | grep 'Subject:'

# Create the chained cert
cat client.crt > client-chained.crt
cat ca.crt >> client-chained.crt
# EOF