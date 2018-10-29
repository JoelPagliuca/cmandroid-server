#!/usr/bin/env bash

if [ "$1" == devices ]; then
	echo "List of devices attached"
	echo -e "F39C7C991AD0\tdevice"
	echo -e "5D05814EF3F1\tdevice"
	echo -e "DF172F869D0C\tdevice"
elif [ "$1" == "-s" ]; then
	CMD=$@
	echo "running 'adb ${CMD[@]:3}' on device '$2'" >> fake-adb.log 
else
	echo "not an expected command"
	exit 1
fi

exit 0