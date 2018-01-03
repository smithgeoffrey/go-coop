# ELK

### Overview

Log aggregation seems essential for a pipeline. I tried https://hub.docker.com/r/sebp/elk/ without luck.

    docker pull sebp/elk
    docker run -d -p 5601:5601 -p 9200:9200 -p 5044:5044 --name elk sebp/elk

So I tried a raspberry pi specific one, at https://github.com/stefanwalther/rpi-docker-elk:

    