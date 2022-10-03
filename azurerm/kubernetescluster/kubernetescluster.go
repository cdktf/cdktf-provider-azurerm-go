package kubernetescluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/kubernetescluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster azurerm_kubernetes_cluster}.
type KubernetesCluster interface {
	cdktf.TerraformResource
	AciConnectorLinux() KubernetesClusterAciConnectorLinuxOutputReference
	AciConnectorLinuxInput() *KubernetesClusterAciConnectorLinux
	ApiServerAuthorizedIpRanges() *[]*string
	SetApiServerAuthorizedIpRanges(val *[]*string)
	ApiServerAuthorizedIpRangesInput() *[]*string
	AutomaticChannelUpgrade() *string
	SetAutomaticChannelUpgrade(val *string)
	AutomaticChannelUpgradeInput() *string
	AutoScalerProfile() KubernetesClusterAutoScalerProfileOutputReference
	AutoScalerProfileInput() *KubernetesClusterAutoScalerProfile
	AzureActiveDirectoryRoleBasedAccessControl() KubernetesClusterAzureActiveDirectoryRoleBasedAccessControlOutputReference
	AzureActiveDirectoryRoleBasedAccessControlInput() *KubernetesClusterAzureActiveDirectoryRoleBasedAccessControl
	AzurePolicyEnabled() interface{}
	SetAzurePolicyEnabled(val interface{})
	AzurePolicyEnabledInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() *float64
	// Experimental.
	SetCount(val *float64)
	DefaultNodePool() KubernetesClusterDefaultNodePoolOutputReference
	DefaultNodePoolInput() *KubernetesClusterDefaultNodePool
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DiskEncryptionSetId() *string
	SetDiskEncryptionSetId(val *string)
	DiskEncryptionSetIdInput() *string
	DnsPrefix() *string
	SetDnsPrefix(val *string)
	DnsPrefixInput() *string
	DnsPrefixPrivateCluster() *string
	SetDnsPrefixPrivateCluster(val *string)
	DnsPrefixPrivateClusterInput() *string
	EdgeZone() *string
	SetEdgeZone(val *string)
	EdgeZoneInput() *string
	EnablePodSecurityPolicy() interface{}
	SetEnablePodSecurityPolicy(val interface{})
	EnablePodSecurityPolicyInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	Fqdn() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HttpApplicationRoutingEnabled() interface{}
	SetHttpApplicationRoutingEnabled(val interface{})
	HttpApplicationRoutingEnabledInput() interface{}
	HttpApplicationRoutingZoneName() *string
	HttpProxyConfig() KubernetesClusterHttpProxyConfigOutputReference
	HttpProxyConfigInput() *KubernetesClusterHttpProxyConfig
	Id() *string
	SetId(val *string)
	Identity() KubernetesClusterIdentityOutputReference
	IdentityInput() *KubernetesClusterIdentity
	IdInput() *string
	IngressApplicationGateway() KubernetesClusterIngressApplicationGatewayOutputReference
	IngressApplicationGatewayInput() *KubernetesClusterIngressApplicationGateway
	KeyVaultSecretsProvider() KubernetesClusterKeyVaultSecretsProviderOutputReference
	KeyVaultSecretsProviderInput() *KubernetesClusterKeyVaultSecretsProvider
	KubeAdminConfig() KubernetesClusterKubeAdminConfigList
	KubeAdminConfigRaw() *string
	KubeConfig() KubernetesClusterKubeConfigList
	KubeConfigRaw() *string
	KubeletIdentity() KubernetesClusterKubeletIdentityOutputReference
	KubeletIdentityInput() *KubernetesClusterKubeletIdentity
	KubernetesVersion() *string
	SetKubernetesVersion(val *string)
	KubernetesVersionInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LinuxProfile() KubernetesClusterLinuxProfileOutputReference
	LinuxProfileInput() *KubernetesClusterLinuxProfile
	LocalAccountDisabled() interface{}
	SetLocalAccountDisabled(val interface{})
	LocalAccountDisabledInput() interface{}
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	MaintenanceWindow() KubernetesClusterMaintenanceWindowOutputReference
	MaintenanceWindowInput() *KubernetesClusterMaintenanceWindow
	MicrosoftDefender() KubernetesClusterMicrosoftDefenderOutputReference
	MicrosoftDefenderInput() *KubernetesClusterMicrosoftDefender
	Name() *string
	SetName(val *string)
	NameInput() *string
	NetworkProfile() KubernetesClusterNetworkProfileOutputReference
	NetworkProfileInput() *KubernetesClusterNetworkProfile
	// The tree node.
	Node() constructs.Node
	NodeResourceGroup() *string
	SetNodeResourceGroup(val *string)
	NodeResourceGroupInput() *string
	OidcIssuerEnabled() interface{}
	SetOidcIssuerEnabled(val interface{})
	OidcIssuerEnabledInput() interface{}
	OidcIssuerUrl() *string
	OmsAgent() KubernetesClusterOmsAgentOutputReference
	OmsAgentInput() *KubernetesClusterOmsAgent
	OpenServiceMeshEnabled() interface{}
	SetOpenServiceMeshEnabled(val interface{})
	OpenServiceMeshEnabledInput() interface{}
	PortalFqdn() *string
	PrivateClusterEnabled() interface{}
	SetPrivateClusterEnabled(val interface{})
	PrivateClusterEnabledInput() interface{}
	PrivateClusterPublicFqdnEnabled() interface{}
	SetPrivateClusterPublicFqdnEnabled(val interface{})
	PrivateClusterPublicFqdnEnabledInput() interface{}
	PrivateDnsZoneId() *string
	SetPrivateDnsZoneId(val *string)
	PrivateDnsZoneIdInput() *string
	PrivateFqdn() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	PublicNetworkAccessEnabled() interface{}
	SetPublicNetworkAccessEnabled(val interface{})
	PublicNetworkAccessEnabledInput() interface{}
	// Experimental.
	RawOverrides() interface{}
	ResourceGroupName() *string
	SetResourceGroupName(val *string)
	ResourceGroupNameInput() *string
	RoleBasedAccessControlEnabled() interface{}
	SetRoleBasedAccessControlEnabled(val interface{})
	RoleBasedAccessControlEnabledInput() interface{}
	RunCommandEnabled() interface{}
	SetRunCommandEnabled(val interface{})
	RunCommandEnabledInput() interface{}
	ServicePrincipal() KubernetesClusterServicePrincipalOutputReference
	ServicePrincipalInput() *KubernetesClusterServicePrincipal
	SkuTier() *string
	SetSkuTier(val *string)
	SkuTierInput() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() KubernetesClusterTimeoutsOutputReference
	TimeoutsInput() interface{}
	WindowsProfile() KubernetesClusterWindowsProfileOutputReference
	WindowsProfileInput() *KubernetesClusterWindowsProfile
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutAciConnectorLinux(value *KubernetesClusterAciConnectorLinux)
	PutAutoScalerProfile(value *KubernetesClusterAutoScalerProfile)
	PutAzureActiveDirectoryRoleBasedAccessControl(value *KubernetesClusterAzureActiveDirectoryRoleBasedAccessControl)
	PutDefaultNodePool(value *KubernetesClusterDefaultNodePool)
	PutHttpProxyConfig(value *KubernetesClusterHttpProxyConfig)
	PutIdentity(value *KubernetesClusterIdentity)
	PutIngressApplicationGateway(value *KubernetesClusterIngressApplicationGateway)
	PutKeyVaultSecretsProvider(value *KubernetesClusterKeyVaultSecretsProvider)
	PutKubeletIdentity(value *KubernetesClusterKubeletIdentity)
	PutLinuxProfile(value *KubernetesClusterLinuxProfile)
	PutMaintenanceWindow(value *KubernetesClusterMaintenanceWindow)
	PutMicrosoftDefender(value *KubernetesClusterMicrosoftDefender)
	PutNetworkProfile(value *KubernetesClusterNetworkProfile)
	PutOmsAgent(value *KubernetesClusterOmsAgent)
	PutServicePrincipal(value *KubernetesClusterServicePrincipal)
	PutTimeouts(value *KubernetesClusterTimeouts)
	PutWindowsProfile(value *KubernetesClusterWindowsProfile)
	ResetAciConnectorLinux()
	ResetApiServerAuthorizedIpRanges()
	ResetAutomaticChannelUpgrade()
	ResetAutoScalerProfile()
	ResetAzureActiveDirectoryRoleBasedAccessControl()
	ResetAzurePolicyEnabled()
	ResetDiskEncryptionSetId()
	ResetDnsPrefix()
	ResetDnsPrefixPrivateCluster()
	ResetEdgeZone()
	ResetEnablePodSecurityPolicy()
	ResetHttpApplicationRoutingEnabled()
	ResetHttpProxyConfig()
	ResetId()
	ResetIdentity()
	ResetIngressApplicationGateway()
	ResetKeyVaultSecretsProvider()
	ResetKubeletIdentity()
	ResetKubernetesVersion()
	ResetLinuxProfile()
	ResetLocalAccountDisabled()
	ResetMaintenanceWindow()
	ResetMicrosoftDefender()
	ResetNetworkProfile()
	ResetNodeResourceGroup()
	ResetOidcIssuerEnabled()
	ResetOmsAgent()
	ResetOpenServiceMeshEnabled()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPrivateClusterEnabled()
	ResetPrivateClusterPublicFqdnEnabled()
	ResetPrivateDnsZoneId()
	ResetPublicNetworkAccessEnabled()
	ResetRoleBasedAccessControlEnabled()
	ResetRunCommandEnabled()
	ResetServicePrincipal()
	ResetSkuTier()
	ResetTags()
	ResetTimeouts()
	ResetWindowsProfile()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for KubernetesCluster
type jsiiProxy_KubernetesCluster struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_KubernetesCluster) AciConnectorLinux() KubernetesClusterAciConnectorLinuxOutputReference {
	var returns KubernetesClusterAciConnectorLinuxOutputReference
	_jsii_.Get(
		j,
		"aciConnectorLinux",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) AciConnectorLinuxInput() *KubernetesClusterAciConnectorLinux {
	var returns *KubernetesClusterAciConnectorLinux
	_jsii_.Get(
		j,
		"aciConnectorLinuxInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) ApiServerAuthorizedIpRanges() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"apiServerAuthorizedIpRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) ApiServerAuthorizedIpRangesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"apiServerAuthorizedIpRangesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) AutomaticChannelUpgrade() *string {
	var returns *string
	_jsii_.Get(
		j,
		"automaticChannelUpgrade",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) AutomaticChannelUpgradeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"automaticChannelUpgradeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) AutoScalerProfile() KubernetesClusterAutoScalerProfileOutputReference {
	var returns KubernetesClusterAutoScalerProfileOutputReference
	_jsii_.Get(
		j,
		"autoScalerProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) AutoScalerProfileInput() *KubernetesClusterAutoScalerProfile {
	var returns *KubernetesClusterAutoScalerProfile
	_jsii_.Get(
		j,
		"autoScalerProfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) AzureActiveDirectoryRoleBasedAccessControl() KubernetesClusterAzureActiveDirectoryRoleBasedAccessControlOutputReference {
	var returns KubernetesClusterAzureActiveDirectoryRoleBasedAccessControlOutputReference
	_jsii_.Get(
		j,
		"azureActiveDirectoryRoleBasedAccessControl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) AzureActiveDirectoryRoleBasedAccessControlInput() *KubernetesClusterAzureActiveDirectoryRoleBasedAccessControl {
	var returns *KubernetesClusterAzureActiveDirectoryRoleBasedAccessControl
	_jsii_.Get(
		j,
		"azureActiveDirectoryRoleBasedAccessControlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) AzurePolicyEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"azurePolicyEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) AzurePolicyEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"azurePolicyEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) DefaultNodePool() KubernetesClusterDefaultNodePoolOutputReference {
	var returns KubernetesClusterDefaultNodePoolOutputReference
	_jsii_.Get(
		j,
		"defaultNodePool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) DefaultNodePoolInput() *KubernetesClusterDefaultNodePool {
	var returns *KubernetesClusterDefaultNodePool
	_jsii_.Get(
		j,
		"defaultNodePoolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) DiskEncryptionSetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskEncryptionSetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) DiskEncryptionSetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskEncryptionSetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) DnsPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) DnsPrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) DnsPrefixPrivateCluster() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsPrefixPrivateCluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) DnsPrefixPrivateClusterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsPrefixPrivateClusterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) EdgeZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"edgeZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) EdgeZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"edgeZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) EnablePodSecurityPolicy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePodSecurityPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) EnablePodSecurityPolicyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePodSecurityPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) Fqdn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqdn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) HttpApplicationRoutingEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"httpApplicationRoutingEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) HttpApplicationRoutingEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"httpApplicationRoutingEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) HttpApplicationRoutingZoneName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpApplicationRoutingZoneName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) HttpProxyConfig() KubernetesClusterHttpProxyConfigOutputReference {
	var returns KubernetesClusterHttpProxyConfigOutputReference
	_jsii_.Get(
		j,
		"httpProxyConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) HttpProxyConfigInput() *KubernetesClusterHttpProxyConfig {
	var returns *KubernetesClusterHttpProxyConfig
	_jsii_.Get(
		j,
		"httpProxyConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) Identity() KubernetesClusterIdentityOutputReference {
	var returns KubernetesClusterIdentityOutputReference
	_jsii_.Get(
		j,
		"identity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) IdentityInput() *KubernetesClusterIdentity {
	var returns *KubernetesClusterIdentity
	_jsii_.Get(
		j,
		"identityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) IngressApplicationGateway() KubernetesClusterIngressApplicationGatewayOutputReference {
	var returns KubernetesClusterIngressApplicationGatewayOutputReference
	_jsii_.Get(
		j,
		"ingressApplicationGateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) IngressApplicationGatewayInput() *KubernetesClusterIngressApplicationGateway {
	var returns *KubernetesClusterIngressApplicationGateway
	_jsii_.Get(
		j,
		"ingressApplicationGatewayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) KeyVaultSecretsProvider() KubernetesClusterKeyVaultSecretsProviderOutputReference {
	var returns KubernetesClusterKeyVaultSecretsProviderOutputReference
	_jsii_.Get(
		j,
		"keyVaultSecretsProvider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) KeyVaultSecretsProviderInput() *KubernetesClusterKeyVaultSecretsProvider {
	var returns *KubernetesClusterKeyVaultSecretsProvider
	_jsii_.Get(
		j,
		"keyVaultSecretsProviderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) KubeAdminConfig() KubernetesClusterKubeAdminConfigList {
	var returns KubernetesClusterKubeAdminConfigList
	_jsii_.Get(
		j,
		"kubeAdminConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) KubeAdminConfigRaw() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kubeAdminConfigRaw",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) KubeConfig() KubernetesClusterKubeConfigList {
	var returns KubernetesClusterKubeConfigList
	_jsii_.Get(
		j,
		"kubeConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) KubeConfigRaw() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kubeConfigRaw",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) KubeletIdentity() KubernetesClusterKubeletIdentityOutputReference {
	var returns KubernetesClusterKubeletIdentityOutputReference
	_jsii_.Get(
		j,
		"kubeletIdentity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) KubeletIdentityInput() *KubernetesClusterKubeletIdentity {
	var returns *KubernetesClusterKubeletIdentity
	_jsii_.Get(
		j,
		"kubeletIdentityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) KubernetesVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kubernetesVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) KubernetesVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kubernetesVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) LinuxProfile() KubernetesClusterLinuxProfileOutputReference {
	var returns KubernetesClusterLinuxProfileOutputReference
	_jsii_.Get(
		j,
		"linuxProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) LinuxProfileInput() *KubernetesClusterLinuxProfile {
	var returns *KubernetesClusterLinuxProfile
	_jsii_.Get(
		j,
		"linuxProfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) LocalAccountDisabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"localAccountDisabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) LocalAccountDisabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"localAccountDisabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) MaintenanceWindow() KubernetesClusterMaintenanceWindowOutputReference {
	var returns KubernetesClusterMaintenanceWindowOutputReference
	_jsii_.Get(
		j,
		"maintenanceWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) MaintenanceWindowInput() *KubernetesClusterMaintenanceWindow {
	var returns *KubernetesClusterMaintenanceWindow
	_jsii_.Get(
		j,
		"maintenanceWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) MicrosoftDefender() KubernetesClusterMicrosoftDefenderOutputReference {
	var returns KubernetesClusterMicrosoftDefenderOutputReference
	_jsii_.Get(
		j,
		"microsoftDefender",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) MicrosoftDefenderInput() *KubernetesClusterMicrosoftDefender {
	var returns *KubernetesClusterMicrosoftDefender
	_jsii_.Get(
		j,
		"microsoftDefenderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) NetworkProfile() KubernetesClusterNetworkProfileOutputReference {
	var returns KubernetesClusterNetworkProfileOutputReference
	_jsii_.Get(
		j,
		"networkProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) NetworkProfileInput() *KubernetesClusterNetworkProfile {
	var returns *KubernetesClusterNetworkProfile
	_jsii_.Get(
		j,
		"networkProfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) NodeResourceGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeResourceGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) NodeResourceGroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeResourceGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) OidcIssuerEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"oidcIssuerEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) OidcIssuerEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"oidcIssuerEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) OidcIssuerUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oidcIssuerUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) OmsAgent() KubernetesClusterOmsAgentOutputReference {
	var returns KubernetesClusterOmsAgentOutputReference
	_jsii_.Get(
		j,
		"omsAgent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) OmsAgentInput() *KubernetesClusterOmsAgent {
	var returns *KubernetesClusterOmsAgent
	_jsii_.Get(
		j,
		"omsAgentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) OpenServiceMeshEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"openServiceMeshEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) OpenServiceMeshEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"openServiceMeshEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) PortalFqdn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"portalFqdn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) PrivateClusterEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"privateClusterEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) PrivateClusterEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"privateClusterEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) PrivateClusterPublicFqdnEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"privateClusterPublicFqdnEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) PrivateClusterPublicFqdnEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"privateClusterPublicFqdnEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) PrivateDnsZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateDnsZoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) PrivateDnsZoneIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateDnsZoneIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) PrivateFqdn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateFqdn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) PublicNetworkAccessEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publicNetworkAccessEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) PublicNetworkAccessEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publicNetworkAccessEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) RoleBasedAccessControlEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"roleBasedAccessControlEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) RoleBasedAccessControlEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"roleBasedAccessControlEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) RunCommandEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runCommandEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) RunCommandEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runCommandEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) ServicePrincipal() KubernetesClusterServicePrincipalOutputReference {
	var returns KubernetesClusterServicePrincipalOutputReference
	_jsii_.Get(
		j,
		"servicePrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) ServicePrincipalInput() *KubernetesClusterServicePrincipal {
	var returns *KubernetesClusterServicePrincipal
	_jsii_.Get(
		j,
		"servicePrincipalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) SkuTier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"skuTier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) SkuTierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"skuTierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) Timeouts() KubernetesClusterTimeoutsOutputReference {
	var returns KubernetesClusterTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) WindowsProfile() KubernetesClusterWindowsProfileOutputReference {
	var returns KubernetesClusterWindowsProfileOutputReference
	_jsii_.Get(
		j,
		"windowsProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) WindowsProfileInput() *KubernetesClusterWindowsProfile {
	var returns *KubernetesClusterWindowsProfile
	_jsii_.Get(
		j,
		"windowsProfileInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster azurerm_kubernetes_cluster} Resource.
func NewKubernetesCluster(scope constructs.Construct, id *string, config *KubernetesClusterConfig) KubernetesCluster {
	_init_.Initialize()

	j := jsiiProxy_KubernetesCluster{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesCluster.KubernetesCluster",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster azurerm_kubernetes_cluster} Resource.
func NewKubernetesCluster_Override(k KubernetesCluster, scope constructs.Construct, id *string, config *KubernetesClusterConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesCluster.KubernetesCluster",
		[]interface{}{scope, id, config},
		k,
	)
}

func (j *jsiiProxy_KubernetesCluster) SetApiServerAuthorizedIpRanges(val *[]*string) {
	_jsii_.Set(
		j,
		"apiServerAuthorizedIpRanges",
		val,
	)
}

func (j *jsiiProxy_KubernetesCluster) SetAutomaticChannelUpgrade(val *string) {
	_jsii_.Set(
		j,
		"automaticChannelUpgrade",
		val,
	)
}

func (j *jsiiProxy_KubernetesCluster) SetAzurePolicyEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"azurePolicyEnabled",
		val,
	)
}

func (j *jsiiProxy_KubernetesCluster) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_KubernetesCluster) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_KubernetesCluster) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_KubernetesCluster) SetDiskEncryptionSetId(val *string) {
	_jsii_.Set(
		j,
		"diskEncryptionSetId",
		val,
	)
}

func (j *jsiiProxy_KubernetesCluster) SetDnsPrefix(val *string) {
	_jsii_.Set(
		j,
		"dnsPrefix",
		val,
	)
}

func (j *jsiiProxy_KubernetesCluster) SetDnsPrefixPrivateCluster(val *string) {
	_jsii_.Set(
		j,
		"dnsPrefixPrivateCluster",
		val,
	)
}

func (j *jsiiProxy_KubernetesCluster) SetEdgeZone(val *string) {
	_jsii_.Set(
		j,
		"edgeZone",
		val,
	)
}

func (j *jsiiProxy_KubernetesCluster) SetEnablePodSecurityPolicy(val interface{}) {
	_jsii_.Set(
		j,
		"enablePodSecurityPolicy",
		val,
	)
}

func (j *jsiiProxy_KubernetesCluster) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_KubernetesCluster) SetHttpApplicationRoutingEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"httpApplicationRoutingEnabled",
		val,
	)
}

func (j *jsiiProxy_KubernetesCluster) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_KubernetesCluster) SetKubernetesVersion(val *string) {
	_jsii_.Set(
		j,
		"kubernetesVersion",
		val,
	)
}

func (j *jsiiProxy_KubernetesCluster) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_KubernetesCluster) SetLocalAccountDisabled(val interface{}) {
	_jsii_.Set(
		j,
		"localAccountDisabled",
		val,
	)
}

func (j *jsiiProxy_KubernetesCluster) SetLocation(val *string) {
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_KubernetesCluster) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_KubernetesCluster) SetNodeResourceGroup(val *string) {
	_jsii_.Set(
		j,
		"nodeResourceGroup",
		val,
	)
}

func (j *jsiiProxy_KubernetesCluster) SetOidcIssuerEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"oidcIssuerEnabled",
		val,
	)
}

func (j *jsiiProxy_KubernetesCluster) SetOpenServiceMeshEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"openServiceMeshEnabled",
		val,
	)
}

func (j *jsiiProxy_KubernetesCluster) SetPrivateClusterEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"privateClusterEnabled",
		val,
	)
}

func (j *jsiiProxy_KubernetesCluster) SetPrivateClusterPublicFqdnEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"privateClusterPublicFqdnEnabled",
		val,
	)
}

func (j *jsiiProxy_KubernetesCluster) SetPrivateDnsZoneId(val *string) {
	_jsii_.Set(
		j,
		"privateDnsZoneId",
		val,
	)
}

func (j *jsiiProxy_KubernetesCluster) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_KubernetesCluster) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_KubernetesCluster) SetPublicNetworkAccessEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"publicNetworkAccessEnabled",
		val,
	)
}

func (j *jsiiProxy_KubernetesCluster) SetResourceGroupName(val *string) {
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_KubernetesCluster) SetRoleBasedAccessControlEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"roleBasedAccessControlEnabled",
		val,
	)
}

func (j *jsiiProxy_KubernetesCluster) SetRunCommandEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"runCommandEnabled",
		val,
	)
}

func (j *jsiiProxy_KubernetesCluster) SetSkuTier(val *string) {
	_jsii_.Set(
		j,
		"skuTier",
		val,
	)
}

func (j *jsiiProxy_KubernetesCluster) SetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func KubernetesCluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.kubernetesCluster.KubernetesCluster",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func KubernetesCluster_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.kubernetesCluster.KubernetesCluster",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (k *jsiiProxy_KubernetesCluster) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		k,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (k *jsiiProxy_KubernetesCluster) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesCluster) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesCluster) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		k,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesCluster) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesCluster) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		k,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesCluster) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		k,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesCluster) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		k,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesCluster) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesCluster) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		k,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesCluster) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesCluster) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		k,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (k *jsiiProxy_KubernetesCluster) PutAciConnectorLinux(value *KubernetesClusterAciConnectorLinux) {
	_jsii_.InvokeVoid(
		k,
		"putAciConnectorLinux",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KubernetesCluster) PutAutoScalerProfile(value *KubernetesClusterAutoScalerProfile) {
	_jsii_.InvokeVoid(
		k,
		"putAutoScalerProfile",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KubernetesCluster) PutAzureActiveDirectoryRoleBasedAccessControl(value *KubernetesClusterAzureActiveDirectoryRoleBasedAccessControl) {
	_jsii_.InvokeVoid(
		k,
		"putAzureActiveDirectoryRoleBasedAccessControl",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KubernetesCluster) PutDefaultNodePool(value *KubernetesClusterDefaultNodePool) {
	_jsii_.InvokeVoid(
		k,
		"putDefaultNodePool",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KubernetesCluster) PutHttpProxyConfig(value *KubernetesClusterHttpProxyConfig) {
	_jsii_.InvokeVoid(
		k,
		"putHttpProxyConfig",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KubernetesCluster) PutIdentity(value *KubernetesClusterIdentity) {
	_jsii_.InvokeVoid(
		k,
		"putIdentity",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KubernetesCluster) PutIngressApplicationGateway(value *KubernetesClusterIngressApplicationGateway) {
	_jsii_.InvokeVoid(
		k,
		"putIngressApplicationGateway",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KubernetesCluster) PutKeyVaultSecretsProvider(value *KubernetesClusterKeyVaultSecretsProvider) {
	_jsii_.InvokeVoid(
		k,
		"putKeyVaultSecretsProvider",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KubernetesCluster) PutKubeletIdentity(value *KubernetesClusterKubeletIdentity) {
	_jsii_.InvokeVoid(
		k,
		"putKubeletIdentity",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KubernetesCluster) PutLinuxProfile(value *KubernetesClusterLinuxProfile) {
	_jsii_.InvokeVoid(
		k,
		"putLinuxProfile",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KubernetesCluster) PutMaintenanceWindow(value *KubernetesClusterMaintenanceWindow) {
	_jsii_.InvokeVoid(
		k,
		"putMaintenanceWindow",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KubernetesCluster) PutMicrosoftDefender(value *KubernetesClusterMicrosoftDefender) {
	_jsii_.InvokeVoid(
		k,
		"putMicrosoftDefender",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KubernetesCluster) PutNetworkProfile(value *KubernetesClusterNetworkProfile) {
	_jsii_.InvokeVoid(
		k,
		"putNetworkProfile",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KubernetesCluster) PutOmsAgent(value *KubernetesClusterOmsAgent) {
	_jsii_.InvokeVoid(
		k,
		"putOmsAgent",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KubernetesCluster) PutServicePrincipal(value *KubernetesClusterServicePrincipal) {
	_jsii_.InvokeVoid(
		k,
		"putServicePrincipal",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KubernetesCluster) PutTimeouts(value *KubernetesClusterTimeouts) {
	_jsii_.InvokeVoid(
		k,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KubernetesCluster) PutWindowsProfile(value *KubernetesClusterWindowsProfile) {
	_jsii_.InvokeVoid(
		k,
		"putWindowsProfile",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KubernetesCluster) ResetAciConnectorLinux() {
	_jsii_.InvokeVoid(
		k,
		"resetAciConnectorLinux",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesCluster) ResetApiServerAuthorizedIpRanges() {
	_jsii_.InvokeVoid(
		k,
		"resetApiServerAuthorizedIpRanges",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesCluster) ResetAutomaticChannelUpgrade() {
	_jsii_.InvokeVoid(
		k,
		"resetAutomaticChannelUpgrade",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesCluster) ResetAutoScalerProfile() {
	_jsii_.InvokeVoid(
		k,
		"resetAutoScalerProfile",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesCluster) ResetAzureActiveDirectoryRoleBasedAccessControl() {
	_jsii_.InvokeVoid(
		k,
		"resetAzureActiveDirectoryRoleBasedAccessControl",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesCluster) ResetAzurePolicyEnabled() {
	_jsii_.InvokeVoid(
		k,
		"resetAzurePolicyEnabled",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesCluster) ResetDiskEncryptionSetId() {
	_jsii_.InvokeVoid(
		k,
		"resetDiskEncryptionSetId",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesCluster) ResetDnsPrefix() {
	_jsii_.InvokeVoid(
		k,
		"resetDnsPrefix",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesCluster) ResetDnsPrefixPrivateCluster() {
	_jsii_.InvokeVoid(
		k,
		"resetDnsPrefixPrivateCluster",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesCluster) ResetEdgeZone() {
	_jsii_.InvokeVoid(
		k,
		"resetEdgeZone",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesCluster) ResetEnablePodSecurityPolicy() {
	_jsii_.InvokeVoid(
		k,
		"resetEnablePodSecurityPolicy",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesCluster) ResetHttpApplicationRoutingEnabled() {
	_jsii_.InvokeVoid(
		k,
		"resetHttpApplicationRoutingEnabled",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesCluster) ResetHttpProxyConfig() {
	_jsii_.InvokeVoid(
		k,
		"resetHttpProxyConfig",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesCluster) ResetId() {
	_jsii_.InvokeVoid(
		k,
		"resetId",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesCluster) ResetIdentity() {
	_jsii_.InvokeVoid(
		k,
		"resetIdentity",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesCluster) ResetIngressApplicationGateway() {
	_jsii_.InvokeVoid(
		k,
		"resetIngressApplicationGateway",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesCluster) ResetKeyVaultSecretsProvider() {
	_jsii_.InvokeVoid(
		k,
		"resetKeyVaultSecretsProvider",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesCluster) ResetKubeletIdentity() {
	_jsii_.InvokeVoid(
		k,
		"resetKubeletIdentity",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesCluster) ResetKubernetesVersion() {
	_jsii_.InvokeVoid(
		k,
		"resetKubernetesVersion",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesCluster) ResetLinuxProfile() {
	_jsii_.InvokeVoid(
		k,
		"resetLinuxProfile",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesCluster) ResetLocalAccountDisabled() {
	_jsii_.InvokeVoid(
		k,
		"resetLocalAccountDisabled",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesCluster) ResetMaintenanceWindow() {
	_jsii_.InvokeVoid(
		k,
		"resetMaintenanceWindow",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesCluster) ResetMicrosoftDefender() {
	_jsii_.InvokeVoid(
		k,
		"resetMicrosoftDefender",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesCluster) ResetNetworkProfile() {
	_jsii_.InvokeVoid(
		k,
		"resetNetworkProfile",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesCluster) ResetNodeResourceGroup() {
	_jsii_.InvokeVoid(
		k,
		"resetNodeResourceGroup",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesCluster) ResetOidcIssuerEnabled() {
	_jsii_.InvokeVoid(
		k,
		"resetOidcIssuerEnabled",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesCluster) ResetOmsAgent() {
	_jsii_.InvokeVoid(
		k,
		"resetOmsAgent",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesCluster) ResetOpenServiceMeshEnabled() {
	_jsii_.InvokeVoid(
		k,
		"resetOpenServiceMeshEnabled",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesCluster) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		k,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesCluster) ResetPrivateClusterEnabled() {
	_jsii_.InvokeVoid(
		k,
		"resetPrivateClusterEnabled",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesCluster) ResetPrivateClusterPublicFqdnEnabled() {
	_jsii_.InvokeVoid(
		k,
		"resetPrivateClusterPublicFqdnEnabled",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesCluster) ResetPrivateDnsZoneId() {
	_jsii_.InvokeVoid(
		k,
		"resetPrivateDnsZoneId",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesCluster) ResetPublicNetworkAccessEnabled() {
	_jsii_.InvokeVoid(
		k,
		"resetPublicNetworkAccessEnabled",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesCluster) ResetRoleBasedAccessControlEnabled() {
	_jsii_.InvokeVoid(
		k,
		"resetRoleBasedAccessControlEnabled",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesCluster) ResetRunCommandEnabled() {
	_jsii_.InvokeVoid(
		k,
		"resetRunCommandEnabled",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesCluster) ResetServicePrincipal() {
	_jsii_.InvokeVoid(
		k,
		"resetServicePrincipal",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesCluster) ResetSkuTier() {
	_jsii_.InvokeVoid(
		k,
		"resetSkuTier",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesCluster) ResetTags() {
	_jsii_.InvokeVoid(
		k,
		"resetTags",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesCluster) ResetTimeouts() {
	_jsii_.InvokeVoid(
		k,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesCluster) ResetWindowsProfile() {
	_jsii_.InvokeVoid(
		k,
		"resetWindowsProfile",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesCluster) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesCluster) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesCluster) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesCluster) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type KubernetesClusterAciConnectorLinux struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#subnet_name KubernetesCluster#subnet_name}.
	SubnetName *string `field:"required" json:"subnetName" yaml:"subnetName"`
}

type KubernetesClusterAciConnectorLinuxOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *KubernetesClusterAciConnectorLinux
	SetInternalValue(val *KubernetesClusterAciConnectorLinux)
	SubnetName() *string
	SetSubnetName(val *string)
	SubnetNameInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KubernetesClusterAciConnectorLinuxOutputReference
type jsiiProxy_KubernetesClusterAciConnectorLinuxOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_KubernetesClusterAciConnectorLinuxOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAciConnectorLinuxOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAciConnectorLinuxOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAciConnectorLinuxOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAciConnectorLinuxOutputReference) InternalValue() *KubernetesClusterAciConnectorLinux {
	var returns *KubernetesClusterAciConnectorLinux
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAciConnectorLinuxOutputReference) SubnetName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAciConnectorLinuxOutputReference) SubnetNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAciConnectorLinuxOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAciConnectorLinuxOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewKubernetesClusterAciConnectorLinuxOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) KubernetesClusterAciConnectorLinuxOutputReference {
	_init_.Initialize()

	j := jsiiProxy_KubernetesClusterAciConnectorLinuxOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesCluster.KubernetesClusterAciConnectorLinuxOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKubernetesClusterAciConnectorLinuxOutputReference_Override(k KubernetesClusterAciConnectorLinuxOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesCluster.KubernetesClusterAciConnectorLinuxOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_KubernetesClusterAciConnectorLinuxOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterAciConnectorLinuxOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterAciConnectorLinuxOutputReference) SetInternalValue(val *KubernetesClusterAciConnectorLinux) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterAciConnectorLinuxOutputReference) SetSubnetName(val *string) {
	_jsii_.Set(
		j,
		"subnetName",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterAciConnectorLinuxOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterAciConnectorLinuxOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (k *jsiiProxy_KubernetesClusterAciConnectorLinuxOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterAciConnectorLinuxOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterAciConnectorLinuxOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterAciConnectorLinuxOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		k,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterAciConnectorLinuxOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterAciConnectorLinuxOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		k,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterAciConnectorLinuxOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		k,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterAciConnectorLinuxOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		k,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterAciConnectorLinuxOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterAciConnectorLinuxOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		k,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterAciConnectorLinuxOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterAciConnectorLinuxOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterAciConnectorLinuxOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterAciConnectorLinuxOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type KubernetesClusterAutoScalerProfile struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#balance_similar_node_groups KubernetesCluster#balance_similar_node_groups}.
	BalanceSimilarNodeGroups interface{} `field:"optional" json:"balanceSimilarNodeGroups" yaml:"balanceSimilarNodeGroups"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#empty_bulk_delete_max KubernetesCluster#empty_bulk_delete_max}.
	EmptyBulkDeleteMax *string `field:"optional" json:"emptyBulkDeleteMax" yaml:"emptyBulkDeleteMax"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#expander KubernetesCluster#expander}.
	Expander *string `field:"optional" json:"expander" yaml:"expander"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#max_graceful_termination_sec KubernetesCluster#max_graceful_termination_sec}.
	MaxGracefulTerminationSec *string `field:"optional" json:"maxGracefulTerminationSec" yaml:"maxGracefulTerminationSec"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#max_node_provisioning_time KubernetesCluster#max_node_provisioning_time}.
	MaxNodeProvisioningTime *string `field:"optional" json:"maxNodeProvisioningTime" yaml:"maxNodeProvisioningTime"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#max_unready_nodes KubernetesCluster#max_unready_nodes}.
	MaxUnreadyNodes *float64 `field:"optional" json:"maxUnreadyNodes" yaml:"maxUnreadyNodes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#max_unready_percentage KubernetesCluster#max_unready_percentage}.
	MaxUnreadyPercentage *float64 `field:"optional" json:"maxUnreadyPercentage" yaml:"maxUnreadyPercentage"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#new_pod_scale_up_delay KubernetesCluster#new_pod_scale_up_delay}.
	NewPodScaleUpDelay *string `field:"optional" json:"newPodScaleUpDelay" yaml:"newPodScaleUpDelay"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#scale_down_delay_after_add KubernetesCluster#scale_down_delay_after_add}.
	ScaleDownDelayAfterAdd *string `field:"optional" json:"scaleDownDelayAfterAdd" yaml:"scaleDownDelayAfterAdd"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#scale_down_delay_after_delete KubernetesCluster#scale_down_delay_after_delete}.
	ScaleDownDelayAfterDelete *string `field:"optional" json:"scaleDownDelayAfterDelete" yaml:"scaleDownDelayAfterDelete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#scale_down_delay_after_failure KubernetesCluster#scale_down_delay_after_failure}.
	ScaleDownDelayAfterFailure *string `field:"optional" json:"scaleDownDelayAfterFailure" yaml:"scaleDownDelayAfterFailure"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#scale_down_unneeded KubernetesCluster#scale_down_unneeded}.
	ScaleDownUnneeded *string `field:"optional" json:"scaleDownUnneeded" yaml:"scaleDownUnneeded"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#scale_down_unready KubernetesCluster#scale_down_unready}.
	ScaleDownUnready *string `field:"optional" json:"scaleDownUnready" yaml:"scaleDownUnready"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#scale_down_utilization_threshold KubernetesCluster#scale_down_utilization_threshold}.
	ScaleDownUtilizationThreshold *string `field:"optional" json:"scaleDownUtilizationThreshold" yaml:"scaleDownUtilizationThreshold"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#scan_interval KubernetesCluster#scan_interval}.
	ScanInterval *string `field:"optional" json:"scanInterval" yaml:"scanInterval"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#skip_nodes_with_local_storage KubernetesCluster#skip_nodes_with_local_storage}.
	SkipNodesWithLocalStorage interface{} `field:"optional" json:"skipNodesWithLocalStorage" yaml:"skipNodesWithLocalStorage"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#skip_nodes_with_system_pods KubernetesCluster#skip_nodes_with_system_pods}.
	SkipNodesWithSystemPods interface{} `field:"optional" json:"skipNodesWithSystemPods" yaml:"skipNodesWithSystemPods"`
}

type KubernetesClusterAutoScalerProfileOutputReference interface {
	cdktf.ComplexObject
	BalanceSimilarNodeGroups() interface{}
	SetBalanceSimilarNodeGroups(val interface{})
	BalanceSimilarNodeGroupsInput() interface{}
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	EmptyBulkDeleteMax() *string
	SetEmptyBulkDeleteMax(val *string)
	EmptyBulkDeleteMaxInput() *string
	Expander() *string
	SetExpander(val *string)
	ExpanderInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *KubernetesClusterAutoScalerProfile
	SetInternalValue(val *KubernetesClusterAutoScalerProfile)
	MaxGracefulTerminationSec() *string
	SetMaxGracefulTerminationSec(val *string)
	MaxGracefulTerminationSecInput() *string
	MaxNodeProvisioningTime() *string
	SetMaxNodeProvisioningTime(val *string)
	MaxNodeProvisioningTimeInput() *string
	MaxUnreadyNodes() *float64
	SetMaxUnreadyNodes(val *float64)
	MaxUnreadyNodesInput() *float64
	MaxUnreadyPercentage() *float64
	SetMaxUnreadyPercentage(val *float64)
	MaxUnreadyPercentageInput() *float64
	NewPodScaleUpDelay() *string
	SetNewPodScaleUpDelay(val *string)
	NewPodScaleUpDelayInput() *string
	ScaleDownDelayAfterAdd() *string
	SetScaleDownDelayAfterAdd(val *string)
	ScaleDownDelayAfterAddInput() *string
	ScaleDownDelayAfterDelete() *string
	SetScaleDownDelayAfterDelete(val *string)
	ScaleDownDelayAfterDeleteInput() *string
	ScaleDownDelayAfterFailure() *string
	SetScaleDownDelayAfterFailure(val *string)
	ScaleDownDelayAfterFailureInput() *string
	ScaleDownUnneeded() *string
	SetScaleDownUnneeded(val *string)
	ScaleDownUnneededInput() *string
	ScaleDownUnready() *string
	SetScaleDownUnready(val *string)
	ScaleDownUnreadyInput() *string
	ScaleDownUtilizationThreshold() *string
	SetScaleDownUtilizationThreshold(val *string)
	ScaleDownUtilizationThresholdInput() *string
	ScanInterval() *string
	SetScanInterval(val *string)
	ScanIntervalInput() *string
	SkipNodesWithLocalStorage() interface{}
	SetSkipNodesWithLocalStorage(val interface{})
	SkipNodesWithLocalStorageInput() interface{}
	SkipNodesWithSystemPods() interface{}
	SetSkipNodesWithSystemPods(val interface{})
	SkipNodesWithSystemPodsInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetBalanceSimilarNodeGroups()
	ResetEmptyBulkDeleteMax()
	ResetExpander()
	ResetMaxGracefulTerminationSec()
	ResetMaxNodeProvisioningTime()
	ResetMaxUnreadyNodes()
	ResetMaxUnreadyPercentage()
	ResetNewPodScaleUpDelay()
	ResetScaleDownDelayAfterAdd()
	ResetScaleDownDelayAfterDelete()
	ResetScaleDownDelayAfterFailure()
	ResetScaleDownUnneeded()
	ResetScaleDownUnready()
	ResetScaleDownUtilizationThreshold()
	ResetScanInterval()
	ResetSkipNodesWithLocalStorage()
	ResetSkipNodesWithSystemPods()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KubernetesClusterAutoScalerProfileOutputReference
type jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) BalanceSimilarNodeGroups() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"balanceSimilarNodeGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) BalanceSimilarNodeGroupsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"balanceSimilarNodeGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) EmptyBulkDeleteMax() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emptyBulkDeleteMax",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) EmptyBulkDeleteMaxInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emptyBulkDeleteMaxInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) Expander() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expander",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) ExpanderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expanderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) InternalValue() *KubernetesClusterAutoScalerProfile {
	var returns *KubernetesClusterAutoScalerProfile
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) MaxGracefulTerminationSec() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxGracefulTerminationSec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) MaxGracefulTerminationSecInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxGracefulTerminationSecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) MaxNodeProvisioningTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxNodeProvisioningTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) MaxNodeProvisioningTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxNodeProvisioningTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) MaxUnreadyNodes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxUnreadyNodes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) MaxUnreadyNodesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxUnreadyNodesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) MaxUnreadyPercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxUnreadyPercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) MaxUnreadyPercentageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxUnreadyPercentageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) NewPodScaleUpDelay() *string {
	var returns *string
	_jsii_.Get(
		j,
		"newPodScaleUpDelay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) NewPodScaleUpDelayInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"newPodScaleUpDelayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) ScaleDownDelayAfterAdd() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scaleDownDelayAfterAdd",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) ScaleDownDelayAfterAddInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scaleDownDelayAfterAddInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) ScaleDownDelayAfterDelete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scaleDownDelayAfterDelete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) ScaleDownDelayAfterDeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scaleDownDelayAfterDeleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) ScaleDownDelayAfterFailure() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scaleDownDelayAfterFailure",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) ScaleDownDelayAfterFailureInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scaleDownDelayAfterFailureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) ScaleDownUnneeded() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scaleDownUnneeded",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) ScaleDownUnneededInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scaleDownUnneededInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) ScaleDownUnready() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scaleDownUnready",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) ScaleDownUnreadyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scaleDownUnreadyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) ScaleDownUtilizationThreshold() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scaleDownUtilizationThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) ScaleDownUtilizationThresholdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scaleDownUtilizationThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) ScanInterval() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scanInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) ScanIntervalInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scanIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) SkipNodesWithLocalStorage() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipNodesWithLocalStorage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) SkipNodesWithLocalStorageInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipNodesWithLocalStorageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) SkipNodesWithSystemPods() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipNodesWithSystemPods",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) SkipNodesWithSystemPodsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipNodesWithSystemPodsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewKubernetesClusterAutoScalerProfileOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) KubernetesClusterAutoScalerProfileOutputReference {
	_init_.Initialize()

	j := jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesCluster.KubernetesClusterAutoScalerProfileOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKubernetesClusterAutoScalerProfileOutputReference_Override(k KubernetesClusterAutoScalerProfileOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesCluster.KubernetesClusterAutoScalerProfileOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) SetBalanceSimilarNodeGroups(val interface{}) {
	_jsii_.Set(
		j,
		"balanceSimilarNodeGroups",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) SetEmptyBulkDeleteMax(val *string) {
	_jsii_.Set(
		j,
		"emptyBulkDeleteMax",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) SetExpander(val *string) {
	_jsii_.Set(
		j,
		"expander",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) SetInternalValue(val *KubernetesClusterAutoScalerProfile) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) SetMaxGracefulTerminationSec(val *string) {
	_jsii_.Set(
		j,
		"maxGracefulTerminationSec",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) SetMaxNodeProvisioningTime(val *string) {
	_jsii_.Set(
		j,
		"maxNodeProvisioningTime",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) SetMaxUnreadyNodes(val *float64) {
	_jsii_.Set(
		j,
		"maxUnreadyNodes",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) SetMaxUnreadyPercentage(val *float64) {
	_jsii_.Set(
		j,
		"maxUnreadyPercentage",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) SetNewPodScaleUpDelay(val *string) {
	_jsii_.Set(
		j,
		"newPodScaleUpDelay",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) SetScaleDownDelayAfterAdd(val *string) {
	_jsii_.Set(
		j,
		"scaleDownDelayAfterAdd",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) SetScaleDownDelayAfterDelete(val *string) {
	_jsii_.Set(
		j,
		"scaleDownDelayAfterDelete",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) SetScaleDownDelayAfterFailure(val *string) {
	_jsii_.Set(
		j,
		"scaleDownDelayAfterFailure",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) SetScaleDownUnneeded(val *string) {
	_jsii_.Set(
		j,
		"scaleDownUnneeded",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) SetScaleDownUnready(val *string) {
	_jsii_.Set(
		j,
		"scaleDownUnready",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) SetScaleDownUtilizationThreshold(val *string) {
	_jsii_.Set(
		j,
		"scaleDownUtilizationThreshold",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) SetScanInterval(val *string) {
	_jsii_.Set(
		j,
		"scanInterval",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) SetSkipNodesWithLocalStorage(val interface{}) {
	_jsii_.Set(
		j,
		"skipNodesWithLocalStorage",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) SetSkipNodesWithSystemPods(val interface{}) {
	_jsii_.Set(
		j,
		"skipNodesWithSystemPods",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (k *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		k,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		k,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		k,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		k,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		k,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) ResetBalanceSimilarNodeGroups() {
	_jsii_.InvokeVoid(
		k,
		"resetBalanceSimilarNodeGroups",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) ResetEmptyBulkDeleteMax() {
	_jsii_.InvokeVoid(
		k,
		"resetEmptyBulkDeleteMax",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) ResetExpander() {
	_jsii_.InvokeVoid(
		k,
		"resetExpander",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) ResetMaxGracefulTerminationSec() {
	_jsii_.InvokeVoid(
		k,
		"resetMaxGracefulTerminationSec",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) ResetMaxNodeProvisioningTime() {
	_jsii_.InvokeVoid(
		k,
		"resetMaxNodeProvisioningTime",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) ResetMaxUnreadyNodes() {
	_jsii_.InvokeVoid(
		k,
		"resetMaxUnreadyNodes",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) ResetMaxUnreadyPercentage() {
	_jsii_.InvokeVoid(
		k,
		"resetMaxUnreadyPercentage",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) ResetNewPodScaleUpDelay() {
	_jsii_.InvokeVoid(
		k,
		"resetNewPodScaleUpDelay",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) ResetScaleDownDelayAfterAdd() {
	_jsii_.InvokeVoid(
		k,
		"resetScaleDownDelayAfterAdd",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) ResetScaleDownDelayAfterDelete() {
	_jsii_.InvokeVoid(
		k,
		"resetScaleDownDelayAfterDelete",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) ResetScaleDownDelayAfterFailure() {
	_jsii_.InvokeVoid(
		k,
		"resetScaleDownDelayAfterFailure",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) ResetScaleDownUnneeded() {
	_jsii_.InvokeVoid(
		k,
		"resetScaleDownUnneeded",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) ResetScaleDownUnready() {
	_jsii_.InvokeVoid(
		k,
		"resetScaleDownUnready",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) ResetScaleDownUtilizationThreshold() {
	_jsii_.InvokeVoid(
		k,
		"resetScaleDownUtilizationThreshold",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) ResetScanInterval() {
	_jsii_.InvokeVoid(
		k,
		"resetScanInterval",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) ResetSkipNodesWithLocalStorage() {
	_jsii_.InvokeVoid(
		k,
		"resetSkipNodesWithLocalStorage",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) ResetSkipNodesWithSystemPods() {
	_jsii_.InvokeVoid(
		k,
		"resetSkipNodesWithSystemPods",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type KubernetesClusterAzureActiveDirectoryRoleBasedAccessControl struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#admin_group_object_ids KubernetesCluster#admin_group_object_ids}.
	AdminGroupObjectIds *[]*string `field:"optional" json:"adminGroupObjectIds" yaml:"adminGroupObjectIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#azure_rbac_enabled KubernetesCluster#azure_rbac_enabled}.
	AzureRbacEnabled interface{} `field:"optional" json:"azureRbacEnabled" yaml:"azureRbacEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#client_app_id KubernetesCluster#client_app_id}.
	ClientAppId *string `field:"optional" json:"clientAppId" yaml:"clientAppId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#managed KubernetesCluster#managed}.
	Managed interface{} `field:"optional" json:"managed" yaml:"managed"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#server_app_id KubernetesCluster#server_app_id}.
	ServerAppId *string `field:"optional" json:"serverAppId" yaml:"serverAppId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#server_app_secret KubernetesCluster#server_app_secret}.
	ServerAppSecret *string `field:"optional" json:"serverAppSecret" yaml:"serverAppSecret"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#tenant_id KubernetesCluster#tenant_id}.
	TenantId *string `field:"optional" json:"tenantId" yaml:"tenantId"`
}

type KubernetesClusterAzureActiveDirectoryRoleBasedAccessControlOutputReference interface {
	cdktf.ComplexObject
	AdminGroupObjectIds() *[]*string
	SetAdminGroupObjectIds(val *[]*string)
	AdminGroupObjectIdsInput() *[]*string
	AzureRbacEnabled() interface{}
	SetAzureRbacEnabled(val interface{})
	AzureRbacEnabledInput() interface{}
	ClientAppId() *string
	SetClientAppId(val *string)
	ClientAppIdInput() *string
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *KubernetesClusterAzureActiveDirectoryRoleBasedAccessControl
	SetInternalValue(val *KubernetesClusterAzureActiveDirectoryRoleBasedAccessControl)
	Managed() interface{}
	SetManaged(val interface{})
	ManagedInput() interface{}
	ServerAppId() *string
	SetServerAppId(val *string)
	ServerAppIdInput() *string
	ServerAppSecret() *string
	SetServerAppSecret(val *string)
	ServerAppSecretInput() *string
	TenantId() *string
	SetTenantId(val *string)
	TenantIdInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetAdminGroupObjectIds()
	ResetAzureRbacEnabled()
	ResetClientAppId()
	ResetManaged()
	ResetServerAppId()
	ResetServerAppSecret()
	ResetTenantId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KubernetesClusterAzureActiveDirectoryRoleBasedAccessControlOutputReference
type jsiiProxy_KubernetesClusterAzureActiveDirectoryRoleBasedAccessControlOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_KubernetesClusterAzureActiveDirectoryRoleBasedAccessControlOutputReference) AdminGroupObjectIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"adminGroupObjectIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAzureActiveDirectoryRoleBasedAccessControlOutputReference) AdminGroupObjectIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"adminGroupObjectIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAzureActiveDirectoryRoleBasedAccessControlOutputReference) AzureRbacEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"azureRbacEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAzureActiveDirectoryRoleBasedAccessControlOutputReference) AzureRbacEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"azureRbacEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAzureActiveDirectoryRoleBasedAccessControlOutputReference) ClientAppId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientAppId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAzureActiveDirectoryRoleBasedAccessControlOutputReference) ClientAppIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientAppIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAzureActiveDirectoryRoleBasedAccessControlOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAzureActiveDirectoryRoleBasedAccessControlOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAzureActiveDirectoryRoleBasedAccessControlOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAzureActiveDirectoryRoleBasedAccessControlOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAzureActiveDirectoryRoleBasedAccessControlOutputReference) InternalValue() *KubernetesClusterAzureActiveDirectoryRoleBasedAccessControl {
	var returns *KubernetesClusterAzureActiveDirectoryRoleBasedAccessControl
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAzureActiveDirectoryRoleBasedAccessControlOutputReference) Managed() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"managed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAzureActiveDirectoryRoleBasedAccessControlOutputReference) ManagedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"managedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAzureActiveDirectoryRoleBasedAccessControlOutputReference) ServerAppId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverAppId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAzureActiveDirectoryRoleBasedAccessControlOutputReference) ServerAppIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverAppIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAzureActiveDirectoryRoleBasedAccessControlOutputReference) ServerAppSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverAppSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAzureActiveDirectoryRoleBasedAccessControlOutputReference) ServerAppSecretInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverAppSecretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAzureActiveDirectoryRoleBasedAccessControlOutputReference) TenantId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenantId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAzureActiveDirectoryRoleBasedAccessControlOutputReference) TenantIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenantIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAzureActiveDirectoryRoleBasedAccessControlOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAzureActiveDirectoryRoleBasedAccessControlOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewKubernetesClusterAzureActiveDirectoryRoleBasedAccessControlOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) KubernetesClusterAzureActiveDirectoryRoleBasedAccessControlOutputReference {
	_init_.Initialize()

	j := jsiiProxy_KubernetesClusterAzureActiveDirectoryRoleBasedAccessControlOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesCluster.KubernetesClusterAzureActiveDirectoryRoleBasedAccessControlOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKubernetesClusterAzureActiveDirectoryRoleBasedAccessControlOutputReference_Override(k KubernetesClusterAzureActiveDirectoryRoleBasedAccessControlOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesCluster.KubernetesClusterAzureActiveDirectoryRoleBasedAccessControlOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_KubernetesClusterAzureActiveDirectoryRoleBasedAccessControlOutputReference) SetAdminGroupObjectIds(val *[]*string) {
	_jsii_.Set(
		j,
		"adminGroupObjectIds",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterAzureActiveDirectoryRoleBasedAccessControlOutputReference) SetAzureRbacEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"azureRbacEnabled",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterAzureActiveDirectoryRoleBasedAccessControlOutputReference) SetClientAppId(val *string) {
	_jsii_.Set(
		j,
		"clientAppId",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterAzureActiveDirectoryRoleBasedAccessControlOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterAzureActiveDirectoryRoleBasedAccessControlOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterAzureActiveDirectoryRoleBasedAccessControlOutputReference) SetInternalValue(val *KubernetesClusterAzureActiveDirectoryRoleBasedAccessControl) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterAzureActiveDirectoryRoleBasedAccessControlOutputReference) SetManaged(val interface{}) {
	_jsii_.Set(
		j,
		"managed",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterAzureActiveDirectoryRoleBasedAccessControlOutputReference) SetServerAppId(val *string) {
	_jsii_.Set(
		j,
		"serverAppId",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterAzureActiveDirectoryRoleBasedAccessControlOutputReference) SetServerAppSecret(val *string) {
	_jsii_.Set(
		j,
		"serverAppSecret",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterAzureActiveDirectoryRoleBasedAccessControlOutputReference) SetTenantId(val *string) {
	_jsii_.Set(
		j,
		"tenantId",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterAzureActiveDirectoryRoleBasedAccessControlOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterAzureActiveDirectoryRoleBasedAccessControlOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (k *jsiiProxy_KubernetesClusterAzureActiveDirectoryRoleBasedAccessControlOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterAzureActiveDirectoryRoleBasedAccessControlOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterAzureActiveDirectoryRoleBasedAccessControlOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterAzureActiveDirectoryRoleBasedAccessControlOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		k,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterAzureActiveDirectoryRoleBasedAccessControlOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterAzureActiveDirectoryRoleBasedAccessControlOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		k,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterAzureActiveDirectoryRoleBasedAccessControlOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		k,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterAzureActiveDirectoryRoleBasedAccessControlOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		k,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterAzureActiveDirectoryRoleBasedAccessControlOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterAzureActiveDirectoryRoleBasedAccessControlOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		k,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterAzureActiveDirectoryRoleBasedAccessControlOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterAzureActiveDirectoryRoleBasedAccessControlOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterAzureActiveDirectoryRoleBasedAccessControlOutputReference) ResetAdminGroupObjectIds() {
	_jsii_.InvokeVoid(
		k,
		"resetAdminGroupObjectIds",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterAzureActiveDirectoryRoleBasedAccessControlOutputReference) ResetAzureRbacEnabled() {
	_jsii_.InvokeVoid(
		k,
		"resetAzureRbacEnabled",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterAzureActiveDirectoryRoleBasedAccessControlOutputReference) ResetClientAppId() {
	_jsii_.InvokeVoid(
		k,
		"resetClientAppId",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterAzureActiveDirectoryRoleBasedAccessControlOutputReference) ResetManaged() {
	_jsii_.InvokeVoid(
		k,
		"resetManaged",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterAzureActiveDirectoryRoleBasedAccessControlOutputReference) ResetServerAppId() {
	_jsii_.InvokeVoid(
		k,
		"resetServerAppId",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterAzureActiveDirectoryRoleBasedAccessControlOutputReference) ResetServerAppSecret() {
	_jsii_.InvokeVoid(
		k,
		"resetServerAppSecret",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterAzureActiveDirectoryRoleBasedAccessControlOutputReference) ResetTenantId() {
	_jsii_.InvokeVoid(
		k,
		"resetTenantId",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterAzureActiveDirectoryRoleBasedAccessControlOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterAzureActiveDirectoryRoleBasedAccessControlOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

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
	// ingress_application_gateway block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#ingress_application_gateway KubernetesCluster#ingress_application_gateway}
	IngressApplicationGateway *KubernetesClusterIngressApplicationGateway `field:"optional" json:"ingressApplicationGateway" yaml:"ingressApplicationGateway"`
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#tags KubernetesCluster#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#timeouts KubernetesCluster#timeouts}
	Timeouts *KubernetesClusterTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// windows_profile block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#windows_profile KubernetesCluster#windows_profile}
	WindowsProfile *KubernetesClusterWindowsProfile `field:"optional" json:"windowsProfile" yaml:"windowsProfile"`
}

type KubernetesClusterDefaultNodePool struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#name KubernetesCluster#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#vm_size KubernetesCluster#vm_size}.
	VmSize *string `field:"required" json:"vmSize" yaml:"vmSize"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#capacity_reservation_group_id KubernetesCluster#capacity_reservation_group_id}.
	CapacityReservationGroupId *string `field:"optional" json:"capacityReservationGroupId" yaml:"capacityReservationGroupId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#enable_auto_scaling KubernetesCluster#enable_auto_scaling}.
	EnableAutoScaling interface{} `field:"optional" json:"enableAutoScaling" yaml:"enableAutoScaling"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#enable_host_encryption KubernetesCluster#enable_host_encryption}.
	EnableHostEncryption interface{} `field:"optional" json:"enableHostEncryption" yaml:"enableHostEncryption"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#enable_node_public_ip KubernetesCluster#enable_node_public_ip}.
	EnableNodePublicIp interface{} `field:"optional" json:"enableNodePublicIp" yaml:"enableNodePublicIp"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#fips_enabled KubernetesCluster#fips_enabled}.
	FipsEnabled interface{} `field:"optional" json:"fipsEnabled" yaml:"fipsEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#host_group_id KubernetesCluster#host_group_id}.
	HostGroupId *string `field:"optional" json:"hostGroupId" yaml:"hostGroupId"`
	// kubelet_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#kubelet_config KubernetesCluster#kubelet_config}
	KubeletConfig *KubernetesClusterDefaultNodePoolKubeletConfig `field:"optional" json:"kubeletConfig" yaml:"kubeletConfig"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#kubelet_disk_type KubernetesCluster#kubelet_disk_type}.
	KubeletDiskType *string `field:"optional" json:"kubeletDiskType" yaml:"kubeletDiskType"`
	// linux_os_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#linux_os_config KubernetesCluster#linux_os_config}
	LinuxOsConfig *KubernetesClusterDefaultNodePoolLinuxOsConfig `field:"optional" json:"linuxOsConfig" yaml:"linuxOsConfig"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#max_count KubernetesCluster#max_count}.
	MaxCount *float64 `field:"optional" json:"maxCount" yaml:"maxCount"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#max_pods KubernetesCluster#max_pods}.
	MaxPods *float64 `field:"optional" json:"maxPods" yaml:"maxPods"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#min_count KubernetesCluster#min_count}.
	MinCount *float64 `field:"optional" json:"minCount" yaml:"minCount"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#node_count KubernetesCluster#node_count}.
	NodeCount *float64 `field:"optional" json:"nodeCount" yaml:"nodeCount"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#node_labels KubernetesCluster#node_labels}.
	NodeLabels *map[string]*string `field:"optional" json:"nodeLabels" yaml:"nodeLabels"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#node_public_ip_prefix_id KubernetesCluster#node_public_ip_prefix_id}.
	NodePublicIpPrefixId *string `field:"optional" json:"nodePublicIpPrefixId" yaml:"nodePublicIpPrefixId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#node_taints KubernetesCluster#node_taints}.
	NodeTaints *[]*string `field:"optional" json:"nodeTaints" yaml:"nodeTaints"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#only_critical_addons_enabled KubernetesCluster#only_critical_addons_enabled}.
	OnlyCriticalAddonsEnabled interface{} `field:"optional" json:"onlyCriticalAddonsEnabled" yaml:"onlyCriticalAddonsEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#orchestrator_version KubernetesCluster#orchestrator_version}.
	OrchestratorVersion *string `field:"optional" json:"orchestratorVersion" yaml:"orchestratorVersion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#os_disk_size_gb KubernetesCluster#os_disk_size_gb}.
	OsDiskSizeGb *float64 `field:"optional" json:"osDiskSizeGb" yaml:"osDiskSizeGb"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#os_disk_type KubernetesCluster#os_disk_type}.
	OsDiskType *string `field:"optional" json:"osDiskType" yaml:"osDiskType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#os_sku KubernetesCluster#os_sku}.
	OsSku *string `field:"optional" json:"osSku" yaml:"osSku"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#pod_subnet_id KubernetesCluster#pod_subnet_id}.
	PodSubnetId *string `field:"optional" json:"podSubnetId" yaml:"podSubnetId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#proximity_placement_group_id KubernetesCluster#proximity_placement_group_id}.
	ProximityPlacementGroupId *string `field:"optional" json:"proximityPlacementGroupId" yaml:"proximityPlacementGroupId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#tags KubernetesCluster#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#type KubernetesCluster#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#ultra_ssd_enabled KubernetesCluster#ultra_ssd_enabled}.
	UltraSsdEnabled interface{} `field:"optional" json:"ultraSsdEnabled" yaml:"ultraSsdEnabled"`
	// upgrade_settings block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#upgrade_settings KubernetesCluster#upgrade_settings}
	UpgradeSettings *KubernetesClusterDefaultNodePoolUpgradeSettings `field:"optional" json:"upgradeSettings" yaml:"upgradeSettings"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#vnet_subnet_id KubernetesCluster#vnet_subnet_id}.
	VnetSubnetId *string `field:"optional" json:"vnetSubnetId" yaml:"vnetSubnetId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#zones KubernetesCluster#zones}.
	Zones *[]*string `field:"optional" json:"zones" yaml:"zones"`
}

type KubernetesClusterDefaultNodePoolKubeletConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#allowed_unsafe_sysctls KubernetesCluster#allowed_unsafe_sysctls}.
	AllowedUnsafeSysctls *[]*string `field:"optional" json:"allowedUnsafeSysctls" yaml:"allowedUnsafeSysctls"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#container_log_max_line KubernetesCluster#container_log_max_line}.
	ContainerLogMaxLine *float64 `field:"optional" json:"containerLogMaxLine" yaml:"containerLogMaxLine"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#container_log_max_size_mb KubernetesCluster#container_log_max_size_mb}.
	ContainerLogMaxSizeMb *float64 `field:"optional" json:"containerLogMaxSizeMb" yaml:"containerLogMaxSizeMb"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#cpu_cfs_quota_enabled KubernetesCluster#cpu_cfs_quota_enabled}.
	CpuCfsQuotaEnabled interface{} `field:"optional" json:"cpuCfsQuotaEnabled" yaml:"cpuCfsQuotaEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#cpu_cfs_quota_period KubernetesCluster#cpu_cfs_quota_period}.
	CpuCfsQuotaPeriod *string `field:"optional" json:"cpuCfsQuotaPeriod" yaml:"cpuCfsQuotaPeriod"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#cpu_manager_policy KubernetesCluster#cpu_manager_policy}.
	CpuManagerPolicy *string `field:"optional" json:"cpuManagerPolicy" yaml:"cpuManagerPolicy"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#image_gc_high_threshold KubernetesCluster#image_gc_high_threshold}.
	ImageGcHighThreshold *float64 `field:"optional" json:"imageGcHighThreshold" yaml:"imageGcHighThreshold"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#image_gc_low_threshold KubernetesCluster#image_gc_low_threshold}.
	ImageGcLowThreshold *float64 `field:"optional" json:"imageGcLowThreshold" yaml:"imageGcLowThreshold"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#pod_max_pid KubernetesCluster#pod_max_pid}.
	PodMaxPid *float64 `field:"optional" json:"podMaxPid" yaml:"podMaxPid"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#topology_manager_policy KubernetesCluster#topology_manager_policy}.
	TopologyManagerPolicy *string `field:"optional" json:"topologyManagerPolicy" yaml:"topologyManagerPolicy"`
}

type KubernetesClusterDefaultNodePoolKubeletConfigOutputReference interface {
	cdktf.ComplexObject
	AllowedUnsafeSysctls() *[]*string
	SetAllowedUnsafeSysctls(val *[]*string)
	AllowedUnsafeSysctlsInput() *[]*string
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	ContainerLogMaxLine() *float64
	SetContainerLogMaxLine(val *float64)
	ContainerLogMaxLineInput() *float64
	ContainerLogMaxSizeMb() *float64
	SetContainerLogMaxSizeMb(val *float64)
	ContainerLogMaxSizeMbInput() *float64
	CpuCfsQuotaEnabled() interface{}
	SetCpuCfsQuotaEnabled(val interface{})
	CpuCfsQuotaEnabledInput() interface{}
	CpuCfsQuotaPeriod() *string
	SetCpuCfsQuotaPeriod(val *string)
	CpuCfsQuotaPeriodInput() *string
	CpuManagerPolicy() *string
	SetCpuManagerPolicy(val *string)
	CpuManagerPolicyInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	ImageGcHighThreshold() *float64
	SetImageGcHighThreshold(val *float64)
	ImageGcHighThresholdInput() *float64
	ImageGcLowThreshold() *float64
	SetImageGcLowThreshold(val *float64)
	ImageGcLowThresholdInput() *float64
	InternalValue() *KubernetesClusterDefaultNodePoolKubeletConfig
	SetInternalValue(val *KubernetesClusterDefaultNodePoolKubeletConfig)
	PodMaxPid() *float64
	SetPodMaxPid(val *float64)
	PodMaxPidInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TopologyManagerPolicy() *string
	SetTopologyManagerPolicy(val *string)
	TopologyManagerPolicyInput() *string
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetAllowedUnsafeSysctls()
	ResetContainerLogMaxLine()
	ResetContainerLogMaxSizeMb()
	ResetCpuCfsQuotaEnabled()
	ResetCpuCfsQuotaPeriod()
	ResetCpuManagerPolicy()
	ResetImageGcHighThreshold()
	ResetImageGcLowThreshold()
	ResetPodMaxPid()
	ResetTopologyManagerPolicy()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KubernetesClusterDefaultNodePoolKubeletConfigOutputReference
type jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference) AllowedUnsafeSysctls() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedUnsafeSysctls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference) AllowedUnsafeSysctlsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedUnsafeSysctlsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference) ContainerLogMaxLine() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"containerLogMaxLine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference) ContainerLogMaxLineInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"containerLogMaxLineInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference) ContainerLogMaxSizeMb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"containerLogMaxSizeMb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference) ContainerLogMaxSizeMbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"containerLogMaxSizeMbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference) CpuCfsQuotaEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cpuCfsQuotaEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference) CpuCfsQuotaEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cpuCfsQuotaEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference) CpuCfsQuotaPeriod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cpuCfsQuotaPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference) CpuCfsQuotaPeriodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cpuCfsQuotaPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference) CpuManagerPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cpuManagerPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference) CpuManagerPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cpuManagerPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference) ImageGcHighThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"imageGcHighThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference) ImageGcHighThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"imageGcHighThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference) ImageGcLowThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"imageGcLowThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference) ImageGcLowThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"imageGcLowThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference) InternalValue() *KubernetesClusterDefaultNodePoolKubeletConfig {
	var returns *KubernetesClusterDefaultNodePoolKubeletConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference) PodMaxPid() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"podMaxPid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference) PodMaxPidInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"podMaxPidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference) TopologyManagerPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"topologyManagerPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference) TopologyManagerPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"topologyManagerPolicyInput",
		&returns,
	)
	return returns
}


func NewKubernetesClusterDefaultNodePoolKubeletConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) KubernetesClusterDefaultNodePoolKubeletConfigOutputReference {
	_init_.Initialize()

	j := jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesCluster.KubernetesClusterDefaultNodePoolKubeletConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKubernetesClusterDefaultNodePoolKubeletConfigOutputReference_Override(k KubernetesClusterDefaultNodePoolKubeletConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesCluster.KubernetesClusterDefaultNodePoolKubeletConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference) SetAllowedUnsafeSysctls(val *[]*string) {
	_jsii_.Set(
		j,
		"allowedUnsafeSysctls",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference) SetContainerLogMaxLine(val *float64) {
	_jsii_.Set(
		j,
		"containerLogMaxLine",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference) SetContainerLogMaxSizeMb(val *float64) {
	_jsii_.Set(
		j,
		"containerLogMaxSizeMb",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference) SetCpuCfsQuotaEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"cpuCfsQuotaEnabled",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference) SetCpuCfsQuotaPeriod(val *string) {
	_jsii_.Set(
		j,
		"cpuCfsQuotaPeriod",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference) SetCpuManagerPolicy(val *string) {
	_jsii_.Set(
		j,
		"cpuManagerPolicy",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference) SetImageGcHighThreshold(val *float64) {
	_jsii_.Set(
		j,
		"imageGcHighThreshold",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference) SetImageGcLowThreshold(val *float64) {
	_jsii_.Set(
		j,
		"imageGcLowThreshold",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference) SetInternalValue(val *KubernetesClusterDefaultNodePoolKubeletConfig) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference) SetPodMaxPid(val *float64) {
	_jsii_.Set(
		j,
		"podMaxPid",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference) SetTopologyManagerPolicy(val *string) {
	_jsii_.Set(
		j,
		"topologyManagerPolicy",
		val,
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		k,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		k,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		k,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		k,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		k,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference) ResetAllowedUnsafeSysctls() {
	_jsii_.InvokeVoid(
		k,
		"resetAllowedUnsafeSysctls",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference) ResetContainerLogMaxLine() {
	_jsii_.InvokeVoid(
		k,
		"resetContainerLogMaxLine",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference) ResetContainerLogMaxSizeMb() {
	_jsii_.InvokeVoid(
		k,
		"resetContainerLogMaxSizeMb",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference) ResetCpuCfsQuotaEnabled() {
	_jsii_.InvokeVoid(
		k,
		"resetCpuCfsQuotaEnabled",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference) ResetCpuCfsQuotaPeriod() {
	_jsii_.InvokeVoid(
		k,
		"resetCpuCfsQuotaPeriod",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference) ResetCpuManagerPolicy() {
	_jsii_.InvokeVoid(
		k,
		"resetCpuManagerPolicy",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference) ResetImageGcHighThreshold() {
	_jsii_.InvokeVoid(
		k,
		"resetImageGcHighThreshold",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference) ResetImageGcLowThreshold() {
	_jsii_.InvokeVoid(
		k,
		"resetImageGcLowThreshold",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference) ResetPodMaxPid() {
	_jsii_.InvokeVoid(
		k,
		"resetPodMaxPid",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference) ResetTopologyManagerPolicy() {
	_jsii_.InvokeVoid(
		k,
		"resetTopologyManagerPolicy",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type KubernetesClusterDefaultNodePoolLinuxOsConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#swap_file_size_mb KubernetesCluster#swap_file_size_mb}.
	SwapFileSizeMb *float64 `field:"optional" json:"swapFileSizeMb" yaml:"swapFileSizeMb"`
	// sysctl_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#sysctl_config KubernetesCluster#sysctl_config}
	SysctlConfig *KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfig `field:"optional" json:"sysctlConfig" yaml:"sysctlConfig"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#transparent_huge_page_defrag KubernetesCluster#transparent_huge_page_defrag}.
	TransparentHugePageDefrag *string `field:"optional" json:"transparentHugePageDefrag" yaml:"transparentHugePageDefrag"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#transparent_huge_page_enabled KubernetesCluster#transparent_huge_page_enabled}.
	TransparentHugePageEnabled *string `field:"optional" json:"transparentHugePageEnabled" yaml:"transparentHugePageEnabled"`
}

type KubernetesClusterDefaultNodePoolLinuxOsConfigOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *KubernetesClusterDefaultNodePoolLinuxOsConfig
	SetInternalValue(val *KubernetesClusterDefaultNodePoolLinuxOsConfig)
	SwapFileSizeMb() *float64
	SetSwapFileSizeMb(val *float64)
	SwapFileSizeMbInput() *float64
	SysctlConfig() KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference
	SysctlConfigInput() *KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfig
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TransparentHugePageDefrag() *string
	SetTransparentHugePageDefrag(val *string)
	TransparentHugePageDefragInput() *string
	TransparentHugePageEnabled() *string
	SetTransparentHugePageEnabled(val *string)
	TransparentHugePageEnabledInput() *string
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	PutSysctlConfig(value *KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfig)
	ResetSwapFileSizeMb()
	ResetSysctlConfig()
	ResetTransparentHugePageDefrag()
	ResetTransparentHugePageEnabled()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KubernetesClusterDefaultNodePoolLinuxOsConfigOutputReference
type jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigOutputReference) InternalValue() *KubernetesClusterDefaultNodePoolLinuxOsConfig {
	var returns *KubernetesClusterDefaultNodePoolLinuxOsConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigOutputReference) SwapFileSizeMb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"swapFileSizeMb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigOutputReference) SwapFileSizeMbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"swapFileSizeMbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigOutputReference) SysctlConfig() KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference {
	var returns KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference
	_jsii_.Get(
		j,
		"sysctlConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigOutputReference) SysctlConfigInput() *KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfig {
	var returns *KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfig
	_jsii_.Get(
		j,
		"sysctlConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigOutputReference) TransparentHugePageDefrag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transparentHugePageDefrag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigOutputReference) TransparentHugePageDefragInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transparentHugePageDefragInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigOutputReference) TransparentHugePageEnabled() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transparentHugePageEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigOutputReference) TransparentHugePageEnabledInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transparentHugePageEnabledInput",
		&returns,
	)
	return returns
}


func NewKubernetesClusterDefaultNodePoolLinuxOsConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) KubernetesClusterDefaultNodePoolLinuxOsConfigOutputReference {
	_init_.Initialize()

	j := jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesCluster.KubernetesClusterDefaultNodePoolLinuxOsConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKubernetesClusterDefaultNodePoolLinuxOsConfigOutputReference_Override(k KubernetesClusterDefaultNodePoolLinuxOsConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesCluster.KubernetesClusterDefaultNodePoolLinuxOsConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigOutputReference) SetInternalValue(val *KubernetesClusterDefaultNodePoolLinuxOsConfig) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigOutputReference) SetSwapFileSizeMb(val *float64) {
	_jsii_.Set(
		j,
		"swapFileSizeMb",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigOutputReference) SetTransparentHugePageDefrag(val *string) {
	_jsii_.Set(
		j,
		"transparentHugePageDefrag",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigOutputReference) SetTransparentHugePageEnabled(val *string) {
	_jsii_.Set(
		j,
		"transparentHugePageEnabled",
		val,
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		k,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		k,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		k,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		k,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		k,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigOutputReference) PutSysctlConfig(value *KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfig) {
	_jsii_.InvokeVoid(
		k,
		"putSysctlConfig",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigOutputReference) ResetSwapFileSizeMb() {
	_jsii_.InvokeVoid(
		k,
		"resetSwapFileSizeMb",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigOutputReference) ResetSysctlConfig() {
	_jsii_.InvokeVoid(
		k,
		"resetSysctlConfig",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigOutputReference) ResetTransparentHugePageDefrag() {
	_jsii_.InvokeVoid(
		k,
		"resetTransparentHugePageDefrag",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigOutputReference) ResetTransparentHugePageEnabled() {
	_jsii_.InvokeVoid(
		k,
		"resetTransparentHugePageEnabled",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#fs_aio_max_nr KubernetesCluster#fs_aio_max_nr}.
	FsAioMaxNr *float64 `field:"optional" json:"fsAioMaxNr" yaml:"fsAioMaxNr"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#fs_file_max KubernetesCluster#fs_file_max}.
	FsFileMax *float64 `field:"optional" json:"fsFileMax" yaml:"fsFileMax"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#fs_inotify_max_user_watches KubernetesCluster#fs_inotify_max_user_watches}.
	FsInotifyMaxUserWatches *float64 `field:"optional" json:"fsInotifyMaxUserWatches" yaml:"fsInotifyMaxUserWatches"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#fs_nr_open KubernetesCluster#fs_nr_open}.
	FsNrOpen *float64 `field:"optional" json:"fsNrOpen" yaml:"fsNrOpen"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#kernel_threads_max KubernetesCluster#kernel_threads_max}.
	KernelThreadsMax *float64 `field:"optional" json:"kernelThreadsMax" yaml:"kernelThreadsMax"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#net_core_netdev_max_backlog KubernetesCluster#net_core_netdev_max_backlog}.
	NetCoreNetdevMaxBacklog *float64 `field:"optional" json:"netCoreNetdevMaxBacklog" yaml:"netCoreNetdevMaxBacklog"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#net_core_optmem_max KubernetesCluster#net_core_optmem_max}.
	NetCoreOptmemMax *float64 `field:"optional" json:"netCoreOptmemMax" yaml:"netCoreOptmemMax"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#net_core_rmem_default KubernetesCluster#net_core_rmem_default}.
	NetCoreRmemDefault *float64 `field:"optional" json:"netCoreRmemDefault" yaml:"netCoreRmemDefault"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#net_core_rmem_max KubernetesCluster#net_core_rmem_max}.
	NetCoreRmemMax *float64 `field:"optional" json:"netCoreRmemMax" yaml:"netCoreRmemMax"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#net_core_somaxconn KubernetesCluster#net_core_somaxconn}.
	NetCoreSomaxconn *float64 `field:"optional" json:"netCoreSomaxconn" yaml:"netCoreSomaxconn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#net_core_wmem_default KubernetesCluster#net_core_wmem_default}.
	NetCoreWmemDefault *float64 `field:"optional" json:"netCoreWmemDefault" yaml:"netCoreWmemDefault"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#net_core_wmem_max KubernetesCluster#net_core_wmem_max}.
	NetCoreWmemMax *float64 `field:"optional" json:"netCoreWmemMax" yaml:"netCoreWmemMax"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#net_ipv4_ip_local_port_range_max KubernetesCluster#net_ipv4_ip_local_port_range_max}.
	NetIpv4IpLocalPortRangeMax *float64 `field:"optional" json:"netIpv4IpLocalPortRangeMax" yaml:"netIpv4IpLocalPortRangeMax"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#net_ipv4_ip_local_port_range_min KubernetesCluster#net_ipv4_ip_local_port_range_min}.
	NetIpv4IpLocalPortRangeMin *float64 `field:"optional" json:"netIpv4IpLocalPortRangeMin" yaml:"netIpv4IpLocalPortRangeMin"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#net_ipv4_neigh_default_gc_thresh1 KubernetesCluster#net_ipv4_neigh_default_gc_thresh1}.
	NetIpv4NeighDefaultGcThresh1 *float64 `field:"optional" json:"netIpv4NeighDefaultGcThresh1" yaml:"netIpv4NeighDefaultGcThresh1"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#net_ipv4_neigh_default_gc_thresh2 KubernetesCluster#net_ipv4_neigh_default_gc_thresh2}.
	NetIpv4NeighDefaultGcThresh2 *float64 `field:"optional" json:"netIpv4NeighDefaultGcThresh2" yaml:"netIpv4NeighDefaultGcThresh2"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#net_ipv4_neigh_default_gc_thresh3 KubernetesCluster#net_ipv4_neigh_default_gc_thresh3}.
	NetIpv4NeighDefaultGcThresh3 *float64 `field:"optional" json:"netIpv4NeighDefaultGcThresh3" yaml:"netIpv4NeighDefaultGcThresh3"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#net_ipv4_tcp_fin_timeout KubernetesCluster#net_ipv4_tcp_fin_timeout}.
	NetIpv4TcpFinTimeout *float64 `field:"optional" json:"netIpv4TcpFinTimeout" yaml:"netIpv4TcpFinTimeout"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#net_ipv4_tcp_keepalive_intvl KubernetesCluster#net_ipv4_tcp_keepalive_intvl}.
	NetIpv4TcpKeepaliveIntvl *float64 `field:"optional" json:"netIpv4TcpKeepaliveIntvl" yaml:"netIpv4TcpKeepaliveIntvl"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#net_ipv4_tcp_keepalive_probes KubernetesCluster#net_ipv4_tcp_keepalive_probes}.
	NetIpv4TcpKeepaliveProbes *float64 `field:"optional" json:"netIpv4TcpKeepaliveProbes" yaml:"netIpv4TcpKeepaliveProbes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#net_ipv4_tcp_keepalive_time KubernetesCluster#net_ipv4_tcp_keepalive_time}.
	NetIpv4TcpKeepaliveTime *float64 `field:"optional" json:"netIpv4TcpKeepaliveTime" yaml:"netIpv4TcpKeepaliveTime"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#net_ipv4_tcp_max_syn_backlog KubernetesCluster#net_ipv4_tcp_max_syn_backlog}.
	NetIpv4TcpMaxSynBacklog *float64 `field:"optional" json:"netIpv4TcpMaxSynBacklog" yaml:"netIpv4TcpMaxSynBacklog"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#net_ipv4_tcp_max_tw_buckets KubernetesCluster#net_ipv4_tcp_max_tw_buckets}.
	NetIpv4TcpMaxTwBuckets *float64 `field:"optional" json:"netIpv4TcpMaxTwBuckets" yaml:"netIpv4TcpMaxTwBuckets"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#net_ipv4_tcp_tw_reuse KubernetesCluster#net_ipv4_tcp_tw_reuse}.
	NetIpv4TcpTwReuse interface{} `field:"optional" json:"netIpv4TcpTwReuse" yaml:"netIpv4TcpTwReuse"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#net_netfilter_nf_conntrack_buckets KubernetesCluster#net_netfilter_nf_conntrack_buckets}.
	NetNetfilterNfConntrackBuckets *float64 `field:"optional" json:"netNetfilterNfConntrackBuckets" yaml:"netNetfilterNfConntrackBuckets"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#net_netfilter_nf_conntrack_max KubernetesCluster#net_netfilter_nf_conntrack_max}.
	NetNetfilterNfConntrackMax *float64 `field:"optional" json:"netNetfilterNfConntrackMax" yaml:"netNetfilterNfConntrackMax"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#vm_max_map_count KubernetesCluster#vm_max_map_count}.
	VmMaxMapCount *float64 `field:"optional" json:"vmMaxMapCount" yaml:"vmMaxMapCount"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#vm_swappiness KubernetesCluster#vm_swappiness}.
	VmSwappiness *float64 `field:"optional" json:"vmSwappiness" yaml:"vmSwappiness"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#vm_vfs_cache_pressure KubernetesCluster#vm_vfs_cache_pressure}.
	VmVfsCachePressure *float64 `field:"optional" json:"vmVfsCachePressure" yaml:"vmVfsCachePressure"`
}

type KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	FsAioMaxNr() *float64
	SetFsAioMaxNr(val *float64)
	FsAioMaxNrInput() *float64
	FsFileMax() *float64
	SetFsFileMax(val *float64)
	FsFileMaxInput() *float64
	FsInotifyMaxUserWatches() *float64
	SetFsInotifyMaxUserWatches(val *float64)
	FsInotifyMaxUserWatchesInput() *float64
	FsNrOpen() *float64
	SetFsNrOpen(val *float64)
	FsNrOpenInput() *float64
	InternalValue() *KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfig
	SetInternalValue(val *KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfig)
	KernelThreadsMax() *float64
	SetKernelThreadsMax(val *float64)
	KernelThreadsMaxInput() *float64
	NetCoreNetdevMaxBacklog() *float64
	SetNetCoreNetdevMaxBacklog(val *float64)
	NetCoreNetdevMaxBacklogInput() *float64
	NetCoreOptmemMax() *float64
	SetNetCoreOptmemMax(val *float64)
	NetCoreOptmemMaxInput() *float64
	NetCoreRmemDefault() *float64
	SetNetCoreRmemDefault(val *float64)
	NetCoreRmemDefaultInput() *float64
	NetCoreRmemMax() *float64
	SetNetCoreRmemMax(val *float64)
	NetCoreRmemMaxInput() *float64
	NetCoreSomaxconn() *float64
	SetNetCoreSomaxconn(val *float64)
	NetCoreSomaxconnInput() *float64
	NetCoreWmemDefault() *float64
	SetNetCoreWmemDefault(val *float64)
	NetCoreWmemDefaultInput() *float64
	NetCoreWmemMax() *float64
	SetNetCoreWmemMax(val *float64)
	NetCoreWmemMaxInput() *float64
	NetIpv4IpLocalPortRangeMax() *float64
	SetNetIpv4IpLocalPortRangeMax(val *float64)
	NetIpv4IpLocalPortRangeMaxInput() *float64
	NetIpv4IpLocalPortRangeMin() *float64
	SetNetIpv4IpLocalPortRangeMin(val *float64)
	NetIpv4IpLocalPortRangeMinInput() *float64
	NetIpv4NeighDefaultGcThresh1() *float64
	SetNetIpv4NeighDefaultGcThresh1(val *float64)
	NetIpv4NeighDefaultGcThresh1Input() *float64
	NetIpv4NeighDefaultGcThresh2() *float64
	SetNetIpv4NeighDefaultGcThresh2(val *float64)
	NetIpv4NeighDefaultGcThresh2Input() *float64
	NetIpv4NeighDefaultGcThresh3() *float64
	SetNetIpv4NeighDefaultGcThresh3(val *float64)
	NetIpv4NeighDefaultGcThresh3Input() *float64
	NetIpv4TcpFinTimeout() *float64
	SetNetIpv4TcpFinTimeout(val *float64)
	NetIpv4TcpFinTimeoutInput() *float64
	NetIpv4TcpKeepaliveIntvl() *float64
	SetNetIpv4TcpKeepaliveIntvl(val *float64)
	NetIpv4TcpKeepaliveIntvlInput() *float64
	NetIpv4TcpKeepaliveProbes() *float64
	SetNetIpv4TcpKeepaliveProbes(val *float64)
	NetIpv4TcpKeepaliveProbesInput() *float64
	NetIpv4TcpKeepaliveTime() *float64
	SetNetIpv4TcpKeepaliveTime(val *float64)
	NetIpv4TcpKeepaliveTimeInput() *float64
	NetIpv4TcpMaxSynBacklog() *float64
	SetNetIpv4TcpMaxSynBacklog(val *float64)
	NetIpv4TcpMaxSynBacklogInput() *float64
	NetIpv4TcpMaxTwBuckets() *float64
	SetNetIpv4TcpMaxTwBuckets(val *float64)
	NetIpv4TcpMaxTwBucketsInput() *float64
	NetIpv4TcpTwReuse() interface{}
	SetNetIpv4TcpTwReuse(val interface{})
	NetIpv4TcpTwReuseInput() interface{}
	NetNetfilterNfConntrackBuckets() *float64
	SetNetNetfilterNfConntrackBuckets(val *float64)
	NetNetfilterNfConntrackBucketsInput() *float64
	NetNetfilterNfConntrackMax() *float64
	SetNetNetfilterNfConntrackMax(val *float64)
	NetNetfilterNfConntrackMaxInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VmMaxMapCount() *float64
	SetVmMaxMapCount(val *float64)
	VmMaxMapCountInput() *float64
	VmSwappiness() *float64
	SetVmSwappiness(val *float64)
	VmSwappinessInput() *float64
	VmVfsCachePressure() *float64
	SetVmVfsCachePressure(val *float64)
	VmVfsCachePressureInput() *float64
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetFsAioMaxNr()
	ResetFsFileMax()
	ResetFsInotifyMaxUserWatches()
	ResetFsNrOpen()
	ResetKernelThreadsMax()
	ResetNetCoreNetdevMaxBacklog()
	ResetNetCoreOptmemMax()
	ResetNetCoreRmemDefault()
	ResetNetCoreRmemMax()
	ResetNetCoreSomaxconn()
	ResetNetCoreWmemDefault()
	ResetNetCoreWmemMax()
	ResetNetIpv4IpLocalPortRangeMax()
	ResetNetIpv4IpLocalPortRangeMin()
	ResetNetIpv4NeighDefaultGcThresh1()
	ResetNetIpv4NeighDefaultGcThresh2()
	ResetNetIpv4NeighDefaultGcThresh3()
	ResetNetIpv4TcpFinTimeout()
	ResetNetIpv4TcpKeepaliveIntvl()
	ResetNetIpv4TcpKeepaliveProbes()
	ResetNetIpv4TcpKeepaliveTime()
	ResetNetIpv4TcpMaxSynBacklog()
	ResetNetIpv4TcpMaxTwBuckets()
	ResetNetIpv4TcpTwReuse()
	ResetNetNetfilterNfConntrackBuckets()
	ResetNetNetfilterNfConntrackMax()
	ResetVmMaxMapCount()
	ResetVmSwappiness()
	ResetVmVfsCachePressure()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference
type jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) FsAioMaxNr() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fsAioMaxNr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) FsAioMaxNrInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fsAioMaxNrInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) FsFileMax() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fsFileMax",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) FsFileMaxInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fsFileMaxInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) FsInotifyMaxUserWatches() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fsInotifyMaxUserWatches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) FsInotifyMaxUserWatchesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fsInotifyMaxUserWatchesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) FsNrOpen() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fsNrOpen",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) FsNrOpenInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fsNrOpenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) InternalValue() *KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfig {
	var returns *KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) KernelThreadsMax() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"kernelThreadsMax",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) KernelThreadsMaxInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"kernelThreadsMaxInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) NetCoreNetdevMaxBacklog() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netCoreNetdevMaxBacklog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) NetCoreNetdevMaxBacklogInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netCoreNetdevMaxBacklogInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) NetCoreOptmemMax() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netCoreOptmemMax",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) NetCoreOptmemMaxInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netCoreOptmemMaxInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) NetCoreRmemDefault() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netCoreRmemDefault",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) NetCoreRmemDefaultInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netCoreRmemDefaultInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) NetCoreRmemMax() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netCoreRmemMax",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) NetCoreRmemMaxInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netCoreRmemMaxInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) NetCoreSomaxconn() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netCoreSomaxconn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) NetCoreSomaxconnInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netCoreSomaxconnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) NetCoreWmemDefault() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netCoreWmemDefault",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) NetCoreWmemDefaultInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netCoreWmemDefaultInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) NetCoreWmemMax() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netCoreWmemMax",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) NetCoreWmemMaxInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netCoreWmemMaxInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) NetIpv4IpLocalPortRangeMax() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netIpv4IpLocalPortRangeMax",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) NetIpv4IpLocalPortRangeMaxInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netIpv4IpLocalPortRangeMaxInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) NetIpv4IpLocalPortRangeMin() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netIpv4IpLocalPortRangeMin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) NetIpv4IpLocalPortRangeMinInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netIpv4IpLocalPortRangeMinInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) NetIpv4NeighDefaultGcThresh1() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netIpv4NeighDefaultGcThresh1",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) NetIpv4NeighDefaultGcThresh1Input() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netIpv4NeighDefaultGcThresh1Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) NetIpv4NeighDefaultGcThresh2() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netIpv4NeighDefaultGcThresh2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) NetIpv4NeighDefaultGcThresh2Input() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netIpv4NeighDefaultGcThresh2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) NetIpv4NeighDefaultGcThresh3() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netIpv4NeighDefaultGcThresh3",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) NetIpv4NeighDefaultGcThresh3Input() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netIpv4NeighDefaultGcThresh3Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) NetIpv4TcpFinTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netIpv4TcpFinTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) NetIpv4TcpFinTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netIpv4TcpFinTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) NetIpv4TcpKeepaliveIntvl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netIpv4TcpKeepaliveIntvl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) NetIpv4TcpKeepaliveIntvlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netIpv4TcpKeepaliveIntvlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) NetIpv4TcpKeepaliveProbes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netIpv4TcpKeepaliveProbes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) NetIpv4TcpKeepaliveProbesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netIpv4TcpKeepaliveProbesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) NetIpv4TcpKeepaliveTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netIpv4TcpKeepaliveTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) NetIpv4TcpKeepaliveTimeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netIpv4TcpKeepaliveTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) NetIpv4TcpMaxSynBacklog() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netIpv4TcpMaxSynBacklog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) NetIpv4TcpMaxSynBacklogInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netIpv4TcpMaxSynBacklogInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) NetIpv4TcpMaxTwBuckets() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netIpv4TcpMaxTwBuckets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) NetIpv4TcpMaxTwBucketsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netIpv4TcpMaxTwBucketsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) NetIpv4TcpTwReuse() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"netIpv4TcpTwReuse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) NetIpv4TcpTwReuseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"netIpv4TcpTwReuseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) NetNetfilterNfConntrackBuckets() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netNetfilterNfConntrackBuckets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) NetNetfilterNfConntrackBucketsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netNetfilterNfConntrackBucketsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) NetNetfilterNfConntrackMax() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netNetfilterNfConntrackMax",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) NetNetfilterNfConntrackMaxInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netNetfilterNfConntrackMaxInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) VmMaxMapCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"vmMaxMapCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) VmMaxMapCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"vmMaxMapCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) VmSwappiness() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"vmSwappiness",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) VmSwappinessInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"vmSwappinessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) VmVfsCachePressure() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"vmVfsCachePressure",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) VmVfsCachePressureInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"vmVfsCachePressureInput",
		&returns,
	)
	return returns
}


func NewKubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference {
	_init_.Initialize()

	j := jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesCluster.KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference_Override(k KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesCluster.KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) SetFsAioMaxNr(val *float64) {
	_jsii_.Set(
		j,
		"fsAioMaxNr",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) SetFsFileMax(val *float64) {
	_jsii_.Set(
		j,
		"fsFileMax",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) SetFsInotifyMaxUserWatches(val *float64) {
	_jsii_.Set(
		j,
		"fsInotifyMaxUserWatches",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) SetFsNrOpen(val *float64) {
	_jsii_.Set(
		j,
		"fsNrOpen",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) SetInternalValue(val *KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfig) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) SetKernelThreadsMax(val *float64) {
	_jsii_.Set(
		j,
		"kernelThreadsMax",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) SetNetCoreNetdevMaxBacklog(val *float64) {
	_jsii_.Set(
		j,
		"netCoreNetdevMaxBacklog",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) SetNetCoreOptmemMax(val *float64) {
	_jsii_.Set(
		j,
		"netCoreOptmemMax",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) SetNetCoreRmemDefault(val *float64) {
	_jsii_.Set(
		j,
		"netCoreRmemDefault",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) SetNetCoreRmemMax(val *float64) {
	_jsii_.Set(
		j,
		"netCoreRmemMax",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) SetNetCoreSomaxconn(val *float64) {
	_jsii_.Set(
		j,
		"netCoreSomaxconn",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) SetNetCoreWmemDefault(val *float64) {
	_jsii_.Set(
		j,
		"netCoreWmemDefault",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) SetNetCoreWmemMax(val *float64) {
	_jsii_.Set(
		j,
		"netCoreWmemMax",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) SetNetIpv4IpLocalPortRangeMax(val *float64) {
	_jsii_.Set(
		j,
		"netIpv4IpLocalPortRangeMax",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) SetNetIpv4IpLocalPortRangeMin(val *float64) {
	_jsii_.Set(
		j,
		"netIpv4IpLocalPortRangeMin",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) SetNetIpv4NeighDefaultGcThresh1(val *float64) {
	_jsii_.Set(
		j,
		"netIpv4NeighDefaultGcThresh1",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) SetNetIpv4NeighDefaultGcThresh2(val *float64) {
	_jsii_.Set(
		j,
		"netIpv4NeighDefaultGcThresh2",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) SetNetIpv4NeighDefaultGcThresh3(val *float64) {
	_jsii_.Set(
		j,
		"netIpv4NeighDefaultGcThresh3",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) SetNetIpv4TcpFinTimeout(val *float64) {
	_jsii_.Set(
		j,
		"netIpv4TcpFinTimeout",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) SetNetIpv4TcpKeepaliveIntvl(val *float64) {
	_jsii_.Set(
		j,
		"netIpv4TcpKeepaliveIntvl",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) SetNetIpv4TcpKeepaliveProbes(val *float64) {
	_jsii_.Set(
		j,
		"netIpv4TcpKeepaliveProbes",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) SetNetIpv4TcpKeepaliveTime(val *float64) {
	_jsii_.Set(
		j,
		"netIpv4TcpKeepaliveTime",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) SetNetIpv4TcpMaxSynBacklog(val *float64) {
	_jsii_.Set(
		j,
		"netIpv4TcpMaxSynBacklog",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) SetNetIpv4TcpMaxTwBuckets(val *float64) {
	_jsii_.Set(
		j,
		"netIpv4TcpMaxTwBuckets",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) SetNetIpv4TcpTwReuse(val interface{}) {
	_jsii_.Set(
		j,
		"netIpv4TcpTwReuse",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) SetNetNetfilterNfConntrackBuckets(val *float64) {
	_jsii_.Set(
		j,
		"netNetfilterNfConntrackBuckets",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) SetNetNetfilterNfConntrackMax(val *float64) {
	_jsii_.Set(
		j,
		"netNetfilterNfConntrackMax",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) SetVmMaxMapCount(val *float64) {
	_jsii_.Set(
		j,
		"vmMaxMapCount",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) SetVmSwappiness(val *float64) {
	_jsii_.Set(
		j,
		"vmSwappiness",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) SetVmVfsCachePressure(val *float64) {
	_jsii_.Set(
		j,
		"vmVfsCachePressure",
		val,
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		k,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		k,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		k,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		k,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		k,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) ResetFsAioMaxNr() {
	_jsii_.InvokeVoid(
		k,
		"resetFsAioMaxNr",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) ResetFsFileMax() {
	_jsii_.InvokeVoid(
		k,
		"resetFsFileMax",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) ResetFsInotifyMaxUserWatches() {
	_jsii_.InvokeVoid(
		k,
		"resetFsInotifyMaxUserWatches",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) ResetFsNrOpen() {
	_jsii_.InvokeVoid(
		k,
		"resetFsNrOpen",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) ResetKernelThreadsMax() {
	_jsii_.InvokeVoid(
		k,
		"resetKernelThreadsMax",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) ResetNetCoreNetdevMaxBacklog() {
	_jsii_.InvokeVoid(
		k,
		"resetNetCoreNetdevMaxBacklog",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) ResetNetCoreOptmemMax() {
	_jsii_.InvokeVoid(
		k,
		"resetNetCoreOptmemMax",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) ResetNetCoreRmemDefault() {
	_jsii_.InvokeVoid(
		k,
		"resetNetCoreRmemDefault",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) ResetNetCoreRmemMax() {
	_jsii_.InvokeVoid(
		k,
		"resetNetCoreRmemMax",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) ResetNetCoreSomaxconn() {
	_jsii_.InvokeVoid(
		k,
		"resetNetCoreSomaxconn",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) ResetNetCoreWmemDefault() {
	_jsii_.InvokeVoid(
		k,
		"resetNetCoreWmemDefault",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) ResetNetCoreWmemMax() {
	_jsii_.InvokeVoid(
		k,
		"resetNetCoreWmemMax",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) ResetNetIpv4IpLocalPortRangeMax() {
	_jsii_.InvokeVoid(
		k,
		"resetNetIpv4IpLocalPortRangeMax",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) ResetNetIpv4IpLocalPortRangeMin() {
	_jsii_.InvokeVoid(
		k,
		"resetNetIpv4IpLocalPortRangeMin",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) ResetNetIpv4NeighDefaultGcThresh1() {
	_jsii_.InvokeVoid(
		k,
		"resetNetIpv4NeighDefaultGcThresh1",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) ResetNetIpv4NeighDefaultGcThresh2() {
	_jsii_.InvokeVoid(
		k,
		"resetNetIpv4NeighDefaultGcThresh2",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) ResetNetIpv4NeighDefaultGcThresh3() {
	_jsii_.InvokeVoid(
		k,
		"resetNetIpv4NeighDefaultGcThresh3",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) ResetNetIpv4TcpFinTimeout() {
	_jsii_.InvokeVoid(
		k,
		"resetNetIpv4TcpFinTimeout",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) ResetNetIpv4TcpKeepaliveIntvl() {
	_jsii_.InvokeVoid(
		k,
		"resetNetIpv4TcpKeepaliveIntvl",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) ResetNetIpv4TcpKeepaliveProbes() {
	_jsii_.InvokeVoid(
		k,
		"resetNetIpv4TcpKeepaliveProbes",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) ResetNetIpv4TcpKeepaliveTime() {
	_jsii_.InvokeVoid(
		k,
		"resetNetIpv4TcpKeepaliveTime",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) ResetNetIpv4TcpMaxSynBacklog() {
	_jsii_.InvokeVoid(
		k,
		"resetNetIpv4TcpMaxSynBacklog",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) ResetNetIpv4TcpMaxTwBuckets() {
	_jsii_.InvokeVoid(
		k,
		"resetNetIpv4TcpMaxTwBuckets",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) ResetNetIpv4TcpTwReuse() {
	_jsii_.InvokeVoid(
		k,
		"resetNetIpv4TcpTwReuse",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) ResetNetNetfilterNfConntrackBuckets() {
	_jsii_.InvokeVoid(
		k,
		"resetNetNetfilterNfConntrackBuckets",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) ResetNetNetfilterNfConntrackMax() {
	_jsii_.InvokeVoid(
		k,
		"resetNetNetfilterNfConntrackMax",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) ResetVmMaxMapCount() {
	_jsii_.InvokeVoid(
		k,
		"resetVmMaxMapCount",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) ResetVmSwappiness() {
	_jsii_.InvokeVoid(
		k,
		"resetVmSwappiness",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) ResetVmVfsCachePressure() {
	_jsii_.InvokeVoid(
		k,
		"resetVmVfsCachePressure",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type KubernetesClusterDefaultNodePoolOutputReference interface {
	cdktf.ComplexObject
	CapacityReservationGroupId() *string
	SetCapacityReservationGroupId(val *string)
	CapacityReservationGroupIdInput() *string
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	EnableAutoScaling() interface{}
	SetEnableAutoScaling(val interface{})
	EnableAutoScalingInput() interface{}
	EnableHostEncryption() interface{}
	SetEnableHostEncryption(val interface{})
	EnableHostEncryptionInput() interface{}
	EnableNodePublicIp() interface{}
	SetEnableNodePublicIp(val interface{})
	EnableNodePublicIpInput() interface{}
	FipsEnabled() interface{}
	SetFipsEnabled(val interface{})
	FipsEnabledInput() interface{}
	// Experimental.
	Fqn() *string
	HostGroupId() *string
	SetHostGroupId(val *string)
	HostGroupIdInput() *string
	InternalValue() *KubernetesClusterDefaultNodePool
	SetInternalValue(val *KubernetesClusterDefaultNodePool)
	KubeletConfig() KubernetesClusterDefaultNodePoolKubeletConfigOutputReference
	KubeletConfigInput() *KubernetesClusterDefaultNodePoolKubeletConfig
	KubeletDiskType() *string
	SetKubeletDiskType(val *string)
	KubeletDiskTypeInput() *string
	LinuxOsConfig() KubernetesClusterDefaultNodePoolLinuxOsConfigOutputReference
	LinuxOsConfigInput() *KubernetesClusterDefaultNodePoolLinuxOsConfig
	MaxCount() *float64
	SetMaxCount(val *float64)
	MaxCountInput() *float64
	MaxPods() *float64
	SetMaxPods(val *float64)
	MaxPodsInput() *float64
	MinCount() *float64
	SetMinCount(val *float64)
	MinCountInput() *float64
	Name() *string
	SetName(val *string)
	NameInput() *string
	NodeCount() *float64
	SetNodeCount(val *float64)
	NodeCountInput() *float64
	NodeLabels() *map[string]*string
	SetNodeLabels(val *map[string]*string)
	NodeLabelsInput() *map[string]*string
	NodePublicIpPrefixId() *string
	SetNodePublicIpPrefixId(val *string)
	NodePublicIpPrefixIdInput() *string
	NodeTaints() *[]*string
	SetNodeTaints(val *[]*string)
	NodeTaintsInput() *[]*string
	OnlyCriticalAddonsEnabled() interface{}
	SetOnlyCriticalAddonsEnabled(val interface{})
	OnlyCriticalAddonsEnabledInput() interface{}
	OrchestratorVersion() *string
	SetOrchestratorVersion(val *string)
	OrchestratorVersionInput() *string
	OsDiskSizeGb() *float64
	SetOsDiskSizeGb(val *float64)
	OsDiskSizeGbInput() *float64
	OsDiskType() *string
	SetOsDiskType(val *string)
	OsDiskTypeInput() *string
	OsSku() *string
	SetOsSku(val *string)
	OsSkuInput() *string
	PodSubnetId() *string
	SetPodSubnetId(val *string)
	PodSubnetIdInput() *string
	ProximityPlacementGroupId() *string
	SetProximityPlacementGroupId(val *string)
	ProximityPlacementGroupIdInput() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Type() *string
	SetType(val *string)
	TypeInput() *string
	UltraSsdEnabled() interface{}
	SetUltraSsdEnabled(val interface{})
	UltraSsdEnabledInput() interface{}
	UpgradeSettings() KubernetesClusterDefaultNodePoolUpgradeSettingsOutputReference
	UpgradeSettingsInput() *KubernetesClusterDefaultNodePoolUpgradeSettings
	VmSize() *string
	SetVmSize(val *string)
	VmSizeInput() *string
	VnetSubnetId() *string
	SetVnetSubnetId(val *string)
	VnetSubnetIdInput() *string
	Zones() *[]*string
	SetZones(val *[]*string)
	ZonesInput() *[]*string
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	PutKubeletConfig(value *KubernetesClusterDefaultNodePoolKubeletConfig)
	PutLinuxOsConfig(value *KubernetesClusterDefaultNodePoolLinuxOsConfig)
	PutUpgradeSettings(value *KubernetesClusterDefaultNodePoolUpgradeSettings)
	ResetCapacityReservationGroupId()
	ResetEnableAutoScaling()
	ResetEnableHostEncryption()
	ResetEnableNodePublicIp()
	ResetFipsEnabled()
	ResetHostGroupId()
	ResetKubeletConfig()
	ResetKubeletDiskType()
	ResetLinuxOsConfig()
	ResetMaxCount()
	ResetMaxPods()
	ResetMinCount()
	ResetNodeCount()
	ResetNodeLabels()
	ResetNodePublicIpPrefixId()
	ResetNodeTaints()
	ResetOnlyCriticalAddonsEnabled()
	ResetOrchestratorVersion()
	ResetOsDiskSizeGb()
	ResetOsDiskType()
	ResetOsSku()
	ResetPodSubnetId()
	ResetProximityPlacementGroupId()
	ResetTags()
	ResetType()
	ResetUltraSsdEnabled()
	ResetUpgradeSettings()
	ResetVnetSubnetId()
	ResetZones()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KubernetesClusterDefaultNodePoolOutputReference
type jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) CapacityReservationGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"capacityReservationGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) CapacityReservationGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"capacityReservationGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) EnableAutoScaling() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableAutoScaling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) EnableAutoScalingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableAutoScalingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) EnableHostEncryption() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableHostEncryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) EnableHostEncryptionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableHostEncryptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) EnableNodePublicIp() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableNodePublicIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) EnableNodePublicIpInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableNodePublicIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) FipsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fipsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) FipsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fipsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) HostGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) HostGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) InternalValue() *KubernetesClusterDefaultNodePool {
	var returns *KubernetesClusterDefaultNodePool
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) KubeletConfig() KubernetesClusterDefaultNodePoolKubeletConfigOutputReference {
	var returns KubernetesClusterDefaultNodePoolKubeletConfigOutputReference
	_jsii_.Get(
		j,
		"kubeletConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) KubeletConfigInput() *KubernetesClusterDefaultNodePoolKubeletConfig {
	var returns *KubernetesClusterDefaultNodePoolKubeletConfig
	_jsii_.Get(
		j,
		"kubeletConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) KubeletDiskType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kubeletDiskType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) KubeletDiskTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kubeletDiskTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) LinuxOsConfig() KubernetesClusterDefaultNodePoolLinuxOsConfigOutputReference {
	var returns KubernetesClusterDefaultNodePoolLinuxOsConfigOutputReference
	_jsii_.Get(
		j,
		"linuxOsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) LinuxOsConfigInput() *KubernetesClusterDefaultNodePoolLinuxOsConfig {
	var returns *KubernetesClusterDefaultNodePoolLinuxOsConfig
	_jsii_.Get(
		j,
		"linuxOsConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) MaxCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) MaxCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) MaxPods() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxPods",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) MaxPodsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxPodsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) MinCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) MinCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) NodeCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nodeCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) NodeCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nodeCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) NodeLabels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"nodeLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) NodeLabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"nodeLabelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) NodePublicIpPrefixId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodePublicIpPrefixId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) NodePublicIpPrefixIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodePublicIpPrefixIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) NodeTaints() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"nodeTaints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) NodeTaintsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"nodeTaintsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) OnlyCriticalAddonsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"onlyCriticalAddonsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) OnlyCriticalAddonsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"onlyCriticalAddonsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) OrchestratorVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orchestratorVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) OrchestratorVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orchestratorVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) OsDiskSizeGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"osDiskSizeGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) OsDiskSizeGbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"osDiskSizeGbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) OsDiskType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"osDiskType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) OsDiskTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"osDiskTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) OsSku() *string {
	var returns *string
	_jsii_.Get(
		j,
		"osSku",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) OsSkuInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"osSkuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) PodSubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"podSubnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) PodSubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"podSubnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) ProximityPlacementGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"proximityPlacementGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) ProximityPlacementGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"proximityPlacementGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) UltraSsdEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ultraSsdEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) UltraSsdEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ultraSsdEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) UpgradeSettings() KubernetesClusterDefaultNodePoolUpgradeSettingsOutputReference {
	var returns KubernetesClusterDefaultNodePoolUpgradeSettingsOutputReference
	_jsii_.Get(
		j,
		"upgradeSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) UpgradeSettingsInput() *KubernetesClusterDefaultNodePoolUpgradeSettings {
	var returns *KubernetesClusterDefaultNodePoolUpgradeSettings
	_jsii_.Get(
		j,
		"upgradeSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) VmSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vmSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) VmSizeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vmSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) VnetSubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vnetSubnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) VnetSubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vnetSubnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) Zones() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"zones",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) ZonesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"zonesInput",
		&returns,
	)
	return returns
}


func NewKubernetesClusterDefaultNodePoolOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) KubernetesClusterDefaultNodePoolOutputReference {
	_init_.Initialize()

	j := jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesCluster.KubernetesClusterDefaultNodePoolOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKubernetesClusterDefaultNodePoolOutputReference_Override(k KubernetesClusterDefaultNodePoolOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesCluster.KubernetesClusterDefaultNodePoolOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) SetCapacityReservationGroupId(val *string) {
	_jsii_.Set(
		j,
		"capacityReservationGroupId",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) SetEnableAutoScaling(val interface{}) {
	_jsii_.Set(
		j,
		"enableAutoScaling",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) SetEnableHostEncryption(val interface{}) {
	_jsii_.Set(
		j,
		"enableHostEncryption",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) SetEnableNodePublicIp(val interface{}) {
	_jsii_.Set(
		j,
		"enableNodePublicIp",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) SetFipsEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"fipsEnabled",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) SetHostGroupId(val *string) {
	_jsii_.Set(
		j,
		"hostGroupId",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) SetInternalValue(val *KubernetesClusterDefaultNodePool) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) SetKubeletDiskType(val *string) {
	_jsii_.Set(
		j,
		"kubeletDiskType",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) SetMaxCount(val *float64) {
	_jsii_.Set(
		j,
		"maxCount",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) SetMaxPods(val *float64) {
	_jsii_.Set(
		j,
		"maxPods",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) SetMinCount(val *float64) {
	_jsii_.Set(
		j,
		"minCount",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) SetNodeCount(val *float64) {
	_jsii_.Set(
		j,
		"nodeCount",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) SetNodeLabels(val *map[string]*string) {
	_jsii_.Set(
		j,
		"nodeLabels",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) SetNodePublicIpPrefixId(val *string) {
	_jsii_.Set(
		j,
		"nodePublicIpPrefixId",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) SetNodeTaints(val *[]*string) {
	_jsii_.Set(
		j,
		"nodeTaints",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) SetOnlyCriticalAddonsEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"onlyCriticalAddonsEnabled",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) SetOrchestratorVersion(val *string) {
	_jsii_.Set(
		j,
		"orchestratorVersion",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) SetOsDiskSizeGb(val *float64) {
	_jsii_.Set(
		j,
		"osDiskSizeGb",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) SetOsDiskType(val *string) {
	_jsii_.Set(
		j,
		"osDiskType",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) SetOsSku(val *string) {
	_jsii_.Set(
		j,
		"osSku",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) SetPodSubnetId(val *string) {
	_jsii_.Set(
		j,
		"podSubnetId",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) SetProximityPlacementGroupId(val *string) {
	_jsii_.Set(
		j,
		"proximityPlacementGroupId",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) SetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) SetType(val *string) {
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) SetUltraSsdEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"ultraSsdEnabled",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) SetVmSize(val *string) {
	_jsii_.Set(
		j,
		"vmSize",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) SetVnetSubnetId(val *string) {
	_jsii_.Set(
		j,
		"vnetSubnetId",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) SetZones(val *[]*string) {
	_jsii_.Set(
		j,
		"zones",
		val,
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		k,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		k,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		k,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		k,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		k,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) PutKubeletConfig(value *KubernetesClusterDefaultNodePoolKubeletConfig) {
	_jsii_.InvokeVoid(
		k,
		"putKubeletConfig",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) PutLinuxOsConfig(value *KubernetesClusterDefaultNodePoolLinuxOsConfig) {
	_jsii_.InvokeVoid(
		k,
		"putLinuxOsConfig",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) PutUpgradeSettings(value *KubernetesClusterDefaultNodePoolUpgradeSettings) {
	_jsii_.InvokeVoid(
		k,
		"putUpgradeSettings",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) ResetCapacityReservationGroupId() {
	_jsii_.InvokeVoid(
		k,
		"resetCapacityReservationGroupId",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) ResetEnableAutoScaling() {
	_jsii_.InvokeVoid(
		k,
		"resetEnableAutoScaling",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) ResetEnableHostEncryption() {
	_jsii_.InvokeVoid(
		k,
		"resetEnableHostEncryption",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) ResetEnableNodePublicIp() {
	_jsii_.InvokeVoid(
		k,
		"resetEnableNodePublicIp",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) ResetFipsEnabled() {
	_jsii_.InvokeVoid(
		k,
		"resetFipsEnabled",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) ResetHostGroupId() {
	_jsii_.InvokeVoid(
		k,
		"resetHostGroupId",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) ResetKubeletConfig() {
	_jsii_.InvokeVoid(
		k,
		"resetKubeletConfig",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) ResetKubeletDiskType() {
	_jsii_.InvokeVoid(
		k,
		"resetKubeletDiskType",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) ResetLinuxOsConfig() {
	_jsii_.InvokeVoid(
		k,
		"resetLinuxOsConfig",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) ResetMaxCount() {
	_jsii_.InvokeVoid(
		k,
		"resetMaxCount",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) ResetMaxPods() {
	_jsii_.InvokeVoid(
		k,
		"resetMaxPods",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) ResetMinCount() {
	_jsii_.InvokeVoid(
		k,
		"resetMinCount",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) ResetNodeCount() {
	_jsii_.InvokeVoid(
		k,
		"resetNodeCount",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) ResetNodeLabels() {
	_jsii_.InvokeVoid(
		k,
		"resetNodeLabels",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) ResetNodePublicIpPrefixId() {
	_jsii_.InvokeVoid(
		k,
		"resetNodePublicIpPrefixId",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) ResetNodeTaints() {
	_jsii_.InvokeVoid(
		k,
		"resetNodeTaints",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) ResetOnlyCriticalAddonsEnabled() {
	_jsii_.InvokeVoid(
		k,
		"resetOnlyCriticalAddonsEnabled",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) ResetOrchestratorVersion() {
	_jsii_.InvokeVoid(
		k,
		"resetOrchestratorVersion",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) ResetOsDiskSizeGb() {
	_jsii_.InvokeVoid(
		k,
		"resetOsDiskSizeGb",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) ResetOsDiskType() {
	_jsii_.InvokeVoid(
		k,
		"resetOsDiskType",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) ResetOsSku() {
	_jsii_.InvokeVoid(
		k,
		"resetOsSku",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) ResetPodSubnetId() {
	_jsii_.InvokeVoid(
		k,
		"resetPodSubnetId",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) ResetProximityPlacementGroupId() {
	_jsii_.InvokeVoid(
		k,
		"resetProximityPlacementGroupId",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) ResetTags() {
	_jsii_.InvokeVoid(
		k,
		"resetTags",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) ResetType() {
	_jsii_.InvokeVoid(
		k,
		"resetType",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) ResetUltraSsdEnabled() {
	_jsii_.InvokeVoid(
		k,
		"resetUltraSsdEnabled",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) ResetUpgradeSettings() {
	_jsii_.InvokeVoid(
		k,
		"resetUpgradeSettings",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) ResetVnetSubnetId() {
	_jsii_.InvokeVoid(
		k,
		"resetVnetSubnetId",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) ResetZones() {
	_jsii_.InvokeVoid(
		k,
		"resetZones",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type KubernetesClusterDefaultNodePoolUpgradeSettings struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#max_surge KubernetesCluster#max_surge}.
	MaxSurge *string `field:"required" json:"maxSurge" yaml:"maxSurge"`
}

type KubernetesClusterDefaultNodePoolUpgradeSettingsOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *KubernetesClusterDefaultNodePoolUpgradeSettings
	SetInternalValue(val *KubernetesClusterDefaultNodePoolUpgradeSettings)
	MaxSurge() *string
	SetMaxSurge(val *string)
	MaxSurgeInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KubernetesClusterDefaultNodePoolUpgradeSettingsOutputReference
type jsiiProxy_KubernetesClusterDefaultNodePoolUpgradeSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolUpgradeSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolUpgradeSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolUpgradeSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolUpgradeSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolUpgradeSettingsOutputReference) InternalValue() *KubernetesClusterDefaultNodePoolUpgradeSettings {
	var returns *KubernetesClusterDefaultNodePoolUpgradeSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolUpgradeSettingsOutputReference) MaxSurge() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxSurge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolUpgradeSettingsOutputReference) MaxSurgeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxSurgeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolUpgradeSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolUpgradeSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewKubernetesClusterDefaultNodePoolUpgradeSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) KubernetesClusterDefaultNodePoolUpgradeSettingsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_KubernetesClusterDefaultNodePoolUpgradeSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesCluster.KubernetesClusterDefaultNodePoolUpgradeSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKubernetesClusterDefaultNodePoolUpgradeSettingsOutputReference_Override(k KubernetesClusterDefaultNodePoolUpgradeSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesCluster.KubernetesClusterDefaultNodePoolUpgradeSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolUpgradeSettingsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolUpgradeSettingsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolUpgradeSettingsOutputReference) SetInternalValue(val *KubernetesClusterDefaultNodePoolUpgradeSettings) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolUpgradeSettingsOutputReference) SetMaxSurge(val *string) {
	_jsii_.Set(
		j,
		"maxSurge",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolUpgradeSettingsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolUpgradeSettingsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolUpgradeSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolUpgradeSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolUpgradeSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolUpgradeSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		k,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolUpgradeSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolUpgradeSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		k,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolUpgradeSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		k,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolUpgradeSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		k,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolUpgradeSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolUpgradeSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		k,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolUpgradeSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolUpgradeSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolUpgradeSettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolUpgradeSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type KubernetesClusterHttpProxyConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#http_proxy KubernetesCluster#http_proxy}.
	HttpProxy *string `field:"optional" json:"httpProxy" yaml:"httpProxy"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#https_proxy KubernetesCluster#https_proxy}.
	HttpsProxy *string `field:"optional" json:"httpsProxy" yaml:"httpsProxy"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#no_proxy KubernetesCluster#no_proxy}.
	NoProxy *[]*string `field:"optional" json:"noProxy" yaml:"noProxy"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#trusted_ca KubernetesCluster#trusted_ca}.
	TrustedCa *string `field:"optional" json:"trustedCa" yaml:"trustedCa"`
}

type KubernetesClusterHttpProxyConfigOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	HttpProxy() *string
	SetHttpProxy(val *string)
	HttpProxyInput() *string
	HttpsProxy() *string
	SetHttpsProxy(val *string)
	HttpsProxyInput() *string
	InternalValue() *KubernetesClusterHttpProxyConfig
	SetInternalValue(val *KubernetesClusterHttpProxyConfig)
	NoProxy() *[]*string
	SetNoProxy(val *[]*string)
	NoProxyInput() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TrustedCa() *string
	SetTrustedCa(val *string)
	TrustedCaInput() *string
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetHttpProxy()
	ResetHttpsProxy()
	ResetNoProxy()
	ResetTrustedCa()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KubernetesClusterHttpProxyConfigOutputReference
type jsiiProxy_KubernetesClusterHttpProxyConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_KubernetesClusterHttpProxyConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterHttpProxyConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterHttpProxyConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterHttpProxyConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterHttpProxyConfigOutputReference) HttpProxy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpProxy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterHttpProxyConfigOutputReference) HttpProxyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpProxyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterHttpProxyConfigOutputReference) HttpsProxy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpsProxy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterHttpProxyConfigOutputReference) HttpsProxyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpsProxyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterHttpProxyConfigOutputReference) InternalValue() *KubernetesClusterHttpProxyConfig {
	var returns *KubernetesClusterHttpProxyConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterHttpProxyConfigOutputReference) NoProxy() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"noProxy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterHttpProxyConfigOutputReference) NoProxyInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"noProxyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterHttpProxyConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterHttpProxyConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterHttpProxyConfigOutputReference) TrustedCa() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trustedCa",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterHttpProxyConfigOutputReference) TrustedCaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trustedCaInput",
		&returns,
	)
	return returns
}


func NewKubernetesClusterHttpProxyConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) KubernetesClusterHttpProxyConfigOutputReference {
	_init_.Initialize()

	j := jsiiProxy_KubernetesClusterHttpProxyConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesCluster.KubernetesClusterHttpProxyConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKubernetesClusterHttpProxyConfigOutputReference_Override(k KubernetesClusterHttpProxyConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesCluster.KubernetesClusterHttpProxyConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_KubernetesClusterHttpProxyConfigOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterHttpProxyConfigOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterHttpProxyConfigOutputReference) SetHttpProxy(val *string) {
	_jsii_.Set(
		j,
		"httpProxy",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterHttpProxyConfigOutputReference) SetHttpsProxy(val *string) {
	_jsii_.Set(
		j,
		"httpsProxy",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterHttpProxyConfigOutputReference) SetInternalValue(val *KubernetesClusterHttpProxyConfig) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterHttpProxyConfigOutputReference) SetNoProxy(val *[]*string) {
	_jsii_.Set(
		j,
		"noProxy",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterHttpProxyConfigOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterHttpProxyConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterHttpProxyConfigOutputReference) SetTrustedCa(val *string) {
	_jsii_.Set(
		j,
		"trustedCa",
		val,
	)
}

func (k *jsiiProxy_KubernetesClusterHttpProxyConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterHttpProxyConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterHttpProxyConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterHttpProxyConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		k,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterHttpProxyConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterHttpProxyConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		k,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterHttpProxyConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		k,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterHttpProxyConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		k,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterHttpProxyConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterHttpProxyConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		k,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterHttpProxyConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterHttpProxyConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterHttpProxyConfigOutputReference) ResetHttpProxy() {
	_jsii_.InvokeVoid(
		k,
		"resetHttpProxy",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterHttpProxyConfigOutputReference) ResetHttpsProxy() {
	_jsii_.InvokeVoid(
		k,
		"resetHttpsProxy",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterHttpProxyConfigOutputReference) ResetNoProxy() {
	_jsii_.InvokeVoid(
		k,
		"resetNoProxy",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterHttpProxyConfigOutputReference) ResetTrustedCa() {
	_jsii_.InvokeVoid(
		k,
		"resetTrustedCa",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterHttpProxyConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterHttpProxyConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type KubernetesClusterIdentity struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#type KubernetesCluster#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#identity_ids KubernetesCluster#identity_ids}.
	IdentityIds *[]*string `field:"optional" json:"identityIds" yaml:"identityIds"`
}

type KubernetesClusterIdentityOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	IdentityIds() *[]*string
	SetIdentityIds(val *[]*string)
	IdentityIdsInput() *[]*string
	InternalValue() *KubernetesClusterIdentity
	SetInternalValue(val *KubernetesClusterIdentity)
	PrincipalId() *string
	TenantId() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Type() *string
	SetType(val *string)
	TypeInput() *string
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetIdentityIds()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KubernetesClusterIdentityOutputReference
type jsiiProxy_KubernetesClusterIdentityOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_KubernetesClusterIdentityOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterIdentityOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterIdentityOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterIdentityOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterIdentityOutputReference) IdentityIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"identityIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterIdentityOutputReference) IdentityIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"identityIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterIdentityOutputReference) InternalValue() *KubernetesClusterIdentity {
	var returns *KubernetesClusterIdentity
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterIdentityOutputReference) PrincipalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"principalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterIdentityOutputReference) TenantId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenantId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterIdentityOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterIdentityOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterIdentityOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterIdentityOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewKubernetesClusterIdentityOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) KubernetesClusterIdentityOutputReference {
	_init_.Initialize()

	j := jsiiProxy_KubernetesClusterIdentityOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesCluster.KubernetesClusterIdentityOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKubernetesClusterIdentityOutputReference_Override(k KubernetesClusterIdentityOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesCluster.KubernetesClusterIdentityOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_KubernetesClusterIdentityOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterIdentityOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterIdentityOutputReference) SetIdentityIds(val *[]*string) {
	_jsii_.Set(
		j,
		"identityIds",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterIdentityOutputReference) SetInternalValue(val *KubernetesClusterIdentity) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterIdentityOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterIdentityOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterIdentityOutputReference) SetType(val *string) {
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (k *jsiiProxy_KubernetesClusterIdentityOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterIdentityOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterIdentityOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterIdentityOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		k,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterIdentityOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterIdentityOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		k,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterIdentityOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		k,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterIdentityOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		k,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterIdentityOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterIdentityOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		k,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterIdentityOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterIdentityOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterIdentityOutputReference) ResetIdentityIds() {
	_jsii_.InvokeVoid(
		k,
		"resetIdentityIds",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterIdentityOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterIdentityOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type KubernetesClusterIngressApplicationGateway struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#gateway_id KubernetesCluster#gateway_id}.
	GatewayId *string `field:"optional" json:"gatewayId" yaml:"gatewayId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#gateway_name KubernetesCluster#gateway_name}.
	GatewayName *string `field:"optional" json:"gatewayName" yaml:"gatewayName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#subnet_cidr KubernetesCluster#subnet_cidr}.
	SubnetCidr *string `field:"optional" json:"subnetCidr" yaml:"subnetCidr"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#subnet_id KubernetesCluster#subnet_id}.
	SubnetId *string `field:"optional" json:"subnetId" yaml:"subnetId"`
}

type KubernetesClusterIngressApplicationGatewayIngressApplicationGatewayIdentity struct {
}

type KubernetesClusterIngressApplicationGatewayIngressApplicationGatewayIdentityList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) KubernetesClusterIngressApplicationGatewayIngressApplicationGatewayIdentityOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KubernetesClusterIngressApplicationGatewayIngressApplicationGatewayIdentityList
type jsiiProxy_KubernetesClusterIngressApplicationGatewayIngressApplicationGatewayIdentityList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_KubernetesClusterIngressApplicationGatewayIngressApplicationGatewayIdentityList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterIngressApplicationGatewayIngressApplicationGatewayIdentityList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterIngressApplicationGatewayIngressApplicationGatewayIdentityList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterIngressApplicationGatewayIngressApplicationGatewayIdentityList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterIngressApplicationGatewayIngressApplicationGatewayIdentityList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewKubernetesClusterIngressApplicationGatewayIngressApplicationGatewayIdentityList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) KubernetesClusterIngressApplicationGatewayIngressApplicationGatewayIdentityList {
	_init_.Initialize()

	j := jsiiProxy_KubernetesClusterIngressApplicationGatewayIngressApplicationGatewayIdentityList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesCluster.KubernetesClusterIngressApplicationGatewayIngressApplicationGatewayIdentityList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewKubernetesClusterIngressApplicationGatewayIngressApplicationGatewayIdentityList_Override(k KubernetesClusterIngressApplicationGatewayIngressApplicationGatewayIdentityList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesCluster.KubernetesClusterIngressApplicationGatewayIngressApplicationGatewayIdentityList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		k,
	)
}

func (j *jsiiProxy_KubernetesClusterIngressApplicationGatewayIngressApplicationGatewayIdentityList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterIngressApplicationGatewayIngressApplicationGatewayIdentityList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterIngressApplicationGatewayIngressApplicationGatewayIdentityList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (k *jsiiProxy_KubernetesClusterIngressApplicationGatewayIngressApplicationGatewayIdentityList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterIngressApplicationGatewayIngressApplicationGatewayIdentityList) Get(index *float64) KubernetesClusterIngressApplicationGatewayIngressApplicationGatewayIdentityOutputReference {
	var returns KubernetesClusterIngressApplicationGatewayIngressApplicationGatewayIdentityOutputReference

	_jsii_.Invoke(
		k,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterIngressApplicationGatewayIngressApplicationGatewayIdentityList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterIngressApplicationGatewayIngressApplicationGatewayIdentityList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type KubernetesClusterIngressApplicationGatewayIngressApplicationGatewayIdentityOutputReference interface {
	cdktf.ComplexObject
	ClientId() *string
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *KubernetesClusterIngressApplicationGatewayIngressApplicationGatewayIdentity
	SetInternalValue(val *KubernetesClusterIngressApplicationGatewayIngressApplicationGatewayIdentity)
	ObjectId() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UserAssignedIdentityId() *string
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KubernetesClusterIngressApplicationGatewayIngressApplicationGatewayIdentityOutputReference
type jsiiProxy_KubernetesClusterIngressApplicationGatewayIngressApplicationGatewayIdentityOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_KubernetesClusterIngressApplicationGatewayIngressApplicationGatewayIdentityOutputReference) ClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterIngressApplicationGatewayIngressApplicationGatewayIdentityOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterIngressApplicationGatewayIngressApplicationGatewayIdentityOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterIngressApplicationGatewayIngressApplicationGatewayIdentityOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterIngressApplicationGatewayIngressApplicationGatewayIdentityOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterIngressApplicationGatewayIngressApplicationGatewayIdentityOutputReference) InternalValue() *KubernetesClusterIngressApplicationGatewayIngressApplicationGatewayIdentity {
	var returns *KubernetesClusterIngressApplicationGatewayIngressApplicationGatewayIdentity
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterIngressApplicationGatewayIngressApplicationGatewayIdentityOutputReference) ObjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterIngressApplicationGatewayIngressApplicationGatewayIdentityOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterIngressApplicationGatewayIngressApplicationGatewayIdentityOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterIngressApplicationGatewayIngressApplicationGatewayIdentityOutputReference) UserAssignedIdentityId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userAssignedIdentityId",
		&returns,
	)
	return returns
}


func NewKubernetesClusterIngressApplicationGatewayIngressApplicationGatewayIdentityOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) KubernetesClusterIngressApplicationGatewayIngressApplicationGatewayIdentityOutputReference {
	_init_.Initialize()

	j := jsiiProxy_KubernetesClusterIngressApplicationGatewayIngressApplicationGatewayIdentityOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesCluster.KubernetesClusterIngressApplicationGatewayIngressApplicationGatewayIdentityOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewKubernetesClusterIngressApplicationGatewayIngressApplicationGatewayIdentityOutputReference_Override(k KubernetesClusterIngressApplicationGatewayIngressApplicationGatewayIdentityOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesCluster.KubernetesClusterIngressApplicationGatewayIngressApplicationGatewayIdentityOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		k,
	)
}

func (j *jsiiProxy_KubernetesClusterIngressApplicationGatewayIngressApplicationGatewayIdentityOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterIngressApplicationGatewayIngressApplicationGatewayIdentityOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterIngressApplicationGatewayIngressApplicationGatewayIdentityOutputReference) SetInternalValue(val *KubernetesClusterIngressApplicationGatewayIngressApplicationGatewayIdentity) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterIngressApplicationGatewayIngressApplicationGatewayIdentityOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterIngressApplicationGatewayIngressApplicationGatewayIdentityOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (k *jsiiProxy_KubernetesClusterIngressApplicationGatewayIngressApplicationGatewayIdentityOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterIngressApplicationGatewayIngressApplicationGatewayIdentityOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterIngressApplicationGatewayIngressApplicationGatewayIdentityOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterIngressApplicationGatewayIngressApplicationGatewayIdentityOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		k,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterIngressApplicationGatewayIngressApplicationGatewayIdentityOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterIngressApplicationGatewayIngressApplicationGatewayIdentityOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		k,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterIngressApplicationGatewayIngressApplicationGatewayIdentityOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		k,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterIngressApplicationGatewayIngressApplicationGatewayIdentityOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		k,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterIngressApplicationGatewayIngressApplicationGatewayIdentityOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterIngressApplicationGatewayIngressApplicationGatewayIdentityOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		k,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterIngressApplicationGatewayIngressApplicationGatewayIdentityOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterIngressApplicationGatewayIngressApplicationGatewayIdentityOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterIngressApplicationGatewayIngressApplicationGatewayIdentityOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterIngressApplicationGatewayIngressApplicationGatewayIdentityOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type KubernetesClusterIngressApplicationGatewayOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	EffectiveGatewayId() *string
	// Experimental.
	Fqn() *string
	GatewayId() *string
	SetGatewayId(val *string)
	GatewayIdInput() *string
	GatewayName() *string
	SetGatewayName(val *string)
	GatewayNameInput() *string
	IngressApplicationGatewayIdentity() KubernetesClusterIngressApplicationGatewayIngressApplicationGatewayIdentityList
	InternalValue() *KubernetesClusterIngressApplicationGateway
	SetInternalValue(val *KubernetesClusterIngressApplicationGateway)
	SubnetCidr() *string
	SetSubnetCidr(val *string)
	SubnetCidrInput() *string
	SubnetId() *string
	SetSubnetId(val *string)
	SubnetIdInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetGatewayId()
	ResetGatewayName()
	ResetSubnetCidr()
	ResetSubnetId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KubernetesClusterIngressApplicationGatewayOutputReference
type jsiiProxy_KubernetesClusterIngressApplicationGatewayOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_KubernetesClusterIngressApplicationGatewayOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterIngressApplicationGatewayOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterIngressApplicationGatewayOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterIngressApplicationGatewayOutputReference) EffectiveGatewayId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"effectiveGatewayId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterIngressApplicationGatewayOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterIngressApplicationGatewayOutputReference) GatewayId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterIngressApplicationGatewayOutputReference) GatewayIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterIngressApplicationGatewayOutputReference) GatewayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterIngressApplicationGatewayOutputReference) GatewayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterIngressApplicationGatewayOutputReference) IngressApplicationGatewayIdentity() KubernetesClusterIngressApplicationGatewayIngressApplicationGatewayIdentityList {
	var returns KubernetesClusterIngressApplicationGatewayIngressApplicationGatewayIdentityList
	_jsii_.Get(
		j,
		"ingressApplicationGatewayIdentity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterIngressApplicationGatewayOutputReference) InternalValue() *KubernetesClusterIngressApplicationGateway {
	var returns *KubernetesClusterIngressApplicationGateway
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterIngressApplicationGatewayOutputReference) SubnetCidr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetCidr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterIngressApplicationGatewayOutputReference) SubnetCidrInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetCidrInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterIngressApplicationGatewayOutputReference) SubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterIngressApplicationGatewayOutputReference) SubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterIngressApplicationGatewayOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterIngressApplicationGatewayOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewKubernetesClusterIngressApplicationGatewayOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) KubernetesClusterIngressApplicationGatewayOutputReference {
	_init_.Initialize()

	j := jsiiProxy_KubernetesClusterIngressApplicationGatewayOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesCluster.KubernetesClusterIngressApplicationGatewayOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKubernetesClusterIngressApplicationGatewayOutputReference_Override(k KubernetesClusterIngressApplicationGatewayOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesCluster.KubernetesClusterIngressApplicationGatewayOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_KubernetesClusterIngressApplicationGatewayOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterIngressApplicationGatewayOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterIngressApplicationGatewayOutputReference) SetGatewayId(val *string) {
	_jsii_.Set(
		j,
		"gatewayId",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterIngressApplicationGatewayOutputReference) SetGatewayName(val *string) {
	_jsii_.Set(
		j,
		"gatewayName",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterIngressApplicationGatewayOutputReference) SetInternalValue(val *KubernetesClusterIngressApplicationGateway) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterIngressApplicationGatewayOutputReference) SetSubnetCidr(val *string) {
	_jsii_.Set(
		j,
		"subnetCidr",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterIngressApplicationGatewayOutputReference) SetSubnetId(val *string) {
	_jsii_.Set(
		j,
		"subnetId",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterIngressApplicationGatewayOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterIngressApplicationGatewayOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (k *jsiiProxy_KubernetesClusterIngressApplicationGatewayOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterIngressApplicationGatewayOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterIngressApplicationGatewayOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterIngressApplicationGatewayOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		k,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterIngressApplicationGatewayOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterIngressApplicationGatewayOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		k,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterIngressApplicationGatewayOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		k,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterIngressApplicationGatewayOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		k,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterIngressApplicationGatewayOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterIngressApplicationGatewayOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		k,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterIngressApplicationGatewayOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterIngressApplicationGatewayOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterIngressApplicationGatewayOutputReference) ResetGatewayId() {
	_jsii_.InvokeVoid(
		k,
		"resetGatewayId",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterIngressApplicationGatewayOutputReference) ResetGatewayName() {
	_jsii_.InvokeVoid(
		k,
		"resetGatewayName",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterIngressApplicationGatewayOutputReference) ResetSubnetCidr() {
	_jsii_.InvokeVoid(
		k,
		"resetSubnetCidr",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterIngressApplicationGatewayOutputReference) ResetSubnetId() {
	_jsii_.InvokeVoid(
		k,
		"resetSubnetId",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterIngressApplicationGatewayOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterIngressApplicationGatewayOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type KubernetesClusterKeyVaultSecretsProvider struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#secret_rotation_enabled KubernetesCluster#secret_rotation_enabled}.
	SecretRotationEnabled interface{} `field:"optional" json:"secretRotationEnabled" yaml:"secretRotationEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#secret_rotation_interval KubernetesCluster#secret_rotation_interval}.
	SecretRotationInterval *string `field:"optional" json:"secretRotationInterval" yaml:"secretRotationInterval"`
}

type KubernetesClusterKeyVaultSecretsProviderOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *KubernetesClusterKeyVaultSecretsProvider
	SetInternalValue(val *KubernetesClusterKeyVaultSecretsProvider)
	SecretIdentity() KubernetesClusterKeyVaultSecretsProviderSecretIdentityList
	SecretRotationEnabled() interface{}
	SetSecretRotationEnabled(val interface{})
	SecretRotationEnabledInput() interface{}
	SecretRotationInterval() *string
	SetSecretRotationInterval(val *string)
	SecretRotationIntervalInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetSecretRotationEnabled()
	ResetSecretRotationInterval()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KubernetesClusterKeyVaultSecretsProviderOutputReference
type jsiiProxy_KubernetesClusterKeyVaultSecretsProviderOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_KubernetesClusterKeyVaultSecretsProviderOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterKeyVaultSecretsProviderOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterKeyVaultSecretsProviderOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterKeyVaultSecretsProviderOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterKeyVaultSecretsProviderOutputReference) InternalValue() *KubernetesClusterKeyVaultSecretsProvider {
	var returns *KubernetesClusterKeyVaultSecretsProvider
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterKeyVaultSecretsProviderOutputReference) SecretIdentity() KubernetesClusterKeyVaultSecretsProviderSecretIdentityList {
	var returns KubernetesClusterKeyVaultSecretsProviderSecretIdentityList
	_jsii_.Get(
		j,
		"secretIdentity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterKeyVaultSecretsProviderOutputReference) SecretRotationEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secretRotationEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterKeyVaultSecretsProviderOutputReference) SecretRotationEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secretRotationEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterKeyVaultSecretsProviderOutputReference) SecretRotationInterval() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretRotationInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterKeyVaultSecretsProviderOutputReference) SecretRotationIntervalInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretRotationIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterKeyVaultSecretsProviderOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterKeyVaultSecretsProviderOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewKubernetesClusterKeyVaultSecretsProviderOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) KubernetesClusterKeyVaultSecretsProviderOutputReference {
	_init_.Initialize()

	j := jsiiProxy_KubernetesClusterKeyVaultSecretsProviderOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesCluster.KubernetesClusterKeyVaultSecretsProviderOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKubernetesClusterKeyVaultSecretsProviderOutputReference_Override(k KubernetesClusterKeyVaultSecretsProviderOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesCluster.KubernetesClusterKeyVaultSecretsProviderOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_KubernetesClusterKeyVaultSecretsProviderOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterKeyVaultSecretsProviderOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterKeyVaultSecretsProviderOutputReference) SetInternalValue(val *KubernetesClusterKeyVaultSecretsProvider) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterKeyVaultSecretsProviderOutputReference) SetSecretRotationEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"secretRotationEnabled",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterKeyVaultSecretsProviderOutputReference) SetSecretRotationInterval(val *string) {
	_jsii_.Set(
		j,
		"secretRotationInterval",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterKeyVaultSecretsProviderOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterKeyVaultSecretsProviderOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (k *jsiiProxy_KubernetesClusterKeyVaultSecretsProviderOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterKeyVaultSecretsProviderOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterKeyVaultSecretsProviderOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterKeyVaultSecretsProviderOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		k,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterKeyVaultSecretsProviderOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterKeyVaultSecretsProviderOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		k,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterKeyVaultSecretsProviderOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		k,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterKeyVaultSecretsProviderOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		k,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterKeyVaultSecretsProviderOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterKeyVaultSecretsProviderOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		k,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterKeyVaultSecretsProviderOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterKeyVaultSecretsProviderOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterKeyVaultSecretsProviderOutputReference) ResetSecretRotationEnabled() {
	_jsii_.InvokeVoid(
		k,
		"resetSecretRotationEnabled",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterKeyVaultSecretsProviderOutputReference) ResetSecretRotationInterval() {
	_jsii_.InvokeVoid(
		k,
		"resetSecretRotationInterval",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterKeyVaultSecretsProviderOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterKeyVaultSecretsProviderOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type KubernetesClusterKeyVaultSecretsProviderSecretIdentity struct {
}

type KubernetesClusterKeyVaultSecretsProviderSecretIdentityList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) KubernetesClusterKeyVaultSecretsProviderSecretIdentityOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KubernetesClusterKeyVaultSecretsProviderSecretIdentityList
type jsiiProxy_KubernetesClusterKeyVaultSecretsProviderSecretIdentityList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_KubernetesClusterKeyVaultSecretsProviderSecretIdentityList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterKeyVaultSecretsProviderSecretIdentityList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterKeyVaultSecretsProviderSecretIdentityList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterKeyVaultSecretsProviderSecretIdentityList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterKeyVaultSecretsProviderSecretIdentityList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewKubernetesClusterKeyVaultSecretsProviderSecretIdentityList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) KubernetesClusterKeyVaultSecretsProviderSecretIdentityList {
	_init_.Initialize()

	j := jsiiProxy_KubernetesClusterKeyVaultSecretsProviderSecretIdentityList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesCluster.KubernetesClusterKeyVaultSecretsProviderSecretIdentityList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewKubernetesClusterKeyVaultSecretsProviderSecretIdentityList_Override(k KubernetesClusterKeyVaultSecretsProviderSecretIdentityList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesCluster.KubernetesClusterKeyVaultSecretsProviderSecretIdentityList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		k,
	)
}

func (j *jsiiProxy_KubernetesClusterKeyVaultSecretsProviderSecretIdentityList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterKeyVaultSecretsProviderSecretIdentityList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterKeyVaultSecretsProviderSecretIdentityList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (k *jsiiProxy_KubernetesClusterKeyVaultSecretsProviderSecretIdentityList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterKeyVaultSecretsProviderSecretIdentityList) Get(index *float64) KubernetesClusterKeyVaultSecretsProviderSecretIdentityOutputReference {
	var returns KubernetesClusterKeyVaultSecretsProviderSecretIdentityOutputReference

	_jsii_.Invoke(
		k,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterKeyVaultSecretsProviderSecretIdentityList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterKeyVaultSecretsProviderSecretIdentityList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type KubernetesClusterKeyVaultSecretsProviderSecretIdentityOutputReference interface {
	cdktf.ComplexObject
	ClientId() *string
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *KubernetesClusterKeyVaultSecretsProviderSecretIdentity
	SetInternalValue(val *KubernetesClusterKeyVaultSecretsProviderSecretIdentity)
	ObjectId() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UserAssignedIdentityId() *string
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KubernetesClusterKeyVaultSecretsProviderSecretIdentityOutputReference
type jsiiProxy_KubernetesClusterKeyVaultSecretsProviderSecretIdentityOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_KubernetesClusterKeyVaultSecretsProviderSecretIdentityOutputReference) ClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterKeyVaultSecretsProviderSecretIdentityOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterKeyVaultSecretsProviderSecretIdentityOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterKeyVaultSecretsProviderSecretIdentityOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterKeyVaultSecretsProviderSecretIdentityOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterKeyVaultSecretsProviderSecretIdentityOutputReference) InternalValue() *KubernetesClusterKeyVaultSecretsProviderSecretIdentity {
	var returns *KubernetesClusterKeyVaultSecretsProviderSecretIdentity
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterKeyVaultSecretsProviderSecretIdentityOutputReference) ObjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterKeyVaultSecretsProviderSecretIdentityOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterKeyVaultSecretsProviderSecretIdentityOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterKeyVaultSecretsProviderSecretIdentityOutputReference) UserAssignedIdentityId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userAssignedIdentityId",
		&returns,
	)
	return returns
}


func NewKubernetesClusterKeyVaultSecretsProviderSecretIdentityOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) KubernetesClusterKeyVaultSecretsProviderSecretIdentityOutputReference {
	_init_.Initialize()

	j := jsiiProxy_KubernetesClusterKeyVaultSecretsProviderSecretIdentityOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesCluster.KubernetesClusterKeyVaultSecretsProviderSecretIdentityOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewKubernetesClusterKeyVaultSecretsProviderSecretIdentityOutputReference_Override(k KubernetesClusterKeyVaultSecretsProviderSecretIdentityOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesCluster.KubernetesClusterKeyVaultSecretsProviderSecretIdentityOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		k,
	)
}

func (j *jsiiProxy_KubernetesClusterKeyVaultSecretsProviderSecretIdentityOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterKeyVaultSecretsProviderSecretIdentityOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterKeyVaultSecretsProviderSecretIdentityOutputReference) SetInternalValue(val *KubernetesClusterKeyVaultSecretsProviderSecretIdentity) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterKeyVaultSecretsProviderSecretIdentityOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterKeyVaultSecretsProviderSecretIdentityOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (k *jsiiProxy_KubernetesClusterKeyVaultSecretsProviderSecretIdentityOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterKeyVaultSecretsProviderSecretIdentityOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterKeyVaultSecretsProviderSecretIdentityOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterKeyVaultSecretsProviderSecretIdentityOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		k,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterKeyVaultSecretsProviderSecretIdentityOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterKeyVaultSecretsProviderSecretIdentityOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		k,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterKeyVaultSecretsProviderSecretIdentityOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		k,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterKeyVaultSecretsProviderSecretIdentityOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		k,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterKeyVaultSecretsProviderSecretIdentityOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterKeyVaultSecretsProviderSecretIdentityOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		k,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterKeyVaultSecretsProviderSecretIdentityOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterKeyVaultSecretsProviderSecretIdentityOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterKeyVaultSecretsProviderSecretIdentityOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterKeyVaultSecretsProviderSecretIdentityOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type KubernetesClusterKubeAdminConfig struct {
}

type KubernetesClusterKubeAdminConfigList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) KubernetesClusterKubeAdminConfigOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KubernetesClusterKubeAdminConfigList
type jsiiProxy_KubernetesClusterKubeAdminConfigList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_KubernetesClusterKubeAdminConfigList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterKubeAdminConfigList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterKubeAdminConfigList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterKubeAdminConfigList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterKubeAdminConfigList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewKubernetesClusterKubeAdminConfigList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) KubernetesClusterKubeAdminConfigList {
	_init_.Initialize()

	j := jsiiProxy_KubernetesClusterKubeAdminConfigList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesCluster.KubernetesClusterKubeAdminConfigList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewKubernetesClusterKubeAdminConfigList_Override(k KubernetesClusterKubeAdminConfigList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesCluster.KubernetesClusterKubeAdminConfigList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		k,
	)
}

func (j *jsiiProxy_KubernetesClusterKubeAdminConfigList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterKubeAdminConfigList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterKubeAdminConfigList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (k *jsiiProxy_KubernetesClusterKubeAdminConfigList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterKubeAdminConfigList) Get(index *float64) KubernetesClusterKubeAdminConfigOutputReference {
	var returns KubernetesClusterKubeAdminConfigOutputReference

	_jsii_.Invoke(
		k,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterKubeAdminConfigList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterKubeAdminConfigList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type KubernetesClusterKubeAdminConfigOutputReference interface {
	cdktf.ComplexObject
	ClientCertificate() *string
	ClientKey() *string
	ClusterCaCertificate() *string
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	Host() *string
	InternalValue() *KubernetesClusterKubeAdminConfig
	SetInternalValue(val *KubernetesClusterKubeAdminConfig)
	Password() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Username() *string
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KubernetesClusterKubeAdminConfigOutputReference
type jsiiProxy_KubernetesClusterKubeAdminConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_KubernetesClusterKubeAdminConfigOutputReference) ClientCertificate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCertificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterKubeAdminConfigOutputReference) ClientKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterKubeAdminConfigOutputReference) ClusterCaCertificate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterCaCertificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterKubeAdminConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterKubeAdminConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterKubeAdminConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterKubeAdminConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterKubeAdminConfigOutputReference) Host() *string {
	var returns *string
	_jsii_.Get(
		j,
		"host",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterKubeAdminConfigOutputReference) InternalValue() *KubernetesClusterKubeAdminConfig {
	var returns *KubernetesClusterKubeAdminConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterKubeAdminConfigOutputReference) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterKubeAdminConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterKubeAdminConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterKubeAdminConfigOutputReference) Username() *string {
	var returns *string
	_jsii_.Get(
		j,
		"username",
		&returns,
	)
	return returns
}


func NewKubernetesClusterKubeAdminConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) KubernetesClusterKubeAdminConfigOutputReference {
	_init_.Initialize()

	j := jsiiProxy_KubernetesClusterKubeAdminConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesCluster.KubernetesClusterKubeAdminConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewKubernetesClusterKubeAdminConfigOutputReference_Override(k KubernetesClusterKubeAdminConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesCluster.KubernetesClusterKubeAdminConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		k,
	)
}

func (j *jsiiProxy_KubernetesClusterKubeAdminConfigOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterKubeAdminConfigOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterKubeAdminConfigOutputReference) SetInternalValue(val *KubernetesClusterKubeAdminConfig) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterKubeAdminConfigOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterKubeAdminConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (k *jsiiProxy_KubernetesClusterKubeAdminConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterKubeAdminConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterKubeAdminConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterKubeAdminConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		k,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterKubeAdminConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterKubeAdminConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		k,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterKubeAdminConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		k,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterKubeAdminConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		k,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterKubeAdminConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterKubeAdminConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		k,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterKubeAdminConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterKubeAdminConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterKubeAdminConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterKubeAdminConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type KubernetesClusterKubeConfig struct {
}

type KubernetesClusterKubeConfigList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) KubernetesClusterKubeConfigOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KubernetesClusterKubeConfigList
type jsiiProxy_KubernetesClusterKubeConfigList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_KubernetesClusterKubeConfigList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterKubeConfigList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterKubeConfigList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterKubeConfigList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterKubeConfigList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewKubernetesClusterKubeConfigList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) KubernetesClusterKubeConfigList {
	_init_.Initialize()

	j := jsiiProxy_KubernetesClusterKubeConfigList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesCluster.KubernetesClusterKubeConfigList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewKubernetesClusterKubeConfigList_Override(k KubernetesClusterKubeConfigList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesCluster.KubernetesClusterKubeConfigList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		k,
	)
}

func (j *jsiiProxy_KubernetesClusterKubeConfigList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterKubeConfigList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterKubeConfigList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (k *jsiiProxy_KubernetesClusterKubeConfigList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterKubeConfigList) Get(index *float64) KubernetesClusterKubeConfigOutputReference {
	var returns KubernetesClusterKubeConfigOutputReference

	_jsii_.Invoke(
		k,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterKubeConfigList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterKubeConfigList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type KubernetesClusterKubeConfigOutputReference interface {
	cdktf.ComplexObject
	ClientCertificate() *string
	ClientKey() *string
	ClusterCaCertificate() *string
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	Host() *string
	InternalValue() *KubernetesClusterKubeConfig
	SetInternalValue(val *KubernetesClusterKubeConfig)
	Password() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Username() *string
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KubernetesClusterKubeConfigOutputReference
type jsiiProxy_KubernetesClusterKubeConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_KubernetesClusterKubeConfigOutputReference) ClientCertificate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCertificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterKubeConfigOutputReference) ClientKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterKubeConfigOutputReference) ClusterCaCertificate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterCaCertificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterKubeConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterKubeConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterKubeConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterKubeConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterKubeConfigOutputReference) Host() *string {
	var returns *string
	_jsii_.Get(
		j,
		"host",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterKubeConfigOutputReference) InternalValue() *KubernetesClusterKubeConfig {
	var returns *KubernetesClusterKubeConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterKubeConfigOutputReference) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterKubeConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterKubeConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterKubeConfigOutputReference) Username() *string {
	var returns *string
	_jsii_.Get(
		j,
		"username",
		&returns,
	)
	return returns
}


func NewKubernetesClusterKubeConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) KubernetesClusterKubeConfigOutputReference {
	_init_.Initialize()

	j := jsiiProxy_KubernetesClusterKubeConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesCluster.KubernetesClusterKubeConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewKubernetesClusterKubeConfigOutputReference_Override(k KubernetesClusterKubeConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesCluster.KubernetesClusterKubeConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		k,
	)
}

func (j *jsiiProxy_KubernetesClusterKubeConfigOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterKubeConfigOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterKubeConfigOutputReference) SetInternalValue(val *KubernetesClusterKubeConfig) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterKubeConfigOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterKubeConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (k *jsiiProxy_KubernetesClusterKubeConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterKubeConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterKubeConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterKubeConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		k,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterKubeConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterKubeConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		k,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterKubeConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		k,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterKubeConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		k,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterKubeConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterKubeConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		k,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterKubeConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterKubeConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterKubeConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterKubeConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type KubernetesClusterKubeletIdentity struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#client_id KubernetesCluster#client_id}.
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#object_id KubernetesCluster#object_id}.
	ObjectId *string `field:"optional" json:"objectId" yaml:"objectId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#user_assigned_identity_id KubernetesCluster#user_assigned_identity_id}.
	UserAssignedIdentityId *string `field:"optional" json:"userAssignedIdentityId" yaml:"userAssignedIdentityId"`
}

type KubernetesClusterKubeletIdentityOutputReference interface {
	cdktf.ComplexObject
	ClientId() *string
	SetClientId(val *string)
	ClientIdInput() *string
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *KubernetesClusterKubeletIdentity
	SetInternalValue(val *KubernetesClusterKubeletIdentity)
	ObjectId() *string
	SetObjectId(val *string)
	ObjectIdInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UserAssignedIdentityId() *string
	SetUserAssignedIdentityId(val *string)
	UserAssignedIdentityIdInput() *string
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetClientId()
	ResetObjectId()
	ResetUserAssignedIdentityId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KubernetesClusterKubeletIdentityOutputReference
type jsiiProxy_KubernetesClusterKubeletIdentityOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_KubernetesClusterKubeletIdentityOutputReference) ClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterKubeletIdentityOutputReference) ClientIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterKubeletIdentityOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterKubeletIdentityOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterKubeletIdentityOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterKubeletIdentityOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterKubeletIdentityOutputReference) InternalValue() *KubernetesClusterKubeletIdentity {
	var returns *KubernetesClusterKubeletIdentity
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterKubeletIdentityOutputReference) ObjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterKubeletIdentityOutputReference) ObjectIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterKubeletIdentityOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterKubeletIdentityOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterKubeletIdentityOutputReference) UserAssignedIdentityId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userAssignedIdentityId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterKubeletIdentityOutputReference) UserAssignedIdentityIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userAssignedIdentityIdInput",
		&returns,
	)
	return returns
}


func NewKubernetesClusterKubeletIdentityOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) KubernetesClusterKubeletIdentityOutputReference {
	_init_.Initialize()

	j := jsiiProxy_KubernetesClusterKubeletIdentityOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesCluster.KubernetesClusterKubeletIdentityOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKubernetesClusterKubeletIdentityOutputReference_Override(k KubernetesClusterKubeletIdentityOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesCluster.KubernetesClusterKubeletIdentityOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_KubernetesClusterKubeletIdentityOutputReference) SetClientId(val *string) {
	_jsii_.Set(
		j,
		"clientId",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterKubeletIdentityOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterKubeletIdentityOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterKubeletIdentityOutputReference) SetInternalValue(val *KubernetesClusterKubeletIdentity) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterKubeletIdentityOutputReference) SetObjectId(val *string) {
	_jsii_.Set(
		j,
		"objectId",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterKubeletIdentityOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterKubeletIdentityOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterKubeletIdentityOutputReference) SetUserAssignedIdentityId(val *string) {
	_jsii_.Set(
		j,
		"userAssignedIdentityId",
		val,
	)
}

func (k *jsiiProxy_KubernetesClusterKubeletIdentityOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterKubeletIdentityOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterKubeletIdentityOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterKubeletIdentityOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		k,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterKubeletIdentityOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterKubeletIdentityOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		k,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterKubeletIdentityOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		k,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterKubeletIdentityOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		k,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterKubeletIdentityOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterKubeletIdentityOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		k,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterKubeletIdentityOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterKubeletIdentityOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterKubeletIdentityOutputReference) ResetClientId() {
	_jsii_.InvokeVoid(
		k,
		"resetClientId",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterKubeletIdentityOutputReference) ResetObjectId() {
	_jsii_.InvokeVoid(
		k,
		"resetObjectId",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterKubeletIdentityOutputReference) ResetUserAssignedIdentityId() {
	_jsii_.InvokeVoid(
		k,
		"resetUserAssignedIdentityId",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterKubeletIdentityOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterKubeletIdentityOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type KubernetesClusterLinuxProfile struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#admin_username KubernetesCluster#admin_username}.
	AdminUsername *string `field:"required" json:"adminUsername" yaml:"adminUsername"`
	// ssh_key block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#ssh_key KubernetesCluster#ssh_key}
	SshKey *KubernetesClusterLinuxProfileSshKey `field:"required" json:"sshKey" yaml:"sshKey"`
}

type KubernetesClusterLinuxProfileOutputReference interface {
	cdktf.ComplexObject
	AdminUsername() *string
	SetAdminUsername(val *string)
	AdminUsernameInput() *string
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *KubernetesClusterLinuxProfile
	SetInternalValue(val *KubernetesClusterLinuxProfile)
	SshKey() KubernetesClusterLinuxProfileSshKeyOutputReference
	SshKeyInput() *KubernetesClusterLinuxProfileSshKey
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	PutSshKey(value *KubernetesClusterLinuxProfileSshKey)
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KubernetesClusterLinuxProfileOutputReference
type jsiiProxy_KubernetesClusterLinuxProfileOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_KubernetesClusterLinuxProfileOutputReference) AdminUsername() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminUsername",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterLinuxProfileOutputReference) AdminUsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminUsernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterLinuxProfileOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterLinuxProfileOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterLinuxProfileOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterLinuxProfileOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterLinuxProfileOutputReference) InternalValue() *KubernetesClusterLinuxProfile {
	var returns *KubernetesClusterLinuxProfile
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterLinuxProfileOutputReference) SshKey() KubernetesClusterLinuxProfileSshKeyOutputReference {
	var returns KubernetesClusterLinuxProfileSshKeyOutputReference
	_jsii_.Get(
		j,
		"sshKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterLinuxProfileOutputReference) SshKeyInput() *KubernetesClusterLinuxProfileSshKey {
	var returns *KubernetesClusterLinuxProfileSshKey
	_jsii_.Get(
		j,
		"sshKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterLinuxProfileOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterLinuxProfileOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewKubernetesClusterLinuxProfileOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) KubernetesClusterLinuxProfileOutputReference {
	_init_.Initialize()

	j := jsiiProxy_KubernetesClusterLinuxProfileOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesCluster.KubernetesClusterLinuxProfileOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKubernetesClusterLinuxProfileOutputReference_Override(k KubernetesClusterLinuxProfileOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesCluster.KubernetesClusterLinuxProfileOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_KubernetesClusterLinuxProfileOutputReference) SetAdminUsername(val *string) {
	_jsii_.Set(
		j,
		"adminUsername",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterLinuxProfileOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterLinuxProfileOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterLinuxProfileOutputReference) SetInternalValue(val *KubernetesClusterLinuxProfile) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterLinuxProfileOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterLinuxProfileOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (k *jsiiProxy_KubernetesClusterLinuxProfileOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterLinuxProfileOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterLinuxProfileOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterLinuxProfileOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		k,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterLinuxProfileOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterLinuxProfileOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		k,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterLinuxProfileOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		k,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterLinuxProfileOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		k,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterLinuxProfileOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterLinuxProfileOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		k,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterLinuxProfileOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterLinuxProfileOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterLinuxProfileOutputReference) PutSshKey(value *KubernetesClusterLinuxProfileSshKey) {
	_jsii_.InvokeVoid(
		k,
		"putSshKey",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KubernetesClusterLinuxProfileOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterLinuxProfileOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type KubernetesClusterLinuxProfileSshKey struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#key_data KubernetesCluster#key_data}.
	KeyData *string `field:"required" json:"keyData" yaml:"keyData"`
}

type KubernetesClusterLinuxProfileSshKeyOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *KubernetesClusterLinuxProfileSshKey
	SetInternalValue(val *KubernetesClusterLinuxProfileSshKey)
	KeyData() *string
	SetKeyData(val *string)
	KeyDataInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KubernetesClusterLinuxProfileSshKeyOutputReference
type jsiiProxy_KubernetesClusterLinuxProfileSshKeyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_KubernetesClusterLinuxProfileSshKeyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterLinuxProfileSshKeyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterLinuxProfileSshKeyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterLinuxProfileSshKeyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterLinuxProfileSshKeyOutputReference) InternalValue() *KubernetesClusterLinuxProfileSshKey {
	var returns *KubernetesClusterLinuxProfileSshKey
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterLinuxProfileSshKeyOutputReference) KeyData() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterLinuxProfileSshKeyOutputReference) KeyDataInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterLinuxProfileSshKeyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterLinuxProfileSshKeyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewKubernetesClusterLinuxProfileSshKeyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) KubernetesClusterLinuxProfileSshKeyOutputReference {
	_init_.Initialize()

	j := jsiiProxy_KubernetesClusterLinuxProfileSshKeyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesCluster.KubernetesClusterLinuxProfileSshKeyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKubernetesClusterLinuxProfileSshKeyOutputReference_Override(k KubernetesClusterLinuxProfileSshKeyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesCluster.KubernetesClusterLinuxProfileSshKeyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_KubernetesClusterLinuxProfileSshKeyOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterLinuxProfileSshKeyOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterLinuxProfileSshKeyOutputReference) SetInternalValue(val *KubernetesClusterLinuxProfileSshKey) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterLinuxProfileSshKeyOutputReference) SetKeyData(val *string) {
	_jsii_.Set(
		j,
		"keyData",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterLinuxProfileSshKeyOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterLinuxProfileSshKeyOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (k *jsiiProxy_KubernetesClusterLinuxProfileSshKeyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterLinuxProfileSshKeyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterLinuxProfileSshKeyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterLinuxProfileSshKeyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		k,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterLinuxProfileSshKeyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterLinuxProfileSshKeyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		k,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterLinuxProfileSshKeyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		k,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterLinuxProfileSshKeyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		k,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterLinuxProfileSshKeyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterLinuxProfileSshKeyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		k,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterLinuxProfileSshKeyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterLinuxProfileSshKeyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterLinuxProfileSshKeyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterLinuxProfileSshKeyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type KubernetesClusterMaintenanceWindow struct {
	// allowed block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#allowed KubernetesCluster#allowed}
	Allowed interface{} `field:"optional" json:"allowed" yaml:"allowed"`
	// not_allowed block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#not_allowed KubernetesCluster#not_allowed}
	NotAllowed interface{} `field:"optional" json:"notAllowed" yaml:"notAllowed"`
}

type KubernetesClusterMaintenanceWindowAllowed struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#day KubernetesCluster#day}.
	Day *string `field:"required" json:"day" yaml:"day"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#hours KubernetesCluster#hours}.
	Hours *[]*float64 `field:"required" json:"hours" yaml:"hours"`
}

type KubernetesClusterMaintenanceWindowAllowedList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) KubernetesClusterMaintenanceWindowAllowedOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KubernetesClusterMaintenanceWindowAllowedList
type jsiiProxy_KubernetesClusterMaintenanceWindowAllowedList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowAllowedList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowAllowedList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowAllowedList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowAllowedList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowAllowedList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowAllowedList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewKubernetesClusterMaintenanceWindowAllowedList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) KubernetesClusterMaintenanceWindowAllowedList {
	_init_.Initialize()

	j := jsiiProxy_KubernetesClusterMaintenanceWindowAllowedList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesCluster.KubernetesClusterMaintenanceWindowAllowedList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewKubernetesClusterMaintenanceWindowAllowedList_Override(k KubernetesClusterMaintenanceWindowAllowedList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesCluster.KubernetesClusterMaintenanceWindowAllowedList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		k,
	)
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowAllowedList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowAllowedList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowAllowedList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowAllowedList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (k *jsiiProxy_KubernetesClusterMaintenanceWindowAllowedList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterMaintenanceWindowAllowedList) Get(index *float64) KubernetesClusterMaintenanceWindowAllowedOutputReference {
	var returns KubernetesClusterMaintenanceWindowAllowedOutputReference

	_jsii_.Invoke(
		k,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterMaintenanceWindowAllowedList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterMaintenanceWindowAllowedList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type KubernetesClusterMaintenanceWindowAllowedOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Day() *string
	SetDay(val *string)
	DayInput() *string
	// Experimental.
	Fqn() *string
	Hours() *[]*float64
	SetHours(val *[]*float64)
	HoursInput() *[]*float64
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KubernetesClusterMaintenanceWindowAllowedOutputReference
type jsiiProxy_KubernetesClusterMaintenanceWindowAllowedOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowAllowedOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowAllowedOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowAllowedOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowAllowedOutputReference) Day() *string {
	var returns *string
	_jsii_.Get(
		j,
		"day",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowAllowedOutputReference) DayInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowAllowedOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowAllowedOutputReference) Hours() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"hours",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowAllowedOutputReference) HoursInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"hoursInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowAllowedOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowAllowedOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowAllowedOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewKubernetesClusterMaintenanceWindowAllowedOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) KubernetesClusterMaintenanceWindowAllowedOutputReference {
	_init_.Initialize()

	j := jsiiProxy_KubernetesClusterMaintenanceWindowAllowedOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesCluster.KubernetesClusterMaintenanceWindowAllowedOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewKubernetesClusterMaintenanceWindowAllowedOutputReference_Override(k KubernetesClusterMaintenanceWindowAllowedOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesCluster.KubernetesClusterMaintenanceWindowAllowedOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		k,
	)
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowAllowedOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowAllowedOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowAllowedOutputReference) SetDay(val *string) {
	_jsii_.Set(
		j,
		"day",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowAllowedOutputReference) SetHours(val *[]*float64) {
	_jsii_.Set(
		j,
		"hours",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowAllowedOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowAllowedOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowAllowedOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (k *jsiiProxy_KubernetesClusterMaintenanceWindowAllowedOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterMaintenanceWindowAllowedOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterMaintenanceWindowAllowedOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterMaintenanceWindowAllowedOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		k,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterMaintenanceWindowAllowedOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterMaintenanceWindowAllowedOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		k,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterMaintenanceWindowAllowedOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		k,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterMaintenanceWindowAllowedOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		k,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterMaintenanceWindowAllowedOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterMaintenanceWindowAllowedOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		k,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterMaintenanceWindowAllowedOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterMaintenanceWindowAllowedOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterMaintenanceWindowAllowedOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterMaintenanceWindowAllowedOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type KubernetesClusterMaintenanceWindowNotAllowed struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#end KubernetesCluster#end}.
	End *string `field:"required" json:"end" yaml:"end"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#start KubernetesCluster#start}.
	Start *string `field:"required" json:"start" yaml:"start"`
}

type KubernetesClusterMaintenanceWindowNotAllowedList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) KubernetesClusterMaintenanceWindowNotAllowedOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KubernetesClusterMaintenanceWindowNotAllowedList
type jsiiProxy_KubernetesClusterMaintenanceWindowNotAllowedList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowNotAllowedList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowNotAllowedList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowNotAllowedList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowNotAllowedList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowNotAllowedList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowNotAllowedList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewKubernetesClusterMaintenanceWindowNotAllowedList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) KubernetesClusterMaintenanceWindowNotAllowedList {
	_init_.Initialize()

	j := jsiiProxy_KubernetesClusterMaintenanceWindowNotAllowedList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesCluster.KubernetesClusterMaintenanceWindowNotAllowedList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewKubernetesClusterMaintenanceWindowNotAllowedList_Override(k KubernetesClusterMaintenanceWindowNotAllowedList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesCluster.KubernetesClusterMaintenanceWindowNotAllowedList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		k,
	)
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowNotAllowedList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowNotAllowedList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowNotAllowedList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowNotAllowedList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (k *jsiiProxy_KubernetesClusterMaintenanceWindowNotAllowedList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterMaintenanceWindowNotAllowedList) Get(index *float64) KubernetesClusterMaintenanceWindowNotAllowedOutputReference {
	var returns KubernetesClusterMaintenanceWindowNotAllowedOutputReference

	_jsii_.Invoke(
		k,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterMaintenanceWindowNotAllowedList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterMaintenanceWindowNotAllowedList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type KubernetesClusterMaintenanceWindowNotAllowedOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	End() *string
	SetEnd(val *string)
	EndInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Start() *string
	SetStart(val *string)
	StartInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KubernetesClusterMaintenanceWindowNotAllowedOutputReference
type jsiiProxy_KubernetesClusterMaintenanceWindowNotAllowedOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowNotAllowedOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowNotAllowedOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowNotAllowedOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowNotAllowedOutputReference) End() *string {
	var returns *string
	_jsii_.Get(
		j,
		"end",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowNotAllowedOutputReference) EndInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowNotAllowedOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowNotAllowedOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowNotAllowedOutputReference) Start() *string {
	var returns *string
	_jsii_.Get(
		j,
		"start",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowNotAllowedOutputReference) StartInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowNotAllowedOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowNotAllowedOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewKubernetesClusterMaintenanceWindowNotAllowedOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) KubernetesClusterMaintenanceWindowNotAllowedOutputReference {
	_init_.Initialize()

	j := jsiiProxy_KubernetesClusterMaintenanceWindowNotAllowedOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesCluster.KubernetesClusterMaintenanceWindowNotAllowedOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewKubernetesClusterMaintenanceWindowNotAllowedOutputReference_Override(k KubernetesClusterMaintenanceWindowNotAllowedOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesCluster.KubernetesClusterMaintenanceWindowNotAllowedOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		k,
	)
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowNotAllowedOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowNotAllowedOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowNotAllowedOutputReference) SetEnd(val *string) {
	_jsii_.Set(
		j,
		"end",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowNotAllowedOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowNotAllowedOutputReference) SetStart(val *string) {
	_jsii_.Set(
		j,
		"start",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowNotAllowedOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowNotAllowedOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (k *jsiiProxy_KubernetesClusterMaintenanceWindowNotAllowedOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterMaintenanceWindowNotAllowedOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterMaintenanceWindowNotAllowedOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterMaintenanceWindowNotAllowedOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		k,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterMaintenanceWindowNotAllowedOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterMaintenanceWindowNotAllowedOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		k,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterMaintenanceWindowNotAllowedOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		k,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterMaintenanceWindowNotAllowedOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		k,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterMaintenanceWindowNotAllowedOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterMaintenanceWindowNotAllowedOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		k,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterMaintenanceWindowNotAllowedOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterMaintenanceWindowNotAllowedOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterMaintenanceWindowNotAllowedOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterMaintenanceWindowNotAllowedOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type KubernetesClusterMaintenanceWindowOutputReference interface {
	cdktf.ComplexObject
	Allowed() KubernetesClusterMaintenanceWindowAllowedList
	AllowedInput() interface{}
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *KubernetesClusterMaintenanceWindow
	SetInternalValue(val *KubernetesClusterMaintenanceWindow)
	NotAllowed() KubernetesClusterMaintenanceWindowNotAllowedList
	NotAllowedInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	PutAllowed(value interface{})
	PutNotAllowed(value interface{})
	ResetAllowed()
	ResetNotAllowed()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KubernetesClusterMaintenanceWindowOutputReference
type jsiiProxy_KubernetesClusterMaintenanceWindowOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowOutputReference) Allowed() KubernetesClusterMaintenanceWindowAllowedList {
	var returns KubernetesClusterMaintenanceWindowAllowedList
	_jsii_.Get(
		j,
		"allowed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowOutputReference) AllowedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowOutputReference) InternalValue() *KubernetesClusterMaintenanceWindow {
	var returns *KubernetesClusterMaintenanceWindow
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowOutputReference) NotAllowed() KubernetesClusterMaintenanceWindowNotAllowedList {
	var returns KubernetesClusterMaintenanceWindowNotAllowedList
	_jsii_.Get(
		j,
		"notAllowed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowOutputReference) NotAllowedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"notAllowedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewKubernetesClusterMaintenanceWindowOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) KubernetesClusterMaintenanceWindowOutputReference {
	_init_.Initialize()

	j := jsiiProxy_KubernetesClusterMaintenanceWindowOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesCluster.KubernetesClusterMaintenanceWindowOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKubernetesClusterMaintenanceWindowOutputReference_Override(k KubernetesClusterMaintenanceWindowOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesCluster.KubernetesClusterMaintenanceWindowOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowOutputReference) SetInternalValue(val *KubernetesClusterMaintenanceWindow) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterMaintenanceWindowOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (k *jsiiProxy_KubernetesClusterMaintenanceWindowOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterMaintenanceWindowOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterMaintenanceWindowOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterMaintenanceWindowOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		k,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterMaintenanceWindowOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterMaintenanceWindowOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		k,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterMaintenanceWindowOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		k,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterMaintenanceWindowOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		k,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterMaintenanceWindowOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterMaintenanceWindowOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		k,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterMaintenanceWindowOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterMaintenanceWindowOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterMaintenanceWindowOutputReference) PutAllowed(value interface{}) {
	_jsii_.InvokeVoid(
		k,
		"putAllowed",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KubernetesClusterMaintenanceWindowOutputReference) PutNotAllowed(value interface{}) {
	_jsii_.InvokeVoid(
		k,
		"putNotAllowed",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KubernetesClusterMaintenanceWindowOutputReference) ResetAllowed() {
	_jsii_.InvokeVoid(
		k,
		"resetAllowed",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterMaintenanceWindowOutputReference) ResetNotAllowed() {
	_jsii_.InvokeVoid(
		k,
		"resetNotAllowed",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterMaintenanceWindowOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterMaintenanceWindowOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type KubernetesClusterMicrosoftDefender struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#log_analytics_workspace_id KubernetesCluster#log_analytics_workspace_id}.
	LogAnalyticsWorkspaceId *string `field:"required" json:"logAnalyticsWorkspaceId" yaml:"logAnalyticsWorkspaceId"`
}

type KubernetesClusterMicrosoftDefenderOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *KubernetesClusterMicrosoftDefender
	SetInternalValue(val *KubernetesClusterMicrosoftDefender)
	LogAnalyticsWorkspaceId() *string
	SetLogAnalyticsWorkspaceId(val *string)
	LogAnalyticsWorkspaceIdInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KubernetesClusterMicrosoftDefenderOutputReference
type jsiiProxy_KubernetesClusterMicrosoftDefenderOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_KubernetesClusterMicrosoftDefenderOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMicrosoftDefenderOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMicrosoftDefenderOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMicrosoftDefenderOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMicrosoftDefenderOutputReference) InternalValue() *KubernetesClusterMicrosoftDefender {
	var returns *KubernetesClusterMicrosoftDefender
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMicrosoftDefenderOutputReference) LogAnalyticsWorkspaceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logAnalyticsWorkspaceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMicrosoftDefenderOutputReference) LogAnalyticsWorkspaceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logAnalyticsWorkspaceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMicrosoftDefenderOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterMicrosoftDefenderOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewKubernetesClusterMicrosoftDefenderOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) KubernetesClusterMicrosoftDefenderOutputReference {
	_init_.Initialize()

	j := jsiiProxy_KubernetesClusterMicrosoftDefenderOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesCluster.KubernetesClusterMicrosoftDefenderOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKubernetesClusterMicrosoftDefenderOutputReference_Override(k KubernetesClusterMicrosoftDefenderOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesCluster.KubernetesClusterMicrosoftDefenderOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_KubernetesClusterMicrosoftDefenderOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterMicrosoftDefenderOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterMicrosoftDefenderOutputReference) SetInternalValue(val *KubernetesClusterMicrosoftDefender) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterMicrosoftDefenderOutputReference) SetLogAnalyticsWorkspaceId(val *string) {
	_jsii_.Set(
		j,
		"logAnalyticsWorkspaceId",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterMicrosoftDefenderOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterMicrosoftDefenderOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (k *jsiiProxy_KubernetesClusterMicrosoftDefenderOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterMicrosoftDefenderOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterMicrosoftDefenderOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterMicrosoftDefenderOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		k,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterMicrosoftDefenderOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterMicrosoftDefenderOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		k,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterMicrosoftDefenderOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		k,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterMicrosoftDefenderOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		k,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterMicrosoftDefenderOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterMicrosoftDefenderOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		k,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterMicrosoftDefenderOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterMicrosoftDefenderOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterMicrosoftDefenderOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterMicrosoftDefenderOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type KubernetesClusterNetworkProfile struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#network_plugin KubernetesCluster#network_plugin}.
	NetworkPlugin *string `field:"required" json:"networkPlugin" yaml:"networkPlugin"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#dns_service_ip KubernetesCluster#dns_service_ip}.
	DnsServiceIp *string `field:"optional" json:"dnsServiceIp" yaml:"dnsServiceIp"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#docker_bridge_cidr KubernetesCluster#docker_bridge_cidr}.
	DockerBridgeCidr *string `field:"optional" json:"dockerBridgeCidr" yaml:"dockerBridgeCidr"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#ip_versions KubernetesCluster#ip_versions}.
	IpVersions *[]*string `field:"optional" json:"ipVersions" yaml:"ipVersions"`
	// load_balancer_profile block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#load_balancer_profile KubernetesCluster#load_balancer_profile}
	LoadBalancerProfile *KubernetesClusterNetworkProfileLoadBalancerProfile `field:"optional" json:"loadBalancerProfile" yaml:"loadBalancerProfile"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#load_balancer_sku KubernetesCluster#load_balancer_sku}.
	LoadBalancerSku *string `field:"optional" json:"loadBalancerSku" yaml:"loadBalancerSku"`
	// nat_gateway_profile block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#nat_gateway_profile KubernetesCluster#nat_gateway_profile}
	NatGatewayProfile *KubernetesClusterNetworkProfileNatGatewayProfile `field:"optional" json:"natGatewayProfile" yaml:"natGatewayProfile"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#network_mode KubernetesCluster#network_mode}.
	NetworkMode *string `field:"optional" json:"networkMode" yaml:"networkMode"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#network_policy KubernetesCluster#network_policy}.
	NetworkPolicy *string `field:"optional" json:"networkPolicy" yaml:"networkPolicy"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#outbound_type KubernetesCluster#outbound_type}.
	OutboundType *string `field:"optional" json:"outboundType" yaml:"outboundType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#pod_cidr KubernetesCluster#pod_cidr}.
	PodCidr *string `field:"optional" json:"podCidr" yaml:"podCidr"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#service_cidr KubernetesCluster#service_cidr}.
	ServiceCidr *string `field:"optional" json:"serviceCidr" yaml:"serviceCidr"`
}

type KubernetesClusterNetworkProfileLoadBalancerProfile struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#idle_timeout_in_minutes KubernetesCluster#idle_timeout_in_minutes}.
	IdleTimeoutInMinutes *float64 `field:"optional" json:"idleTimeoutInMinutes" yaml:"idleTimeoutInMinutes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#managed_outbound_ip_count KubernetesCluster#managed_outbound_ip_count}.
	ManagedOutboundIpCount *float64 `field:"optional" json:"managedOutboundIpCount" yaml:"managedOutboundIpCount"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#outbound_ip_address_ids KubernetesCluster#outbound_ip_address_ids}.
	OutboundIpAddressIds *[]*string `field:"optional" json:"outboundIpAddressIds" yaml:"outboundIpAddressIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#outbound_ip_prefix_ids KubernetesCluster#outbound_ip_prefix_ids}.
	OutboundIpPrefixIds *[]*string `field:"optional" json:"outboundIpPrefixIds" yaml:"outboundIpPrefixIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#outbound_ports_allocated KubernetesCluster#outbound_ports_allocated}.
	OutboundPortsAllocated *float64 `field:"optional" json:"outboundPortsAllocated" yaml:"outboundPortsAllocated"`
}

type KubernetesClusterNetworkProfileLoadBalancerProfileOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	EffectiveOutboundIps() *[]*string
	// Experimental.
	Fqn() *string
	IdleTimeoutInMinutes() *float64
	SetIdleTimeoutInMinutes(val *float64)
	IdleTimeoutInMinutesInput() *float64
	InternalValue() *KubernetesClusterNetworkProfileLoadBalancerProfile
	SetInternalValue(val *KubernetesClusterNetworkProfileLoadBalancerProfile)
	ManagedOutboundIpCount() *float64
	SetManagedOutboundIpCount(val *float64)
	ManagedOutboundIpCountInput() *float64
	OutboundIpAddressIds() *[]*string
	SetOutboundIpAddressIds(val *[]*string)
	OutboundIpAddressIdsInput() *[]*string
	OutboundIpPrefixIds() *[]*string
	SetOutboundIpPrefixIds(val *[]*string)
	OutboundIpPrefixIdsInput() *[]*string
	OutboundPortsAllocated() *float64
	SetOutboundPortsAllocated(val *float64)
	OutboundPortsAllocatedInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetIdleTimeoutInMinutes()
	ResetManagedOutboundIpCount()
	ResetOutboundIpAddressIds()
	ResetOutboundIpPrefixIds()
	ResetOutboundPortsAllocated()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KubernetesClusterNetworkProfileLoadBalancerProfileOutputReference
type jsiiProxy_KubernetesClusterNetworkProfileLoadBalancerProfileOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileLoadBalancerProfileOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileLoadBalancerProfileOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileLoadBalancerProfileOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileLoadBalancerProfileOutputReference) EffectiveOutboundIps() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"effectiveOutboundIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileLoadBalancerProfileOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileLoadBalancerProfileOutputReference) IdleTimeoutInMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"idleTimeoutInMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileLoadBalancerProfileOutputReference) IdleTimeoutInMinutesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"idleTimeoutInMinutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileLoadBalancerProfileOutputReference) InternalValue() *KubernetesClusterNetworkProfileLoadBalancerProfile {
	var returns *KubernetesClusterNetworkProfileLoadBalancerProfile
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileLoadBalancerProfileOutputReference) ManagedOutboundIpCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"managedOutboundIpCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileLoadBalancerProfileOutputReference) ManagedOutboundIpCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"managedOutboundIpCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileLoadBalancerProfileOutputReference) OutboundIpAddressIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"outboundIpAddressIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileLoadBalancerProfileOutputReference) OutboundIpAddressIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"outboundIpAddressIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileLoadBalancerProfileOutputReference) OutboundIpPrefixIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"outboundIpPrefixIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileLoadBalancerProfileOutputReference) OutboundIpPrefixIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"outboundIpPrefixIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileLoadBalancerProfileOutputReference) OutboundPortsAllocated() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"outboundPortsAllocated",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileLoadBalancerProfileOutputReference) OutboundPortsAllocatedInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"outboundPortsAllocatedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileLoadBalancerProfileOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileLoadBalancerProfileOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewKubernetesClusterNetworkProfileLoadBalancerProfileOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) KubernetesClusterNetworkProfileLoadBalancerProfileOutputReference {
	_init_.Initialize()

	j := jsiiProxy_KubernetesClusterNetworkProfileLoadBalancerProfileOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesCluster.KubernetesClusterNetworkProfileLoadBalancerProfileOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKubernetesClusterNetworkProfileLoadBalancerProfileOutputReference_Override(k KubernetesClusterNetworkProfileLoadBalancerProfileOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesCluster.KubernetesClusterNetworkProfileLoadBalancerProfileOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileLoadBalancerProfileOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileLoadBalancerProfileOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileLoadBalancerProfileOutputReference) SetIdleTimeoutInMinutes(val *float64) {
	_jsii_.Set(
		j,
		"idleTimeoutInMinutes",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileLoadBalancerProfileOutputReference) SetInternalValue(val *KubernetesClusterNetworkProfileLoadBalancerProfile) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileLoadBalancerProfileOutputReference) SetManagedOutboundIpCount(val *float64) {
	_jsii_.Set(
		j,
		"managedOutboundIpCount",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileLoadBalancerProfileOutputReference) SetOutboundIpAddressIds(val *[]*string) {
	_jsii_.Set(
		j,
		"outboundIpAddressIds",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileLoadBalancerProfileOutputReference) SetOutboundIpPrefixIds(val *[]*string) {
	_jsii_.Set(
		j,
		"outboundIpPrefixIds",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileLoadBalancerProfileOutputReference) SetOutboundPortsAllocated(val *float64) {
	_jsii_.Set(
		j,
		"outboundPortsAllocated",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileLoadBalancerProfileOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileLoadBalancerProfileOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (k *jsiiProxy_KubernetesClusterNetworkProfileLoadBalancerProfileOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNetworkProfileLoadBalancerProfileOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNetworkProfileLoadBalancerProfileOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNetworkProfileLoadBalancerProfileOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		k,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNetworkProfileLoadBalancerProfileOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNetworkProfileLoadBalancerProfileOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		k,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNetworkProfileLoadBalancerProfileOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		k,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNetworkProfileLoadBalancerProfileOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		k,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNetworkProfileLoadBalancerProfileOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNetworkProfileLoadBalancerProfileOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		k,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNetworkProfileLoadBalancerProfileOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNetworkProfileLoadBalancerProfileOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNetworkProfileLoadBalancerProfileOutputReference) ResetIdleTimeoutInMinutes() {
	_jsii_.InvokeVoid(
		k,
		"resetIdleTimeoutInMinutes",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNetworkProfileLoadBalancerProfileOutputReference) ResetManagedOutboundIpCount() {
	_jsii_.InvokeVoid(
		k,
		"resetManagedOutboundIpCount",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNetworkProfileLoadBalancerProfileOutputReference) ResetOutboundIpAddressIds() {
	_jsii_.InvokeVoid(
		k,
		"resetOutboundIpAddressIds",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNetworkProfileLoadBalancerProfileOutputReference) ResetOutboundIpPrefixIds() {
	_jsii_.InvokeVoid(
		k,
		"resetOutboundIpPrefixIds",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNetworkProfileLoadBalancerProfileOutputReference) ResetOutboundPortsAllocated() {
	_jsii_.InvokeVoid(
		k,
		"resetOutboundPortsAllocated",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNetworkProfileLoadBalancerProfileOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNetworkProfileLoadBalancerProfileOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type KubernetesClusterNetworkProfileNatGatewayProfile struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#idle_timeout_in_minutes KubernetesCluster#idle_timeout_in_minutes}.
	IdleTimeoutInMinutes *float64 `field:"optional" json:"idleTimeoutInMinutes" yaml:"idleTimeoutInMinutes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#managed_outbound_ip_count KubernetesCluster#managed_outbound_ip_count}.
	ManagedOutboundIpCount *float64 `field:"optional" json:"managedOutboundIpCount" yaml:"managedOutboundIpCount"`
}

type KubernetesClusterNetworkProfileNatGatewayProfileOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	EffectiveOutboundIps() *[]*string
	// Experimental.
	Fqn() *string
	IdleTimeoutInMinutes() *float64
	SetIdleTimeoutInMinutes(val *float64)
	IdleTimeoutInMinutesInput() *float64
	InternalValue() *KubernetesClusterNetworkProfileNatGatewayProfile
	SetInternalValue(val *KubernetesClusterNetworkProfileNatGatewayProfile)
	ManagedOutboundIpCount() *float64
	SetManagedOutboundIpCount(val *float64)
	ManagedOutboundIpCountInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetIdleTimeoutInMinutes()
	ResetManagedOutboundIpCount()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KubernetesClusterNetworkProfileNatGatewayProfileOutputReference
type jsiiProxy_KubernetesClusterNetworkProfileNatGatewayProfileOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileNatGatewayProfileOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileNatGatewayProfileOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileNatGatewayProfileOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileNatGatewayProfileOutputReference) EffectiveOutboundIps() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"effectiveOutboundIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileNatGatewayProfileOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileNatGatewayProfileOutputReference) IdleTimeoutInMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"idleTimeoutInMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileNatGatewayProfileOutputReference) IdleTimeoutInMinutesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"idleTimeoutInMinutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileNatGatewayProfileOutputReference) InternalValue() *KubernetesClusterNetworkProfileNatGatewayProfile {
	var returns *KubernetesClusterNetworkProfileNatGatewayProfile
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileNatGatewayProfileOutputReference) ManagedOutboundIpCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"managedOutboundIpCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileNatGatewayProfileOutputReference) ManagedOutboundIpCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"managedOutboundIpCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileNatGatewayProfileOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileNatGatewayProfileOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewKubernetesClusterNetworkProfileNatGatewayProfileOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) KubernetesClusterNetworkProfileNatGatewayProfileOutputReference {
	_init_.Initialize()

	j := jsiiProxy_KubernetesClusterNetworkProfileNatGatewayProfileOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesCluster.KubernetesClusterNetworkProfileNatGatewayProfileOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKubernetesClusterNetworkProfileNatGatewayProfileOutputReference_Override(k KubernetesClusterNetworkProfileNatGatewayProfileOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesCluster.KubernetesClusterNetworkProfileNatGatewayProfileOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileNatGatewayProfileOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileNatGatewayProfileOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileNatGatewayProfileOutputReference) SetIdleTimeoutInMinutes(val *float64) {
	_jsii_.Set(
		j,
		"idleTimeoutInMinutes",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileNatGatewayProfileOutputReference) SetInternalValue(val *KubernetesClusterNetworkProfileNatGatewayProfile) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileNatGatewayProfileOutputReference) SetManagedOutboundIpCount(val *float64) {
	_jsii_.Set(
		j,
		"managedOutboundIpCount",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileNatGatewayProfileOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileNatGatewayProfileOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (k *jsiiProxy_KubernetesClusterNetworkProfileNatGatewayProfileOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNetworkProfileNatGatewayProfileOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNetworkProfileNatGatewayProfileOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNetworkProfileNatGatewayProfileOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		k,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNetworkProfileNatGatewayProfileOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNetworkProfileNatGatewayProfileOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		k,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNetworkProfileNatGatewayProfileOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		k,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNetworkProfileNatGatewayProfileOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		k,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNetworkProfileNatGatewayProfileOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNetworkProfileNatGatewayProfileOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		k,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNetworkProfileNatGatewayProfileOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNetworkProfileNatGatewayProfileOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNetworkProfileNatGatewayProfileOutputReference) ResetIdleTimeoutInMinutes() {
	_jsii_.InvokeVoid(
		k,
		"resetIdleTimeoutInMinutes",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNetworkProfileNatGatewayProfileOutputReference) ResetManagedOutboundIpCount() {
	_jsii_.InvokeVoid(
		k,
		"resetManagedOutboundIpCount",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNetworkProfileNatGatewayProfileOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNetworkProfileNatGatewayProfileOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type KubernetesClusterNetworkProfileOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DnsServiceIp() *string
	SetDnsServiceIp(val *string)
	DnsServiceIpInput() *string
	DockerBridgeCidr() *string
	SetDockerBridgeCidr(val *string)
	DockerBridgeCidrInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *KubernetesClusterNetworkProfile
	SetInternalValue(val *KubernetesClusterNetworkProfile)
	IpVersions() *[]*string
	SetIpVersions(val *[]*string)
	IpVersionsInput() *[]*string
	LoadBalancerProfile() KubernetesClusterNetworkProfileLoadBalancerProfileOutputReference
	LoadBalancerProfileInput() *KubernetesClusterNetworkProfileLoadBalancerProfile
	LoadBalancerSku() *string
	SetLoadBalancerSku(val *string)
	LoadBalancerSkuInput() *string
	NatGatewayProfile() KubernetesClusterNetworkProfileNatGatewayProfileOutputReference
	NatGatewayProfileInput() *KubernetesClusterNetworkProfileNatGatewayProfile
	NetworkMode() *string
	SetNetworkMode(val *string)
	NetworkModeInput() *string
	NetworkPlugin() *string
	SetNetworkPlugin(val *string)
	NetworkPluginInput() *string
	NetworkPolicy() *string
	SetNetworkPolicy(val *string)
	NetworkPolicyInput() *string
	OutboundType() *string
	SetOutboundType(val *string)
	OutboundTypeInput() *string
	PodCidr() *string
	SetPodCidr(val *string)
	PodCidrInput() *string
	ServiceCidr() *string
	SetServiceCidr(val *string)
	ServiceCidrInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	PutLoadBalancerProfile(value *KubernetesClusterNetworkProfileLoadBalancerProfile)
	PutNatGatewayProfile(value *KubernetesClusterNetworkProfileNatGatewayProfile)
	ResetDnsServiceIp()
	ResetDockerBridgeCidr()
	ResetIpVersions()
	ResetLoadBalancerProfile()
	ResetLoadBalancerSku()
	ResetNatGatewayProfile()
	ResetNetworkMode()
	ResetNetworkPolicy()
	ResetOutboundType()
	ResetPodCidr()
	ResetServiceCidr()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KubernetesClusterNetworkProfileOutputReference
type jsiiProxy_KubernetesClusterNetworkProfileOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) DnsServiceIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsServiceIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) DnsServiceIpInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsServiceIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) DockerBridgeCidr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dockerBridgeCidr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) DockerBridgeCidrInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dockerBridgeCidrInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) InternalValue() *KubernetesClusterNetworkProfile {
	var returns *KubernetesClusterNetworkProfile
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) IpVersions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipVersions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) IpVersionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipVersionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) LoadBalancerProfile() KubernetesClusterNetworkProfileLoadBalancerProfileOutputReference {
	var returns KubernetesClusterNetworkProfileLoadBalancerProfileOutputReference
	_jsii_.Get(
		j,
		"loadBalancerProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) LoadBalancerProfileInput() *KubernetesClusterNetworkProfileLoadBalancerProfile {
	var returns *KubernetesClusterNetworkProfileLoadBalancerProfile
	_jsii_.Get(
		j,
		"loadBalancerProfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) LoadBalancerSku() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancerSku",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) LoadBalancerSkuInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancerSkuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) NatGatewayProfile() KubernetesClusterNetworkProfileNatGatewayProfileOutputReference {
	var returns KubernetesClusterNetworkProfileNatGatewayProfileOutputReference
	_jsii_.Get(
		j,
		"natGatewayProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) NatGatewayProfileInput() *KubernetesClusterNetworkProfileNatGatewayProfile {
	var returns *KubernetesClusterNetworkProfileNatGatewayProfile
	_jsii_.Get(
		j,
		"natGatewayProfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) NetworkMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) NetworkModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) NetworkPlugin() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkPlugin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) NetworkPluginInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkPluginInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) NetworkPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) NetworkPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) OutboundType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outboundType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) OutboundTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outboundTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) PodCidr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"podCidr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) PodCidrInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"podCidrInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) ServiceCidr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceCidr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) ServiceCidrInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceCidrInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewKubernetesClusterNetworkProfileOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) KubernetesClusterNetworkProfileOutputReference {
	_init_.Initialize()

	j := jsiiProxy_KubernetesClusterNetworkProfileOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesCluster.KubernetesClusterNetworkProfileOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKubernetesClusterNetworkProfileOutputReference_Override(k KubernetesClusterNetworkProfileOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesCluster.KubernetesClusterNetworkProfileOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) SetDnsServiceIp(val *string) {
	_jsii_.Set(
		j,
		"dnsServiceIp",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) SetDockerBridgeCidr(val *string) {
	_jsii_.Set(
		j,
		"dockerBridgeCidr",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) SetInternalValue(val *KubernetesClusterNetworkProfile) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) SetIpVersions(val *[]*string) {
	_jsii_.Set(
		j,
		"ipVersions",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) SetLoadBalancerSku(val *string) {
	_jsii_.Set(
		j,
		"loadBalancerSku",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) SetNetworkMode(val *string) {
	_jsii_.Set(
		j,
		"networkMode",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) SetNetworkPlugin(val *string) {
	_jsii_.Set(
		j,
		"networkPlugin",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) SetNetworkPolicy(val *string) {
	_jsii_.Set(
		j,
		"networkPolicy",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) SetOutboundType(val *string) {
	_jsii_.Set(
		j,
		"outboundType",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) SetPodCidr(val *string) {
	_jsii_.Set(
		j,
		"podCidr",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) SetServiceCidr(val *string) {
	_jsii_.Set(
		j,
		"serviceCidr",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (k *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		k,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		k,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		k,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		k,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		k,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) PutLoadBalancerProfile(value *KubernetesClusterNetworkProfileLoadBalancerProfile) {
	_jsii_.InvokeVoid(
		k,
		"putLoadBalancerProfile",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) PutNatGatewayProfile(value *KubernetesClusterNetworkProfileNatGatewayProfile) {
	_jsii_.InvokeVoid(
		k,
		"putNatGatewayProfile",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) ResetDnsServiceIp() {
	_jsii_.InvokeVoid(
		k,
		"resetDnsServiceIp",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) ResetDockerBridgeCidr() {
	_jsii_.InvokeVoid(
		k,
		"resetDockerBridgeCidr",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) ResetIpVersions() {
	_jsii_.InvokeVoid(
		k,
		"resetIpVersions",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) ResetLoadBalancerProfile() {
	_jsii_.InvokeVoid(
		k,
		"resetLoadBalancerProfile",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) ResetLoadBalancerSku() {
	_jsii_.InvokeVoid(
		k,
		"resetLoadBalancerSku",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) ResetNatGatewayProfile() {
	_jsii_.InvokeVoid(
		k,
		"resetNatGatewayProfile",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) ResetNetworkMode() {
	_jsii_.InvokeVoid(
		k,
		"resetNetworkMode",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) ResetNetworkPolicy() {
	_jsii_.InvokeVoid(
		k,
		"resetNetworkPolicy",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) ResetOutboundType() {
	_jsii_.InvokeVoid(
		k,
		"resetOutboundType",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) ResetPodCidr() {
	_jsii_.InvokeVoid(
		k,
		"resetPodCidr",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) ResetServiceCidr() {
	_jsii_.InvokeVoid(
		k,
		"resetServiceCidr",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNetworkProfileOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type KubernetesClusterOmsAgent struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#log_analytics_workspace_id KubernetesCluster#log_analytics_workspace_id}.
	LogAnalyticsWorkspaceId *string `field:"required" json:"logAnalyticsWorkspaceId" yaml:"logAnalyticsWorkspaceId"`
}

type KubernetesClusterOmsAgentOmsAgentIdentity struct {
}

type KubernetesClusterOmsAgentOmsAgentIdentityList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) KubernetesClusterOmsAgentOmsAgentIdentityOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KubernetesClusterOmsAgentOmsAgentIdentityList
type jsiiProxy_KubernetesClusterOmsAgentOmsAgentIdentityList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_KubernetesClusterOmsAgentOmsAgentIdentityList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterOmsAgentOmsAgentIdentityList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterOmsAgentOmsAgentIdentityList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterOmsAgentOmsAgentIdentityList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterOmsAgentOmsAgentIdentityList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewKubernetesClusterOmsAgentOmsAgentIdentityList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) KubernetesClusterOmsAgentOmsAgentIdentityList {
	_init_.Initialize()

	j := jsiiProxy_KubernetesClusterOmsAgentOmsAgentIdentityList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesCluster.KubernetesClusterOmsAgentOmsAgentIdentityList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewKubernetesClusterOmsAgentOmsAgentIdentityList_Override(k KubernetesClusterOmsAgentOmsAgentIdentityList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesCluster.KubernetesClusterOmsAgentOmsAgentIdentityList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		k,
	)
}

func (j *jsiiProxy_KubernetesClusterOmsAgentOmsAgentIdentityList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterOmsAgentOmsAgentIdentityList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterOmsAgentOmsAgentIdentityList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (k *jsiiProxy_KubernetesClusterOmsAgentOmsAgentIdentityList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterOmsAgentOmsAgentIdentityList) Get(index *float64) KubernetesClusterOmsAgentOmsAgentIdentityOutputReference {
	var returns KubernetesClusterOmsAgentOmsAgentIdentityOutputReference

	_jsii_.Invoke(
		k,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterOmsAgentOmsAgentIdentityList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterOmsAgentOmsAgentIdentityList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type KubernetesClusterOmsAgentOmsAgentIdentityOutputReference interface {
	cdktf.ComplexObject
	ClientId() *string
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *KubernetesClusterOmsAgentOmsAgentIdentity
	SetInternalValue(val *KubernetesClusterOmsAgentOmsAgentIdentity)
	ObjectId() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UserAssignedIdentityId() *string
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KubernetesClusterOmsAgentOmsAgentIdentityOutputReference
type jsiiProxy_KubernetesClusterOmsAgentOmsAgentIdentityOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_KubernetesClusterOmsAgentOmsAgentIdentityOutputReference) ClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterOmsAgentOmsAgentIdentityOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterOmsAgentOmsAgentIdentityOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterOmsAgentOmsAgentIdentityOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterOmsAgentOmsAgentIdentityOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterOmsAgentOmsAgentIdentityOutputReference) InternalValue() *KubernetesClusterOmsAgentOmsAgentIdentity {
	var returns *KubernetesClusterOmsAgentOmsAgentIdentity
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterOmsAgentOmsAgentIdentityOutputReference) ObjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterOmsAgentOmsAgentIdentityOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterOmsAgentOmsAgentIdentityOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterOmsAgentOmsAgentIdentityOutputReference) UserAssignedIdentityId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userAssignedIdentityId",
		&returns,
	)
	return returns
}


func NewKubernetesClusterOmsAgentOmsAgentIdentityOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) KubernetesClusterOmsAgentOmsAgentIdentityOutputReference {
	_init_.Initialize()

	j := jsiiProxy_KubernetesClusterOmsAgentOmsAgentIdentityOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesCluster.KubernetesClusterOmsAgentOmsAgentIdentityOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewKubernetesClusterOmsAgentOmsAgentIdentityOutputReference_Override(k KubernetesClusterOmsAgentOmsAgentIdentityOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesCluster.KubernetesClusterOmsAgentOmsAgentIdentityOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		k,
	)
}

func (j *jsiiProxy_KubernetesClusterOmsAgentOmsAgentIdentityOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterOmsAgentOmsAgentIdentityOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterOmsAgentOmsAgentIdentityOutputReference) SetInternalValue(val *KubernetesClusterOmsAgentOmsAgentIdentity) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterOmsAgentOmsAgentIdentityOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterOmsAgentOmsAgentIdentityOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (k *jsiiProxy_KubernetesClusterOmsAgentOmsAgentIdentityOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterOmsAgentOmsAgentIdentityOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterOmsAgentOmsAgentIdentityOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterOmsAgentOmsAgentIdentityOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		k,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterOmsAgentOmsAgentIdentityOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterOmsAgentOmsAgentIdentityOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		k,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterOmsAgentOmsAgentIdentityOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		k,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterOmsAgentOmsAgentIdentityOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		k,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterOmsAgentOmsAgentIdentityOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterOmsAgentOmsAgentIdentityOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		k,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterOmsAgentOmsAgentIdentityOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterOmsAgentOmsAgentIdentityOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterOmsAgentOmsAgentIdentityOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterOmsAgentOmsAgentIdentityOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type KubernetesClusterOmsAgentOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *KubernetesClusterOmsAgent
	SetInternalValue(val *KubernetesClusterOmsAgent)
	LogAnalyticsWorkspaceId() *string
	SetLogAnalyticsWorkspaceId(val *string)
	LogAnalyticsWorkspaceIdInput() *string
	OmsAgentIdentity() KubernetesClusterOmsAgentOmsAgentIdentityList
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KubernetesClusterOmsAgentOutputReference
type jsiiProxy_KubernetesClusterOmsAgentOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_KubernetesClusterOmsAgentOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterOmsAgentOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterOmsAgentOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterOmsAgentOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterOmsAgentOutputReference) InternalValue() *KubernetesClusterOmsAgent {
	var returns *KubernetesClusterOmsAgent
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterOmsAgentOutputReference) LogAnalyticsWorkspaceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logAnalyticsWorkspaceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterOmsAgentOutputReference) LogAnalyticsWorkspaceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logAnalyticsWorkspaceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterOmsAgentOutputReference) OmsAgentIdentity() KubernetesClusterOmsAgentOmsAgentIdentityList {
	var returns KubernetesClusterOmsAgentOmsAgentIdentityList
	_jsii_.Get(
		j,
		"omsAgentIdentity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterOmsAgentOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterOmsAgentOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewKubernetesClusterOmsAgentOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) KubernetesClusterOmsAgentOutputReference {
	_init_.Initialize()

	j := jsiiProxy_KubernetesClusterOmsAgentOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesCluster.KubernetesClusterOmsAgentOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKubernetesClusterOmsAgentOutputReference_Override(k KubernetesClusterOmsAgentOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesCluster.KubernetesClusterOmsAgentOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_KubernetesClusterOmsAgentOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterOmsAgentOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterOmsAgentOutputReference) SetInternalValue(val *KubernetesClusterOmsAgent) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterOmsAgentOutputReference) SetLogAnalyticsWorkspaceId(val *string) {
	_jsii_.Set(
		j,
		"logAnalyticsWorkspaceId",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterOmsAgentOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterOmsAgentOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (k *jsiiProxy_KubernetesClusterOmsAgentOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterOmsAgentOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterOmsAgentOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterOmsAgentOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		k,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterOmsAgentOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterOmsAgentOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		k,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterOmsAgentOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		k,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterOmsAgentOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		k,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterOmsAgentOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterOmsAgentOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		k,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterOmsAgentOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterOmsAgentOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterOmsAgentOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterOmsAgentOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type KubernetesClusterServicePrincipal struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#client_id KubernetesCluster#client_id}.
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#client_secret KubernetesCluster#client_secret}.
	ClientSecret *string `field:"required" json:"clientSecret" yaml:"clientSecret"`
}

type KubernetesClusterServicePrincipalOutputReference interface {
	cdktf.ComplexObject
	ClientId() *string
	SetClientId(val *string)
	ClientIdInput() *string
	ClientSecret() *string
	SetClientSecret(val *string)
	ClientSecretInput() *string
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *KubernetesClusterServicePrincipal
	SetInternalValue(val *KubernetesClusterServicePrincipal)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KubernetesClusterServicePrincipalOutputReference
type jsiiProxy_KubernetesClusterServicePrincipalOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_KubernetesClusterServicePrincipalOutputReference) ClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterServicePrincipalOutputReference) ClientIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterServicePrincipalOutputReference) ClientSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterServicePrincipalOutputReference) ClientSecretInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterServicePrincipalOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterServicePrincipalOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterServicePrincipalOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterServicePrincipalOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterServicePrincipalOutputReference) InternalValue() *KubernetesClusterServicePrincipal {
	var returns *KubernetesClusterServicePrincipal
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterServicePrincipalOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterServicePrincipalOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewKubernetesClusterServicePrincipalOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) KubernetesClusterServicePrincipalOutputReference {
	_init_.Initialize()

	j := jsiiProxy_KubernetesClusterServicePrincipalOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesCluster.KubernetesClusterServicePrincipalOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKubernetesClusterServicePrincipalOutputReference_Override(k KubernetesClusterServicePrincipalOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesCluster.KubernetesClusterServicePrincipalOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_KubernetesClusterServicePrincipalOutputReference) SetClientId(val *string) {
	_jsii_.Set(
		j,
		"clientId",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterServicePrincipalOutputReference) SetClientSecret(val *string) {
	_jsii_.Set(
		j,
		"clientSecret",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterServicePrincipalOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterServicePrincipalOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterServicePrincipalOutputReference) SetInternalValue(val *KubernetesClusterServicePrincipal) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterServicePrincipalOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterServicePrincipalOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (k *jsiiProxy_KubernetesClusterServicePrincipalOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterServicePrincipalOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterServicePrincipalOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterServicePrincipalOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		k,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterServicePrincipalOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterServicePrincipalOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		k,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterServicePrincipalOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		k,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterServicePrincipalOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		k,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterServicePrincipalOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterServicePrincipalOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		k,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterServicePrincipalOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterServicePrincipalOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterServicePrincipalOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterServicePrincipalOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type KubernetesClusterTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#create KubernetesCluster#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#delete KubernetesCluster#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#read KubernetesCluster#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#update KubernetesCluster#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

type KubernetesClusterTimeoutsOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	Create() *string
	SetCreate(val *string)
	CreateInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Delete() *string
	SetDelete(val *string)
	DeleteInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Read() *string
	SetRead(val *string)
	ReadInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Update() *string
	SetUpdate(val *string)
	UpdateInput() *string
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetCreate()
	ResetDelete()
	ResetRead()
	ResetUpdate()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KubernetesClusterTimeoutsOutputReference
type jsiiProxy_KubernetesClusterTimeoutsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_KubernetesClusterTimeoutsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterTimeoutsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterTimeoutsOutputReference) Create() *string {
	var returns *string
	_jsii_.Get(
		j,
		"create",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterTimeoutsOutputReference) CreateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterTimeoutsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterTimeoutsOutputReference) Delete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterTimeoutsOutputReference) DeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterTimeoutsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterTimeoutsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterTimeoutsOutputReference) Read() *string {
	var returns *string
	_jsii_.Get(
		j,
		"read",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterTimeoutsOutputReference) ReadInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterTimeoutsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterTimeoutsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterTimeoutsOutputReference) Update() *string {
	var returns *string
	_jsii_.Get(
		j,
		"update",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterTimeoutsOutputReference) UpdateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateInput",
		&returns,
	)
	return returns
}


func NewKubernetesClusterTimeoutsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) KubernetesClusterTimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_KubernetesClusterTimeoutsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesCluster.KubernetesClusterTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKubernetesClusterTimeoutsOutputReference_Override(k KubernetesClusterTimeoutsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesCluster.KubernetesClusterTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_KubernetesClusterTimeoutsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterTimeoutsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterTimeoutsOutputReference) SetCreate(val *string) {
	_jsii_.Set(
		j,
		"create",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterTimeoutsOutputReference) SetDelete(val *string) {
	_jsii_.Set(
		j,
		"delete",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterTimeoutsOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterTimeoutsOutputReference) SetRead(val *string) {
	_jsii_.Set(
		j,
		"read",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterTimeoutsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterTimeoutsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterTimeoutsOutputReference) SetUpdate(val *string) {
	_jsii_.Set(
		j,
		"update",
		val,
	)
}

func (k *jsiiProxy_KubernetesClusterTimeoutsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterTimeoutsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterTimeoutsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		k,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterTimeoutsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterTimeoutsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		k,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterTimeoutsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		k,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterTimeoutsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		k,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterTimeoutsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterTimeoutsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		k,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterTimeoutsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterTimeoutsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterTimeoutsOutputReference) ResetCreate() {
	_jsii_.InvokeVoid(
		k,
		"resetCreate",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterTimeoutsOutputReference) ResetDelete() {
	_jsii_.InvokeVoid(
		k,
		"resetDelete",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterTimeoutsOutputReference) ResetRead() {
	_jsii_.InvokeVoid(
		k,
		"resetRead",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterTimeoutsOutputReference) ResetUpdate() {
	_jsii_.InvokeVoid(
		k,
		"resetUpdate",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterTimeoutsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterTimeoutsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type KubernetesClusterWindowsProfile struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#admin_username KubernetesCluster#admin_username}.
	AdminUsername *string `field:"required" json:"adminUsername" yaml:"adminUsername"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#admin_password KubernetesCluster#admin_password}.
	AdminPassword *string `field:"optional" json:"adminPassword" yaml:"adminPassword"`
	// gmsa block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#gmsa KubernetesCluster#gmsa}
	Gmsa *KubernetesClusterWindowsProfileGmsa `field:"optional" json:"gmsa" yaml:"gmsa"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#license KubernetesCluster#license}.
	License *string `field:"optional" json:"license" yaml:"license"`
}

type KubernetesClusterWindowsProfileGmsa struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#dns_server KubernetesCluster#dns_server}.
	DnsServer *string `field:"required" json:"dnsServer" yaml:"dnsServer"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#root_domain KubernetesCluster#root_domain}.
	RootDomain *string `field:"required" json:"rootDomain" yaml:"rootDomain"`
}

type KubernetesClusterWindowsProfileGmsaOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DnsServer() *string
	SetDnsServer(val *string)
	DnsServerInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *KubernetesClusterWindowsProfileGmsa
	SetInternalValue(val *KubernetesClusterWindowsProfileGmsa)
	RootDomain() *string
	SetRootDomain(val *string)
	RootDomainInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KubernetesClusterWindowsProfileGmsaOutputReference
type jsiiProxy_KubernetesClusterWindowsProfileGmsaOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_KubernetesClusterWindowsProfileGmsaOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterWindowsProfileGmsaOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterWindowsProfileGmsaOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterWindowsProfileGmsaOutputReference) DnsServer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsServer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterWindowsProfileGmsaOutputReference) DnsServerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsServerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterWindowsProfileGmsaOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterWindowsProfileGmsaOutputReference) InternalValue() *KubernetesClusterWindowsProfileGmsa {
	var returns *KubernetesClusterWindowsProfileGmsa
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterWindowsProfileGmsaOutputReference) RootDomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rootDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterWindowsProfileGmsaOutputReference) RootDomainInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rootDomainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterWindowsProfileGmsaOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterWindowsProfileGmsaOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewKubernetesClusterWindowsProfileGmsaOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) KubernetesClusterWindowsProfileGmsaOutputReference {
	_init_.Initialize()

	j := jsiiProxy_KubernetesClusterWindowsProfileGmsaOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesCluster.KubernetesClusterWindowsProfileGmsaOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKubernetesClusterWindowsProfileGmsaOutputReference_Override(k KubernetesClusterWindowsProfileGmsaOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesCluster.KubernetesClusterWindowsProfileGmsaOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_KubernetesClusterWindowsProfileGmsaOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterWindowsProfileGmsaOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterWindowsProfileGmsaOutputReference) SetDnsServer(val *string) {
	_jsii_.Set(
		j,
		"dnsServer",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterWindowsProfileGmsaOutputReference) SetInternalValue(val *KubernetesClusterWindowsProfileGmsa) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterWindowsProfileGmsaOutputReference) SetRootDomain(val *string) {
	_jsii_.Set(
		j,
		"rootDomain",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterWindowsProfileGmsaOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterWindowsProfileGmsaOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (k *jsiiProxy_KubernetesClusterWindowsProfileGmsaOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterWindowsProfileGmsaOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterWindowsProfileGmsaOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterWindowsProfileGmsaOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		k,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterWindowsProfileGmsaOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterWindowsProfileGmsaOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		k,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterWindowsProfileGmsaOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		k,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterWindowsProfileGmsaOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		k,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterWindowsProfileGmsaOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterWindowsProfileGmsaOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		k,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterWindowsProfileGmsaOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterWindowsProfileGmsaOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterWindowsProfileGmsaOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterWindowsProfileGmsaOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type KubernetesClusterWindowsProfileOutputReference interface {
	cdktf.ComplexObject
	AdminPassword() *string
	SetAdminPassword(val *string)
	AdminPasswordInput() *string
	AdminUsername() *string
	SetAdminUsername(val *string)
	AdminUsernameInput() *string
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	Gmsa() KubernetesClusterWindowsProfileGmsaOutputReference
	GmsaInput() *KubernetesClusterWindowsProfileGmsa
	InternalValue() *KubernetesClusterWindowsProfile
	SetInternalValue(val *KubernetesClusterWindowsProfile)
	License() *string
	SetLicense(val *string)
	LicenseInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	PutGmsa(value *KubernetesClusterWindowsProfileGmsa)
	ResetAdminPassword()
	ResetGmsa()
	ResetLicense()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KubernetesClusterWindowsProfileOutputReference
type jsiiProxy_KubernetesClusterWindowsProfileOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_KubernetesClusterWindowsProfileOutputReference) AdminPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterWindowsProfileOutputReference) AdminPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterWindowsProfileOutputReference) AdminUsername() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminUsername",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterWindowsProfileOutputReference) AdminUsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminUsernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterWindowsProfileOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterWindowsProfileOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterWindowsProfileOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterWindowsProfileOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterWindowsProfileOutputReference) Gmsa() KubernetesClusterWindowsProfileGmsaOutputReference {
	var returns KubernetesClusterWindowsProfileGmsaOutputReference
	_jsii_.Get(
		j,
		"gmsa",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterWindowsProfileOutputReference) GmsaInput() *KubernetesClusterWindowsProfileGmsa {
	var returns *KubernetesClusterWindowsProfileGmsa
	_jsii_.Get(
		j,
		"gmsaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterWindowsProfileOutputReference) InternalValue() *KubernetesClusterWindowsProfile {
	var returns *KubernetesClusterWindowsProfile
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterWindowsProfileOutputReference) License() *string {
	var returns *string
	_jsii_.Get(
		j,
		"license",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterWindowsProfileOutputReference) LicenseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterWindowsProfileOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterWindowsProfileOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewKubernetesClusterWindowsProfileOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) KubernetesClusterWindowsProfileOutputReference {
	_init_.Initialize()

	j := jsiiProxy_KubernetesClusterWindowsProfileOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesCluster.KubernetesClusterWindowsProfileOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKubernetesClusterWindowsProfileOutputReference_Override(k KubernetesClusterWindowsProfileOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesCluster.KubernetesClusterWindowsProfileOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_KubernetesClusterWindowsProfileOutputReference) SetAdminPassword(val *string) {
	_jsii_.Set(
		j,
		"adminPassword",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterWindowsProfileOutputReference) SetAdminUsername(val *string) {
	_jsii_.Set(
		j,
		"adminUsername",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterWindowsProfileOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterWindowsProfileOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterWindowsProfileOutputReference) SetInternalValue(val *KubernetesClusterWindowsProfile) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterWindowsProfileOutputReference) SetLicense(val *string) {
	_jsii_.Set(
		j,
		"license",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterWindowsProfileOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterWindowsProfileOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (k *jsiiProxy_KubernetesClusterWindowsProfileOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterWindowsProfileOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterWindowsProfileOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterWindowsProfileOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		k,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterWindowsProfileOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterWindowsProfileOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		k,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterWindowsProfileOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		k,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterWindowsProfileOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		k,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterWindowsProfileOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterWindowsProfileOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		k,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterWindowsProfileOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterWindowsProfileOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterWindowsProfileOutputReference) PutGmsa(value *KubernetesClusterWindowsProfileGmsa) {
	_jsii_.InvokeVoid(
		k,
		"putGmsa",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KubernetesClusterWindowsProfileOutputReference) ResetAdminPassword() {
	_jsii_.InvokeVoid(
		k,
		"resetAdminPassword",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterWindowsProfileOutputReference) ResetGmsa() {
	_jsii_.InvokeVoid(
		k,
		"resetGmsa",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterWindowsProfileOutputReference) ResetLicense() {
	_jsii_.InvokeVoid(
		k,
		"resetLicense",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterWindowsProfileOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterWindowsProfileOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

