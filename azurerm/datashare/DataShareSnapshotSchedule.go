package datashare


type DataShareSnapshotSchedule struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_share#name DataShare#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_share#recurrence DataShare#recurrence}.
	Recurrence *string `field:"required" json:"recurrence" yaml:"recurrence"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_share#start_time DataShare#start_time}.
	StartTime *string `field:"required" json:"startTime" yaml:"startTime"`
}

