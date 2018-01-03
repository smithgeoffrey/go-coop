# Monitoring

### Overview

Monitoring seems essential for a good pipeline.  I tried https://hub.docker.com/r/prom/prometheus/ without luck.

    docker pull prom/prometheus
    docker run -p 9090:9090 -it -v /prometheus-data prom/prometheus --config.file=/prometheus-data/prometheus.yml
    
So I tried one specific to raspberry pi at ___.