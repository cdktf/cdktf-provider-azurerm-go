package automationmodule


type AutomationModuleModuleLinkHash struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/automation_module#algorithm AutomationModule#algorithm}.
	Algorithm *string `field:"required" json:"algorithm" yaml:"algorithm"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/automation_module#value AutomationModule#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

