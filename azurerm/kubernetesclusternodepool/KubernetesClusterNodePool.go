package kubernetesclusternodepool

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v9/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v9/kubernetesclusternodepool/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/kubernetes_cluster_node_pool azurerm_kubernetes_cluster_node_pool}.
type KubernetesClusterNodePool interface {
	cdktf.TerraformResource
	CapacityReservationGroupId() *string
	SetCapacityReservationGroupId(val *string)
	CapacityReservationGroupIdInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CustomCaTrustEnabled() interface{}
	SetCustomCaTrustEnabled(val interface{})
	CustomCaTrustEnabledInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EnableAutoScaling() interface{}
	SetEnableAutoScaling(val interface{})
	EnableAutoScalingInput() interface{}
	EnableHostEncryption() interface{}
	SetEnableHostEncryption(val interface{})
	EnableHostEncryptionInput() interface{}
	EnableNodePublicIp() interface{}
	SetEnableNodePublicIp(val interface{})
	EnableNodePublicIpInput() interface{}
	EvictionPolicy() *string
	SetEvictionPolicy(val *string)
	EvictionPolicyInput() *string
	FipsEnabled() interface{}
	SetFipsEnabled(val interface{})
	FipsEnabledInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HostGroupId() *string
	SetHostGroupId(val *string)
	HostGroupIdInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	KubeletConfig() KubernetesClusterNodePoolKubeletConfigOutputReference
	KubeletConfigInput() *KubernetesClusterNodePoolKubeletConfig
	KubeletDiskType() *string
	SetKubeletDiskType(val *string)
	KubeletDiskTypeInput() *string
	KubernetesClusterId() *string
	SetKubernetesClusterId(val *string)
	KubernetesClusterIdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LinuxOsConfig() KubernetesClusterNodePoolLinuxOsConfigOutputReference
	LinuxOsConfigInput() *KubernetesClusterNodePoolLinuxOsConfig
	MaxCount() *float64
	SetMaxCount(val *float64)
	MaxCountInput() *float64
	MaxPods() *float64
	SetMaxPods(val *float64)
	MaxPodsInput() *float64
	MessageOfTheDay() *string
	SetMessageOfTheDay(val *string)
	MessageOfTheDayInput() *string
	MinCount() *float64
	SetMinCount(val *float64)
	MinCountInput() *float64
	Mode() *string
	SetMode(val *string)
	ModeInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	NodeCount() *float64
	SetNodeCount(val *float64)
	NodeCountInput() *float64
	NodeLabels() *map[string]*string
	SetNodeLabels(val *map[string]*string)
	NodeLabelsInput() *map[string]*string
	NodeNetworkProfile() KubernetesClusterNodePoolNodeNetworkProfileOutputReference
	NodeNetworkProfileInput() *KubernetesClusterNodePoolNodeNetworkProfile
	NodePublicIpPrefixId() *string
	SetNodePublicIpPrefixId(val *string)
	NodePublicIpPrefixIdInput() *string
	NodeTaints() *[]*string
	SetNodeTaints(val *[]*string)
	NodeTaintsInput() *[]*string
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
	OsType() *string
	SetOsType(val *string)
	OsTypeInput() *string
	PodSubnetId() *string
	SetPodSubnetId(val *string)
	PodSubnetIdInput() *string
	Priority() *string
	SetPriority(val *string)
	PriorityInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	ProximityPlacementGroupId() *string
	SetProximityPlacementGroupId(val *string)
	ProximityPlacementGroupIdInput() *string
	// Experimental.
	RawOverrides() interface{}
	ScaleDownMode() *string
	SetScaleDownMode(val *string)
	ScaleDownModeInput() *string
	SnapshotId() *string
	SetSnapshotId(val *string)
	SnapshotIdInput() *string
	SpotMaxPrice() *float64
	SetSpotMaxPrice(val *float64)
	SpotMaxPriceInput() *float64
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() KubernetesClusterNodePoolTimeoutsOutputReference
	TimeoutsInput() interface{}
	UltraSsdEnabled() interface{}
	SetUltraSsdEnabled(val interface{})
	UltraSsdEnabledInput() interface{}
	UpgradeSettings() KubernetesClusterNodePoolUpgradeSettingsOutputReference
	UpgradeSettingsInput() *KubernetesClusterNodePoolUpgradeSettings
	VmSize() *string
	SetVmSize(val *string)
	VmSizeInput() *string
	VnetSubnetId() *string
	SetVnetSubnetId(val *string)
	VnetSubnetIdInput() *string
	WindowsProfile() KubernetesClusterNodePoolWindowsProfileOutputReference
	WindowsProfileInput() *KubernetesClusterNodePoolWindowsProfile
	WorkloadRuntime() *string
	SetWorkloadRuntime(val *string)
	WorkloadRuntimeInput() *string
	Zones() *[]*string
	SetZones(val *[]*string)
	ZonesInput() *[]*string
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
	PutKubeletConfig(value *KubernetesClusterNodePoolKubeletConfig)
	PutLinuxOsConfig(value *KubernetesClusterNodePoolLinuxOsConfig)
	PutNodeNetworkProfile(value *KubernetesClusterNodePoolNodeNetworkProfile)
	PutTimeouts(value *KubernetesClusterNodePoolTimeouts)
	PutUpgradeSettings(value *KubernetesClusterNodePoolUpgradeSettings)
	PutWindowsProfile(value *KubernetesClusterNodePoolWindowsProfile)
	ResetCapacityReservationGroupId()
	ResetCustomCaTrustEnabled()
	ResetEnableAutoScaling()
	ResetEnableHostEncryption()
	ResetEnableNodePublicIp()
	ResetEvictionPolicy()
	ResetFipsEnabled()
	ResetHostGroupId()
	ResetId()
	ResetKubeletConfig()
	ResetKubeletDiskType()
	ResetLinuxOsConfig()
	ResetMaxCount()
	ResetMaxPods()
	ResetMessageOfTheDay()
	ResetMinCount()
	ResetMode()
	ResetNodeCount()
	ResetNodeLabels()
	ResetNodeNetworkProfile()
	ResetNodePublicIpPrefixId()
	ResetNodeTaints()
	ResetOrchestratorVersion()
	ResetOsDiskSizeGb()
	ResetOsDiskType()
	ResetOsSku()
	ResetOsType()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPodSubnetId()
	ResetPriority()
	ResetProximityPlacementGroupId()
	ResetScaleDownMode()
	ResetSnapshotId()
	ResetSpotMaxPrice()
	ResetTags()
	ResetTimeouts()
	ResetUltraSsdEnabled()
	ResetUpgradeSettings()
	ResetVnetSubnetId()
	ResetWindowsProfile()
	ResetWorkloadRuntime()
	ResetZones()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for KubernetesClusterNodePool
type jsiiProxy_KubernetesClusterNodePool struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_KubernetesClusterNodePool) CapacityReservationGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"capacityReservationGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePool) CapacityReservationGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"capacityReservationGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePool) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePool) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePool) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePool) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePool) CustomCaTrustEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customCaTrustEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePool) CustomCaTrustEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customCaTrustEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePool) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePool) EnableAutoScaling() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableAutoScaling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePool) EnableAutoScalingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableAutoScalingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePool) EnableHostEncryption() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableHostEncryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePool) EnableHostEncryptionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableHostEncryptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePool) EnableNodePublicIp() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableNodePublicIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePool) EnableNodePublicIpInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableNodePublicIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePool) EvictionPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"evictionPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePool) EvictionPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"evictionPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePool) FipsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fipsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePool) FipsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fipsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePool) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePool) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePool) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePool) HostGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePool) HostGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePool) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePool) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePool) KubeletConfig() KubernetesClusterNodePoolKubeletConfigOutputReference {
	var returns KubernetesClusterNodePoolKubeletConfigOutputReference
	_jsii_.Get(
		j,
		"kubeletConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePool) KubeletConfigInput() *KubernetesClusterNodePoolKubeletConfig {
	var returns *KubernetesClusterNodePoolKubeletConfig
	_jsii_.Get(
		j,
		"kubeletConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePool) KubeletDiskType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kubeletDiskType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePool) KubeletDiskTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kubeletDiskTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePool) KubernetesClusterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kubernetesClusterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePool) KubernetesClusterIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kubernetesClusterIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePool) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePool) LinuxOsConfig() KubernetesClusterNodePoolLinuxOsConfigOutputReference {
	var returns KubernetesClusterNodePoolLinuxOsConfigOutputReference
	_jsii_.Get(
		j,
		"linuxOsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePool) LinuxOsConfigInput() *KubernetesClusterNodePoolLinuxOsConfig {
	var returns *KubernetesClusterNodePoolLinuxOsConfig
	_jsii_.Get(
		j,
		"linuxOsConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePool) MaxCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePool) MaxCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePool) MaxPods() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxPods",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePool) MaxPodsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxPodsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePool) MessageOfTheDay() *string {
	var returns *string
	_jsii_.Get(
		j,
		"messageOfTheDay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePool) MessageOfTheDayInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"messageOfTheDayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePool) MinCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePool) MinCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePool) Mode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePool) ModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePool) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePool) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePool) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePool) NodeCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nodeCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePool) NodeCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nodeCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePool) NodeLabels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"nodeLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePool) NodeLabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"nodeLabelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePool) NodeNetworkProfile() KubernetesClusterNodePoolNodeNetworkProfileOutputReference {
	var returns KubernetesClusterNodePoolNodeNetworkProfileOutputReference
	_jsii_.Get(
		j,
		"nodeNetworkProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePool) NodeNetworkProfileInput() *KubernetesClusterNodePoolNodeNetworkProfile {
	var returns *KubernetesClusterNodePoolNodeNetworkProfile
	_jsii_.Get(
		j,
		"nodeNetworkProfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePool) NodePublicIpPrefixId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodePublicIpPrefixId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePool) NodePublicIpPrefixIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodePublicIpPrefixIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePool) NodeTaints() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"nodeTaints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePool) NodeTaintsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"nodeTaintsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePool) OrchestratorVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orchestratorVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePool) OrchestratorVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orchestratorVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePool) OsDiskSizeGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"osDiskSizeGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePool) OsDiskSizeGbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"osDiskSizeGbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePool) OsDiskType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"osDiskType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePool) OsDiskTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"osDiskTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePool) OsSku() *string {
	var returns *string
	_jsii_.Get(
		j,
		"osSku",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePool) OsSkuInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"osSkuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePool) OsType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"osType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePool) OsTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"osTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePool) PodSubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"podSubnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePool) PodSubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"podSubnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePool) Priority() *string {
	var returns *string
	_jsii_.Get(
		j,
		"priority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePool) PriorityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"priorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePool) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePool) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePool) ProximityPlacementGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"proximityPlacementGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePool) ProximityPlacementGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"proximityPlacementGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePool) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePool) ScaleDownMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scaleDownMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePool) ScaleDownModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scaleDownModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePool) SnapshotId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePool) SnapshotIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePool) SpotMaxPrice() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"spotMaxPrice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePool) SpotMaxPriceInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"spotMaxPriceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePool) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePool) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePool) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePool) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePool) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePool) Timeouts() KubernetesClusterNodePoolTimeoutsOutputReference {
	var returns KubernetesClusterNodePoolTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePool) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePool) UltraSsdEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ultraSsdEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePool) UltraSsdEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ultraSsdEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePool) UpgradeSettings() KubernetesClusterNodePoolUpgradeSettingsOutputReference {
	var returns KubernetesClusterNodePoolUpgradeSettingsOutputReference
	_jsii_.Get(
		j,
		"upgradeSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePool) UpgradeSettingsInput() *KubernetesClusterNodePoolUpgradeSettings {
	var returns *KubernetesClusterNodePoolUpgradeSettings
	_jsii_.Get(
		j,
		"upgradeSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePool) VmSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vmSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePool) VmSizeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vmSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePool) VnetSubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vnetSubnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePool) VnetSubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vnetSubnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePool) WindowsProfile() KubernetesClusterNodePoolWindowsProfileOutputReference {
	var returns KubernetesClusterNodePoolWindowsProfileOutputReference
	_jsii_.Get(
		j,
		"windowsProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePool) WindowsProfileInput() *KubernetesClusterNodePoolWindowsProfile {
	var returns *KubernetesClusterNodePoolWindowsProfile
	_jsii_.Get(
		j,
		"windowsProfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePool) WorkloadRuntime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workloadRuntime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePool) WorkloadRuntimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workloadRuntimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePool) Zones() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"zones",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePool) ZonesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"zonesInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/kubernetes_cluster_node_pool azurerm_kubernetes_cluster_node_pool} Resource.
func NewKubernetesClusterNodePool(scope constructs.Construct, id *string, config *KubernetesClusterNodePoolConfig) KubernetesClusterNodePool {
	_init_.Initialize()

	if err := validateNewKubernetesClusterNodePoolParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_KubernetesClusterNodePool{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesClusterNodePool.KubernetesClusterNodePool",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/kubernetes_cluster_node_pool azurerm_kubernetes_cluster_node_pool} Resource.
func NewKubernetesClusterNodePool_Override(k KubernetesClusterNodePool, scope constructs.Construct, id *string, config *KubernetesClusterNodePoolConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesClusterNodePool.KubernetesClusterNodePool",
		[]interface{}{scope, id, config},
		k,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePool)SetCapacityReservationGroupId(val *string) {
	if err := j.validateSetCapacityReservationGroupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"capacityReservationGroupId",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePool)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePool)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePool)SetCustomCaTrustEnabled(val interface{}) {
	if err := j.validateSetCustomCaTrustEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customCaTrustEnabled",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePool)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePool)SetEnableAutoScaling(val interface{}) {
	if err := j.validateSetEnableAutoScalingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableAutoScaling",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePool)SetEnableHostEncryption(val interface{}) {
	if err := j.validateSetEnableHostEncryptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableHostEncryption",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePool)SetEnableNodePublicIp(val interface{}) {
	if err := j.validateSetEnableNodePublicIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableNodePublicIp",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePool)SetEvictionPolicy(val *string) {
	if err := j.validateSetEvictionPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"evictionPolicy",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePool)SetFipsEnabled(val interface{}) {
	if err := j.validateSetFipsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fipsEnabled",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePool)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePool)SetHostGroupId(val *string) {
	if err := j.validateSetHostGroupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostGroupId",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePool)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePool)SetKubeletDiskType(val *string) {
	if err := j.validateSetKubeletDiskTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kubeletDiskType",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePool)SetKubernetesClusterId(val *string) {
	if err := j.validateSetKubernetesClusterIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kubernetesClusterId",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePool)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePool)SetMaxCount(val *float64) {
	if err := j.validateSetMaxCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxCount",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePool)SetMaxPods(val *float64) {
	if err := j.validateSetMaxPodsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxPods",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePool)SetMessageOfTheDay(val *string) {
	if err := j.validateSetMessageOfTheDayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"messageOfTheDay",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePool)SetMinCount(val *float64) {
	if err := j.validateSetMinCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minCount",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePool)SetMode(val *string) {
	if err := j.validateSetModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mode",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePool)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePool)SetNodeCount(val *float64) {
	if err := j.validateSetNodeCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeCount",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePool)SetNodeLabels(val *map[string]*string) {
	if err := j.validateSetNodeLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeLabels",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePool)SetNodePublicIpPrefixId(val *string) {
	if err := j.validateSetNodePublicIpPrefixIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodePublicIpPrefixId",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePool)SetNodeTaints(val *[]*string) {
	if err := j.validateSetNodeTaintsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeTaints",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePool)SetOrchestratorVersion(val *string) {
	if err := j.validateSetOrchestratorVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"orchestratorVersion",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePool)SetOsDiskSizeGb(val *float64) {
	if err := j.validateSetOsDiskSizeGbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"osDiskSizeGb",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePool)SetOsDiskType(val *string) {
	if err := j.validateSetOsDiskTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"osDiskType",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePool)SetOsSku(val *string) {
	if err := j.validateSetOsSkuParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"osSku",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePool)SetOsType(val *string) {
	if err := j.validateSetOsTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"osType",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePool)SetPodSubnetId(val *string) {
	if err := j.validateSetPodSubnetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"podSubnetId",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePool)SetPriority(val *string) {
	if err := j.validateSetPriorityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"priority",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePool)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePool)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePool)SetProximityPlacementGroupId(val *string) {
	if err := j.validateSetProximityPlacementGroupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"proximityPlacementGroupId",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePool)SetScaleDownMode(val *string) {
	if err := j.validateSetScaleDownModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scaleDownMode",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePool)SetSnapshotId(val *string) {
	if err := j.validateSetSnapshotIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"snapshotId",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePool)SetSpotMaxPrice(val *float64) {
	if err := j.validateSetSpotMaxPriceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"spotMaxPrice",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePool)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePool)SetUltraSsdEnabled(val interface{}) {
	if err := j.validateSetUltraSsdEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ultraSsdEnabled",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePool)SetVmSize(val *string) {
	if err := j.validateSetVmSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vmSize",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePool)SetVnetSubnetId(val *string) {
	if err := j.validateSetVnetSubnetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vnetSubnetId",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePool)SetWorkloadRuntime(val *string) {
	if err := j.validateSetWorkloadRuntimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workloadRuntime",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePool)SetZones(val *[]*string) {
	if err := j.validateSetZonesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zones",
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
func KubernetesClusterNodePool_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateKubernetesClusterNodePool_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.kubernetesClusterNodePool.KubernetesClusterNodePool",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func KubernetesClusterNodePool_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateKubernetesClusterNodePool_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.kubernetesClusterNodePool.KubernetesClusterNodePool",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func KubernetesClusterNodePool_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateKubernetesClusterNodePool_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.kubernetesClusterNodePool.KubernetesClusterNodePool",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func KubernetesClusterNodePool_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.kubernetesClusterNodePool.KubernetesClusterNodePool",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (k *jsiiProxy_KubernetesClusterNodePool) AddOverride(path *string, value interface{}) {
	if err := k.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (k *jsiiProxy_KubernetesClusterNodePool) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := k.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNodePool) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := k.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNodePool) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := k.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		k,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNodePool) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := k.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNodePool) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := k.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		k,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNodePool) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := k.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		k,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNodePool) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := k.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		k,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNodePool) GetStringAttribute(terraformAttribute *string) *string {
	if err := k.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		k,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNodePool) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := k.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		k,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNodePool) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := k.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNodePool) OverrideLogicalId(newLogicalId *string) {
	if err := k.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (k *jsiiProxy_KubernetesClusterNodePool) PutKubeletConfig(value *KubernetesClusterNodePoolKubeletConfig) {
	if err := k.validatePutKubeletConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putKubeletConfig",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KubernetesClusterNodePool) PutLinuxOsConfig(value *KubernetesClusterNodePoolLinuxOsConfig) {
	if err := k.validatePutLinuxOsConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putLinuxOsConfig",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KubernetesClusterNodePool) PutNodeNetworkProfile(value *KubernetesClusterNodePoolNodeNetworkProfile) {
	if err := k.validatePutNodeNetworkProfileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putNodeNetworkProfile",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KubernetesClusterNodePool) PutTimeouts(value *KubernetesClusterNodePoolTimeouts) {
	if err := k.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KubernetesClusterNodePool) PutUpgradeSettings(value *KubernetesClusterNodePoolUpgradeSettings) {
	if err := k.validatePutUpgradeSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putUpgradeSettings",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KubernetesClusterNodePool) PutWindowsProfile(value *KubernetesClusterNodePoolWindowsProfile) {
	if err := k.validatePutWindowsProfileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putWindowsProfile",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KubernetesClusterNodePool) ResetCapacityReservationGroupId() {
	_jsii_.InvokeVoid(
		k,
		"resetCapacityReservationGroupId",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNodePool) ResetCustomCaTrustEnabled() {
	_jsii_.InvokeVoid(
		k,
		"resetCustomCaTrustEnabled",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNodePool) ResetEnableAutoScaling() {
	_jsii_.InvokeVoid(
		k,
		"resetEnableAutoScaling",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNodePool) ResetEnableHostEncryption() {
	_jsii_.InvokeVoid(
		k,
		"resetEnableHostEncryption",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNodePool) ResetEnableNodePublicIp() {
	_jsii_.InvokeVoid(
		k,
		"resetEnableNodePublicIp",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNodePool) ResetEvictionPolicy() {
	_jsii_.InvokeVoid(
		k,
		"resetEvictionPolicy",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNodePool) ResetFipsEnabled() {
	_jsii_.InvokeVoid(
		k,
		"resetFipsEnabled",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNodePool) ResetHostGroupId() {
	_jsii_.InvokeVoid(
		k,
		"resetHostGroupId",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNodePool) ResetId() {
	_jsii_.InvokeVoid(
		k,
		"resetId",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNodePool) ResetKubeletConfig() {
	_jsii_.InvokeVoid(
		k,
		"resetKubeletConfig",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNodePool) ResetKubeletDiskType() {
	_jsii_.InvokeVoid(
		k,
		"resetKubeletDiskType",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNodePool) ResetLinuxOsConfig() {
	_jsii_.InvokeVoid(
		k,
		"resetLinuxOsConfig",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNodePool) ResetMaxCount() {
	_jsii_.InvokeVoid(
		k,
		"resetMaxCount",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNodePool) ResetMaxPods() {
	_jsii_.InvokeVoid(
		k,
		"resetMaxPods",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNodePool) ResetMessageOfTheDay() {
	_jsii_.InvokeVoid(
		k,
		"resetMessageOfTheDay",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNodePool) ResetMinCount() {
	_jsii_.InvokeVoid(
		k,
		"resetMinCount",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNodePool) ResetMode() {
	_jsii_.InvokeVoid(
		k,
		"resetMode",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNodePool) ResetNodeCount() {
	_jsii_.InvokeVoid(
		k,
		"resetNodeCount",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNodePool) ResetNodeLabels() {
	_jsii_.InvokeVoid(
		k,
		"resetNodeLabels",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNodePool) ResetNodeNetworkProfile() {
	_jsii_.InvokeVoid(
		k,
		"resetNodeNetworkProfile",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNodePool) ResetNodePublicIpPrefixId() {
	_jsii_.InvokeVoid(
		k,
		"resetNodePublicIpPrefixId",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNodePool) ResetNodeTaints() {
	_jsii_.InvokeVoid(
		k,
		"resetNodeTaints",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNodePool) ResetOrchestratorVersion() {
	_jsii_.InvokeVoid(
		k,
		"resetOrchestratorVersion",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNodePool) ResetOsDiskSizeGb() {
	_jsii_.InvokeVoid(
		k,
		"resetOsDiskSizeGb",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNodePool) ResetOsDiskType() {
	_jsii_.InvokeVoid(
		k,
		"resetOsDiskType",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNodePool) ResetOsSku() {
	_jsii_.InvokeVoid(
		k,
		"resetOsSku",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNodePool) ResetOsType() {
	_jsii_.InvokeVoid(
		k,
		"resetOsType",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNodePool) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		k,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNodePool) ResetPodSubnetId() {
	_jsii_.InvokeVoid(
		k,
		"resetPodSubnetId",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNodePool) ResetPriority() {
	_jsii_.InvokeVoid(
		k,
		"resetPriority",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNodePool) ResetProximityPlacementGroupId() {
	_jsii_.InvokeVoid(
		k,
		"resetProximityPlacementGroupId",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNodePool) ResetScaleDownMode() {
	_jsii_.InvokeVoid(
		k,
		"resetScaleDownMode",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNodePool) ResetSnapshotId() {
	_jsii_.InvokeVoid(
		k,
		"resetSnapshotId",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNodePool) ResetSpotMaxPrice() {
	_jsii_.InvokeVoid(
		k,
		"resetSpotMaxPrice",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNodePool) ResetTags() {
	_jsii_.InvokeVoid(
		k,
		"resetTags",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNodePool) ResetTimeouts() {
	_jsii_.InvokeVoid(
		k,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNodePool) ResetUltraSsdEnabled() {
	_jsii_.InvokeVoid(
		k,
		"resetUltraSsdEnabled",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNodePool) ResetUpgradeSettings() {
	_jsii_.InvokeVoid(
		k,
		"resetUpgradeSettings",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNodePool) ResetVnetSubnetId() {
	_jsii_.InvokeVoid(
		k,
		"resetVnetSubnetId",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNodePool) ResetWindowsProfile() {
	_jsii_.InvokeVoid(
		k,
		"resetWindowsProfile",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNodePool) ResetWorkloadRuntime() {
	_jsii_.InvokeVoid(
		k,
		"resetWorkloadRuntime",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNodePool) ResetZones() {
	_jsii_.InvokeVoid(
		k,
		"resetZones",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNodePool) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNodePool) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNodePool) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNodePool) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

