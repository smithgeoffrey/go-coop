# Logging

### Overview

Log aggregation seems all but essential for, well, most things.  I tried a raspberry pi specific ELK at https://github.com/stefanwalther/rpi-docker-elk, using docker-compose:

    cd /opt/docker 
    git clone https://github.com/stefanwalther/rpi-docker-elk.git
    cd rpi-docker-elk/
    docker-compose up -d

Browsed to Kibana and clicked the `create` button:

    http://<ip>:5601

But it was running separate containers for E, L & K, and the night after starting it, my pi hung trying to SSH in.  I started shopping for something smaller, I could always come back.

I looked at https://blog.treasuredata.com/blog/2016/08/03/distributed-logging-architecture-in-the-container-era/ and wanted to try fluentd/hadoop.

