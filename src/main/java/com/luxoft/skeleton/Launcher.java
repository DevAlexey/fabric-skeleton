package com.luxoft.skeleton;


import com.luxoft.skeleton.fabric.SkeletonBlockchainConnector;
import com.luxoft.skeleton.fabric.SkeletonBlockchainConnectorFactory;
import com.luxoft.skeleton.fabric.proto.EventFromChaincode;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

/**
 * Entry point of app
 */
public class Launcher {

    private static String CONFIG = "network/fabric-devnet.yaml";

    private static final Logger logger = LoggerFactory.getLogger(Launcher.class);

    public static void main(String[] args) throws Exception {
        EventFromChaincode.Entity entity = EventFromChaincode.Entity.newBuilder()
                .setName("Luxoft")
                .setDescription("Test Sketelon for Fabric")
                .setType(EventFromChaincode.Type.COMPANY)
                .build();

        SkeletonBlockchainConnector blockchain = new SkeletonBlockchainConnectorFactory(CONFIG).getAdminBlockchainConnector();

        EventFromChaincode.GetEntity getEntity = blockchain.putEntity(entity).get();

        EventFromChaincode.Entity receivedEntity = blockchain.getEntity(getEntity).get();

        logger.info("Received entity name: " + receivedEntity.getName());

    }
}
