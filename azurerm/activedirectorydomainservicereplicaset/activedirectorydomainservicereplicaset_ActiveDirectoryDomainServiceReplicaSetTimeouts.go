package activedirectorydomainservicereplicaset


type ActiveDirectoryDomainServiceReplicaSetTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/active_directory_domain_service_replica_set#create ActiveDirectoryDomainServiceReplicaSet#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/active_directory_domain_service_replica_set#delete ActiveDirectoryDomainServiceReplicaSet#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/active_directory_domain_service_replica_set#read ActiveDirectoryDomainServiceReplicaSet#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/active_directory_domain_service_replica_set#update ActiveDirectoryDomainServiceReplicaSet#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}
