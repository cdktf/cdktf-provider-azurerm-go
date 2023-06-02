package automationmodule


type AutomationModuleModuleLink struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.59.0/docs/resources/automation_module#uri AutomationModule#uri}.
	Uri *string `field:"required" json:"uri" yaml:"uri"`
	// hash block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.59.0/docs/resources/automation_module#hash AutomationModule#hash}
	Hash *AutomationModuleModuleLinkHash `field:"optional" json:"hash" yaml:"hash"`
}

