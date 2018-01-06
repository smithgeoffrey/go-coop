# Monitoring

Prometheus is an open source monitoring system that uses a multi-dimensional data-model and a powerful query language, allowing fine-tuned control and more accurate reporting.  It can be frontended with Grafana if the native PromDash isn't enough, and it has AlertManager for sending email or slack. [1]

Here on the pi, run a Prometheus `server` as a container.  Then run `node_exporter` agents that expose host metrics.  Run the exporter on the pi itself and on containers running on the pi.

### Expose Metrics of the Pi

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

Our Dockerfile can do the node_exporter similarly:

    # create the dockerfile
    ...
    ADD    node_exporter /bin/
    
    EXPOSE 8081 9100
    CMD ["/app/gobinary", "nohup /bin/node_exporter &"]
    EOF

Run jenkins again and replace the coop container.  Note the port translation when we instantiate the container, the standard 9100 is the container's and 9101 is the hosts.  This way it can coexist with node_exporter serving 9100 for the pi itself. 

    docker container run -d -p 9101:9100 -p 8081:8081 --name coop coop

### References

[1] https://www.digitalocean.com/community/tutorials/how-to-use-prometheus-to-monitor-your-centos-7-server

[2] See, e.g., https://coreos.com/os/docs/latest/getting-started-with-systemd.html#unit-file
