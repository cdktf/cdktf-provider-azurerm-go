package lighthouseassignment


type LighthouseAssignmentTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/lighthouse_assignment#create LighthouseAssignment#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/lighthouse_assignment#delete LighthouseAssignment#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/lighthouse_assignment#read LighthouseAssignment#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}
