package orbitalspacecraft


type OrbitalSpacecraftLinks struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/orbital_spacecraft#bandwidth_mhz OrbitalSpacecraft#bandwidth_mhz}.
	BandwidthMhz *float64 `field:"required" json:"bandwidthMhz" yaml:"bandwidthMhz"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/orbital_spacecraft#center_frequency_mhz OrbitalSpacecraft#center_frequency_mhz}.
	CenterFrequencyMhz *float64 `field:"required" json:"centerFrequencyMhz" yaml:"centerFrequencyMhz"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/orbital_spacecraft#direction OrbitalSpacecraft#direction}.
	Direction *string `field:"required" json:"direction" yaml:"direction"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/orbital_spacecraft#name OrbitalSpacecraft#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/orbital_spacecraft#polarization OrbitalSpacecraft#polarization}.
	Polarization *string `field:"required" json:"polarization" yaml:"polarization"`
}

