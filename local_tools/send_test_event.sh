#!/bin/bash

curl -X POST --header "Content-Type: application/json"\
 --data-ascii '{"date": "2016-01-15T10:15:00Z", "type": "Homerun", "stadium": {"name": "Fenway Park"}}'\
 http://localhost:8888/events
