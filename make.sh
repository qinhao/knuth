#!/bin/sh

interview="knuth-interview-2006"
sections="sections"

if [ ! -d "$interview" ]; then
    git clone https://github.com/kragen/knuth-interview-2006
fi

if [ ! -d "$sections" ]; then
    mkdir "$sections"
fi

go run proccess.go

node generate.js
