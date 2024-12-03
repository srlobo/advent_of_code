#!/bin/bash
# vim: ai:ts=8:sw=8:noet
set -EeufCo pipefail
export SHELLOPTS        # propagate set to children by default
IFS=$'\t
'
umask 0077

command -v egrep >/dev/null 2>&1 || { echo 'please install egrep or use image that has it'; exit 1; }
command -v awk >/dev/null 2>&1 || { echo 'please install awk or use image that has it'; exit 1; }

egrep -o 'mul\([0-9]{1,3},[0-9]{1,3}\)' "$1" | awk -F'[\(\),]' '{sum+=($2*$3)}END{print sum}'
