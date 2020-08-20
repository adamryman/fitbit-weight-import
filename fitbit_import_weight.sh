#!/usr/bin/env bash

# Fail script and exit if any command exit's non-zero
set -e
# Add a DEBUG=* to the env and see all commands run. Useful for debugging
# DEBUG ./base
[[ -z $DEBUG ]] || set -x

# .env provides $FITBIT_ACCESS_TOKEN
source .env

curl -i -H "Authorization: Bearer $FITBIT_ACCESS_TOKEN"   https://api.fitbit.com/1/user/-/profile.json
