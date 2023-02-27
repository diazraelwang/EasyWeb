#!/bin/bash

#
kill -9 $(pgrep webserver)
cd /tmp/EasyWeb
git pull https://github.com/diazraelwang/EasyWeb.git
cd webserver/
chmod 755 。/webserver
./webserver &