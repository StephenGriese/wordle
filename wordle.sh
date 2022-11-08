#!/bin/bash
# sort wordle.txt | grep -v '[solarft]' | grep e | grep i | grep n | grep '.e[^i][^n]t' 

#DEBUG echo "$#" arguments were provided

if [ "$#" -ne 6 ]; then
    echo "You must provide 6 parameters"
    exit
fi 

goodletters=""
grep=""
missedletters=""

missedletters="$1"; shift

for arg; do
    # DEBUG echo "$i" "$arg"
    if [ $arg == "." ]; then
        #DEBUG echo "Letter number $i is unknown"
        grep="$grep."
        continue
    fi
    if [ ${arg:0:1} == "?" ]; then
        #DEBUG echo "Letter number $i is not ${arg:1}"
        goodletters="$goodletters${arg:1}"
        grep="$grep[^${arg:1}]"
        continue
    fi
    #DEBUG echo "Letter number $i is $arg"
    goodletters="${goodletters}${arg}"
    grep="$grep$arg"
done
#DEBUG echo "goodletters=$goodletters"
#DEBUG echo "grep=$grep"

command="sort wordle.txt | grep -v '[$missedletters]'"

if [ ${#goodletters} -gt 0 ]; then
    for (( i=0; i<${#goodletters}; i++ )); do 
        command="$command | grep ${goodletters:$i:1}"
        # echo "$i ${goodletters:$i:1}"
    done
fi

command="$command | grep '$grep' | tr '\n' ' ' | xargs -n15"
#command="sort wordle.txt"

#echo command="$command"

bash -c "$command"
echo
