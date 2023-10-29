FROM golang:alpine as build-env

ARG SERVICE_NAME=rod-browser

RUN mkdir /service
ADD . /service/

WORKDIR /service
RUN apk add --no-cache git
RUN go build  -o ${SERVICE_NAME} .


FROM debian:stable-slim
WORKDIR /service
COPY --from=build-env /service/${SERVICE_NAME}    /service/${SERVICE_NAME}

EXPOSE 7779

# CMD ["sh", "-c", "Xvfb :99 -screen 0 1024x768x16 & DISPLAY=:99 /service/traffic-farm-crawler"]

ENTRYPOINT ["/service/rod-browser"]