# Containers 

Docker we're using in the pipeline component for the App and for the monitoring.  You can see the readme's for ~/pipeline and ~/monitoring.  

But here I wanted to have a place to gather breadcrumbs and cheatsheet type stuff re Docker.

It will start really really simple and likely progress toward more complex concepts.

### Performance 

More sophisticated container monitoring solutions are expected.  But first consider some simpler avenues:

- docker container top
- docker container stats
- docker container inspect

### Interactive

- docker container run -it 
- docker container start -ai 
- docker container exec -it 

### Networking

Overview:

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

If you create a non-default network, for free you get 

    automatic DNS resolution for all containers on the newtork using their container names

E.g., if on a non-default network:

    docker container exec -it <source container name> ping <dest container name>

