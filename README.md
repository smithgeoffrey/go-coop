# Chicken Coop Automation with Jenkins, Docker & Go

### Pipeline

Running on a raspberry pi is a nicety for my use case (and it's fun), but it's not overly important.  The project can run on any platform.

- jenkins running locally on the pi  
  - polls version control of an app for each commit
    - builds app into a binary
    - builds docker image housing the binary
    - deploys container from the image 
    - tests
    - publishes image if tests pass

### Logging & Monitoring

Add a couple more containers:

- see ~/logging
- see ~/monitoring

### App

Python? Java? Super great languages. I'd been wanting something in between.  Go's a small, modern, self-contained ecosystem that compiles into a fast binary including dependencies for ease of deployment.  After I mature with it, I can play with best-in-class concurrency.  It has tendrils in dev and ops, and cloud programming. Here are the go books I've used, in order of my getting them:

- `The Go Programming Language` by Donovan and Kernighan
- `Go in Practice` by Butcher & Farina

I needed a subject for the app.  I recently added a chicken coop at my house. It has a 12-inch door allowing access to an enclosed run during the day. Manually setting the door each morning and night was a chore, so I automated it with hardware. [1]  I hooked a few things together, and the door just does its thing.  But I wanted to remotely verify coop status, particularly in the winter.  Just add a raspberry pi, a couple types of sensors and a network camera: [2]

- is the door really up or down as expected
- what are the temps outside versus inside the coop
- a live video stream of the run, where the birds spend nearly all of their awake time (when not free ranging in the backyard on weekends)

I used an IDE called GoLand. [3] I loosely followed:

- (organization) https://golang.org/doc/code.html#Organization 
- (vendoring) http://lucasfcosta.com/2017/02/07/Understanding-Go-Dependency-Management.html and https://github.com/golang/dep

Even more loosely, I browsed some tutorials on webapps using go/gin. [4]  I wanted just a few basics:

    GENERALLY
    - keep everything broken out and modular so the structure looks simple and clean even as the app grows
    - use dependency managment [5]
    - use a debugger [6]
    
    TESTING
    - include testing as a top-level package, a first-class citizen
    - play with mocking
    - play with continuously building/testing the app
    
    DATABASE
    - run postgres if any database is needed [7]
    - use GORM to interact with it (http://jinzhu.me/gorm)
    
    UI
    - keep UI as a top-level package
    - HTML templates
    - reusable components like header, footer, menu and sidebar
    - auth for home page
    
    API
    - keep API as a top-level package
    - json

### Raspberry Pi

I installed jenkins as follows.  Make sure you have java8 installed first.  My pi already had 7 and 8 arm versions of jre available, e.g., `ln -s /usr/lib/jvm/jdk-8-oracle-arm32-vfp-hflt/bin/java /etc/alternatives/java`.  Jenkins couldn't fetch plugins throwing a java trace relating to an SSL error, until I changed the update URL from https to http at `Manage Plugins > Advanced tab > Update Site URL`.  I added a few plugins. [8]  

    sudo wget -q -O - https://jenkins-ci.org/debian/jenkins-ci.org.key | sudo apt-key add -
    sudo deb http://pkg.jenkins-ci.org/debian binary/ > /etc/apt/sources.list.d/jenkins.list'
    sudo apt-get update
    sudo apt-get install jenkins
    vi /etc/default/jenkins # change the listening port
      HTTP_HOST=0.0.0.0
      AJP_HOST=0.0.0.0
    systemctl restart jenkins
    <do setup at http://ip:8080>

I installed docker via `https://store.docker.com/editions/community/docker-ce-desktop-mac` (laptop) and `curl -sSL https://get.docker.com | sh` (raspberry pi).  In Docker, I had to edit systemd for the docker service as shown, then in Jenkins, I set `Docker Builder > Docker URL` to `tcp://localhost:2375` instead of using `http://`.

    #ExecStart=/usr/bin/dockerd -H fd://
    ExecStart=/usr/bin/dockerd -H tcp://0.0.0.0:2375

I installed docker-compose using their `Alternative Install Options` method, via pip.  There was some complexity that I stumbled through, but this got me up and running:

    apt-get remove python-pip
    easy_install pip
    ln -s /usr/local/bin/pip /usr/bin/pip
    pip install docker-compose
    docker-compose --version

I installed go at /usr/local/go but you could put it anywhere. Just download the `arm` version and unzip it there. That is GOROOT, not to be confused with GOPATH.  GOPATH sets your `workspace` having three subdirs `bin`, `pkg`, `src`, with your code under `src`. You also want to add the GOROOT binary to your PATH so that you can run `go <options>` at the command line.  Here's my bashrc for all of this. [9] The top-level config/ sets environment variables consumed by a startup script for the service in systemd that I created. [10]

### References

[1] It's a 12 volt system on a marine battery. A standard battery maintainer charges the battery that powers a 12-volt relay that powers a linear actuator that moves the door.  A solar sensor acts as input to the relay, and the relay flips the polarity of its output when triggered.  When the sun rises then sets, the door vertically slides open then shut.  Here's a parts list.

Battery: https://shop.hamiltonmarine.com/products/battery-deep-cycle--80-amp-hours-mca-500-35925.html

Charger: https://www.amazon.com/BLACK-DECKER-BM3B-Battery-Maintainer/dp/B0051D3MP6/ref=sr_1_13?ie=UTF8&qid=1504353447&sr=8-13&keywords=battery+charger

Solar sensor: https://www.amazon.com/HIGHROCK-Photocell-Switch-Photoswitch-Sensor/dp/B019BR5Y3U/ref=sr_1_fkmr0_1?ie=UTF8&qid=1511712391&sr=8-1-fkmr0&keywords=HIGHROCK+Ac+Dc+12v+10a+Auto+on+Off+Photocell+Light+Switch+Photoswitch+Light+Sensor+Switch+Roll+over+image+to+zoom+in+HIGHROCK+HIGHROCK+Ac+Dc+12v+10a+Auto

Relay: http://www.modellingelectronics.co.uk/products/reverse-polarity-switch.php (purchased via http://www.ebay.co.uk)

Actuator Arm: https://www.ebay.com/itm/16-inch-Linear-Actuator-Stroke-12-Volt-DC-200-Pound-Max-Lift-12V-Heavy-Duty-/361075336287?hash=item5411c4645f:g:l-4AAOxyg6BR0PJR

[2] IP camera: http://www.vivotek.com/ip8332-c/#views:view=jplist-grid-view

POE injector to power camera: https://www.amazon.com/WT-GPOE-4-48v48w-Gigabit-Passive-Ethernet-Injector/dp/B015S8397E

Door position sensors: https://www.amazon.com/gp/product/B0009SUF08/ref=oh_aui_detailpage_o02_s00?ie=UTF8&psc=1

Temperature sensors: https://www.amazon.com/gp/product/B01IOK40DA/ref=oh_aui_detailpage_o02_s01?ie=UTF8&psc=1

[3] https://www.jetbrains.com/go/

[4] I started with https://github.com/gin-gonic/gin.  Then I found three articles, here.  I barely finished browsing the first before I just started tinkering with my setup.  I do hope to return to these for things like examples of auth, DB conns and testing.  
https://semaphoreci.com/community/tutorials/building-go-web-applications-and-microservices-using-gin
https://semaphoreci.com/community/tutorials/test-driven-development-of-go-web-applications-with-gin
http://cgrant.io/tutorials/go/simple-crud-api-with-go-gin-and-gorm/

[5] I want to try vendoring with `https://github.com/golang/dep`.  I installed it using `brew install dep` (laptop) and `go get -u github.com/golang/dep/...` (raspberry pi). 

[6] See https://lincolnloop.com/blog/debugging-go-code/.  I want to try delv versus godebug at https://github.com/derekparker/delve and https://github.com/mailgun/godebug, respectively, and whatever my IDE has if anything.

[7] I installed postgres on the pi via:

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

But here I want to try running it in a container.  I did that via:

    insert

[8] Jenkins plugins I installed beyond the default suite:

    Git Plugin
    Go Plugin
    Packer
    Terraform Plugin
    Pipeline
    Slack Notification Plugin
    Hudson Post build task
    Show Build Parameters plugin
    Timestamper
    Workspace Cleanup Plugin

[9] Bashrc:

    export GOROOT=/usr/local/go
    export GOPATH=$HOME
    mkdir -p $GOPATH/bin $GOPATH/pkg $GOPATH/src 
    export PATH+=:$GOROOT/bin

[10] It lives at /etc/systemd/system/coop.service as shown.  It lets me do `systemctl start coop`:
    
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