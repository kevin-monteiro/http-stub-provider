#!/bin/bash

if [ "$1" = "" ]; then
	echo "Version is empty"
	exit 0
fi

docker build -t "kaytm/http-stub-provider:$1" .