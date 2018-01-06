# Monitoring

Prometheus is an open source monitoring system that uses a multi-dimensional data-model and a powerful query language, allowing fine-tuned control and more accurate reporting.  It can be frontended with Grafana if the native PromDash isn't enough, and it has AlertManager for sending email or slack. [1]

Here on the pi, run a Prometheus `server` as a container.  Then run `node_exporter` agents that expose host metrics.  Run the exporter on the pi itself and on containers running on the pi.

### Expose Metrics on the Pi

I tried a blog specific to raspberry pi at https://blog.alexellis.io/prometheus-nodeexporter-rpi/ (rpi setup) and https://blog.alexellis.io/prometheus-monitoring/ (prom generally).

    curl -SL https://github.com/prometheus/node_exporter/releases/download/v0.14.0/node_exporter-0.14.0.linux-armv7.tar.gz > node_exporter.tar.gz && \
    sudo tar -xvf node_exporter.tar.gz -C /usr/local/bin/ --strip-components=1

Now create a systemd startup for it at /etc/systemd/system/prom_node_exporter.service [2]:

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

Verify locally:

    # systemd log
    journalctl -u prom_node_exporter
    
    # netstat
    netstat -antup | grep exporter    
    tcp6       0      0 :::9100                 :::*                    LISTEN      2816/node_exporter

Browse to it remotely from your laptop:

    http://<rpi>:9100/metrics

### Expose Metrics of Containers Running on The Pi

The above applies node_exporter to the pi itself and listens on the standard port 9100.  It would be nice to apply it to containers too, although they'd have to use a different, e.g., 9101, 9102, etc.  For example, the go app's docker file could extend for this:

    # create the dockerfile
    ...
    COPY    node_exporter /bin/
    
    EXPOSE 8081 9100
    CMD ["/app/gobinary", "/bin/node_exporter"]
    EOF

Note that the port exposed is the standard 9100, but then we increment it when instantiating the container from the image:

    docker container run -d -p 8081:8081 -p 9101:9100 --name coop coop

### References

[1] https://www.digitalocean.com/community/tutorials/how-to-use-prometheus-to-monitor-your-centos-7-server

[2] See, e.g., https://coreos.com/os/docs/latest/getting-started-with-systemd.html#unit-file
