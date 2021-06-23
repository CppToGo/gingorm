#!/bin/sh

function compile(){
    echo ">>>> building $1 <<<<"
    echo building...
    cd $1/src
    go build -o ../bin/$1
    cd - > /dev/null
    echo ">>>> build $1 done <<<<"
}

function build(){
    if [ $# -ne 0 ]; then
        for d in $@; do
            local d=${d%%/*}
            echo $d
            compile $d
        done
    else
        echo "please input [Service_Name]"
    fi
}

function sigout(){
    echo
    echo ">>>> job has been terminated by `whoami`"
    echo
    exit 0
}

trap sigout SIGINT

case $1 in
	"gz")
	tar -czvf gingorm.tar.gz src/ vendor/ template/ go.mod go.sum build.sh
	;;
	*)
    build $@
	;;
esac
