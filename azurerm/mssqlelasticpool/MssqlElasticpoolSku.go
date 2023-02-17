package mssqlelasticpool


type MssqlElasticpoolSku struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_elasticpool#capacity MssqlElasticpool#capacity}.
	Capacity *float64 `field:"required" json:"capacity" yaml:"capacity"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_elasticpool#name MssqlElasticpool#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_elasticpool#tier MssqlElasticpool#tier}.
	Tier *string `field:"required" json:"tier" yaml:"tier"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_elasticpool#family MssqlElasticpool#family}.
	Family *string `field:"optional" json:"family" yaml:"family"`
}

