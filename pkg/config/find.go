package config

func Find(configurationName string) NodeConfiguration {
	for i := range Vars.NodeConfigurations {
		if Vars.NodeConfigurations[i].ConfigurationName == configurationName {
			return Vars.NodeConfigurations[i]
		}
	}
	panic("Configuration does not exist.")
}
