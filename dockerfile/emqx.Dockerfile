# FROM ubuntu:latest AS sslgenerator

# RUN apt update -y && \
#     apt upgrade -y

# # RUN certbot certonly --standalone --non-interactive --agree-tos --email "mail" -d "domain"

FROM emqx:latest AS deployed

# COPY --from=sslgenerator /etc/letsencrypt /etc/letsencrypt