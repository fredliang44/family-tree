FROM alpine@sha256:5a32c65954238c8f70fb9652ba7af86dbd3a10c053258578efea39897c127bf4
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