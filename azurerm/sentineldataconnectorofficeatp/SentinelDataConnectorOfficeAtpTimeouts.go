package sentineldataconnectorofficeatp


type SentinelDataConnectorOfficeAtpTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sentinel_data_connector_office_atp#create SentinelDataConnectorOfficeAtp#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sentinel_data_connector_office_atp#delete SentinelDataConnectorOfficeAtp#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sentinel_data_connector_office_atp#read SentinelDataConnectorOfficeAtp#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}
