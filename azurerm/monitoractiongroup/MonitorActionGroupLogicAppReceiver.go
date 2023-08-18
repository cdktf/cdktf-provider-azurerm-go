package monitoractiongroup


type MonitorActionGroupLogicAppReceiver struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/monitor_action_group#callback_url MonitorActionGroup#callback_url}.
	CallbackUrl *string `field:"required" json:"callbackUrl" yaml:"callbackUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/monitor_action_group#name MonitorActionGroup#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/monitor_action_group#resource_id MonitorActionGroup#resource_id}.
	ResourceId *string `field:"required" json:"resourceId" yaml:"resourceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/monitor_action_group#use_common_alert_schema MonitorActionGroup#use_common_alert_schema}.
	UseCommonAlertSchema interface{} `field:"optional" json:"useCommonAlertSchema" yaml:"useCommonAlertSchema"`
}

