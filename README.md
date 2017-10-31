[![Build Status](https://travis-ci.org/TobiEiss/freeboardBackend.svg?branch=master)](https://travis-ci.org/TobiEiss/freeboardBackend)
[![Coverage Status](https://coveralls.io/repos/github/TobiEiss/freeboardBackend/badge.svg?branch=master)](https://coveralls.io/github/TobiEiss/freeboardBackend?branch=master)

WIP: freeboard backend

## Setup
- create your personal new repository
- add this repository as remote `git remote add freshboard git@github.com:TobiEiss/freshboard.git`
- pull from this remote `git pull freshboard master --allow-unrelated-histories`
- init git submodules (the original freeboard) `git submodule init`
- update git submodules (the original freeboard) `git submodule update`
- (optional) test installation:
    - if you have installed go: `go run server.go`
    - browse: `http://localhost:8080/#source=config.json`

### build Dockerfile
- Raspberrypi (rasbian): `docker build -t freshboard -f Dockerfile-raspberrypi .`

## RaspberryPi-config
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