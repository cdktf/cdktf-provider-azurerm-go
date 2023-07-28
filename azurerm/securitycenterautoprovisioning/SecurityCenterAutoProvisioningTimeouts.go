package securitycenterautoprovisioning


type SecurityCenterAutoProvisioningTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/security_center_auto_provisioning#create SecurityCenterAutoProvisioning#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/security_center_auto_provisioning#delete SecurityCenterAutoProvisioning#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/security_center_auto_provisioning#read SecurityCenterAutoProvisioning#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/security_center_auto_provisioning#update SecurityCenterAutoProvisioning#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

