package securitycenterassessmentpolicy


type SecurityCenterAssessmentPolicyTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/security_center_assessment_policy#create SecurityCenterAssessmentPolicy#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/security_center_assessment_policy#delete SecurityCenterAssessmentPolicy#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/security_center_assessment_policy#read SecurityCenterAssessmentPolicy#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/security_center_assessment_policy#update SecurityCenterAssessmentPolicy#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}
