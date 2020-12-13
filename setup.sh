#!/bin/bash

if [ -z "$1" ]
    then
        echo "no Args"
    else
       mkdir 2020/$1
        if [ -s 2020/$1/$1.go ]
            then
                echo "file exists"
            else
                cp templates/template.go 2021/$1/$1.go
                cp templates/aochelpers.go 2021/$1/helpers.go
                touch 2020/$1/input.txt
                touch 2020/$1/input2.txt
                touch 2020/$1/input3.txt
        fi
fi