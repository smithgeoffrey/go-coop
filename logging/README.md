# Logging

### Fluentd to Mongo

It'd be nice to keep life small and simple for running on the pi. This is just a chicken coop after all.  What about this:

- https://www.mongodb.com/post/27619817959/fluentd-mongodb-the-easiest-way-to-log-your-data.

<insert>

### Fluentd to ELK

If you're thinking more like, "what if I had to run this at work or something more enterprisy?"  ELK is pretty common.  Assuming an ELK stack were already up, all we'd need is a small client to run on containers.  Docker has a native logdriver for fluentd.  And I found a blog showing its use with springboot containers pushing to ELK: https://programmaticponderings.com/2017/04/10/streaming-docker-logs-to-the-elastic-stack-using-fluentd/.

For me to test though, I'd need an ELK.  Running it on the pi caused a crash first try. [1]  But as a development exercise, I should be able to run ELK smoothly as a container on my laptop and test containers on the pi pushing logs to it. After that, I could try just pushing to mongo or anything better suited for running on the pi and keeping life simple.

<insert>

### References

[1] I tried a raspberry pi specific ELK at https://github.com/stefanwalther/rpi-docker-elk, using docker-compose:

    cd /opt/docker 
    git clone https://github.com/stefanwalther/rpi-docker-elk.git
    cd rpi-docker-elk/
    docker-compose up -d
    http://<ip>:5601

But it was running separate containers for E, L & K that felt heavy, and the night after starting it, my pi hung trying to SSH in, forcing me to reboot.  