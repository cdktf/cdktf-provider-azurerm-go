// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package signalrservice

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SignalrServiceConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/signalr_service#location SignalrService#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/signalr_service#name SignalrService#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/signalr_service#resource_group_name SignalrService#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// sku block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/signalr_service#sku SignalrService#sku}
	Sku *SignalrServiceSku `field:"required" json:"sku" yaml:"sku"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/signalr_service#aad_auth_enabled SignalrService#aad_auth_enabled}.
	AadAuthEnabled interface{} `field:"optional" json:"aadAuthEnabled" yaml:"aadAuthEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/signalr_service#connectivity_logs_enabled SignalrService#connectivity_logs_enabled}.
	ConnectivityLogsEnabled interface{} `field:"optional" json:"connectivityLogsEnabled" yaml:"connectivityLogsEnabled"`
	// cors block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/signalr_service#cors SignalrService#cors}
	Cors interface{} `field:"optional" json:"cors" yaml:"cors"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/signalr_service#http_request_logs_enabled SignalrService#http_request_logs_enabled}.
	HttpRequestLogsEnabled interface{} `field:"optional" json:"httpRequestLogsEnabled" yaml:"httpRequestLogsEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/signalr_service#id SignalrService#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// identity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/signalr_service#identity SignalrService#identity}
	Identity *SignalrServiceIdentity `field:"optional" json:"identity" yaml:"identity"`
	// live_trace block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/signalr_service#live_trace SignalrService#live_trace}
	LiveTrace *SignalrServiceLiveTrace `field:"optional" json:"liveTrace" yaml:"liveTrace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/signalr_service#live_trace_enabled SignalrService#live_trace_enabled}.
	LiveTraceEnabled interface{} `field:"optional" json:"liveTraceEnabled" yaml:"liveTraceEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/signalr_service#local_auth_enabled SignalrService#local_auth_enabled}.
	LocalAuthEnabled interface{} `field:"optional" json:"localAuthEnabled" yaml:"localAuthEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/signalr_service#messaging_logs_enabled SignalrService#messaging_logs_enabled}.
	MessagingLogsEnabled interface{} `field:"optional" json:"messagingLogsEnabled" yaml:"messagingLogsEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/signalr_service#public_network_access_enabled SignalrService#public_network_access_enabled}.
	PublicNetworkAccessEnabled interface{} `field:"optional" json:"publicNetworkAccessEnabled" yaml:"publicNetworkAccessEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/signalr_service#serverless_connection_timeout_in_seconds SignalrService#serverless_connection_timeout_in_seconds}.
	ServerlessConnectionTimeoutInSeconds *float64 `field:"optional" json:"serverlessConnectionTimeoutInSeconds" yaml:"serverlessConnectionTimeoutInSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/signalr_service#service_mode SignalrService#service_mode}.
	ServiceMode *string `field:"optional" json:"serviceMode" yaml:"serviceMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/signalr_service#tags SignalrService#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/signalr_service#timeouts SignalrService#timeouts}
	Timeouts *SignalrServiceTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/signalr_service#tls_client_cert_enabled SignalrService#tls_client_cert_enabled}.
	TlsClientCertEnabled interface{} `field:"optional" json:"tlsClientCertEnabled" yaml:"tlsClientCertEnabled"`
	// upstream_endpoint block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/signalr_service#upstream_endpoint SignalrService#upstream_endpoint}
	UpstreamEndpoint interface{} `field:"optional" json:"upstreamEndpoint" yaml:"upstreamEndpoint"`
}

