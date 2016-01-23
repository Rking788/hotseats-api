#!/bin/bash

curl -X POST --header "Content-Type: application/json"\
 --data-ascii '{"date": "2016-01-15T10:15:00GMT-0700", "type": "Homerun"}'\
 http://localhost:8080/events
