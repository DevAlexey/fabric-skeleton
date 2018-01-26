This PoC built on top of [fabric-ops-blueprint]()

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
1. Create separate folder in your project (lets call it *ops*)
2. Copy files from this repository to *ops*
3. Install requirements and set up environment variables on the host machine of deploy process:

- Install python or virtualenv with python - virtualenv is recommended
- ```pip install -r requirements.txt``` (with activated virtualenv)
- set AWS_ACCESS_KEY_ID, AWS_SECRET_ACCESS_KEY, AWS_REGION env variables
 ([docs from AWS](https://docs.aws.amazon.com/general/latest/gr/managing-aws-access-keys.html)).
- Install cryptogen and configtx from Hyberledger Fabric 

4. Read and follow *aws-project-starter.yaml* to continue customization for your project
5. Use *ops-cli* script to manage deployment (run or use doc below)

This ops stuff is just boilerplate, feel free to edit and customize every single part of it in your project. 

#### Examples

See eVote project (*ops* dir) TODO LINK to see how this boilerplate can be customized for real project

