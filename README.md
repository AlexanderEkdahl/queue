/qr/<random>

Shows the QR code that should be scanned. If the new link is changing this has to adapt.

This website could be hidden behind a random url to prevent users from accessing it remotely.

/new/<rotating-random>

Redirects the user to a new ticket

The new link could be ever changing as to prevent users from accessing it remotely. The last couple of random links should work should a client be slow and not have time to access the link.

/t/<ticket-slug>

Eventually

import "github.com/boombuler/barcode"

the printed ticket should show the slug should a user want to access the ticket remotely

/

Root path could show average queue length

## Stuff
Vibrations
http://www.sitepoint.com/use-html5-vibration-api/

## Estimating remaining time

http://en.wikipedia.org/wiki/Moving_average#Exponential_moving_average

5     0.2
10    0.2
20    0.1

1/5-0 = 0.2

1/10-5 = 0.2

1/20-10 = 0.1

Could use an interval instead between a best case scenario and worst case scenario

## Usage

    go install && queue -host="192.168.1.5:8080"
