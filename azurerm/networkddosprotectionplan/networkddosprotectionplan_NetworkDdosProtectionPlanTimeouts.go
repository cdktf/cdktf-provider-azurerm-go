package networkddosprotectionplan


type NetworkDdosProtectionPlanTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/network_ddos_protection_plan#create NetworkDdosProtectionPlan#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/network_ddos_protection_plan#delete NetworkDdosProtectionPlan#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/network_ddos_protection_plan#read NetworkDdosProtectionPlan#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/network_ddos_protection_plan#update NetworkDdosProtectionPlan#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}
