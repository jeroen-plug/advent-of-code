#!/bin/bash

awk -F, 'BEGIN {
    print "<svg xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 1000 1000\">"
    printf "<polygon points=\""
}
{
    printf "%.1f,%.1f ", $1/100, $2/100
}
END {
    print "\" fill=\"green\" stroke=\"red\" stroke-width=\"2.0\"/>"
    print "</svg>"
}' input/9.txt > input/9.svg
