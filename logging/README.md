# Logging

### Overview

Log aggregation seems essential.  I tried a raspberry pi specific ELK at https://github.com/stefanwalther/rpi-docker-elk, using docker-compose:

    cd /opt/docker 
    git clone https://github.com/stefanwalther/rpi-docker-elk.git
    cd rpi-docker-elk/
    docker-compose up -d
    http://<ip>:5601

But it was running separate containers for E, L & K that feels heavy, and the night after starting it, my pi hung trying to SSH in, forcing me to reboot.  I started shopping for something simpler and smaller.  I looked at https://blog.treasuredata.com/blog/2016/08/03/distributed-logging-architecture-in-the-container-era/ which raised fluentd/hadoop.  Hadoop may be overkill here, but it looks like fluentd is a great fit.

    # see https://hub.docker.com/r/fluent/fluentd/
    docker run -d -p 24224:24224 -p 24224:24224/udp -v /data:/fluentd/log fluent/fluentd

I have to develop this more ^^ to get it working.