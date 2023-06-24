package automanageconfiguration


type AutomanageConfigurationBackupRetentionPolicyDailySchedule struct {
	// retention_duration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/automanage_configuration#retention_duration AutomanageConfiguration#retention_duration}
	RetentionDuration *AutomanageConfigurationBackupRetentionPolicyDailyScheduleRetentionDuration `field:"optional" json:"retentionDuration" yaml:"retentionDuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/automanage_configuration#retention_times AutomanageConfiguration#retention_times}.
	RetentionTimes *[]*string `field:"optional" json:"retentionTimes" yaml:"retentionTimes"`
}

