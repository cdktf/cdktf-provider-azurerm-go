package synapsesqlpoolsecurityalertpolicy


type SynapseSqlPoolSecurityAlertPolicyTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/synapse_sql_pool_security_alert_policy#create SynapseSqlPoolSecurityAlertPolicy#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/synapse_sql_pool_security_alert_policy#delete SynapseSqlPoolSecurityAlertPolicy#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/synapse_sql_pool_security_alert_policy#read SynapseSqlPoolSecurityAlertPolicy#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/synapse_sql_pool_security_alert_policy#update SynapseSqlPoolSecurityAlertPolicy#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}
