package automationmodule


type AutomationModuleModuleLinkHash struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/automation_module#algorithm AutomationModule#algorithm}.
	Algorithm *string `field:"required" json:"algorithm" yaml:"algorithm"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/automation_module#value AutomationModule#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

