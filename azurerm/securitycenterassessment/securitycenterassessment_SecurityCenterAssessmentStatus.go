package securitycenterassessment


type SecurityCenterAssessmentStatus struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/security_center_assessment#code SecurityCenterAssessment#code}.
	Code *string `field:"required" json:"code" yaml:"code"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/security_center_assessment#cause SecurityCenterAssessment#cause}.
	Cause *string `field:"optional" json:"cause" yaml:"cause"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/security_center_assessment#description SecurityCenterAssessment#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

