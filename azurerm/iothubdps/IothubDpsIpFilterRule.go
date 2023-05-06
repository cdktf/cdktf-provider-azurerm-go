package iothubdps


type IothubDpsIpFilterRule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/iothub_dps#action IothubDps#action}.
	Action *string `field:"required" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/iothub_dps#ip_mask IothubDps#ip_mask}.
	IpMask *string `field:"required" json:"ipMask" yaml:"ipMask"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/iothub_dps#name IothubDps#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/iothub_dps#target IothubDps#target}.
	Target *string `field:"optional" json:"target" yaml:"target"`
}

