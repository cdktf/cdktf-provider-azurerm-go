package kustoeventgriddataconnection


type KustoEventgridDataConnectionTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/kusto_eventgrid_data_connection#create KustoEventgridDataConnection#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/kusto_eventgrid_data_connection#delete KustoEventgridDataConnection#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/kusto_eventgrid_data_connection#read KustoEventgridDataConnection#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/kusto_eventgrid_data_connection#update KustoEventgridDataConnection#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

