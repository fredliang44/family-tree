FROM alpine@sha256:6e6778d41552b2d73b437e3e07c8e8299bd6903e9560419b1dd19e7a590fd670
LABEL maintainer="fredliang"

RUN apk --no-cache add tzdata  ca-certificates && \
    ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" > /etc/timezone

WORKDIR /app
ENV GIN_MODE=release
ADD config/ /app/config/
ADD main /app/main
ADD docs /app/docs/
RUN chmod +x ./main
CMD ["./main"]