#!/bin/bash

git pull
path=$PWD

cd $path/service
num=$(netstat -nlpt | grep ":::3000" | grep -v "grep" | wc -l)
if [ $num -eq 1 ]; then
    kill -9 $(ps -ef | grep "app" | grep -v "grep" | awk '{print $2}')
    sleep 3
fi
nohup go run app.go &
sleep 3

num=$(netstat -nlpt | grep ":::3000" | grep -v "grep" | wc -l)
if [ $num -eq 1 ]; then
    echo "run app success"
else
    cd $path/app
    tail -n 10 nohup.out
    echo "run app fail"
fi