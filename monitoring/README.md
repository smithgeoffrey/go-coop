# Monitoring

### Overview

Monitoring seems essential for a good pipeline.  I tried https://hub.docker.com/r/prom/prometheus/ without luck.

    docker pull prom/prometheus
    docker run -p 9090:9090 -it -v /prometheus-data prom/prometheus --config.file=/prometheus-data/prometheus.yml
    
So I tried a blog specific to raspberry pi at https://github.com/ajeetraina/prometheus-armv7 (rpi setup) and https://blog.alexellis.io/prometheus-monitoring/ (prom generally).

    curl -SL https://github.com/prometheus/node_exporter/releases/download/v0.14.0/node_exporter-0.14.0.linux-armv7.tar.gz > node_exporter.tar.gz && \
    sudo tar -xvf node_exporter.tar.gz -C /usr/local/bin/ --strip-components=1

Now create a systemd startup for it [1]:

    [Unit]
    Description=Prometheus Node Exporter
    After=docker.service
    Requires=docker.service
    
    [Service]
    TimeoutStartSec=0
    ExecStart=<insert>
    
    [Install]
    WantedBy=multi-user.target
    Alias=prom.service

### References

[1] See, e.g., https://coreos.com/os/docs/latest/getting-started-with-systemd.html#unit-file
