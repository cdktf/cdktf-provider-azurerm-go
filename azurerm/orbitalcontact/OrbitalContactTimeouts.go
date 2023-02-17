package orbitalcontact


type OrbitalContactTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/orbital_contact#create OrbitalContact#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/orbital_contact#delete OrbitalContact#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/orbital_contact#read OrbitalContact#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

