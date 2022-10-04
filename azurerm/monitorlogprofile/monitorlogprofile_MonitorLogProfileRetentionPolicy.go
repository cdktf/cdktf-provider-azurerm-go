package monitorlogprofile


type MonitorLogProfileRetentionPolicy struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_log_profile#enabled MonitorLogProfile#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_log_profile#days MonitorLogProfile#days}.
	Days *float64 `field:"optional" json:"days" yaml:"days"`
}

