#!/bin/bash

#
kill -9 $(pgrep EasyWeb)
cd /tmp/EasyWeb
git pull https://github.com/diazraelwang/EasyWeb.git
./deployserver/deployserver &