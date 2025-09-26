// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oracleautonomousdatabase


type OracleAutonomousDatabaseLongTermBackupSchedule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.46.0/docs/resources/oracle_autonomous_database#enabled OracleAutonomousDatabase#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.46.0/docs/resources/oracle_autonomous_database#repeat_cadence OracleAutonomousDatabase#repeat_cadence}.
	RepeatCadence *string `field:"required" json:"repeatCadence" yaml:"repeatCadence"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.46.0/docs/resources/oracle_autonomous_database#retention_period_in_days OracleAutonomousDatabase#retention_period_in_days}.
	RetentionPeriodInDays *float64 `field:"required" json:"retentionPeriodInDays" yaml:"retentionPeriodInDays"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.46.0/docs/resources/oracle_autonomous_database#time_of_backup OracleAutonomousDatabase#time_of_backup}.
	TimeOfBackup *string `field:"required" json:"timeOfBackup" yaml:"timeOfBackup"`
}

