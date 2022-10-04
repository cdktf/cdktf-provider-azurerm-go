package automationmodule


type AutomationModuleModuleLink struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/automation_module#uri AutomationModule#uri}.
	Uri *string `field:"required" json:"uri" yaml:"uri"`
	// hash block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/automation_module#hash AutomationModule#hash}
	Hash *AutomationModuleModuleLinkHash `field:"optional" json:"hash" yaml:"hash"`
}

