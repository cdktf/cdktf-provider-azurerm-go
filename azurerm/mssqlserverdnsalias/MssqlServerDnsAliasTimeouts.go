package mssqlserverdnsalias


type MssqlServerDnsAliasTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.68.0/docs/resources/mssql_server_dns_alias#create MssqlServerDnsAlias#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.68.0/docs/resources/mssql_server_dns_alias#delete MssqlServerDnsAlias#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.68.0/docs/resources/mssql_server_dns_alias#read MssqlServerDnsAlias#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

