package databricksaccessconnector


type DatabricksAccessConnectorTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/databricks_access_connector#create DatabricksAccessConnector#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/databricks_access_connector#delete DatabricksAccessConnector#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/databricks_access_connector#read DatabricksAccessConnector#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/databricks_access_connector#update DatabricksAccessConnector#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}
