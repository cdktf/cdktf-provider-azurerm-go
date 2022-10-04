package sentinelalertrulenrt


type SentinelAlertRuleNrtAlertDetailsOverride struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sentinel_alert_rule_nrt#description_format SentinelAlertRuleNrt#description_format}.
	DescriptionFormat *string `field:"optional" json:"descriptionFormat" yaml:"descriptionFormat"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sentinel_alert_rule_nrt#display_name_format SentinelAlertRuleNrt#display_name_format}.
	DisplayNameFormat *string `field:"optional" json:"displayNameFormat" yaml:"displayNameFormat"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sentinel_alert_rule_nrt#severity_column_name SentinelAlertRuleNrt#severity_column_name}.
	SeverityColumnName *string `field:"optional" json:"severityColumnName" yaml:"severityColumnName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sentinel_alert_rule_nrt#tactics_column_name SentinelAlertRuleNrt#tactics_column_name}.
	TacticsColumnName *string `field:"optional" json:"tacticsColumnName" yaml:"tacticsColumnName"`
}

