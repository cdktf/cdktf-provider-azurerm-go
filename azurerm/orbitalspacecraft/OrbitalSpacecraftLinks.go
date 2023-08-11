package orbitalspacecraft


type OrbitalSpacecraftLinks struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/orbital_spacecraft#bandwidth_mhz OrbitalSpacecraft#bandwidth_mhz}.
	BandwidthMhz *float64 `field:"required" json:"bandwidthMhz" yaml:"bandwidthMhz"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/orbital_spacecraft#center_frequency_mhz OrbitalSpacecraft#center_frequency_mhz}.
	CenterFrequencyMhz *float64 `field:"required" json:"centerFrequencyMhz" yaml:"centerFrequencyMhz"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/orbital_spacecraft#direction OrbitalSpacecraft#direction}.
	Direction *string `field:"required" json:"direction" yaml:"direction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/orbital_spacecraft#name OrbitalSpacecraft#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/orbital_spacecraft#polarization OrbitalSpacecraft#polarization}.
	Polarization *string `field:"required" json:"polarization" yaml:"polarization"`
}

