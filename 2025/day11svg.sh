#!/bin/bash

{
    echo "digraph reactor {"
    echo "    rankdir=LR;"
    echo "    node [shape=box, style=rounded];"
    echo ""
    echo "    \"you\" [fillcolor=lightblue, style=\"rounded,filled\"];"
    echo "    \"svr\" [fillcolor=lightgreen, style=\"rounded,filled\"];"
    echo "    \"dac\" [fillcolor=orange, style=\"rounded,filled\"];"
    echo "    \"fft\" [fillcolor=yellow, style=\"rounded,filled\"];"
    echo "    \"out\" [fillcolor=red, style=\"rounded,filled\"];"
    echo ""

    awk -F: '{
        device = $1
        outputs = $2
        n = split(outputs, arr, " ")
        for (i = 1; i <= n; i++) {
            if (arr[i] != "") {
                printf "    \"%s\" -> \"%s\";\n", device, arr[i]
            }
        }
    }' "input/11.txt"

    echo "}"
} | dot -Tsvg -o "input/11.svg"
