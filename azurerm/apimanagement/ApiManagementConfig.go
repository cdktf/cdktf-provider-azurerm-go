// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apimanagement

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApiManagementConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.49.0/docs/resources/api_management#location ApiManagement#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.49.0/docs/resources/api_management#name ApiManagement#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.49.0/docs/resources/api_management#publisher_email ApiManagement#publisher_email}.
	PublisherEmail *string `field:"required" json:"publisherEmail" yaml:"publisherEmail"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.49.0/docs/resources/api_management#publisher_name ApiManagement#publisher_name}.
	PublisherName *string `field:"required" json:"publisherName" yaml:"publisherName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.49.0/docs/resources/api_management#resource_group_name ApiManagement#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.49.0/docs/resources/api_management#sku_name ApiManagement#sku_name}.
	SkuName *string `field:"required" json:"skuName" yaml:"skuName"`
	// additional_location block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.49.0/docs/resources/api_management#additional_location ApiManagement#additional_location}
	AdditionalLocation interface{} `field:"optional" json:"additionalLocation" yaml:"additionalLocation"`
	// certificate block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.49.0/docs/resources/api_management#certificate ApiManagement#certificate}
	Certificate interface{} `field:"optional" json:"certificate" yaml:"certificate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.49.0/docs/resources/api_management#client_certificate_enabled ApiManagement#client_certificate_enabled}.
	ClientCertificateEnabled interface{} `field:"optional" json:"clientCertificateEnabled" yaml:"clientCertificateEnabled"`
	// delegation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.49.0/docs/resources/api_management#delegation ApiManagement#delegation}
	Delegation *ApiManagementDelegation `field:"optional" json:"delegation" yaml:"delegation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.49.0/docs/resources/api_management#gateway_disabled ApiManagement#gateway_disabled}.
	GatewayDisabled interface{} `field:"optional" json:"gatewayDisabled" yaml:"gatewayDisabled"`
	// hostname_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.49.0/docs/resources/api_management#hostname_configuration ApiManagement#hostname_configuration}
	HostnameConfiguration *ApiManagementHostnameConfiguration `field:"optional" json:"hostnameConfiguration" yaml:"hostnameConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.49.0/docs/resources/api_management#id ApiManagement#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// identity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.49.0/docs/resources/api_management#identity ApiManagement#identity}
	Identity *ApiManagementIdentity `field:"optional" json:"identity" yaml:"identity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.49.0/docs/resources/api_management#min_api_version ApiManagement#min_api_version}.
	MinApiVersion *string `field:"optional" json:"minApiVersion" yaml:"minApiVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.49.0/docs/resources/api_management#notification_sender_email ApiManagement#notification_sender_email}.
	NotificationSenderEmail *string `field:"optional" json:"notificationSenderEmail" yaml:"notificationSenderEmail"`
	// protocols block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.49.0/docs/resources/api_management#protocols ApiManagement#protocols}
	Protocols *ApiManagementProtocols `field:"optional" json:"protocols" yaml:"protocols"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.49.0/docs/resources/api_management#public_ip_address_id ApiManagement#public_ip_address_id}.
	PublicIpAddressId *string `field:"optional" json:"publicIpAddressId" yaml:"publicIpAddressId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.49.0/docs/resources/api_management#public_network_access_enabled ApiManagement#public_network_access_enabled}.
	PublicNetworkAccessEnabled interface{} `field:"optional" json:"publicNetworkAccessEnabled" yaml:"publicNetworkAccessEnabled"`
	// security block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.49.0/docs/resources/api_management#security ApiManagement#security}
	Security *ApiManagementSecurity `field:"optional" json:"security" yaml:"security"`
	// sign_in block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.49.0/docs/resources/api_management#sign_in ApiManagement#sign_in}
	SignIn *ApiManagementSignIn `field:"optional" json:"signIn" yaml:"signIn"`
	// sign_up block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.49.0/docs/resources/api_management#sign_up ApiManagement#sign_up}
	SignUp *ApiManagementSignUp `field:"optional" json:"signUp" yaml:"signUp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.49.0/docs/resources/api_management#tags ApiManagement#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// tenant_access block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.49.0/docs/resources/api_management#tenant_access ApiManagement#tenant_access}
	TenantAccess *ApiManagementTenantAccess `field:"optional" json:"tenantAccess" yaml:"tenantAccess"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.49.0/docs/resources/api_management#timeouts ApiManagement#timeouts}
	Timeouts *ApiManagementTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// virtual_network_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.49.0/docs/resources/api_management#virtual_network_configuration ApiManagement#virtual_network_configuration}
	VirtualNetworkConfiguration *ApiManagementVirtualNetworkConfiguration `field:"optional" json:"virtualNetworkConfiguration" yaml:"virtualNetworkConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.49.0/docs/resources/api_management#virtual_network_type ApiManagement#virtual_network_type}.
	VirtualNetworkType *string `field:"optional" json:"virtualNetworkType" yaml:"virtualNetworkType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.49.0/docs/resources/api_management#zones ApiManagement#zones}.
	Zones *[]*string `field:"optional" json:"zones" yaml:"zones"`
}

