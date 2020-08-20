#!/usr/bin/env bash

# Fail script and exit if any command exit's non-zero
set -e
# Add a DEBUG=* to the env and see all commands run. Useful for debugging
# DEBUG ./base
[[ -z $DEBUG ]] || set -x

# .env provides $FITBIT_ACCESS_TOKEN
if [[ -z $FITBIT_ACCESS_TOKEN ]]; then
	source .env;
fi

# Note, fitbit has a super tiny API rate limit of 150 per hour, be prepared to run this every hour if you have a lot of data to import
# Note, this will upload in KG unless you use en_US
curl -i \
  --header 'Content-Type: application/json' \
  --header "Authorization: Bearer $FITBIT_ACCESS_TOKEN" \
  --header "Accept-Language: en_US" \
  --header "Accept-Locale: en_US" \
  --request POST \
  --data '' \
  "https://api.fitbit.com/1/user/-/body/log/weight.json?weight=$2&date=$1"

#curl -i \
  #--header 'Content-Type: application/json' \
  #--header "Authorization: Bearer $FITBIT_ACCESS_TOKEN" \
  #--request DELETE \
  #--data '' \
  #"https://api.fitbit.com/1/user/-/body/log/weight/1386892799000.json"

