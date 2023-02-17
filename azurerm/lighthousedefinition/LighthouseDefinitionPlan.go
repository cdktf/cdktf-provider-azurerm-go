package lighthousedefinition


type LighthouseDefinitionPlan struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/lighthouse_definition#name LighthouseDefinition#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/lighthouse_definition#product LighthouseDefinition#product}.
	Product *string `field:"required" json:"product" yaml:"product"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/lighthouse_definition#publisher LighthouseDefinition#publisher}.
	Publisher *string `field:"required" json:"publisher" yaml:"publisher"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/lighthouse_definition#version LighthouseDefinition#version}.
	Version *string `field:"required" json:"version" yaml:"version"`
}

