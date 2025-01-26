# build a tiny docker image
FROM alpine:latest

# Установим bash, curl и другие необходимые утилиты
RUN apk add --no-cache bash curl postgresql-client

RUN mkdir /app

COPY authApp /app

CMD ["/app/authApp"]
