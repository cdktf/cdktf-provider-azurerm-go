package mssqlmanagedinstance

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MssqlManagedInstanceConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/mssql_managed_instance#administrator_login MssqlManagedInstance#administrator_login}.
	AdministratorLogin *string `field:"required" json:"administratorLogin" yaml:"administratorLogin"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/mssql_managed_instance#administrator_login_password MssqlManagedInstance#administrator_login_password}.
	AdministratorLoginPassword *string `field:"required" json:"administratorLoginPassword" yaml:"administratorLoginPassword"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/mssql_managed_instance#license_type MssqlManagedInstance#license_type}.
	LicenseType *string `field:"required" json:"licenseType" yaml:"licenseType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/mssql_managed_instance#location MssqlManagedInstance#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/mssql_managed_instance#name MssqlManagedInstance#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/mssql_managed_instance#resource_group_name MssqlManagedInstance#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/mssql_managed_instance#sku_name MssqlManagedInstance#sku_name}.
	SkuName *string `field:"required" json:"skuName" yaml:"skuName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/mssql_managed_instance#storage_size_in_gb MssqlManagedInstance#storage_size_in_gb}.
	StorageSizeInGb *float64 `field:"required" json:"storageSizeInGb" yaml:"storageSizeInGb"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/mssql_managed_instance#subnet_id MssqlManagedInstance#subnet_id}.
	SubnetId *string `field:"required" json:"subnetId" yaml:"subnetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/mssql_managed_instance#vcores MssqlManagedInstance#vcores}.
	Vcores *float64 `field:"required" json:"vcores" yaml:"vcores"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/mssql_managed_instance#collation MssqlManagedInstance#collation}.
	Collation *string `field:"optional" json:"collation" yaml:"collation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/mssql_managed_instance#dns_zone_partner_id MssqlManagedInstance#dns_zone_partner_id}.
	DnsZonePartnerId *string `field:"optional" json:"dnsZonePartnerId" yaml:"dnsZonePartnerId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/mssql_managed_instance#id MssqlManagedInstance#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// identity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/mssql_managed_instance#identity MssqlManagedInstance#identity}
	Identity *MssqlManagedInstanceIdentity `field:"optional" json:"identity" yaml:"identity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/mssql_managed_instance#maintenance_configuration_name MssqlManagedInstance#maintenance_configuration_name}.
	MaintenanceConfigurationName *string `field:"optional" json:"maintenanceConfigurationName" yaml:"maintenanceConfigurationName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/mssql_managed_instance#minimum_tls_version MssqlManagedInstance#minimum_tls_version}.
	MinimumTlsVersion *string `field:"optional" json:"minimumTlsVersion" yaml:"minimumTlsVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/mssql_managed_instance#proxy_override MssqlManagedInstance#proxy_override}.
	ProxyOverride *string `field:"optional" json:"proxyOverride" yaml:"proxyOverride"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/mssql_managed_instance#public_data_endpoint_enabled MssqlManagedInstance#public_data_endpoint_enabled}.
	PublicDataEndpointEnabled interface{} `field:"optional" json:"publicDataEndpointEnabled" yaml:"publicDataEndpointEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/mssql_managed_instance#storage_account_type MssqlManagedInstance#storage_account_type}.
	StorageAccountType *string `field:"optional" json:"storageAccountType" yaml:"storageAccountType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/mssql_managed_instance#tags MssqlManagedInstance#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/mssql_managed_instance#timeouts MssqlManagedInstance#timeouts}
	Timeouts *MssqlManagedInstanceTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/mssql_managed_instance#timezone_id MssqlManagedInstance#timezone_id}.
	TimezoneId *string `field:"optional" json:"timezoneId" yaml:"timezoneId"`
}

