package mysqlactivedirectoryadministrator


type MysqlActiveDirectoryAdministratorTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mysql_active_directory_administrator#create MysqlActiveDirectoryAdministrator#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mysql_active_directory_administrator#delete MysqlActiveDirectoryAdministrator#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mysql_active_directory_administrator#read MysqlActiveDirectoryAdministrator#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mysql_active_directory_administrator#update MysqlActiveDirectoryAdministrator#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}
