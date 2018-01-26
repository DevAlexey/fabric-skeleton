#!/usr/bin/env bash

./loadUsersAndPlans.sh
java -cp fabric-insurance-0.1-rest-client.jar com.luxoft.uhg.rest.Main --simulate data/test-batch-org0.yaml