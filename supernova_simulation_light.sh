#!/bin/bash
while true; do
	/usr/local/bin/supernova -sub-accounts 5 -transactions 2 -url http://gnoland:26657 -mnemonic "source bonus chronic canvas draft south burst lottery vacant surface solve popular case indicate oppose farm nothing bullet exhibit title speed wink action roast" -output REALM_DEPLOYMENT_result.json  -mode REALM_DEPLOYMENT
	sleep 1
	/usr/local/bin/supernova -sub-accounts 4 -transactions 2 -url http://gnoland:26657 -mnemonic "walnut secret absent call depend clown bunker cram drift catch congress afraid enforce awesome talent guitar leaf clump buffalo adult modify shoe chief fork" -output PACKAGE_DEPLOYMENT_result.json  -mode PACKAGE_DEPLOYMENT
	sleep 1
	/usr/local/bin/supernova -sub-accounts 6 -transactions 2 -url http://gnoland:26657 -mnemonic "walnut secret absent call depend clown bunker cram drift catch congress afraid enforce awesome talent guitar leaf clump buffalo adult modify shoe chief fork" -output REALM_CALL_result.json  -mode REALM_CALL
	sleep 1
done
