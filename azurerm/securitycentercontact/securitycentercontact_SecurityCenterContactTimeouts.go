package securitycentercontact


type SecurityCenterContactTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/security_center_contact#create SecurityCenterContact#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/security_center_contact#delete SecurityCenterContact#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/security_center_contact#read SecurityCenterContact#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/security_center_contact#update SecurityCenterContact#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}
