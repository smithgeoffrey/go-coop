# Chicken Coop Automation with Docker and Go

### Overview

I had been playing with go a little, and I knew almost nothing about docker.  I saw opportunity to combine the two as a good fit.  I was looking for a real-world project to help me learn docker and use go.

I recently added a chicken coop at my house. It has a 12-inch door allowing access to an enclosed run during the day, while closing them up in the coop at night. Manually setting the door each morning and night was a chore, so I automated it with hardware. [1]

Avoiding software was nice: no bugs or releases, no patching or upgrades. I hooked a few things together, and the door just does its thing.  But I wanted to remotely verify coop status, particularly in the winter. I had a network camera lying around from a past project. All I needed was to add a raspberry pi and a couple types of sensors.

- install go on the pi and serve a webapp displaying coop sensors and video
- is the door really up or down as expected
- what are the temps outside versus inside the coop
- a live video stream of the run, where the birds spend nearly all of their awake time (when not freeranging in the backyard on weekends)
- build the app into a binary and run it in a container
- run another container that serves a resource consumed by the app, maybe a database
- try to pick up knowledge along the way regarding inter-container networking, shared storage, etc. 

Here's a parts list. [2]

### Books

In order of my getting them:

- `The Go Programming Language` by Donovan and Kernighan
- `Go in Practice` by Butcher & Farina
- `My First Docker Book`

### Design

I loosely followed some tutorials on webapps using go/gin. [3]  I wanted just a few basics:

    GENERALLY
    - keep everything broken out and modular so the structure looks simple and clean even as the app grows
    - use dependency managment [4]
    - use a debugger [5]
    
    TESTING
    - include testing as a top-level package, a first-class citizen
    - play with mocking
    - play with continuously building/testing the app
    
    DATABASE
    - run postgres. [6]
    - use GORM to interact with it (http://jinzhu.me/gorm/)
    
    UI
    - keep UI as a top-level package
    - HTML templates
    - reusable components like header, footer, menu and sidebar
    - auth for home page
    
    API
    - keep API as a top-level package
    - json

### Project Setup

I loosely followed:

- (organization) https://golang.org/doc/code.html#Organization 
- (vendoring) http://lucasfcosta.com/2017/02/07/Understanding-Go-Dependency-Management.html and https://github.com/golang/dep
# TODO: add some docker forums here

# TODO: add docker install here
On the raspberry pi, I install go at /usr/local/go but you could put it anywhere. Just download the `arm` version and unzip it there. That is GOROOT, not to be confused with GOPATH.  GOPATH sets your `workspace` having three subdirs `bin`, `pkg`, `src`, with your code under `src`. You also want to add the GOROOT binary to your PATH so that you can run `go <options>` at the command line.  Here's my bashrc for all of this. [7]

The top-level config/ sets environment variables consumed by a startup script for the service in systemd that I created. [8]

I used an IDE called GoLand. [9] I developed on my laptop and pushed to the pi over many iterations.

### References

[1] It's a 12 volt system on a marine battery. A standard batter maintainer charges the battery that powers a 12-volt relay that powers a linear actuator that moves the door.  A solar sensor acts as input to the relay, and the relay flips the polarity of its output when triggered.  When the sun rises then sets, the door vertically slides open then shut.  Here's a parts list.

Battery: https://shop.hamiltonmarine.com/products/battery-deep-cycle--80-amp-hours-mca-500-35925.html

Charger: https://www.amazon.com/BLACK-DECKER-BM3B-Battery-Maintainer/dp/B0051D3MP6/ref=sr_1_13?ie=UTF8&qid=1504353447&sr=8-13&keywords=battery+charger

Solar sensor: https://www.amazon.com/HIGHROCK-Photocell-Switch-Photoswitch-Sensor/dp/B019BR5Y3U/ref=sr_1_fkmr0_1?ie=UTF8&qid=1511712391&sr=8-1-fkmr0&keywords=HIGHROCK+Ac+Dc+12v+10a+Auto+on+Off+Photocell+Light+Switch+Photoswitch+Light+Sensor+Switch+Roll+over+image+to+zoom+in+HIGHROCK+HIGHROCK+Ac+Dc+12v+10a+Auto

Relay: http://www.modellingelectronics.co.uk/products/reverse-polarity-switch.php (purchased via http://www.ebay.co.uk)

Actuator Arm: https://www.ebay.com/itm/16-inch-Linear-Actuator-Stroke-12-Volt-DC-200-Pound-Max-Lift-12V-Heavy-Duty-/361075336287?hash=item5411c4645f:g:l-4AAOxyg6BR0PJR

[2] IP camera: http://www.vivotek.com/ip8332-c/#views:view=jplist-grid-view

POE injector to power camera: https://www.amazon.com/WT-GPOE-4-48v48w-Gigabit-Passive-Ethernet-Injector/dp/B015S8397E

Door position sensors: https://www.amazon.com/gp/product/B0009SUF08/ref=oh_aui_detailpage_o02_s00?ie=UTF8&psc=1

Temperature sensors: https://www.amazon.com/gp/product/B01IOK40DA/ref=oh_aui_detailpage_o02_s01?ie=UTF8&psc=1

[3] https://github.com/gin-gonic/gin
https://semaphoreci.com/community/tutorials/building-go-web-applications-and-microservices-using-gin
https://semaphoreci.com/community/tutorials/test-driven-development-of-go-web-applications-with-gin
http://cgrant.io/tutorials/go/simple-crud-api-with-go-gin-and-gorm/

[4] I want to try vendoring with `https://github.com/golang/dep`.  I installed it using `brew install dep` (laptop) and `go get -u github.com/golang/dep/...` (raspberry pi). 

[5] See https://lincolnloop.com/blog/debugging-go-code/.  I want to try delv versus godebug at https://github.com/derekparker/delve and https://github.com/mailgun/godebug, respectively, and whatever my IDE has if anything.

[6] Install postgres on the pi:

    apt-get update && apt-get install postgresql-9.4
    echo "host  all  all  172.16.1.0/24 md5" >> /etc/postgresql/9.4/main/pg_hba.conf
    echo "local coop coop md5" >> /etc/postgresql/9.4/main/pg_hba.conf
    <comment out the `local all all peer`> line in pg_hba.conf
    vi /etc/postgresql/9.4/main/postgresql.conf
      <change listen_addresses = 'localhost' to listen_addresses = '*'>
    systemctl restart postgresql
    sudo -u postgres psql
    > create role youradminusernameofchoice with login superuser password 'changeme';
    > create role coop with login password 'changeme';
    > create database coop;
    > grant all privileges on database coop to coop;
    > grant all privileges on database coop to youradminusernameofchoice;
    > \q

[7] Bashrc:

    export GOROOT=/usr/local/go
    export GOPATH=$HOME
    mkdir -p $GOPATH/bin $GOPATH/pkg $GOPATH/src 
    export PATH+=:$GOROOT/bin

[8] It lives at /etc/systemd/system/coop.service as shown.  It lets me do `systemctl start coop`:
    
    [Unit]
    Description=Golang Chicken Coop Web Service
    After=network.target auditd.service
    
    [Service]
    WorkingDirectory=/home/gsmith/src/github.com/smithgeoffrey/go-coop
    EnvironmentFile=/home/gsmith/src/github.com/smithgeoffrey/go-coop/config/environment.vars
    ExecStart=/usr/local/go/bin/go run /home/gsmith/src/github.com/smithgeoffrey/go-coop/main.go
    
    [Install]
    WantedBy=multi-user.target
    Alias=coop.service

[9] https://www.jetbrains.com/go/
