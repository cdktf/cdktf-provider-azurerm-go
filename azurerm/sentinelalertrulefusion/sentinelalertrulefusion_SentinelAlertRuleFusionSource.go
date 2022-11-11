package sentinelalertrulefusion


type SentinelAlertRuleFusionSource struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sentinel_alert_rule_fusion#name SentinelAlertRuleFusion#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sentinel_alert_rule_fusion#enabled SentinelAlertRuleFusion#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// sub_type block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sentinel_alert_rule_fusion#sub_type SentinelAlertRuleFusion#sub_type}
	SubType interface{} `field:"optional" json:"subType" yaml:"subType"`
}

