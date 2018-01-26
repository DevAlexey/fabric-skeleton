### Fabric Ops blueprint

Basic boilerplate of devops for Hyberledger Fabric based projects.

#### Why and What

If you have Hyberledger Fabric based project, question of devops probably will be the one of biggest technical
 challenges - you will need to have easy and stable ways to deploy/restart/kill your network, especially in cloud environment (like AWS). 
 We developed several independent Fabric based projects, this repository is our try to use some of our experience to 
 speed up development process in future projects from DevOps side.
 
This repo contains useful files to handle Fabric network with Ansible. 
It also contains boilerplate for command line interface.  
 
#### How to use it  

To start devops automation of your Hyberledger Fabric project, follow next steps:
1. Install requirements and set up environment variables on the host machine of deploy process:

- Install python or virtualenv with python - virtualenv is recommended
- ```pip install -r requirements.txt``` (with activated virtualenv)
- set AWS_ACCESS_KEY_ID, AWS_SECRET_ACCESS_KEY, AWS_REGION env variables
 ([docs from AWS](https://docs.aws.amazon.com/general/latest/gr/managing-aws-access-keys.html)).
- Install cryptogen and configtx from Hyberledger Fabric 

2. Prepare cluster config - copy cluster_configs/cluster-config.yaml.source to <cluster_id>.yaml and
set numbers as you want
3. Read and follow *aws-project-starter.yaml* to continue customization for your project business layer
4. Use *ops-cli* script to manage deployment (documented below)

This ops stuff is just boilerplate, feel free to edit and customize every single part of it in your project. 

#### ops-cli

ops-cli is simple bash script to run whole ansible 

There are some flags you can set:
- *-c |--cluster_id <cluster_id>* - cluster id, used to differ clusters between each other. If not set - will be asked on start.
- *-k* - (optional) to kill cluster 
- *-r* - (optional) to restart cluster 
- *-i* - path to pem file for vms access 
- *-u* - (optional) remote user name ('ubuntu' by default) 
- *-n |--network_dir <network_dir>* - (optional) to set path to hyberledger fabric artifacts (crypto etc). 
    By default, new artifacts will be generated automatically

Examples:
```
# start cluster with id cluster1
./ops-cli -c cluster1 -i ~/apitester.pem

# kill cluster with id cluster1
./ops-cli -c cluster1 -i ~/apitester.pem -k

# restart cluster with id cluster1
./ops-cli -c cluster1 -i ~/apitester.pem -r
```




