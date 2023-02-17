package monitoractivitylogalert


type MonitorActivityLogAlertCriteriaResourceHealth struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_activity_log_alert#current MonitorActivityLogAlert#current}.
	Current *[]*string `field:"optional" json:"current" yaml:"current"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_activity_log_alert#previous MonitorActivityLogAlert#previous}.
	Previous *[]*string `field:"optional" json:"previous" yaml:"previous"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_activity_log_alert#reason MonitorActivityLogAlert#reason}.
	Reason *[]*string `field:"optional" json:"reason" yaml:"reason"`
}

