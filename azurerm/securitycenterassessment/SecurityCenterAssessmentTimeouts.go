package securitycenterassessment


type SecurityCenterAssessmentTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/security_center_assessment#create SecurityCenterAssessment#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/security_center_assessment#delete SecurityCenterAssessment#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/security_center_assessment#read SecurityCenterAssessment#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/security_center_assessment#update SecurityCenterAssessment#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

