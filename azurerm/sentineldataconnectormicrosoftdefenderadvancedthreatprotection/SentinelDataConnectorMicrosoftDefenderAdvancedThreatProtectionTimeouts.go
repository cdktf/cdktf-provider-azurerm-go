package sentineldataconnectormicrosoftdefenderadvancedthreatprotection


type SentinelDataConnectorMicrosoftDefenderAdvancedThreatProtectionTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sentinel_data_connector_microsoft_defender_advanced_threat_protection#create SentinelDataConnectorMicrosoftDefenderAdvancedThreatProtection#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sentinel_data_connector_microsoft_defender_advanced_threat_protection#delete SentinelDataConnectorMicrosoftDefenderAdvancedThreatProtection#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sentinel_data_connector_microsoft_defender_advanced_threat_protection#read SentinelDataConnectorMicrosoftDefenderAdvancedThreatProtection#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}
