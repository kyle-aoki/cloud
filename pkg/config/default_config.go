package config

const DEFAULT_CONFIG = 
`---
ShowTerminatedNodes: false
Configurations:
  -
    Name: vpn
    VPCNameTag: cloudlab
    SubnetNameTag: public
    AMI: ami-091130e4e91d7bb45
    KeyPair: kp1
    InstanceType: t2.nano
    StorageSize: 8gb
    PrivateIp: 10.0.1.100
    SecurityGroupNames:
      - ssh
      - udp_500
      - udp_4500
  - 
    Name: public
    VPCNameTag: cloudlab
    SubnetNameTag: public
    AMI: ami-091130e4e91d7bb45
    KeyPair: kp1
    InstanceType: t2.nano
    StorageSize: 8gb
    SecurityGroupNames:
      - ssh_only
  - 
    Name: private
    VPCNameTag: cloudlab
    SubnetNameTag: private
    AMI: ami-091130e4e91d7bb45
    KeyPair: kp1
    InstanceType: t2.nano
    StorageSize: 8gb
    SecurityGroupNames:
      - all_traffic
`
