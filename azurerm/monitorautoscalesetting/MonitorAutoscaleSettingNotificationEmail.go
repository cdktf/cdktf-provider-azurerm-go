package monitorautoscalesetting


type MonitorAutoscaleSettingNotificationEmail struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/monitor_autoscale_setting#custom_emails MonitorAutoscaleSetting#custom_emails}.
	CustomEmails *[]*string `field:"optional" json:"customEmails" yaml:"customEmails"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/monitor_autoscale_setting#send_to_subscription_administrator MonitorAutoscaleSetting#send_to_subscription_administrator}.
	SendToSubscriptionAdministrator interface{} `field:"optional" json:"sendToSubscriptionAdministrator" yaml:"sendToSubscriptionAdministrator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/monitor_autoscale_setting#send_to_subscription_co_administrator MonitorAutoscaleSetting#send_to_subscription_co_administrator}.
	SendToSubscriptionCoAdministrator interface{} `field:"optional" json:"sendToSubscriptionCoAdministrator" yaml:"sendToSubscriptionCoAdministrator"`
}

