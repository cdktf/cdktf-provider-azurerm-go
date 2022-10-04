package iothub


type IothubNetworkRuleSet struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/iothub#apply_to_builtin_eventhub_endpoint Iothub#apply_to_builtin_eventhub_endpoint}.
	ApplyToBuiltinEventhubEndpoint interface{} `field:"optional" json:"applyToBuiltinEventhubEndpoint" yaml:"applyToBuiltinEventhubEndpoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/iothub#default_action Iothub#default_action}.
	DefaultAction *string `field:"optional" json:"defaultAction" yaml:"defaultAction"`
	// ip_rule block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/iothub#ip_rule Iothub#ip_rule}
	IpRule interface{} `field:"optional" json:"ipRule" yaml:"ipRule"`
}

