#!/bin/bash
if [ -z "$1" ]
    then
        echo "No args"
    else 
        cd 2019/$1
        go build && cat input$2.txt | ./$1
fi