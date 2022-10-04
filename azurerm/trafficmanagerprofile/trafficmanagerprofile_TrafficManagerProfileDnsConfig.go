package trafficmanagerprofile


type TrafficManagerProfileDnsConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/traffic_manager_profile#relative_name TrafficManagerProfile#relative_name}.
	RelativeName *string `field:"required" json:"relativeName" yaml:"relativeName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/traffic_manager_profile#ttl TrafficManagerProfile#ttl}.
	Ttl *float64 `field:"required" json:"ttl" yaml:"ttl"`
}

