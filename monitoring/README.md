# Monitoring

### Overview

Monitoring seems essential for a good pipeline.  I tried https://hub.docker.com/r/prom/prometheus/ without luck.

    docker pull prom/prometheus
    docker run -p 9090:9090 -it -v /prometheus-data prom/prometheus --config.file=/prometheus-data/prometheus.yml
    
So I tried a blog specific to raspberry pi at https://blog.alexellis.io/prometheus-nodeexporter-rpi/ (rpi setup) and https://blog.alexellis.io/prometheus-monitoring/ (prom generally).

    curl -SL https://github.com/prometheus/node_exporter/releases/download/v0.14.0/node_exporter-0.14.0.linux-armv7.tar.gz > node_exporter.tar.gz && \
    sudo tar -xvf node_exporter.tar.gz -C /usr/local/bin/ --strip-components=1

Now create a systemd startup for it at /etc/systemd/system/prom_node_exporter.service [1]:

    [Unit]
    Description=Prometheus Node Exporter
    After=docker.service
    Requires=docker.service
    
    [Service]
    TimeoutStartSec=0
    ExecStart=/usr/local/bin/node_exporter
    
    [Install]
    WantedBy=multi-user.target
    Alias=prom.service

Check the service logs:

    journalctl -u prom_node_exporter
    
    Jan 05 19:14:44 pi1 systemd[1]: Starting Prometheus Node Exporter...
    Jan 05 19:14:44 pi1 systemd[1]: Started Prometheus Node Exporter.
    Jan 05 19:14:45 pi1 node_exporter[2816]: time="2018-01-05T19:14:45-05:00" level=info msg="Starting node_exporter (version=0.14.0
    ...
    Jan 05 19:14:45 pi1 node_exporter[2816]: time="2018-01-05T19:14:45-05:00" level=info msg="Listening on :9100" source="node_exporter.go:186"

Check that the port is up:

    root@pi1:/etc/systemd/system# netstat -antup | grep exporter    
    tcp6       0      0 :::9100                 :::*                    LISTEN      2816/node_exporter

### References

[1] See, e.g., https://coreos.com/os/docs/latest/getting-started-with-systemd.html#unit-file
