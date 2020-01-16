FROM go:1.13-buster AS build

ENV GO111MODULE=on

COPY . .

RUN ./build.sh

FROM scratch AS release

