package tenanttemplatedeployment


type TenantTemplateDeploymentTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/tenant_template_deployment#create TenantTemplateDeployment#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/tenant_template_deployment#delete TenantTemplateDeployment#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/tenant_template_deployment#read TenantTemplateDeployment#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/tenant_template_deployment#update TenantTemplateDeployment#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}
