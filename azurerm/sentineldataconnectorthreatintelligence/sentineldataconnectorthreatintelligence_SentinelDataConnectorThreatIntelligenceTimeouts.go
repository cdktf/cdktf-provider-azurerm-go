package sentineldataconnectorthreatintelligence


type SentinelDataConnectorThreatIntelligenceTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sentinel_data_connector_threat_intelligence#create SentinelDataConnectorThreatIntelligence#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sentinel_data_connector_threat_intelligence#delete SentinelDataConnectorThreatIntelligence#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sentinel_data_connector_threat_intelligence#read SentinelDataConnectorThreatIntelligence#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}
