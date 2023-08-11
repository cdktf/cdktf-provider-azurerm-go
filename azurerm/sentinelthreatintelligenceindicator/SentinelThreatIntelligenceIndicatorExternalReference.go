package sentinelthreatintelligenceindicator


type SentinelThreatIntelligenceIndicatorExternalReference struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/sentinel_threat_intelligence_indicator#description SentinelThreatIntelligenceIndicator#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/sentinel_threat_intelligence_indicator#hashes SentinelThreatIntelligenceIndicator#hashes}.
	Hashes *map[string]*string `field:"optional" json:"hashes" yaml:"hashes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/sentinel_threat_intelligence_indicator#source_name SentinelThreatIntelligenceIndicator#source_name}.
	SourceName *string `field:"optional" json:"sourceName" yaml:"sourceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/sentinel_threat_intelligence_indicator#url SentinelThreatIntelligenceIndicator#url}.
	Url *string `field:"optional" json:"url" yaml:"url"`
}

