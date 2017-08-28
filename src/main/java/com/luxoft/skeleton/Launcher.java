package com.luxoft.skeleton;


import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

/**
 * Entry point of app
 */
public class Launcher {

    public static String CONFIG = "network/fabric-devnet.yaml";

    private static final Logger logger = LoggerFactory.getLogger(Launcher.class);

    public static void main(String[] args) throws Exception {

        logger.info("Application started");

    }
}
