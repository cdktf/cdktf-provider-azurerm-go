package networkwatcherflowlog


type NetworkWatcherFlowLogRetentionPolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/network_watcher_flow_log#days NetworkWatcherFlowLog#days}.
	Days *float64 `field:"required" json:"days" yaml:"days"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/network_watcher_flow_log#enabled NetworkWatcherFlowLog#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

