FROM nginx:stable-alpine
RUN apk add --no-cache \
    gettext \
    openssl

RUN mkdir /etc/my-ssl

RUN rm /etc/nginx/conf.d/*.conf

ENV CERT_CITY ""
ENV CERT_DEPT ""
ENV CERT_DOMAIN ""
ENV CERT_ORG ""
ENV CERT_STATE ""
ENV NGINX_ENV production
ENV PRODUCTION_ALLOWLIST 127.0.0.1
ENV PRODUCTION_DENYLIST all
ENV SSL_CERT_FILE /etc/my-ssl/my.crt
ENV SSL_KEY_FILE /etc/my-ssl/my.key
ENV STAGING_ALLOWLIST 127.0.0.1
ENV STAGING_DENYLIST all

COPY cf-origin.crt /etc/ssl/cf-origin.crt
COPY cf-origin.key /etc/ssl/cf-origin.key

COPY generate-nginx-ssl-cert.sh /usr/local/bin/generate-nginx-ssl-cert
RUN chmod +x /usr/local/bin/generate-nginx-ssl-cert

COPY cf-origin.crt /etc/ssl/cf-origin.crt
COPY cf-origin.key /etc/ssl/cf-origin.key


COPY my-gzip.conf my-proxy.conf my-ssl.conf /etc/nginx/
COPY nginx.conf /etc/nginx/conf.d/my.conf

COPY conf.sh /usr/local/bin/conf.sh
RUN chmod +x /usr/local/bin/conf.sh
ENTRYPOINT ["conf.sh"]
CMD ["nginx", "-g", "'daemon off;'"]
