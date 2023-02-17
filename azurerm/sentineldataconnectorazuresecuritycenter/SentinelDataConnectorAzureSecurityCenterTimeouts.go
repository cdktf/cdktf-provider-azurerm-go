package sentineldataconnectorazuresecuritycenter


type SentinelDataConnectorAzureSecurityCenterTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sentinel_data_connector_azure_security_center#create SentinelDataConnectorAzureSecurityCenter#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sentinel_data_connector_azure_security_center#delete SentinelDataConnectorAzureSecurityCenter#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sentinel_data_connector_azure_security_center#read SentinelDataConnectorAzureSecurityCenter#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

