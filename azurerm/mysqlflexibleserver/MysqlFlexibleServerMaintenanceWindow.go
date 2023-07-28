package mysqlflexibleserver


type MysqlFlexibleServerMaintenanceWindow struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/mysql_flexible_server#day_of_week MysqlFlexibleServer#day_of_week}.
	DayOfWeek *float64 `field:"optional" json:"dayOfWeek" yaml:"dayOfWeek"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/mysql_flexible_server#start_hour MysqlFlexibleServer#start_hour}.
	StartHour *float64 `field:"optional" json:"startHour" yaml:"startHour"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/mysql_flexible_server#start_minute MysqlFlexibleServer#start_minute}.
	StartMinute *float64 `field:"optional" json:"startMinute" yaml:"startMinute"`
}

