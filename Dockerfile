FROM alpine@sha256:b3dbf31b77fd99d9c08f780ce6f5282aba076d70a513a8be859d8d3a4d0c92b8
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