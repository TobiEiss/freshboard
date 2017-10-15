[![Build Status](https://travis-ci.org/TobiEiss/freeboardBackend.svg?branch=master)](https://travis-ci.org/TobiEiss/freeboardBackend)
[![Coverage Status](https://coveralls.io/repos/github/TobiEiss/freeboardBackend/badge.svg?branch=master)](https://coveralls.io/github/TobiEiss/freeboardBackend?branch=master)

WIP: freeboard backend

RaspberryPi-config
`/home/pi/.config/lxsession/LXDE-pi/autostart`

```
#@xscreensaver -no-splash
@xset s off
@xset -dpms
@xset s noblank

@sh /home/pi/freshboard/autostart.sh

```

Without mouse  
`sudo apt-get install unclutter`  

```
@unclutter
```