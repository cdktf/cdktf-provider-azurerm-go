package mssqlserversecurityalertpolicy


type MssqlServerSecurityAlertPolicyTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_server_security_alert_policy#create MssqlServerSecurityAlertPolicy#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_server_security_alert_policy#delete MssqlServerSecurityAlertPolicy#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_server_security_alert_policy#read MssqlServerSecurityAlertPolicy#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_server_security_alert_policy#update MssqlServerSecurityAlertPolicy#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}
