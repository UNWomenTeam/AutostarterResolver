#!/bin/bash
# chkconfig: 2345 20 80
# description: Description comes here....

# Source function library.
. /etc/init.d/functions

service_name="autostarterResolver"

start() {
 echo "Starting autostarterResolver 9..."
 /bin/su -s /home/xela/Projects/back-encore-1/AutostarterResolver/bin/AutostarterResolver
}
stop() {
 echo "Stopping autostarterResolver 9..."
 /bin/su -s /bin/bash killall -9 AutostarterResolver
}
status() {
 if (( $(ps -ef | grep -v grep | grep $service_name | wc -l) > 0 )); then
     echo "$service_name is running!!!"
 else
     echo "$service_name is down!!!"
 fi
}
case $1 in
  start|stop|status) $1;;
  restart) stop; start;;
  *) echo "Usage : $0 <start|stop|restart>"; exit 1;;
esac

exit 0
