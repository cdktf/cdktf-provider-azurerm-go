package cognitivedeployment


type CognitiveDeploymentModel struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cognitive_deployment#format CognitiveDeployment#format}.
	Format *string `field:"required" json:"format" yaml:"format"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cognitive_deployment#name CognitiveDeployment#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cognitive_deployment#version CognitiveDeployment#version}.
	Version *string `field:"required" json:"version" yaml:"version"`
}

