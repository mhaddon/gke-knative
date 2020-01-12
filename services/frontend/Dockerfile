FROM nginx:alpine

RUN apk add --no-cache ca-certificates git

RUN mkdir -p /app/html

COPY server.conf /etc/nginx/conf.d/default.conf

COPY resources /app/html

# https://github.com/knative/serving/issues/3809
CMD mkdir -p /var/log/nginx && nginx -g 'daemon off;'
