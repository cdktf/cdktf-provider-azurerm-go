package logzsubaccounttagrule


type LogzSubAccountTagRuleTagFilter struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/logz_sub_account_tag_rule#action LogzSubAccountTagRule#action}.
	Action *string `field:"required" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/logz_sub_account_tag_rule#name LogzSubAccountTagRule#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/logz_sub_account_tag_rule#value LogzSubAccountTagRule#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

