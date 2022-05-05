package config

const DEFAULT_CONFIG = 
`---
ShowTerminatedNodes: false
Configurations:
- 
    Name: public
    VPCNameTag: cloudlab
    SubnetNameTag: public
    AMI: ami-091130e4e91d7bb45
    KeyPair: kp1
    InstanceType: t2.nano
    StorageSize: 8gb
    SecurityGroupNames:
      - all_traffic
`
