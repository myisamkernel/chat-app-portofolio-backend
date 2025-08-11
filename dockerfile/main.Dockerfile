FROM golang:latest AS builder

# Set working directory
WORKDIR /simo-machine-iot

# Copy all files
COPY . .

# Install dependencies and nginx
RUN apt-get update && \
    apt-get upgrade -y && \
    apt-get install -y software-properties-common && \
    apt-get update && \
    apt-get upgrade -y && \
    apt-get clean

RUN go build

FROM nginx:latest AS deployer

COPY --from=builder /simo-machine-iot/seiki-iot /root/seiki-iot
COPY --from=builder /simo-machine-iot/.env /root/.env

WORKDIR /root

RUN apt update -y && \
    apt upgrade -y && \
    apt install -y certbot python3-certbot-nginx

# RUN certbot --nginx --non-interactive --agree-tos --email "mail" -d "domain"

RUN chmod +x /root/seiki-iot

ENTRYPOINT sh -c "nginx && nginx -s reload && ./seiki-iot"

EXPOSE 80 443