# Monitoring

Prometheus is an open source monitoring system that lets you [1]:

  - scrape metrics, counters, gauges and histograms over HTTP
  - communicate to opsgenie, email or slack (alert_manager)
  - agent for OS-level metrics like cpu/mem/disk (node_exporter)
  - traditional nagios-like checks (blackbox_exporter) 

You can frontend with Grafana which has native suppport for Prometheus as a datasource.  I tried https://hub.docker.com/r/joseba/rpi-grafana/:

    docker volume create grafana_data
    docker run -d -p 3000:3000 -v grafana_data:/data joseba/rpi-grafana
    http://<rpi>:9090

For Prometheus, I found a suggested Dockerfile at https://hub.docker.com/r/prom/prometheus/~/dockerfile/.  I had to use the new flag format in the CMD section per https://prometheus.io/blog/2017/06/21/prometheus-20-alpha3-new-rule-format/.  I also had to add some management of /etc/prometheus.yml dicussed at https://prometheus.io/docs/prometheus/latest/installation/.  ARM versions of the COPY sources I found at https://prometheus.io/download/#prometheus. Here's a summary of the Jenkins job I'm using. 

    SOURCE CODE MANAGEMENT
        
        NONE
                
    BUILD ENVIRONMENT
    
        Delete workspace before build starts
        Add timestamps to console output
            
    BUILD

        EXECUTE SHELL
        # prep a docker buildir        
        mkdir $WORKSPACE/docker && \
        curl -SL https://github.com/prometheus/prometheus/releases/download/v2.0.0/prometheus-2.0.0.linux-armv7.tar.gz > $WORKSPACE/prometheus.tar.gz && \
        tar -xvf $WORKSPACE/prometheus.tar.gz -C $WORKSPACE/docker/ --strip-components=1
        
        EXECUTE SHELL
        # handle config file in Jenkins for now
        cat > $WORKSPACE/docker/prometheus.yml << EOF
        global:
          scrape_interval:     15s 
          evaluation_interval: 15s 

        alerting:
          alertmanagers:
          - static_configs:
            - targets:
              # - alertmanager:9093
        
        rule_files:
          # - "first_rules.yml"
          # - "second_rules.yml"
        
        scrape_configs:
          - job_name: 'prometheus'
            static_configs:
              - targets: ['localhost:9090']
        EOF
        
        EXECUTE SHELL
        # create the dockerfile
        cd $WORKSPACE/docker && \
        cat > Dockerfile << EOF
        FROM        quay.io/prometheus/busybox:latest
        MAINTAINER  The Prometheus Authors <prometheus-developers@googlegroups.com>
        
        COPY prometheus                             /bin/prometheus
        COPY promtool                               /bin/promtool
        COPY prometheus.yml                         /etc/prometheus/
        COPY console_libraries/                     /etc/prometheus/
        COPY consoles/                              /etc/prometheus/
        
        EXPOSE     9090
        VOLUME     [ "/prometheus" ]
        WORKDIR    /prometheus
        ENTRYPOINT [ "/bin/prometheus" ]
        CMD        [ "--config.file=/etc/prometheus/prometheus.yml", \
                     "--web.console.libraries=/etc/prometheus/console_libraries", \
                     "--web.console.templates=/etc/prometheus/consoles" ]
        EOF
        
        EXECUTE DOCKER COMMAND
        Docker command: Create/build image
        Build context folder: $WORKSPACE/docker
        Tag of the resulting docker image: geoff-prometheus
    
    POST-BUILD ACTIONS
        
        SLACK NOTIFICATIONS
        notify failure, success & back to normal

We can run an agent called node_exporter per host, that lets Prometheus poll the agent. I followed https://blog.alexellis.io/prometheus-nodeexporter-rpi/ for installing on the pi itself:

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

Browse to node exporter's port (9090) from your laptop and you should see data available to be scraped by Prometheus server:

    http://<rpi>:9100/metrics

For the containers themselves (not just the rpi), our Dockerfile can do the node_exporter similarly:

    # create the dockerfile
    ...
    ADD    node_exporter /bin/
    
    EXPOSE 8081 9100
    CMD ["/app/gobinary", "nohup /bin/node_exporter &"]
    EOF

Run jenkins again and replace the coop container.  Note the port translation when we instantiate the container, the standard 9100 is the container's and 9101 is the hosts.  This way it can coexist with node_exporter serving 9100 for the pi itself. 

    docker container run -d -p 9101:9100 -p 8081:8081 --name coop coop

<insert discussion about issues with this approach>

### References

[1] See https://blog.alexellis.io/prometheus-monitoring/ and
https://www.digitalocean.com/community/tutorials/how-to-use-prometheus-to-monitor-your-centos-7-server

[2] See, e.g., https://coreos.com/os/docs/latest/getting-started-with-systemd.html#unit-file
