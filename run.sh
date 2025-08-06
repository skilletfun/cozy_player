#!/bin/sh
/app/main &
nginx -g 'daemon off;' &
wait -n
exit $?