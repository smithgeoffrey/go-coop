# Logging

### Overview

Log aggregation seems essential.  I tried a raspberry pi specific ELK at https://github.com/stefanwalther/rpi-docker-elk, using docker-compose:

    cd /opt/docker 
    git clone https://github.com/stefanwalther/rpi-docker-elk.git
    cd rpi-docker-elk/
    docker-compose up -d
    http://<ip>:5601

But it was running separate containers for E, L & K that feels heavy, and the night after starting it, my pi hung trying to SSH in, forcing me to reboot.  I started shopping for something simpler and smaller.  I found https://blog.treasuredata.com/blog/2016/08/03/distributed-logging-architecture-in-the-container-era/ which raised fluentd for me. 

Rather than running ELK on the pi, just use a native logdriver in Docker to push logs to ELK. Docker has a fluentd logdriver, and I found a blog showing it's use with springboot containers pushing logs to ELK: https://programmaticponderings.com/2017/04/10/streaming-docker-logs-to-the-elastic-stack-using-fluentd/.

<dev out fluentd pushing to ELK here>
