# Containers 

We're using Docker containers to run the App and Monitoring components of this project.  See ~/pipeline and ~/monitoring for more on that.  But here I wanted to gather breadcrumbs and a cheatsheet aspect re Docker.

### Performance 

More sophisticated container monitoring solutions are expected.  But first consider some simpler avenues:

- docker container top
- docker container stats
- docker container inspect

### Interactive

- docker container run -it 
- docker container start -ia 
- docker container exec -it 

### Networking

Generally:

- private virtual network bridge (docker0) per container
- host NIC does outbound NAT per VN
- no -p needed for containers to talk to each other if they're on the same VN
- -p needed if you want inbound requests to NIC forwarding to the VN
- good to create new VN per app/db pairing
- possible to attach multiple VNs to containers OR to bypass VNs and use host NIC directly

Commands:

- docker network create --driver
- docker network ls
- docker container run -d --name foo --network mynet foo
- docker network inspect
- docker network connect <net id> <host id>
- docker network inspect
- docker network disconnect <net id> <host id>
- docker network inspect

### DNS

If you create a non-default network, for free you get: 

    automatic DNS resolution for all containers on the newtork using their container names

E.g., if on a non-default network:

    docker container exec -it <source container name> ping <dest container name>

### Images

- docker image history
- docker image inspect
- docker pull <repo>
- docker pull <repo>:<tag> # tag defaults to latest
- docker image tag --help
- docker image tag smithgeoffrey/go-coop smithgeoffrey/go-coop:testing # same image id
- docker push smithgeoffrey/go-coop

Dockerfile:

Generally put the the things that change the most at the bottom:

- FROM
- ENV
- RUN
- WORKDIR # change dirs
- COPY|ADD
- EXPOSE 80 443
- CMD

Build & Run:

- docker image build -t mytag .
- docker image ls
- docker container run -p 80:80 --rm mytag

### Logging

Don't log to a log file.  Log to stdout:

- RUN ln -sf /dev/stdout /var/log/app/access.log && ln -sf /dev/stderr /var/log/app/error.log

### Data Volumes

Create in Dockerfile:

- VOLUME /var/lib/mydb

Create at Runtime:

- docker image inspect mydb     # shows "Volumes": {"/var/lib/mydb": {}}
- docker container inspect mydb	# dito
- docker container run -d --name mydb -v friendlyname-vol:/var/lib/mydb mydb
- docker volume ls || docker volume inspect

Create Ahead of Time:

- docker volume create --foo

### Bind Mounts

Mapping from host-level file/dir to a container's file/dir.

- only Runtime not Dockerfile
- docker container run -d --name foo -p 443:443 -v /host/path:/container/path 

