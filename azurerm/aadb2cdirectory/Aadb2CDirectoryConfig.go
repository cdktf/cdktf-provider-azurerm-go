package aadb2cdirectory

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Aadb2CDirectoryConfig struct {
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
	// Location in which the B2C tenant is hosted and data resides. See https://aka.ms/B2CDataResidency for more information.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.66.0/docs/resources/aadb2c_directory#data_residency_location Aadb2CDirectory#data_residency_location}
	DataResidencyLocation *string `field:"required" json:"dataResidencyLocation" yaml:"dataResidencyLocation"`
	// Domain name of the B2C tenant, including onmicrosoft.com suffix.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.66.0/docs/resources/aadb2c_directory#domain_name Aadb2CDirectory#domain_name}
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.66.0/docs/resources/aadb2c_directory#resource_group_name Aadb2CDirectory#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Billing SKU for the B2C tenant. See https://aka.ms/b2cBilling for more information.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.66.0/docs/resources/aadb2c_directory#sku_name Aadb2CDirectory#sku_name}
	SkuName *string `field:"required" json:"skuName" yaml:"skuName"`
	// Country code of the B2C tenant. See https://aka.ms/B2CDataResidency for valid country codes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.66.0/docs/resources/aadb2c_directory#country_code Aadb2CDirectory#country_code}
	CountryCode *string `field:"optional" json:"countryCode" yaml:"countryCode"`
	// The initial display name of the B2C tenant.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.66.0/docs/resources/aadb2c_directory#display_name Aadb2CDirectory#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.66.0/docs/resources/aadb2c_directory#id Aadb2CDirectory#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.66.0/docs/resources/aadb2c_directory#tags Aadb2CDirectory#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.66.0/docs/resources/aadb2c_directory#timeouts Aadb2CDirectory#timeouts}
	Timeouts *Aadb2CDirectoryTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

