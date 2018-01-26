#!/usr/bin/env bash
#java -cp fabric-insurance-0.1-fat.jar com.luxoft.fabric.utils.Configurator --config org0-fabric-utils-config.yaml

java -cp fabric-insurance-0.1-fat.jar com.luxoft.fabric.utils.Configurator --config org1-fabric-utils-config.yaml

#java -cp fabric-insurance-0.1-fat.jar com.luxoft.uhg.Loader --config org0-fabric-utils-config.yaml --add-member data/data-org0.xlsx --add-plan data/data-org0.xlsx --add-plan-member data/data-org0.xlsx

java -cp fabric-insurance-0.1-fat.jar com.luxoft.uhg.Loader --config org1-fabric-utils-config.yaml --add-member data/data-org1.xlsx --add-plan data/data-org1.xlsx --add-plan-member data/data-org1.xlsx