ARG GO_VERSION=1.18

FROM golang:${GO_VERSION}-alpine as builder

ARG SERVICE_NAME
RUN echo $SERVICE_NAME

RUN apk add make

WORKDIR ./src/tipee/$SERVICE_NAME
COPY . .

RUN make update && \
    make build

RUN mkdir -p /app
RUN chmod +x /app
RUN cp -p ./$SERVICE_NAME /app

WORKDIR /app
RUN touch entrypoint.sh
RUN echo -e "#!/bin/sh \n ./$SERVICE_NAME" >> ./entrypoint.sh
RUN chmod +x ./entrypoint.sh
RUN chmod +x ./$SERVICE_NAME
CMD ["./entrypoint.sh"]
