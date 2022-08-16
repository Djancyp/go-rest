FROM golang:1.16 as base

FROM base as dev

RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin




WORKDIR /etc/wait
WORKDIR /opt/app/api


COPY .env /opt/app

COPY ./docker/golang/wait-for-it.sh /etc/wait/wait-for-it.sh
RUN chmod +x /etc/wait/wait-for-it.sh

# COPY ./docker/golang/env-clone.sh /etc/wait/env-clone.sh
# RUN chmod +x /etc/wait/env-clone.sh
# RUN ./etc/wait/env-clone.sh

EXPOSE 8080

CMD ["air"]

