package sqlmanagedinstancefailovergroup


type SqlManagedInstanceFailoverGroupTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.61.0/docs/resources/sql_managed_instance_failover_group#create SqlManagedInstanceFailoverGroup#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.61.0/docs/resources/sql_managed_instance_failover_group#delete SqlManagedInstanceFailoverGroup#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.61.0/docs/resources/sql_managed_instance_failover_group#read SqlManagedInstanceFailoverGroup#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.61.0/docs/resources/sql_managed_instance_failover_group#update SqlManagedInstanceFailoverGroup#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

