package orbitalcontactprofile


type OrbitalContactProfileLinks struct {
	// channels block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/orbital_contact_profile#channels OrbitalContactProfile#channels}
	Channels interface{} `field:"required" json:"channels" yaml:"channels"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/orbital_contact_profile#direction OrbitalContactProfile#direction}.
	Direction *string `field:"required" json:"direction" yaml:"direction"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/orbital_contact_profile#name OrbitalContactProfile#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/orbital_contact_profile#polarization OrbitalContactProfile#polarization}.
	Polarization *string `field:"required" json:"polarization" yaml:"polarization"`
}
