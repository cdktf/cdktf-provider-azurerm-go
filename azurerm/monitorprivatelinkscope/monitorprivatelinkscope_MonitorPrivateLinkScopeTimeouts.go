package monitorprivatelinkscope


type MonitorPrivateLinkScopeTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_private_link_scope#create MonitorPrivateLinkScope#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_private_link_scope#delete MonitorPrivateLinkScope#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_private_link_scope#read MonitorPrivateLinkScope#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_private_link_scope#update MonitorPrivateLinkScope#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

