#!/bin/sh


case $1 in
	"build")
	go build src/*.go
	;;
	"gz")
	tar -czvf gingorm.tar.gz src/ vendor/ template/ go.mod go.sum build.sh
	;;
	*)
	echo "./buiild.sh [build | tar]"
	;;
esac
