# Chicken Coop Monitor

I started a coop project by automating the coop door using hardware [1].

An improvement would be a small web app running on a rasberry pi having a few sensors for coop for monitoring [2].

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


