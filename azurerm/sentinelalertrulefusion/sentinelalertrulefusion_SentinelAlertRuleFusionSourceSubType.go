package sentinelalertrulefusion


type SentinelAlertRuleFusionSourceSubType struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sentinel_alert_rule_fusion#name SentinelAlertRuleFusion#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sentinel_alert_rule_fusion#severities_allowed SentinelAlertRuleFusion#severities_allowed}.
	SeveritiesAllowed *[]*string `field:"required" json:"severitiesAllowed" yaml:"severitiesAllowed"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sentinel_alert_rule_fusion#enabled SentinelAlertRuleFusion#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

