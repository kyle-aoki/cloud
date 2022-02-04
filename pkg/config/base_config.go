package config

const configFileCreationMessage = `
This is your first time running this program.
Therefore, a config file has been created at ~/.cloud
Please open this file and update it with your security groups names and subnet name tags.

IMPORTANT: Security group names are **not** their name tags but their actual *names*.
IMPORTANT: Subnet names are their *name tags* AND not their IDs.

Use 'cloud config' to print out your config.

`

const baseConfig = `{
  "ShowTerminatedNodes": true,
  "NodeConfigurations": [
    {
      "ConfigurationName": "public-node",
      "SubnetNameTag": "PUBLIC",
      "SecurityGroupNames": ["SSH", "UDP_500", "UDP_4500"],
      "AMI": "ami-091130e4e91d7bb45",
      "KeyPair": "kp1",
      "InstanceType": "t2.nano",
      "StorageSize": "8gb"
    },
    {
      "ConfigurationName": "private-node",
      "SubnetNameTag": "private",
      "SecurityGroupNames": ["ALL_TRAFFIC"],
      "AMI": "ami-091130e4e91d7bb45",
      "KeyPair": "kp1",
      "InstanceType": "t2.nano",
      "StorageSize": "8gb"
    }
  ]
}
`
