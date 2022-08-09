#!/usr/bin/env sh

set -eo pipefail

generate-nginx-ssl-cert
REPLACE='$API_DOMAIN $APP_DOMAIN'
REPLACE="${REPLACE} \$SSL_CERT_FILE \$SSL_KEY_FILE"
REPLACE="${REPLACE} \$PRODUCTION_ALLOWLIST \$PRODUCTION_DENYLIST"
REPLACE="${REPLACE} \$STAGING_ALLOWLIST \$STAGING_DENYLIST"

MYCONF=$(envsubst "${REPLACE}" < /etc/nginx/conf.d/my.conf)
case ${NGINX_ENV} in
    production|staging)
        MYCONF=$(echo "${MYCONF}" | sed "s/#${NGINX_ENV}#//g")
        ;;
esac

echo "${MYCONF}" > /etc/nginx/conf.d/my.conf
echo "127.0.0.1 ${API_DOMAIN}" >> /etc/hosts
# patch DNS internally to route to domains
if [ ! "${API_IS_REMOTE}" = 1 ]
then
    echo "172.18.0.100 ${API_DOMAIN}" >> /etc/hosts
fi

eval "${@}"

