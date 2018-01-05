# ELK

### Overview

Log aggregation seems essential for a pipeline. I tried https://hub.docker.com/r/sebp/elk/ without luck.

    docker pull sebp/elk
    docker run -d -p 5601:5601 -p 9200:9200 -p 5044:5044 --name elk sebp/elk

So I tried a raspberry pi specific one, at https://github.com/stefanwalther/rpi-docker-elk, using docker-compose:

    cd /opt/docker 
    git clone https://github.com/stefanwalther/rpi-docker-elk.git
    cd rpi-docker-elk/
    docker-compose up -d

Getting some data in there is said to be a good start:

    nc localhost 5000 < /var/log/jenkins/jenkins.log 

Browse to Kibana and click the `create` button, to enable indexing:

    http://<ip>:5601
    http://<ip>:5601/app/sense
