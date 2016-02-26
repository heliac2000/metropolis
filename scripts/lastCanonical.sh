#!/bin/bash

##
## Usage:
##  ./lastCanonical.sh Cout.csv nparallel > initCanon.json
##

for i in $(seq $((${2}-1)) -1 0)
do
  tail $((${i}*-3000-3)) "$1" | head -3
done | {
echo '['
sed -n '
  1~3{s/,/],[/g;s/:/,/g;s/^/{\n"pos": [[/;s/$/]],/;p};
  2~3{s/,/],[/g;s/:/,/g;s/^/"chr": [[/;s/$/]],/;p};
  3~3{s/,/],[/g;s/:/,/g;s/^/"ori": [[/;s/$/]]\n},/;p};
'
} |
{
  head -n -1; echo '}]'
}
