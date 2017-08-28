package com.luxoft.skeleton;


import com.luxoft.skeleton.fabric.SkeletonBlockchainConnector;
import com.luxoft.skeleton.fabric.SkeletonBlockchainConnectorFactory;
import com.luxoft.skeleton.fabric.proto.TestChaincode;
import org.junit.Test;
import org.junit.runner.RunWith;
import org.mockito.runners.MockitoJUnitRunner;

import static org.junit.Assert.assertEquals;
import static org.junit.Assert.fail;

@RunWith(MockitoJUnitRunner.class)
public class BlockchainConnectorTest {

    @Test
    public void testGetPutEntity() throws Exception {

        String NAME = "name:" + System.currentTimeMillis();


        TestChaincode.Entity entity = TestChaincode.Entity.newBuilder()
                .setName(NAME)
                .setDescription("description1")
                .setType(TestChaincode.Type.COMPANY)
                .build();

        SkeletonBlockchainConnector blockchain = new SkeletonBlockchainConnectorFactory(Launcher.CONFIG).getAdminBlockchainConnector();

        TestChaincode.GetEntity entityRef = blockchain.putEntity(entity).get();

        if (entityRef == null) {
            fail("Null response received");
        } else {
            assertEquals(NAME, entityRef.getName());
        }

        TestChaincode.Entity receivedEntity = blockchain.getEntity(entityRef).get();

        if (receivedEntity == null) {
            fail("Null response received");
        } else {
            assertEquals(entity.getName(), receivedEntity.getName());
            assertEquals(entity.getDescription(), receivedEntity.getDescription());
            assertEquals(entity.getType(), receivedEntity.getType());
        }

        TestChaincode.History history = blockchain.getBalanceHistory(entityRef).get();

        if (history == null) {
            fail("Null response received");
        } else {
            assertEquals(entityRef.getName(), history.getKey());
            assertEquals(1, history.getHistoryCount());
        }
    }
}
