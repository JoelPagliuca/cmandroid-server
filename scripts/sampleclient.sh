#!/usr/bin/env bash

devicelist() {
	curl http://127.0.0.1:8080/devicelist
}

tap() {
	curl -X POST http://127.0.0.1:8080/tap \
		--data '{"deviceId":"device1234", "x":56, "y":56}'
}

startpackage() {
	curl -X POST http://127.0.0.1:8080/startpackage \
		--data '{"deviceId":"device1234, "packageName":"com.google.maps"}'
}

$1