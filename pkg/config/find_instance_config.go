package config

func FindInstanceConfiguration(name string) InstanceConfiguration {
	for _, ic := range Vars.InstanceConfigurations {
		if ic.Name == name {
			return ic
		}
	}
	panic("Instance configuration not found.\nEdit ~/.cloud/config.json to add more instance configs.")
}
