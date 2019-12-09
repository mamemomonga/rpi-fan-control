#!/bin/bash
set -eu

GPIO_FAN=27
TEMP_TH="50.0"

fan_on() {
	echo "FAN ON"
	if [ ! -d "/sys/class/gpio/gpio$GPIO_FAN" ]; then
		echo $GPIO_FAN > /sys/class/gpio/export
	fi
	echo out > /sys/class/gpio/gpio$GPIO_FAN/direction
	echo 1 > /sys/class/gpio/gpio$GPIO_FAN/value
	sleep 30
}

fan_off() {
	echo "FAN OFF"
	if [ -d "/sys/class/gpio/gpio$GPIO_FAN" ]; then
		echo $GPIO_FAN > /sys/class/gpio/unexport
	fi
}

PREV_FLAG=1
fan_on
sleep 5

while true; do
	FLAG=$( vcgencmd measure_temp | perl -nE 'if(/temp=([\d\.]+)/) { my $t=$1; say (($t > '$TEMP_TH') ? "1":"0") }' )
	if [ "$FLAG" != "$PREV_FLAG" ]; then
		case "$FLAG" in
			"1" ) fan_on  ;;
			"0" ) fan_off ;;
		esac
		PREV_FLAG=$FLAG
	fi
	sleep 1
done

