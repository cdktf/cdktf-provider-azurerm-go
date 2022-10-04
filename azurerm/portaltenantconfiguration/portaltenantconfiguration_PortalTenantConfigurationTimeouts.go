package portaltenantconfiguration


type PortalTenantConfigurationTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/portal_tenant_configuration#create PortalTenantConfiguration#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/portal_tenant_configuration#delete PortalTenantConfiguration#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/portal_tenant_configuration#read PortalTenantConfiguration#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/portal_tenant_configuration#update PortalTenantConfiguration#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

