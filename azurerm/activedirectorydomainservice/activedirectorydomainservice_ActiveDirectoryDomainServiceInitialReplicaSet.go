package activedirectorydomainservice


type ActiveDirectoryDomainServiceInitialReplicaSet struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/active_directory_domain_service#subnet_id ActiveDirectoryDomainService#subnet_id}.
	SubnetId *string `field:"required" json:"subnetId" yaml:"subnetId"`
}

