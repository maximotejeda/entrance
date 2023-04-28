FROM golang:1.20 as build

WORKDIR /go/src/app

COPY . .

RUN go mod download && go mod tidy
ARG OS $OS
ARG ARCH $ARCH
ENV PLACE="entrance-$(OS)-(ARCH)"
RUN go build -o /go/bin/entrance ./main/

FROM gcr.io/distroless/cc-debian11
COPY --from=build /go/bin/entrance /

EXPOSE 8083

ENTRYPOINT ["/entrance"]