// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kubernetescluster

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type KubernetesClusterConfig struct {
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
	// default_node_pool block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.84.0/docs/resources/kubernetes_cluster#default_node_pool KubernetesCluster#default_node_pool}
	DefaultNodePool *KubernetesClusterDefaultNodePool `field:"required" json:"defaultNodePool" yaml:"defaultNodePool"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.84.0/docs/resources/kubernetes_cluster#location KubernetesCluster#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.84.0/docs/resources/kubernetes_cluster#name KubernetesCluster#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.84.0/docs/resources/kubernetes_cluster#resource_group_name KubernetesCluster#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// aci_connector_linux block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.84.0/docs/resources/kubernetes_cluster#aci_connector_linux KubernetesCluster#aci_connector_linux}
	AciConnectorLinux *KubernetesClusterAciConnectorLinux `field:"optional" json:"aciConnectorLinux" yaml:"aciConnectorLinux"`
	// api_server_access_profile block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.84.0/docs/resources/kubernetes_cluster#api_server_access_profile KubernetesCluster#api_server_access_profile}
	ApiServerAccessProfile *KubernetesClusterApiServerAccessProfile `field:"optional" json:"apiServerAccessProfile" yaml:"apiServerAccessProfile"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.84.0/docs/resources/kubernetes_cluster#api_server_authorized_ip_ranges KubernetesCluster#api_server_authorized_ip_ranges}.
	ApiServerAuthorizedIpRanges *[]*string `field:"optional" json:"apiServerAuthorizedIpRanges" yaml:"apiServerAuthorizedIpRanges"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.84.0/docs/resources/kubernetes_cluster#automatic_channel_upgrade KubernetesCluster#automatic_channel_upgrade}.
	AutomaticChannelUpgrade *string `field:"optional" json:"automaticChannelUpgrade" yaml:"automaticChannelUpgrade"`
	// auto_scaler_profile block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.84.0/docs/resources/kubernetes_cluster#auto_scaler_profile KubernetesCluster#auto_scaler_profile}
	AutoScalerProfile *KubernetesClusterAutoScalerProfile `field:"optional" json:"autoScalerProfile" yaml:"autoScalerProfile"`
	// azure_active_directory_role_based_access_control block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.84.0/docs/resources/kubernetes_cluster#azure_active_directory_role_based_access_control KubernetesCluster#azure_active_directory_role_based_access_control}
	AzureActiveDirectoryRoleBasedAccessControl *KubernetesClusterAzureActiveDirectoryRoleBasedAccessControl `field:"optional" json:"azureActiveDirectoryRoleBasedAccessControl" yaml:"azureActiveDirectoryRoleBasedAccessControl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.84.0/docs/resources/kubernetes_cluster#azure_policy_enabled KubernetesCluster#azure_policy_enabled}.
	AzurePolicyEnabled interface{} `field:"optional" json:"azurePolicyEnabled" yaml:"azurePolicyEnabled"`
	// confidential_computing block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.84.0/docs/resources/kubernetes_cluster#confidential_computing KubernetesCluster#confidential_computing}
	ConfidentialComputing *KubernetesClusterConfidentialComputing `field:"optional" json:"confidentialComputing" yaml:"confidentialComputing"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.84.0/docs/resources/kubernetes_cluster#custom_ca_trust_certificates_base64 KubernetesCluster#custom_ca_trust_certificates_base64}.
	CustomCaTrustCertificatesBase64 *[]*string `field:"optional" json:"customCaTrustCertificatesBase64" yaml:"customCaTrustCertificatesBase64"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.84.0/docs/resources/kubernetes_cluster#disk_encryption_set_id KubernetesCluster#disk_encryption_set_id}.
	DiskEncryptionSetId *string `field:"optional" json:"diskEncryptionSetId" yaml:"diskEncryptionSetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.84.0/docs/resources/kubernetes_cluster#dns_prefix KubernetesCluster#dns_prefix}.
	DnsPrefix *string `field:"optional" json:"dnsPrefix" yaml:"dnsPrefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.84.0/docs/resources/kubernetes_cluster#dns_prefix_private_cluster KubernetesCluster#dns_prefix_private_cluster}.
	DnsPrefixPrivateCluster *string `field:"optional" json:"dnsPrefixPrivateCluster" yaml:"dnsPrefixPrivateCluster"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.84.0/docs/resources/kubernetes_cluster#edge_zone KubernetesCluster#edge_zone}.
	EdgeZone *string `field:"optional" json:"edgeZone" yaml:"edgeZone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.84.0/docs/resources/kubernetes_cluster#enable_pod_security_policy KubernetesCluster#enable_pod_security_policy}.
	EnablePodSecurityPolicy interface{} `field:"optional" json:"enablePodSecurityPolicy" yaml:"enablePodSecurityPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.84.0/docs/resources/kubernetes_cluster#http_application_routing_enabled KubernetesCluster#http_application_routing_enabled}.
	HttpApplicationRoutingEnabled interface{} `field:"optional" json:"httpApplicationRoutingEnabled" yaml:"httpApplicationRoutingEnabled"`
	// http_proxy_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.84.0/docs/resources/kubernetes_cluster#http_proxy_config KubernetesCluster#http_proxy_config}
	HttpProxyConfig *KubernetesClusterHttpProxyConfig `field:"optional" json:"httpProxyConfig" yaml:"httpProxyConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.84.0/docs/resources/kubernetes_cluster#id KubernetesCluster#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// identity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.84.0/docs/resources/kubernetes_cluster#identity KubernetesCluster#identity}
	Identity *KubernetesClusterIdentity `field:"optional" json:"identity" yaml:"identity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.84.0/docs/resources/kubernetes_cluster#image_cleaner_enabled KubernetesCluster#image_cleaner_enabled}.
	ImageCleanerEnabled interface{} `field:"optional" json:"imageCleanerEnabled" yaml:"imageCleanerEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.84.0/docs/resources/kubernetes_cluster#image_cleaner_interval_hours KubernetesCluster#image_cleaner_interval_hours}.
	ImageCleanerIntervalHours *float64 `field:"optional" json:"imageCleanerIntervalHours" yaml:"imageCleanerIntervalHours"`
	// ingress_application_gateway block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.84.0/docs/resources/kubernetes_cluster#ingress_application_gateway KubernetesCluster#ingress_application_gateway}
	IngressApplicationGateway *KubernetesClusterIngressApplicationGateway `field:"optional" json:"ingressApplicationGateway" yaml:"ingressApplicationGateway"`
	// key_management_service block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.84.0/docs/resources/kubernetes_cluster#key_management_service KubernetesCluster#key_management_service}
	KeyManagementService *KubernetesClusterKeyManagementService `field:"optional" json:"keyManagementService" yaml:"keyManagementService"`
	// key_vault_secrets_provider block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.84.0/docs/resources/kubernetes_cluster#key_vault_secrets_provider KubernetesCluster#key_vault_secrets_provider}
	KeyVaultSecretsProvider *KubernetesClusterKeyVaultSecretsProvider `field:"optional" json:"keyVaultSecretsProvider" yaml:"keyVaultSecretsProvider"`
	// kubelet_identity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.84.0/docs/resources/kubernetes_cluster#kubelet_identity KubernetesCluster#kubelet_identity}
	KubeletIdentity *KubernetesClusterKubeletIdentity `field:"optional" json:"kubeletIdentity" yaml:"kubeletIdentity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.84.0/docs/resources/kubernetes_cluster#kubernetes_version KubernetesCluster#kubernetes_version}.
	KubernetesVersion *string `field:"optional" json:"kubernetesVersion" yaml:"kubernetesVersion"`
	// linux_profile block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.84.0/docs/resources/kubernetes_cluster#linux_profile KubernetesCluster#linux_profile}
	LinuxProfile *KubernetesClusterLinuxProfile `field:"optional" json:"linuxProfile" yaml:"linuxProfile"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.84.0/docs/resources/kubernetes_cluster#local_account_disabled KubernetesCluster#local_account_disabled}.
	LocalAccountDisabled interface{} `field:"optional" json:"localAccountDisabled" yaml:"localAccountDisabled"`
	// maintenance_window block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.84.0/docs/resources/kubernetes_cluster#maintenance_window KubernetesCluster#maintenance_window}
	MaintenanceWindow *KubernetesClusterMaintenanceWindow `field:"optional" json:"maintenanceWindow" yaml:"maintenanceWindow"`
	// maintenance_window_auto_upgrade block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.84.0/docs/resources/kubernetes_cluster#maintenance_window_auto_upgrade KubernetesCluster#maintenance_window_auto_upgrade}
	MaintenanceWindowAutoUpgrade *KubernetesClusterMaintenanceWindowAutoUpgrade `field:"optional" json:"maintenanceWindowAutoUpgrade" yaml:"maintenanceWindowAutoUpgrade"`
	// maintenance_window_node_os block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.84.0/docs/resources/kubernetes_cluster#maintenance_window_node_os KubernetesCluster#maintenance_window_node_os}
	MaintenanceWindowNodeOs *KubernetesClusterMaintenanceWindowNodeOs `field:"optional" json:"maintenanceWindowNodeOs" yaml:"maintenanceWindowNodeOs"`
	// microsoft_defender block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.84.0/docs/resources/kubernetes_cluster#microsoft_defender KubernetesCluster#microsoft_defender}
	MicrosoftDefender *KubernetesClusterMicrosoftDefender `field:"optional" json:"microsoftDefender" yaml:"microsoftDefender"`
	// monitor_metrics block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.84.0/docs/resources/kubernetes_cluster#monitor_metrics KubernetesCluster#monitor_metrics}
	MonitorMetrics *KubernetesClusterMonitorMetrics `field:"optional" json:"monitorMetrics" yaml:"monitorMetrics"`
	// network_profile block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.84.0/docs/resources/kubernetes_cluster#network_profile KubernetesCluster#network_profile}
	NetworkProfile *KubernetesClusterNetworkProfile `field:"optional" json:"networkProfile" yaml:"networkProfile"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.84.0/docs/resources/kubernetes_cluster#node_os_channel_upgrade KubernetesCluster#node_os_channel_upgrade}.
	NodeOsChannelUpgrade *string `field:"optional" json:"nodeOsChannelUpgrade" yaml:"nodeOsChannelUpgrade"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.84.0/docs/resources/kubernetes_cluster#node_resource_group KubernetesCluster#node_resource_group}.
	NodeResourceGroup *string `field:"optional" json:"nodeResourceGroup" yaml:"nodeResourceGroup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.84.0/docs/resources/kubernetes_cluster#oidc_issuer_enabled KubernetesCluster#oidc_issuer_enabled}.
	OidcIssuerEnabled interface{} `field:"optional" json:"oidcIssuerEnabled" yaml:"oidcIssuerEnabled"`
	// oms_agent block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.84.0/docs/resources/kubernetes_cluster#oms_agent KubernetesCluster#oms_agent}
	OmsAgent *KubernetesClusterOmsAgent `field:"optional" json:"omsAgent" yaml:"omsAgent"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.84.0/docs/resources/kubernetes_cluster#open_service_mesh_enabled KubernetesCluster#open_service_mesh_enabled}.
	OpenServiceMeshEnabled interface{} `field:"optional" json:"openServiceMeshEnabled" yaml:"openServiceMeshEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.84.0/docs/resources/kubernetes_cluster#private_cluster_enabled KubernetesCluster#private_cluster_enabled}.
	PrivateClusterEnabled interface{} `field:"optional" json:"privateClusterEnabled" yaml:"privateClusterEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.84.0/docs/resources/kubernetes_cluster#private_cluster_public_fqdn_enabled KubernetesCluster#private_cluster_public_fqdn_enabled}.
	PrivateClusterPublicFqdnEnabled interface{} `field:"optional" json:"privateClusterPublicFqdnEnabled" yaml:"privateClusterPublicFqdnEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.84.0/docs/resources/kubernetes_cluster#private_dns_zone_id KubernetesCluster#private_dns_zone_id}.
	PrivateDnsZoneId *string `field:"optional" json:"privateDnsZoneId" yaml:"privateDnsZoneId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.84.0/docs/resources/kubernetes_cluster#public_network_access_enabled KubernetesCluster#public_network_access_enabled}.
	PublicNetworkAccessEnabled interface{} `field:"optional" json:"publicNetworkAccessEnabled" yaml:"publicNetworkAccessEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.84.0/docs/resources/kubernetes_cluster#role_based_access_control_enabled KubernetesCluster#role_based_access_control_enabled}.
	RoleBasedAccessControlEnabled interface{} `field:"optional" json:"roleBasedAccessControlEnabled" yaml:"roleBasedAccessControlEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.84.0/docs/resources/kubernetes_cluster#run_command_enabled KubernetesCluster#run_command_enabled}.
	RunCommandEnabled interface{} `field:"optional" json:"runCommandEnabled" yaml:"runCommandEnabled"`
	// service_mesh_profile block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.84.0/docs/resources/kubernetes_cluster#service_mesh_profile KubernetesCluster#service_mesh_profile}
	ServiceMeshProfile *KubernetesClusterServiceMeshProfile `field:"optional" json:"serviceMeshProfile" yaml:"serviceMeshProfile"`
	// service_principal block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.84.0/docs/resources/kubernetes_cluster#service_principal KubernetesCluster#service_principal}
	ServicePrincipal *KubernetesClusterServicePrincipal `field:"optional" json:"servicePrincipal" yaml:"servicePrincipal"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.84.0/docs/resources/kubernetes_cluster#sku_tier KubernetesCluster#sku_tier}.
	SkuTier *string `field:"optional" json:"skuTier" yaml:"skuTier"`
	// storage_profile block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.84.0/docs/resources/kubernetes_cluster#storage_profile KubernetesCluster#storage_profile}
	StorageProfile *KubernetesClusterStorageProfile `field:"optional" json:"storageProfile" yaml:"storageProfile"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.84.0/docs/resources/kubernetes_cluster#support_plan KubernetesCluster#support_plan}.
	SupportPlan *string `field:"optional" json:"supportPlan" yaml:"supportPlan"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.84.0/docs/resources/kubernetes_cluster#tags KubernetesCluster#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.84.0/docs/resources/kubernetes_cluster#timeouts KubernetesCluster#timeouts}
	Timeouts *KubernetesClusterTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// web_app_routing block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.84.0/docs/resources/kubernetes_cluster#web_app_routing KubernetesCluster#web_app_routing}
	WebAppRouting *KubernetesClusterWebAppRouting `field:"optional" json:"webAppRouting" yaml:"webAppRouting"`
	// windows_profile block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.84.0/docs/resources/kubernetes_cluster#windows_profile KubernetesCluster#windows_profile}
	WindowsProfile *KubernetesClusterWindowsProfile `field:"optional" json:"windowsProfile" yaml:"windowsProfile"`
	// workload_autoscaler_profile block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.84.0/docs/resources/kubernetes_cluster#workload_autoscaler_profile KubernetesCluster#workload_autoscaler_profile}
	WorkloadAutoscalerProfile *KubernetesClusterWorkloadAutoscalerProfile `field:"optional" json:"workloadAutoscalerProfile" yaml:"workloadAutoscalerProfile"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.84.0/docs/resources/kubernetes_cluster#workload_identity_enabled KubernetesCluster#workload_identity_enabled}.
	WorkloadIdentityEnabled interface{} `field:"optional" json:"workloadIdentityEnabled" yaml:"workloadIdentityEnabled"`
}

