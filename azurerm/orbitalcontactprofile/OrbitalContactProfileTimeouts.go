package orbitalcontactprofile


type OrbitalContactProfileTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/orbital_contact_profile#create OrbitalContactProfile#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/orbital_contact_profile#delete OrbitalContactProfile#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/orbital_contact_profile#read OrbitalContactProfile#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/orbital_contact_profile#update OrbitalContactProfile#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}
