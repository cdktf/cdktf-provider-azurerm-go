package sentinelalertrulemssecurityincident


type SentinelAlertRuleMsSecurityIncidentTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sentinel_alert_rule_ms_security_incident#create SentinelAlertRuleMsSecurityIncident#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sentinel_alert_rule_ms_security_incident#delete SentinelAlertRuleMsSecurityIncident#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sentinel_alert_rule_ms_security_incident#read SentinelAlertRuleMsSecurityIncident#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sentinel_alert_rule_ms_security_incident#update SentinelAlertRuleMsSecurityIncident#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}
