#!/usr/bin/env sh

if [ ! -f ${SSL_CERT_FILE} ]
then
    openssl req -x509 -nodes -days 365 -newkey rsa:2048 \
        -keyout ${SSL_KEY_FILE} \
        -out ${SSL_CERT_FILE} \
        -subj "/C=ES/ST=${CERT_STATE}/L=${CERT_CITY}/O=${CERT_ORG}/OU=${CERT_DEPT}/CN=${CERT_DOMAIN}"
fi

if [ -f /etc/ssl/cf-origin.crt ]
then
    mv /etc/ssl/cf-origin.crt /etc/my-ssl/cf-origin.crt
fi

if [ -f /etc/ssl/cf-origin.key ]
then
    mv /etc/ssl/cf-origin.key /etc/my-ssl/cf-origin.key
fi