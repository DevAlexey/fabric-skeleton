package com.luxoft.skeleton.fabric;


import com.luxoft.fabric.FabricConfig;
import com.luxoft.fabric.FabricConnector;
import com.luxoft.skeleton.fabric.proto.EventFromChaincode;

import static com.luxoft.fabric.utils.Configurator.getConfigReader;

/**
 * Entry point of app
 */
public class Launcher {

    private static String CHANNEL = "test-channel";
    private static String CONFIG = "network/fabric-devnet.yaml";

    public static void main(String[] args) throws Exception {
        EventFromChaincode.PutEntity putEntity = EventFromChaincode.PutEntity.newBuilder()
                .setName("Luxoft")
                .setDescription("Test Sketelon for Fabric")
                .setType(EventFromChaincode.Type.COMPANY)
                .build();

        EventFromChaincode.GetEntity getEntity = EventFromChaincode.GetEntity.newBuilder()
                .setName("Luxoft")
                .build();

        FabricConfig fabricConfig = new FabricConfig(getConfigReader(CONFIG));

        FabricConnector fabric = new FabricConnector(CHANNEL, fabricConfig);

        fabric.invoke("add", CHANNEL, putEntity.toByteArray()).thenRun(()->
            fabric.query("find", CHANNEL, getEntity.toByteArray())).get();
    }
}
