FROM golang:1.16 as base

FROM base as dev

RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

WORKDIR /opt/app/api
# COPY docker/golang/conf.sh /usr/local/bin/conf.sh
# RUN chmod +x /usr/local/bin/conf.sh
# ENTRYPOINT ["conf.sh"]

EXPOSE 8080
CMD ["air"]
