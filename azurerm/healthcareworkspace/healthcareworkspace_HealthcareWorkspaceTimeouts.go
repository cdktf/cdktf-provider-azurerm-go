package healthcareworkspace


type HealthcareWorkspaceTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/healthcare_workspace#create HealthcareWorkspace#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/healthcare_workspace#delete HealthcareWorkspace#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/healthcare_workspace#read HealthcareWorkspace#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/healthcare_workspace#update HealthcareWorkspace#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}
