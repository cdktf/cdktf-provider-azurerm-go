package cognitivedeployment


type CognitiveDeploymentScale struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/cognitive_deployment#type CognitiveDeployment#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/cognitive_deployment#capacity CognitiveDeployment#capacity}.
	Capacity *float64 `field:"optional" json:"capacity" yaml:"capacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/cognitive_deployment#family CognitiveDeployment#family}.
	Family *string `field:"optional" json:"family" yaml:"family"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/cognitive_deployment#size CognitiveDeployment#size}.
	Size *string `field:"optional" json:"size" yaml:"size"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/cognitive_deployment#tier CognitiveDeployment#tier}.
	Tier *string `field:"optional" json:"tier" yaml:"tier"`
}

