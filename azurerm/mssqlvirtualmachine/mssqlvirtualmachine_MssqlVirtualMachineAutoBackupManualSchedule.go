package mssqlvirtualmachine


type MssqlVirtualMachineAutoBackupManualSchedule struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_virtual_machine#full_backup_frequency MssqlVirtualMachine#full_backup_frequency}.
	FullBackupFrequency *string `field:"required" json:"fullBackupFrequency" yaml:"fullBackupFrequency"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_virtual_machine#full_backup_start_hour MssqlVirtualMachine#full_backup_start_hour}.
	FullBackupStartHour *float64 `field:"required" json:"fullBackupStartHour" yaml:"fullBackupStartHour"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_virtual_machine#full_backup_window_in_hours MssqlVirtualMachine#full_backup_window_in_hours}.
	FullBackupWindowInHours *float64 `field:"required" json:"fullBackupWindowInHours" yaml:"fullBackupWindowInHours"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_virtual_machine#log_backup_frequency_in_minutes MssqlVirtualMachine#log_backup_frequency_in_minutes}.
	LogBackupFrequencyInMinutes *float64 `field:"required" json:"logBackupFrequencyInMinutes" yaml:"logBackupFrequencyInMinutes"`
}

