package digitaltwinstimeseriesdatabaseconnection


type DigitalTwinsTimeSeriesDatabaseConnectionTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/digital_twins_time_series_database_connection#create DigitalTwinsTimeSeriesDatabaseConnection#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/digital_twins_time_series_database_connection#delete DigitalTwinsTimeSeriesDatabaseConnection#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/digital_twins_time_series_database_connection#read DigitalTwinsTimeSeriesDatabaseConnection#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}
