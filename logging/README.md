# Logging

### Overview

Log aggregation for the pipeline feels like a good idea.  Assuming an ELK stack were already up, all we'd need is a way to push to it with a small client.  Docker has a native logdriver for fluentd.  I found a blog showing it's use with springboot containers pushing logs to ELK: https://programmaticponderings.com/2017/04/10/streaming-docker-logs-to-the-elastic-stack-using-fluentd/.

Running ELK on the pi caused a crash on first try. [1] As a development exercise, I should be able to run ELK more smoothly as a container on my mac, and test containers on the pi push their logs to it.

<dev out fluentd pushing to ELK here>

### References

[1] I tried a raspberry pi specific ELK at https://github.com/stefanwalther/rpi-docker-elk, using docker-compose:

    cd /opt/docker 
    git clone https://github.com/stefanwalther/rpi-docker-elk.git
    cd rpi-docker-elk/
    docker-compose up -d
    http://<ip>:5601

But it was running separate containers for E, L & K that felt heavy, and the night after starting it, my pi hung trying to SSH in, forcing me to reboot.  