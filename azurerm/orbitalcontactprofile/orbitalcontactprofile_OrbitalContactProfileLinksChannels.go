package orbitalcontactprofile


type OrbitalContactProfileLinksChannels struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/orbital_contact_profile#bandwidth_mhz OrbitalContactProfile#bandwidth_mhz}.
	BandwidthMhz *float64 `field:"required" json:"bandwidthMhz" yaml:"bandwidthMhz"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/orbital_contact_profile#center_frequency_mhz OrbitalContactProfile#center_frequency_mhz}.
	CenterFrequencyMhz *float64 `field:"required" json:"centerFrequencyMhz" yaml:"centerFrequencyMhz"`
	// end_point block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/orbital_contact_profile#end_point OrbitalContactProfile#end_point}
	EndPoint interface{} `field:"required" json:"endPoint" yaml:"endPoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/orbital_contact_profile#name OrbitalContactProfile#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/orbital_contact_profile#demodulation_configuration OrbitalContactProfile#demodulation_configuration}.
	DemodulationConfiguration *string `field:"optional" json:"demodulationConfiguration" yaml:"demodulationConfiguration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/orbital_contact_profile#modulation_configuration OrbitalContactProfile#modulation_configuration}.
	ModulationConfiguration *string `field:"optional" json:"modulationConfiguration" yaml:"modulationConfiguration"`
}
