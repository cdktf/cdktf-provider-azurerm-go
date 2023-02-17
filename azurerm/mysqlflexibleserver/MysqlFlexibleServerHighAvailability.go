package mysqlflexibleserver


type MysqlFlexibleServerHighAvailability struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mysql_flexible_server#mode MysqlFlexibleServer#mode}.
	Mode *string `field:"required" json:"mode" yaml:"mode"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mysql_flexible_server#standby_availability_zone MysqlFlexibleServer#standby_availability_zone}.
	StandbyAvailabilityZone *string `field:"optional" json:"standbyAvailabilityZone" yaml:"standbyAvailabilityZone"`
}

