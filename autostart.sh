#!/bin/sh

# run docker file
docker run -d --publish 8080:8080 --name freshboard freshboard:latest

# run epiphany
chromium-browser --incognito --kiosk http://localhost:8080/\#source\=config.json