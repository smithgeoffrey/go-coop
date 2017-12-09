# Chicken Coop Automation Using Golang 

### Overview

I had been looking for a real-world project to help me learn go.  I recently added a chicken coop at my house. It has a 12-inch door allowing access to an enclosed run during the day, while closing them up in the coop at night. Manually setting the door each morning and night was a chore, so I automated it with hardware. [1] 

Avoiding software was nice: no bugs or releases, no security patching or upgrades, less to go wrong. I hooked a few things together, and the door just does its thing.  But I wanted to remotely verify coop status, particularly in the winter:

- is the door really up or down as expected
- what are the temps outside versus inside the coop
- how about a live video stream of the run, where the birds spend nearly all of their awake time

I had a network camera lying around from a past project. All I needed was to add a raspberry pi and a couple types of sensors: 

- install go on the pi and serve a webapp displaying coop sensors and video. Here's a parts list. [2]

### Webapp design

I followed a tutorial on webapps using go/gin. [3]  I wanted some basic features:

    UI
    - basic auth for home page
    - HTML templates
    - reusable components like header, footer, menu and sidebar
    
    API
    - json

### Project Setup

I loosely followed:

- (organization) https://golang.org/doc/code.html#Organization 
- (vendoring) http://lucasfcosta.com/2017/02/07/Understanding-Go-Dependency-Management.html and https://github.com/golang/dep
- (web framework) https://github.com/gin-gonic/gin

On the raspberry pi, I install go at /usr/local/go but you could put it anywhere. Just download the `arm` version and unzip it there. That is GOROOT, not to be confused with GOPATH.  GOPATH sets your `workspace` having three subdirs `bin`, `pkg`, `src`, with your code under `src`. You also want to add the GOROOT binary to your PATH so that you can run `go <options>` at the command line.  Here's my bashrc for all of this. [4]

The app is broken out into packages for api, ui and test.  Vendor is for dependency managment. [5]  Config sets environment variables consumed by a startup script for the service in systemd that I created. [6]

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

[3] https://semaphoreci.com/community/tutorials/building-go-web-applications-and-microservices-using-gin and
https://semaphoreci.com/community/tutorials/test-driven-development-of-go-web-applications-with-gin
http://cgrant.io/tutorials/go/simple-crud-api-with-go-gin-and-gorm/

[4] Bashrc:

    export GOROOT=/usr/local/go
    export GOPATH=$HOME
    mkdir -p $GOPATH/bin $GOPATH/pkg $GOPATH/src 
    export PATH+=:$GOROOT/bin

[5] I wanted to try vendoring with `https://github.com/golang/dep`.  I installed it using `brew install dep` (laptop) and `go get -u github.com/golang/dep/...` (raspberry pi).

[6] It lives at /etc/systemd/system/coop.service as shown.  It lets me do `systemctl start coop`:
    
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
