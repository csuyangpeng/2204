#!/bin/bash

start(){
	echo "== Start $1 =="
	cd "$1"
	nohup ./"$1" >> logs/out.log 2>&1 &
	if [ "$1" = "amf" ]; then
	  nohup ./"n2proc" >> logs/out.log 2>&1 &
	fi
}

stop(){
	ps -ef | grep -v grep | grep "/$1" | awk '{print $2}' | while read pid
	do
		kill -9 $pid;
	done

	if [ "$1" = "amf" ]; then
	  ps -ef | grep -v grep | grep "n2proc" | awk '{print $2}' | while read pid
	  do
		  kill -9 $pid;
	  done
	fi

	echo "== Stop $1 =="
}

restart(){
	stop $1
	start $1
}

## main ##
if [ "$1" = "start" ]; then
	start "$2"
elif [ "$1" = "stop" ]; then
	stop "$2"
elif [ "$1" = "restart" ]; then
	restart "$2"
elif [ "$1" = "show" ]; then
	ps -ef | grep -v grep | grep "amf" | awk '{print $2,$7,$8}'
	ps -ef | grep -v grep | grep "n2proc" | awk '{print $2,$7,$8}'
	ps -ef | grep -v grep | grep "udm" | awk '{print $2,$7,$8}'
	ps -ef | grep -v grep | grep "upf" | awk '{print $2,$7,$8}'
else
	echo "Usage: ./core_mgr.sh <start|stop|restart|show> <amf|udm|upf>"
fi
