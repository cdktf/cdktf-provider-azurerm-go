package activedirectorydomainservicetrust

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ActiveDirectoryDomainServiceTrustConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/active_directory_domain_service_trust#domain_service_id ActiveDirectoryDomainServiceTrust#domain_service_id}.
	DomainServiceId *string `field:"required" json:"domainServiceId" yaml:"domainServiceId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/active_directory_domain_service_trust#name ActiveDirectoryDomainServiceTrust#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/active_directory_domain_service_trust#password ActiveDirectoryDomainServiceTrust#password}.
	Password *string `field:"required" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/active_directory_domain_service_trust#trusted_domain_dns_ips ActiveDirectoryDomainServiceTrust#trusted_domain_dns_ips}.
	TrustedDomainDnsIps *[]*string `field:"required" json:"trustedDomainDnsIps" yaml:"trustedDomainDnsIps"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/active_directory_domain_service_trust#trusted_domain_fqdn ActiveDirectoryDomainServiceTrust#trusted_domain_fqdn}.
	TrustedDomainFqdn *string `field:"required" json:"trustedDomainFqdn" yaml:"trustedDomainFqdn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/active_directory_domain_service_trust#id ActiveDirectoryDomainServiceTrust#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/active_directory_domain_service_trust#timeouts ActiveDirectoryDomainServiceTrust#timeouts}
	Timeouts *ActiveDirectoryDomainServiceTrustTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

