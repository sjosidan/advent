#!/bin/bash

if [ -z "$1" ]
    then
        echo "no Args"
    else
       mkdir 2015/$1
        if [ -s 2015/$1/$1.go ]
            then
                echo "file exists"
            else
                cp templates/template.go 2015/$1/$1.go
                cp templates/aochelpers.go 2015/$1/helpers.go
                touch 2015/$1/input.txt
                touch 2015/$1/input2.txt
                touch 2015/$1/input3.txt
        fi
fi