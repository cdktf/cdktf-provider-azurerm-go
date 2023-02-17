package iothub


type IothubNetworkRuleSetIpRule struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/iothub#ip_mask Iothub#ip_mask}.
	IpMask *string `field:"required" json:"ipMask" yaml:"ipMask"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/iothub#name Iothub#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/iothub#action Iothub#action}.
	Action *string `field:"optional" json:"action" yaml:"action"`
}

