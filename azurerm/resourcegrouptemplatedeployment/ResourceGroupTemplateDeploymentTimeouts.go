package resourcegrouptemplatedeployment


type ResourceGroupTemplateDeploymentTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/resource_group_template_deployment#create ResourceGroupTemplateDeployment#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/resource_group_template_deployment#delete ResourceGroupTemplateDeployment#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/resource_group_template_deployment#read ResourceGroupTemplateDeployment#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/resource_group_template_deployment#update ResourceGroupTemplateDeployment#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

