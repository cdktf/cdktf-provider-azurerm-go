// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package activedirectorydomainservice

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ActiveDirectoryDomainServiceConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/active_directory_domain_service#domain_name ActiveDirectoryDomainService#domain_name}.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// initial_replica_set block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/active_directory_domain_service#initial_replica_set ActiveDirectoryDomainService#initial_replica_set}
	InitialReplicaSet *ActiveDirectoryDomainServiceInitialReplicaSet `field:"required" json:"initialReplicaSet" yaml:"initialReplicaSet"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/active_directory_domain_service#location ActiveDirectoryDomainService#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/active_directory_domain_service#name ActiveDirectoryDomainService#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/active_directory_domain_service#resource_group_name ActiveDirectoryDomainService#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/active_directory_domain_service#sku ActiveDirectoryDomainService#sku}.
	Sku *string `field:"required" json:"sku" yaml:"sku"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/active_directory_domain_service#domain_configuration_type ActiveDirectoryDomainService#domain_configuration_type}.
	DomainConfigurationType *string `field:"optional" json:"domainConfigurationType" yaml:"domainConfigurationType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/active_directory_domain_service#filtered_sync_enabled ActiveDirectoryDomainService#filtered_sync_enabled}.
	FilteredSyncEnabled interface{} `field:"optional" json:"filteredSyncEnabled" yaml:"filteredSyncEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/active_directory_domain_service#id ActiveDirectoryDomainService#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// notifications block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/active_directory_domain_service#notifications ActiveDirectoryDomainService#notifications}
	Notifications *ActiveDirectoryDomainServiceNotifications `field:"optional" json:"notifications" yaml:"notifications"`
	// secure_ldap block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/active_directory_domain_service#secure_ldap ActiveDirectoryDomainService#secure_ldap}
	SecureLdap *ActiveDirectoryDomainServiceSecureLdap `field:"optional" json:"secureLdap" yaml:"secureLdap"`
	// security block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/active_directory_domain_service#security ActiveDirectoryDomainService#security}
	Security *ActiveDirectoryDomainServiceSecurity `field:"optional" json:"security" yaml:"security"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/active_directory_domain_service#tags ActiveDirectoryDomainService#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/active_directory_domain_service#timeouts ActiveDirectoryDomainService#timeouts}
	Timeouts *ActiveDirectoryDomainServiceTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

