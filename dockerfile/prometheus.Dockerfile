FROM prom/prometheus:latest

COPY ./prometheus.yml /etc/prometheus/prometheus.yml
COPY ./redisrule.yml /etc/prometheus/redisrule.yml