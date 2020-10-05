#!/bin/bash

declare -a args=$@

#if they don't specify build or run, assume run
if [[ $1 != "build" && $1 != "run" ]]; then
    args=( "run" "${args[@]}" )
fi

#go to the current working directory
cd "$PWD"

$(dirname "$0")/oatstart.out ${args[@]}
