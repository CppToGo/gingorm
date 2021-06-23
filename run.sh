#!/bin/sh

if [ $# -eq 0 ]; then
    echo "please input params"
    exit 0
fi

proc=${1%%/*}

echo ">>>> running $proc "
cd $proc/bin

./$proc &

cd - > /dev/null

echo