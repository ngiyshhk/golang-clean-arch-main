FROM golang:1.12-alpine as builder
RUN apk add git gcc libc-dev make
WORKDIR "/go/src/github.com/ngiyshhk/golang-clean-arch-main"
COPY . .
RUN make all

FROM alpine
COPY --from=builder /go/src/github.com/ngiyshhk/golang-clean-arch-main/bin/clean_arch_api ./
CMD ["./clean_arch_api"]