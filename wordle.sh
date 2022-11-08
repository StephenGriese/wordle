#!/bin/bash

if [ "$#" -ne 6 ]; then
    echo "You must provide 6 parameters"
    exit
fi 

goodletters=""
grep=""
missedletters=""

missedletters="$1"; shift

for arg; do
    if [ $arg == "." ]; then
        grep="$grep."
        continue
    fi
    if [ ${arg:0:1} == "?" ]; then
        goodletters="$goodletters${arg:1}"
        grep="$grep[^${arg:1}]"
        continue
    fi
    goodletters="${goodletters}${arg}"
    grep="$grep$arg"
done

command="sort dict | grep -v '[$missedletters]'"

if [ ${#goodletters} -gt 0 ]; then
    for (( i=0; i<${#goodletters}; i++ )); do 
        command="$command | grep ${goodletters:$i:1}"
    done
fi

command="$command | grep '$grep' | tr '\n' ' ' | xargs -n15"

bash -c "$command"
echo
