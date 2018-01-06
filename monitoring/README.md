# Monitoring

Prometheus is an open source monitoring system [1]:
  - frontend with Grafana if native PromDash isn't enough
  - send email or slack with native AlertManager
  - node_exporter host agent for OS-level metrics like cpu/mem/disk
  - blackbox_exporter for traditional nagios-like checks  

### Expose Host Metrics with Node Exporter

To expose metrics of the pi itself, install node_exporter directly. I followed https://blog.alexellis.io/prometheus-nodeexporter-rpi/ (rpi setup).

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

Browse to it remotely from your laptop:

    http://<rpi>:9100/metrics

### Harvest Host Metrics with Prometheus Server

Run the server as a container.  Dockerhub at https://hub.docker.com/r/prom/prometheus/~/dockerfile/ suggested a dockerfile to use.  

    FROM        quay.io/prometheus/busybox:latest
    MAINTAINER  The Prometheus Authors <prometheus-developers@googlegroups.com>
    
    COPY prometheus                             /bin/prometheus
    COPY promtool                               /bin/promtool
    COPY documentation/examples/prometheus.yml  /etc/prometheus/prometheus.yml
    COPY console_libraries/                     /etc/prometheus/
    COPY consoles/                              /etc/prometheus/
    
    EXPOSE     9090
    VOLUME     [ "/prometheus" ]
    WORKDIR    /prometheus
    ENTRYPOINT [ "/bin/prometheus" ]
    CMD        [ "-config.file=/etc/prometheus/prometheus.yml", \
                 "-storage.local.path=/prometheus", \
                 "-web.console.libraries=/etc/prometheus/console_libraries", \
                 "-web.console.templates=/etc/prometheus/consoles" ]

I created a Jenkins job doing that, first downloading their repo:

    SOURCE CODE MANAGEMENT
        
        REPOSITORIES
        https://github.com/prometheus/prometheus
        */master
                
    BUILD ENVIRONMENT
    
        Delete workspace before build starts
        Add timestamps to console output
            
    BUILD
    
        EXECUTE SHELL
        # create the dockerfile
        cd $WORKSPACE/docker && \
        cat > Dockerfile << EOF
        FROM        quay.io/prometheus/busybox:latest
        MAINTAINER  The Prometheus Authors <prometheus-developers@googlegroups.com>
        
        COPY cmd/prometheus                             /bin/prometheus
        COPY cmd/promtool                               /bin/promtool
        COPY documentation/examples/prometheus.yml  /etc/prometheus/prometheus.yml
        COPY console_libraries/                     /etc/prometheus/
        COPY consoles/                              /etc/prometheus/
        
        EXPOSE     9090
        VOLUME     [ "/prometheus" ]
        WORKDIR    /prometheus
        ENTRYPOINT [ "/bin/prometheus" ]
        CMD        [ "-config.file=/etc/prometheus/prometheus.yml", \
                     "-storage.local.path=/prometheus", \
                     "-web.console.libraries=/etc/prometheus/console_libraries", \
                     "-web.console.templates=/etc/prometheus/consoles" ]
        EOF
    
        EXECUTE DOCKER COMMAND
        Docker command: Create/build image
        Build context folder: $WORKSPACE
        Tag of the resulting docker image: geoff-prometheus
    
    POST-BUILD ACTIONS
        
        SLACK NOTIFICATIONS
        notify failure, success & back to normal

### Expose Container Metrics

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

[1] See https://blog.alexellis.io/prometheus-monitoring/ and
https://www.digitalocean.com/community/tutorials/how-to-use-prometheus-to-monitor-your-centos-7-server

[2] See, e.g., https://coreos.com/os/docs/latest/getting-started-with-systemd.html#unit-file
