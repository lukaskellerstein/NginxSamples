FROM nginx
COPY nginx.conf /etc/nginx/nginx.conf
COPY ssl/ /etc/nginx/ssl

RUN pwd
RUN ls -ls

WORKDIR /etc/nginx

RUN pwd
RUN ls -ls

WORKDIR /etc/nginx/ssl

RUN pwd
RUN ls -ls

WORKDIR /

EXPOSE 80