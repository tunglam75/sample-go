ARG IMAGE_NAME

FROM ${IMAGE_NAME}

# COPY conf/error /etc/nginx/error
COPY conf/nginx.conf.template /etc/nginx/

# forward request and error logs to docker log collector
RUN ln -sf /dev/stdout /var/log/nginx/access.log \
    && ln -sf /dev/stderr /var/log/nginx/error.log

EXPOSE 80
STOPSIGNAL SIGTERM

CMD envsubst '$$DOMAIN_UPSTREAM$$NAMESERVERS' < /etc/nginx/nginx.conf.template > /etc/nginx/nginx.conf && \
    exec nginx -g 'daemon off;'
