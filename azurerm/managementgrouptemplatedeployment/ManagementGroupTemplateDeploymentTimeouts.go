package managementgrouptemplatedeployment


type ManagementGroupTemplateDeploymentTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/management_group_template_deployment#create ManagementGroupTemplateDeployment#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/management_group_template_deployment#delete ManagementGroupTemplateDeployment#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/management_group_template_deployment#read ManagementGroupTemplateDeployment#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/management_group_template_deployment#update ManagementGroupTemplateDeployment#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

