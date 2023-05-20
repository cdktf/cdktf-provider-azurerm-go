package sentinelalertrulenrt


type SentinelAlertRuleNrtIncident struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.57.0/docs/resources/sentinel_alert_rule_nrt#create_incident_enabled SentinelAlertRuleNrt#create_incident_enabled}.
	CreateIncidentEnabled interface{} `field:"required" json:"createIncidentEnabled" yaml:"createIncidentEnabled"`
	// grouping block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.57.0/docs/resources/sentinel_alert_rule_nrt#grouping SentinelAlertRuleNrt#grouping}
	Grouping *SentinelAlertRuleNrtIncidentGrouping `field:"required" json:"grouping" yaml:"grouping"`
}

