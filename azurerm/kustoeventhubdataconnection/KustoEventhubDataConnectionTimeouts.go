package kustoeventhubdataconnection


type KustoEventhubDataConnectionTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/kusto_eventhub_data_connection#create KustoEventhubDataConnection#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/kusto_eventhub_data_connection#delete KustoEventhubDataConnection#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/kusto_eventhub_data_connection#read KustoEventhubDataConnection#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/kusto_eventhub_data_connection#update KustoEventhubDataConnection#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

