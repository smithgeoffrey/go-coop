# Monitoring

### Raspberry Pi

I tried a blog specific to raspberry pi at https://blog.alexellis.io/prometheus-nodeexporter-rpi/ (rpi setup) and https://blog.alexellis.io/prometheus-monitoring/ (prom generally).

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

### Containers on the Pi

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

[1] See, e.g., https://coreos.com/os/docs/latest/getting-started-with-systemd.html#unit-file
