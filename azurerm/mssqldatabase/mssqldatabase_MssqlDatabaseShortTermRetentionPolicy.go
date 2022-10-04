package mssqldatabase


type MssqlDatabaseShortTermRetentionPolicy struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_database#retention_days MssqlDatabase#retention_days}.
	RetentionDays *float64 `field:"required" json:"retentionDays" yaml:"retentionDays"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_database#backup_interval_in_hours MssqlDatabase#backup_interval_in_hours}.
	BackupIntervalInHours *float64 `field:"optional" json:"backupIntervalInHours" yaml:"backupIntervalInHours"`
}

