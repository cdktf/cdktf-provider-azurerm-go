package kubernetescluster

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type KubernetesClusterConfig struct {
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
	// default_node_pool block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#default_node_pool KubernetesCluster#default_node_pool}
	DefaultNodePool *KubernetesClusterDefaultNodePool `field:"required" json:"defaultNodePool" yaml:"defaultNodePool"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#location KubernetesCluster#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#name KubernetesCluster#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#resource_group_name KubernetesCluster#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// aci_connector_linux block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#aci_connector_linux KubernetesCluster#aci_connector_linux}
	AciConnectorLinux *KubernetesClusterAciConnectorLinux `field:"optional" json:"aciConnectorLinux" yaml:"aciConnectorLinux"`
	// api_server_access_profile block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#api_server_access_profile KubernetesCluster#api_server_access_profile}
	ApiServerAccessProfile *KubernetesClusterApiServerAccessProfile `field:"optional" json:"apiServerAccessProfile" yaml:"apiServerAccessProfile"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#api_server_authorized_ip_ranges KubernetesCluster#api_server_authorized_ip_ranges}.
	ApiServerAuthorizedIpRanges *[]*string `field:"optional" json:"apiServerAuthorizedIpRanges" yaml:"apiServerAuthorizedIpRanges"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#automatic_channel_upgrade KubernetesCluster#automatic_channel_upgrade}.
	AutomaticChannelUpgrade *string `field:"optional" json:"automaticChannelUpgrade" yaml:"automaticChannelUpgrade"`
	// auto_scaler_profile block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#auto_scaler_profile KubernetesCluster#auto_scaler_profile}
	AutoScalerProfile *KubernetesClusterAutoScalerProfile `field:"optional" json:"autoScalerProfile" yaml:"autoScalerProfile"`
	// azure_active_directory_role_based_access_control block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#azure_active_directory_role_based_access_control KubernetesCluster#azure_active_directory_role_based_access_control}
	AzureActiveDirectoryRoleBasedAccessControl *KubernetesClusterAzureActiveDirectoryRoleBasedAccessControl `field:"optional" json:"azureActiveDirectoryRoleBasedAccessControl" yaml:"azureActiveDirectoryRoleBasedAccessControl"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#azure_policy_enabled KubernetesCluster#azure_policy_enabled}.
	AzurePolicyEnabled interface{} `field:"optional" json:"azurePolicyEnabled" yaml:"azurePolicyEnabled"`
	// confidential_computing block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#confidential_computing KubernetesCluster#confidential_computing}
	ConfidentialComputing *KubernetesClusterConfidentialComputing `field:"optional" json:"confidentialComputing" yaml:"confidentialComputing"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#disk_encryption_set_id KubernetesCluster#disk_encryption_set_id}.
	DiskEncryptionSetId *string `field:"optional" json:"diskEncryptionSetId" yaml:"diskEncryptionSetId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#dns_prefix KubernetesCluster#dns_prefix}.
	DnsPrefix *string `field:"optional" json:"dnsPrefix" yaml:"dnsPrefix"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#dns_prefix_private_cluster KubernetesCluster#dns_prefix_private_cluster}.
	DnsPrefixPrivateCluster *string `field:"optional" json:"dnsPrefixPrivateCluster" yaml:"dnsPrefixPrivateCluster"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#edge_zone KubernetesCluster#edge_zone}.
	EdgeZone *string `field:"optional" json:"edgeZone" yaml:"edgeZone"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#enable_pod_security_policy KubernetesCluster#enable_pod_security_policy}.
	EnablePodSecurityPolicy interface{} `field:"optional" json:"enablePodSecurityPolicy" yaml:"enablePodSecurityPolicy"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#http_application_routing_enabled KubernetesCluster#http_application_routing_enabled}.
	HttpApplicationRoutingEnabled interface{} `field:"optional" json:"httpApplicationRoutingEnabled" yaml:"httpApplicationRoutingEnabled"`
	// http_proxy_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#http_proxy_config KubernetesCluster#http_proxy_config}
	HttpProxyConfig *KubernetesClusterHttpProxyConfig `field:"optional" json:"httpProxyConfig" yaml:"httpProxyConfig"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#id KubernetesCluster#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// identity block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#identity KubernetesCluster#identity}
	Identity *KubernetesClusterIdentity `field:"optional" json:"identity" yaml:"identity"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#image_cleaner_enabled KubernetesCluster#image_cleaner_enabled}.
	ImageCleanerEnabled interface{} `field:"optional" json:"imageCleanerEnabled" yaml:"imageCleanerEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#image_cleaner_interval_hours KubernetesCluster#image_cleaner_interval_hours}.
	ImageCleanerIntervalHours *float64 `field:"optional" json:"imageCleanerIntervalHours" yaml:"imageCleanerIntervalHours"`
	// ingress_application_gateway block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#ingress_application_gateway KubernetesCluster#ingress_application_gateway}
	IngressApplicationGateway *KubernetesClusterIngressApplicationGateway `field:"optional" json:"ingressApplicationGateway" yaml:"ingressApplicationGateway"`
	// key_management_service block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#key_management_service KubernetesCluster#key_management_service}
	KeyManagementService *KubernetesClusterKeyManagementService `field:"optional" json:"keyManagementService" yaml:"keyManagementService"`
	// key_vault_secrets_provider block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#key_vault_secrets_provider KubernetesCluster#key_vault_secrets_provider}
	KeyVaultSecretsProvider *KubernetesClusterKeyVaultSecretsProvider `field:"optional" json:"keyVaultSecretsProvider" yaml:"keyVaultSecretsProvider"`
	// kubelet_identity block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#kubelet_identity KubernetesCluster#kubelet_identity}
	KubeletIdentity *KubernetesClusterKubeletIdentity `field:"optional" json:"kubeletIdentity" yaml:"kubeletIdentity"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#kubernetes_version KubernetesCluster#kubernetes_version}.
	KubernetesVersion *string `field:"optional" json:"kubernetesVersion" yaml:"kubernetesVersion"`
	// linux_profile block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#linux_profile KubernetesCluster#linux_profile}
	LinuxProfile *KubernetesClusterLinuxProfile `field:"optional" json:"linuxProfile" yaml:"linuxProfile"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#local_account_disabled KubernetesCluster#local_account_disabled}.
	LocalAccountDisabled interface{} `field:"optional" json:"localAccountDisabled" yaml:"localAccountDisabled"`
	// maintenance_window block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#maintenance_window KubernetesCluster#maintenance_window}
	MaintenanceWindow *KubernetesClusterMaintenanceWindow `field:"optional" json:"maintenanceWindow" yaml:"maintenanceWindow"`
	// microsoft_defender block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#microsoft_defender KubernetesCluster#microsoft_defender}
	MicrosoftDefender *KubernetesClusterMicrosoftDefender `field:"optional" json:"microsoftDefender" yaml:"microsoftDefender"`
	// monitor_metrics block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#monitor_metrics KubernetesCluster#monitor_metrics}
	MonitorMetrics *KubernetesClusterMonitorMetrics `field:"optional" json:"monitorMetrics" yaml:"monitorMetrics"`
	// network_profile block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#network_profile KubernetesCluster#network_profile}
	NetworkProfile *KubernetesClusterNetworkProfile `field:"optional" json:"networkProfile" yaml:"networkProfile"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#node_resource_group KubernetesCluster#node_resource_group}.
	NodeResourceGroup *string `field:"optional" json:"nodeResourceGroup" yaml:"nodeResourceGroup"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#oidc_issuer_enabled KubernetesCluster#oidc_issuer_enabled}.
	OidcIssuerEnabled interface{} `field:"optional" json:"oidcIssuerEnabled" yaml:"oidcIssuerEnabled"`
	// oms_agent block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#oms_agent KubernetesCluster#oms_agent}
	OmsAgent *KubernetesClusterOmsAgent `field:"optional" json:"omsAgent" yaml:"omsAgent"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#open_service_mesh_enabled KubernetesCluster#open_service_mesh_enabled}.
	OpenServiceMeshEnabled interface{} `field:"optional" json:"openServiceMeshEnabled" yaml:"openServiceMeshEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#private_cluster_enabled KubernetesCluster#private_cluster_enabled}.
	PrivateClusterEnabled interface{} `field:"optional" json:"privateClusterEnabled" yaml:"privateClusterEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#private_cluster_public_fqdn_enabled KubernetesCluster#private_cluster_public_fqdn_enabled}.
	PrivateClusterPublicFqdnEnabled interface{} `field:"optional" json:"privateClusterPublicFqdnEnabled" yaml:"privateClusterPublicFqdnEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#private_dns_zone_id KubernetesCluster#private_dns_zone_id}.
	PrivateDnsZoneId *string `field:"optional" json:"privateDnsZoneId" yaml:"privateDnsZoneId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#public_network_access_enabled KubernetesCluster#public_network_access_enabled}.
	PublicNetworkAccessEnabled interface{} `field:"optional" json:"publicNetworkAccessEnabled" yaml:"publicNetworkAccessEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#role_based_access_control_enabled KubernetesCluster#role_based_access_control_enabled}.
	RoleBasedAccessControlEnabled interface{} `field:"optional" json:"roleBasedAccessControlEnabled" yaml:"roleBasedAccessControlEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#run_command_enabled KubernetesCluster#run_command_enabled}.
	RunCommandEnabled interface{} `field:"optional" json:"runCommandEnabled" yaml:"runCommandEnabled"`
	// service_principal block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#service_principal KubernetesCluster#service_principal}
	ServicePrincipal *KubernetesClusterServicePrincipal `field:"optional" json:"servicePrincipal" yaml:"servicePrincipal"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#sku_tier KubernetesCluster#sku_tier}.
	SkuTier *string `field:"optional" json:"skuTier" yaml:"skuTier"`
	// storage_profile block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#storage_profile KubernetesCluster#storage_profile}
	StorageProfile *KubernetesClusterStorageProfile `field:"optional" json:"storageProfile" yaml:"storageProfile"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#tags KubernetesCluster#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#timeouts KubernetesCluster#timeouts}
	Timeouts *KubernetesClusterTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// web_app_routing block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#web_app_routing KubernetesCluster#web_app_routing}
	WebAppRouting *KubernetesClusterWebAppRouting `field:"optional" json:"webAppRouting" yaml:"webAppRouting"`
	// windows_profile block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#windows_profile KubernetesCluster#windows_profile}
	WindowsProfile *KubernetesClusterWindowsProfile `field:"optional" json:"windowsProfile" yaml:"windowsProfile"`
	// workload_autoscaler_profile block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#workload_autoscaler_profile KubernetesCluster#workload_autoscaler_profile}
	WorkloadAutoscalerProfile *KubernetesClusterWorkloadAutoscalerProfile `field:"optional" json:"workloadAutoscalerProfile" yaml:"workloadAutoscalerProfile"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#workload_identity_enabled KubernetesCluster#workload_identity_enabled}.
	WorkloadIdentityEnabled interface{} `field:"optional" json:"workloadIdentityEnabled" yaml:"workloadIdentityEnabled"`
}

