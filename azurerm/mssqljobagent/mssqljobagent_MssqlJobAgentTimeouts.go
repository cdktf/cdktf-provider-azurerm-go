package mssqljobagent


type MssqlJobAgentTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_job_agent#create MssqlJobAgent#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_job_agent#delete MssqlJobAgent#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_job_agent#read MssqlJobAgent#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_job_agent#update MssqlJobAgent#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}
