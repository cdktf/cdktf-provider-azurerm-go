package logzsubaccounttagrule


type LogzSubAccountTagRuleTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/logz_sub_account_tag_rule#create LogzSubAccountTagRule#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/logz_sub_account_tag_rule#delete LogzSubAccountTagRule#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/logz_sub_account_tag_rule#read LogzSubAccountTagRule#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/logz_sub_account_tag_rule#update LogzSubAccountTagRule#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}
