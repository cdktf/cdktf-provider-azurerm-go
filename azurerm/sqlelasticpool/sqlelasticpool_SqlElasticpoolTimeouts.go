package sqlelasticpool


type SqlElasticpoolTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sql_elasticpool#create SqlElasticpool#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sql_elasticpool#delete SqlElasticpool#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sql_elasticpool#read SqlElasticpool#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sql_elasticpool#update SqlElasticpool#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}
