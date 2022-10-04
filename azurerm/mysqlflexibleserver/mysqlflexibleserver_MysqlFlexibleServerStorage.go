package mysqlflexibleserver


type MysqlFlexibleServerStorage struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mysql_flexible_server#auto_grow_enabled MysqlFlexibleServer#auto_grow_enabled}.
	AutoGrowEnabled interface{} `field:"optional" json:"autoGrowEnabled" yaml:"autoGrowEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mysql_flexible_server#iops MysqlFlexibleServer#iops}.
	Iops *float64 `field:"optional" json:"iops" yaml:"iops"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mysql_flexible_server#size_gb MysqlFlexibleServer#size_gb}.
	SizeGb *float64 `field:"optional" json:"sizeGb" yaml:"sizeGb"`
}

