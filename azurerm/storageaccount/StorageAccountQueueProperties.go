package storageaccount


type StorageAccountQueueProperties struct {
	// cors_rule block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#cors_rule StorageAccount#cors_rule}
	CorsRule interface{} `field:"optional" json:"corsRule" yaml:"corsRule"`
	// hour_metrics block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#hour_metrics StorageAccount#hour_metrics}
	HourMetrics *StorageAccountQueuePropertiesHourMetrics `field:"optional" json:"hourMetrics" yaml:"hourMetrics"`
	// logging block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#logging StorageAccount#logging}
	Logging *StorageAccountQueuePropertiesLogging `field:"optional" json:"logging" yaml:"logging"`
	// minute_metrics block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#minute_metrics StorageAccount#minute_metrics}
	MinuteMetrics *StorageAccountQueuePropertiesMinuteMetrics `field:"optional" json:"minuteMetrics" yaml:"minuteMetrics"`
}
