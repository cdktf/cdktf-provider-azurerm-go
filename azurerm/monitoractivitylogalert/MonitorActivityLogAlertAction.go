package monitoractivitylogalert


type MonitorActivityLogAlertAction struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/monitor_activity_log_alert#action_group_id MonitorActivityLogAlert#action_group_id}.
	ActionGroupId *string `field:"required" json:"actionGroupId" yaml:"actionGroupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/monitor_activity_log_alert#webhook_properties MonitorActivityLogAlert#webhook_properties}.
	WebhookProperties *map[string]*string `field:"optional" json:"webhookProperties" yaml:"webhookProperties"`
}

