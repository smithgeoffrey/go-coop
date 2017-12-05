# Chicken Coop Automation Using Golang 

### Overview

I have been looking for a real-world project to help me learn go. I recently added a chicken coop at my house. It has a 12-inch door allowing the chickens access to an enclosed run during the day while closing them up in the coop at night. Manually setting the door each morning and night is a chore, so I automated it with hardware. [1] Avoiding software was nice -- no bugs or releases, no security patching or upgrades, less to go wrong. I hooked a few things together, and the door just does its thing.   

But I wanted to remotely verify coop status, particularly in the winter:

- is the door really up or down as expected
- what are the temps outside versus inside the coop
- how about a live video stream of the run, where the birds spend nearly all of their awake time

I already have a network camera from a past project. If we add a raspberry pi and a couple types of sensors, we'd have it: 

- install go on the pi and serve a webapp displaying coop sensors and video. Here's a parts list. [2]

### Project Structure

Read https://golang.org/doc/code.html#Organization. On the raspberry pi, I install go at /usr/local/go but you could put it anywhere. Just download the `arm` version and unzip it there. That is GOROOT, not to be confused with GOPATH.  GOPATH sets your so-called `workspace` where you develop your code, having three subdirs: `bin`, `pkg`, `src`. You also want to add the GOROOT binary to your PATH so that you can run `go <options>` at the command line.  Here's my bashrc for all of this:

    export GOROOT=/usr/local/go
    export GOPATH=$HOME
    mkdir -p $GOPATH/bin $GOPATH/pkg $GOPATH/src 
    export PATH+=:$GOROOT/bin

Most of the time you don't need to worry about the `bin` or `pkg` under GOPATH.  But `src` contains your code as well as external libraries downloaded via `go get`.  Each of your projects goes under GOPATH/src/<project> and has a main.go (having a main() function), that acts as the entry point and also means the interpreter knows it will be compilable.  Within, you can have subfolders of other packages that are more like shared libraries, being non executable b/c they lack any main().

    GOPATH
    - bin/
    - pkg/
    - src/
      - coop/
        - main.go
        - config/
          - environment.vars
        - constants/
          - constants.go
        - static/
          - css
          - html
          - images
          - js
        - templates/
          - index.html
        - vendor/

I keep coop/ as its own repo in version control.  Static and templates are straightforward. Vendor is for dependency managment.  I wanted to use `https://github.com/golang/dep` which worked for local development but it wouldn't run on the rpi I think because it lacks an `arm` version of the installer.  Config sets environment variables consumed by a startup script for the service in systemd that I created at /etc/systemd/system/coop.service, below.  It lets me start the app like `systemctl start coop`:

    [Unit]
    Description=Golang Chicken Coop Web Service
    After=network.target auditd.service
    
    [Service]
    WorkingDirectory=/home/gsmith/src/coop
    EnvironmentFile=/home/gsmith/src/coop/config/environment.vars
    ExecStart=/usr/local/go/bin/go run /home/gsmith/src/coop/main.go
    
    [Install]
    WantedBy=multi-user.target
    Alias=coop.service

### References

[1]
It's a 12 volt system on a marine battery. A standard batter maintainer charges the battery that powers a 12-volt relay that powers a linear actuator.  A solar sensor acts as input to the relay, and the relay flips the polarity of its output when triggered.  When the sun rises then sets, the door vertically slides open then shut.  Here's a parts list.

Battery: https://shop.hamiltonmarine.com/products/battery-deep-cycle--80-amp-hours-mca-500-35925.html
Charger: https://www.amazon.com/BLACK-DECKER-BM3B-Battery-Maintainer/dp/B0051D3MP6/ref=sr_1_13?ie=UTF8&qid=1504353447&sr=8-13&keywords=battery+charger
Solar sensor: https://www.amazon.com/HIGHROCK-Photocell-Switch-Photoswitch-Sensor/dp/B019BR5Y3U/ref=sr_1_fkmr0_1?ie=UTF8&qid=1511712391&sr=8-1-fkmr0&keywords=HIGHROCK+Ac+Dc+12v+10a+Auto+on+Off+Photocell+Light+Switch+Photoswitch+Light+Sensor+Switch+Roll+over+image+to+zoom+in+HIGHROCK+HIGHROCK+Ac+Dc+12v+10a+Auto
Relay: http://www.modellingelectronics.co.uk/products/reverse-polarity-switch.php (purchased via http://www.ebay.co.uk)
Actuator Arm: https://www.ebay.com/itm/16-inch-Linear-Actuator-Stroke-12-Volt-DC-200-Pound-Max-Lift-12V-Heavy-Duty-/361075336287?hash=item5411c4645f:g:l-4AAOxyg6BR0PJR

[2]
IP camera: http://www.vivotek.com/ip8332-c/#views:view=jplist-grid-view
POE injector to power camera: <insert>
Raspberry Pi 3: <insert>
Door position sensors: <insert>
Temperature sensors: <insert>
