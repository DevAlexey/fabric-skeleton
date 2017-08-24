package com.luxoft.skeleton.fabric;

import com.luxoft.fabric.FabricConfig;
import org.hyperledger.fabric.sdk.User;
import java.io.FileNotFoundException;
import java.io.FileReader;
import java.io.IOException;
import java.io.Reader;


/**
 * TODO 1) cache connectors to further improve performance
 * TODO 2) generify class to don't depend on Paradox
 * TODO 3) add multi-channelism
 * TODO 4) make static/singleton for non-DI environment
 */
public class SkeletonBlockchainConnectorFactory {


    public final static String channelId = "paradox-channel";

    private final FabricConfig fabricConfig;
    private String configPath;

    public SkeletonBlockchainConnectorFactory(String configPath) throws IOException {
        this.configPath = configPath;
        fabricConfig = new FabricConfig(getConfigReader());
    }


    /**
     * Null user stands for Admin which will be taken automatically from FabricConfig
     * @return
     * @throws Exception
     */
    public SkeletonBlockchainConnector getAdminBlockchainConnector() throws Exception {
        return new SkeletonBlockchainConnector(null, channelId, fabricConfig);
    }


    public SkeletonBlockchainConnector getBlockchainConnector(User user) throws Exception {
        if (user == null) throw new NullPointerException("User can't be null");
        return new SkeletonBlockchainConnector(user, channelId, fabricConfig);
    }


    private Reader getConfigReader() {
        try {
            return new FileReader(configPath);
        } catch (FileNotFoundException e) {
            return null;
        }
    }
}
