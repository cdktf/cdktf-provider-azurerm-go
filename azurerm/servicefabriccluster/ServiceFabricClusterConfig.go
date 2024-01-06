// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package servicefabriccluster

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ServiceFabricClusterConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.86.0/docs/resources/service_fabric_cluster#location ServiceFabricCluster#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.86.0/docs/resources/service_fabric_cluster#management_endpoint ServiceFabricCluster#management_endpoint}.
	ManagementEndpoint *string `field:"required" json:"managementEndpoint" yaml:"managementEndpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.86.0/docs/resources/service_fabric_cluster#name ServiceFabricCluster#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// node_type block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.86.0/docs/resources/service_fabric_cluster#node_type ServiceFabricCluster#node_type}
	NodeType interface{} `field:"required" json:"nodeType" yaml:"nodeType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.86.0/docs/resources/service_fabric_cluster#reliability_level ServiceFabricCluster#reliability_level}.
	ReliabilityLevel *string `field:"required" json:"reliabilityLevel" yaml:"reliabilityLevel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.86.0/docs/resources/service_fabric_cluster#resource_group_name ServiceFabricCluster#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.86.0/docs/resources/service_fabric_cluster#upgrade_mode ServiceFabricCluster#upgrade_mode}.
	UpgradeMode *string `field:"required" json:"upgradeMode" yaml:"upgradeMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.86.0/docs/resources/service_fabric_cluster#vm_image ServiceFabricCluster#vm_image}.
	VmImage *string `field:"required" json:"vmImage" yaml:"vmImage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.86.0/docs/resources/service_fabric_cluster#add_on_features ServiceFabricCluster#add_on_features}.
	AddOnFeatures *[]*string `field:"optional" json:"addOnFeatures" yaml:"addOnFeatures"`
	// azure_active_directory block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.86.0/docs/resources/service_fabric_cluster#azure_active_directory ServiceFabricCluster#azure_active_directory}
	AzureActiveDirectory *ServiceFabricClusterAzureActiveDirectory `field:"optional" json:"azureActiveDirectory" yaml:"azureActiveDirectory"`
	// certificate block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.86.0/docs/resources/service_fabric_cluster#certificate ServiceFabricCluster#certificate}
	Certificate *ServiceFabricClusterCertificate `field:"optional" json:"certificate" yaml:"certificate"`
	// certificate_common_names block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.86.0/docs/resources/service_fabric_cluster#certificate_common_names ServiceFabricCluster#certificate_common_names}
	CertificateCommonNames *ServiceFabricClusterCertificateCommonNames `field:"optional" json:"certificateCommonNames" yaml:"certificateCommonNames"`
	// client_certificate_common_name block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.86.0/docs/resources/service_fabric_cluster#client_certificate_common_name ServiceFabricCluster#client_certificate_common_name}
	ClientCertificateCommonName interface{} `field:"optional" json:"clientCertificateCommonName" yaml:"clientCertificateCommonName"`
	// client_certificate_thumbprint block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.86.0/docs/resources/service_fabric_cluster#client_certificate_thumbprint ServiceFabricCluster#client_certificate_thumbprint}
	ClientCertificateThumbprint interface{} `field:"optional" json:"clientCertificateThumbprint" yaml:"clientCertificateThumbprint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.86.0/docs/resources/service_fabric_cluster#cluster_code_version ServiceFabricCluster#cluster_code_version}.
	ClusterCodeVersion *string `field:"optional" json:"clusterCodeVersion" yaml:"clusterCodeVersion"`
	// diagnostics_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.86.0/docs/resources/service_fabric_cluster#diagnostics_config ServiceFabricCluster#diagnostics_config}
	DiagnosticsConfig *ServiceFabricClusterDiagnosticsConfig `field:"optional" json:"diagnosticsConfig" yaml:"diagnosticsConfig"`
	// fabric_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.86.0/docs/resources/service_fabric_cluster#fabric_settings ServiceFabricCluster#fabric_settings}
	FabricSettings interface{} `field:"optional" json:"fabricSettings" yaml:"fabricSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.86.0/docs/resources/service_fabric_cluster#id ServiceFabricCluster#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// reverse_proxy_certificate block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.86.0/docs/resources/service_fabric_cluster#reverse_proxy_certificate ServiceFabricCluster#reverse_proxy_certificate}
	ReverseProxyCertificate *ServiceFabricClusterReverseProxyCertificate `field:"optional" json:"reverseProxyCertificate" yaml:"reverseProxyCertificate"`
	// reverse_proxy_certificate_common_names block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.86.0/docs/resources/service_fabric_cluster#reverse_proxy_certificate_common_names ServiceFabricCluster#reverse_proxy_certificate_common_names}
	ReverseProxyCertificateCommonNames *ServiceFabricClusterReverseProxyCertificateCommonNames `field:"optional" json:"reverseProxyCertificateCommonNames" yaml:"reverseProxyCertificateCommonNames"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.86.0/docs/resources/service_fabric_cluster#service_fabric_zonal_upgrade_mode ServiceFabricCluster#service_fabric_zonal_upgrade_mode}.
	ServiceFabricZonalUpgradeMode *string `field:"optional" json:"serviceFabricZonalUpgradeMode" yaml:"serviceFabricZonalUpgradeMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.86.0/docs/resources/service_fabric_cluster#tags ServiceFabricCluster#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.86.0/docs/resources/service_fabric_cluster#timeouts ServiceFabricCluster#timeouts}
	Timeouts *ServiceFabricClusterTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// upgrade_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.86.0/docs/resources/service_fabric_cluster#upgrade_policy ServiceFabricCluster#upgrade_policy}
	UpgradePolicy *ServiceFabricClusterUpgradePolicy `field:"optional" json:"upgradePolicy" yaml:"upgradePolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.86.0/docs/resources/service_fabric_cluster#vmss_zonal_upgrade_mode ServiceFabricCluster#vmss_zonal_upgrade_mode}.
	VmssZonalUpgradeMode *string `field:"optional" json:"vmssZonalUpgradeMode" yaml:"vmssZonalUpgradeMode"`
}

