package kubernetesclusternodepool

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/kubernetesclusternodepool/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster_node_pool azurerm_kubernetes_cluster_node_pool}.
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
	Count() *float64
	// Experimental.
	SetCount(val *float64)
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
	PutTimeouts(value *KubernetesClusterNodePoolTimeouts)
	PutUpgradeSettings(value *KubernetesClusterNodePoolUpgradeSettings)
	ResetCapacityReservationGroupId()
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
	ResetMinCount()
	ResetMode()
	ResetNodeCount()
	ResetNodeLabels()
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
	ResetSpotMaxPrice()
	ResetTags()
	ResetTimeouts()
	ResetUltraSsdEnabled()
	ResetUpgradeSettings()
	ResetVnetSubnetId()
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

func (j *jsiiProxy_KubernetesClusterNodePool) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
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


// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster_node_pool azurerm_kubernetes_cluster_node_pool} Resource.
func NewKubernetesClusterNodePool(scope constructs.Construct, id *string, config *KubernetesClusterNodePoolConfig) KubernetesClusterNodePool {
	_init_.Initialize()

	j := jsiiProxy_KubernetesClusterNodePool{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesClusterNodePool.KubernetesClusterNodePool",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster_node_pool azurerm_kubernetes_cluster_node_pool} Resource.
func NewKubernetesClusterNodePool_Override(k KubernetesClusterNodePool, scope constructs.Construct, id *string, config *KubernetesClusterNodePoolConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesClusterNodePool.KubernetesClusterNodePool",
		[]interface{}{scope, id, config},
		k,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePool) SetCapacityReservationGroupId(val *string) {
	_jsii_.Set(
		j,
		"capacityReservationGroupId",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePool) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePool) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePool) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePool) SetEnableAutoScaling(val interface{}) {
	_jsii_.Set(
		j,
		"enableAutoScaling",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePool) SetEnableHostEncryption(val interface{}) {
	_jsii_.Set(
		j,
		"enableHostEncryption",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePool) SetEnableNodePublicIp(val interface{}) {
	_jsii_.Set(
		j,
		"enableNodePublicIp",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePool) SetEvictionPolicy(val *string) {
	_jsii_.Set(
		j,
		"evictionPolicy",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePool) SetFipsEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"fipsEnabled",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePool) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePool) SetHostGroupId(val *string) {
	_jsii_.Set(
		j,
		"hostGroupId",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePool) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePool) SetKubeletDiskType(val *string) {
	_jsii_.Set(
		j,
		"kubeletDiskType",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePool) SetKubernetesClusterId(val *string) {
	_jsii_.Set(
		j,
		"kubernetesClusterId",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePool) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePool) SetMaxCount(val *float64) {
	_jsii_.Set(
		j,
		"maxCount",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePool) SetMaxPods(val *float64) {
	_jsii_.Set(
		j,
		"maxPods",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePool) SetMinCount(val *float64) {
	_jsii_.Set(
		j,
		"minCount",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePool) SetMode(val *string) {
	_jsii_.Set(
		j,
		"mode",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePool) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePool) SetNodeCount(val *float64) {
	_jsii_.Set(
		j,
		"nodeCount",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePool) SetNodeLabels(val *map[string]*string) {
	_jsii_.Set(
		j,
		"nodeLabels",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePool) SetNodePublicIpPrefixId(val *string) {
	_jsii_.Set(
		j,
		"nodePublicIpPrefixId",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePool) SetNodeTaints(val *[]*string) {
	_jsii_.Set(
		j,
		"nodeTaints",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePool) SetOrchestratorVersion(val *string) {
	_jsii_.Set(
		j,
		"orchestratorVersion",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePool) SetOsDiskSizeGb(val *float64) {
	_jsii_.Set(
		j,
		"osDiskSizeGb",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePool) SetOsDiskType(val *string) {
	_jsii_.Set(
		j,
		"osDiskType",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePool) SetOsSku(val *string) {
	_jsii_.Set(
		j,
		"osSku",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePool) SetOsType(val *string) {
	_jsii_.Set(
		j,
		"osType",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePool) SetPodSubnetId(val *string) {
	_jsii_.Set(
		j,
		"podSubnetId",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePool) SetPriority(val *string) {
	_jsii_.Set(
		j,
		"priority",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePool) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePool) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePool) SetProximityPlacementGroupId(val *string) {
	_jsii_.Set(
		j,
		"proximityPlacementGroupId",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePool) SetScaleDownMode(val *string) {
	_jsii_.Set(
		j,
		"scaleDownMode",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePool) SetSpotMaxPrice(val *float64) {
	_jsii_.Set(
		j,
		"spotMaxPrice",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePool) SetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePool) SetUltraSsdEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"ultraSsdEnabled",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePool) SetVmSize(val *string) {
	_jsii_.Set(
		j,
		"vmSize",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePool) SetVnetSubnetId(val *string) {
	_jsii_.Set(
		j,
		"vnetSubnetId",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePool) SetWorkloadRuntime(val *string) {
	_jsii_.Set(
		j,
		"workloadRuntime",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePool) SetZones(val *[]*string) {
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

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.kubernetesClusterNodePool.KubernetesClusterNodePool",
		"isConstruct",
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
	_jsii_.InvokeVoid(
		k,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (k *jsiiProxy_KubernetesClusterNodePool) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
	_jsii_.InvokeVoid(
		k,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (k *jsiiProxy_KubernetesClusterNodePool) PutKubeletConfig(value *KubernetesClusterNodePoolKubeletConfig) {
	_jsii_.InvokeVoid(
		k,
		"putKubeletConfig",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KubernetesClusterNodePool) PutLinuxOsConfig(value *KubernetesClusterNodePoolLinuxOsConfig) {
	_jsii_.InvokeVoid(
		k,
		"putLinuxOsConfig",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KubernetesClusterNodePool) PutTimeouts(value *KubernetesClusterNodePoolTimeouts) {
	_jsii_.InvokeVoid(
		k,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KubernetesClusterNodePool) PutUpgradeSettings(value *KubernetesClusterNodePoolUpgradeSettings) {
	_jsii_.InvokeVoid(
		k,
		"putUpgradeSettings",
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

type KubernetesClusterNodePoolConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster_node_pool#kubernetes_cluster_id KubernetesClusterNodePool#kubernetes_cluster_id}.
	KubernetesClusterId *string `field:"required" json:"kubernetesClusterId" yaml:"kubernetesClusterId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster_node_pool#name KubernetesClusterNodePool#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster_node_pool#vm_size KubernetesClusterNodePool#vm_size}.
	VmSize *string `field:"required" json:"vmSize" yaml:"vmSize"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster_node_pool#capacity_reservation_group_id KubernetesClusterNodePool#capacity_reservation_group_id}.
	CapacityReservationGroupId *string `field:"optional" json:"capacityReservationGroupId" yaml:"capacityReservationGroupId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster_node_pool#enable_auto_scaling KubernetesClusterNodePool#enable_auto_scaling}.
	EnableAutoScaling interface{} `field:"optional" json:"enableAutoScaling" yaml:"enableAutoScaling"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster_node_pool#enable_host_encryption KubernetesClusterNodePool#enable_host_encryption}.
	EnableHostEncryption interface{} `field:"optional" json:"enableHostEncryption" yaml:"enableHostEncryption"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster_node_pool#enable_node_public_ip KubernetesClusterNodePool#enable_node_public_ip}.
	EnableNodePublicIp interface{} `field:"optional" json:"enableNodePublicIp" yaml:"enableNodePublicIp"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster_node_pool#eviction_policy KubernetesClusterNodePool#eviction_policy}.
	EvictionPolicy *string `field:"optional" json:"evictionPolicy" yaml:"evictionPolicy"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster_node_pool#fips_enabled KubernetesClusterNodePool#fips_enabled}.
	FipsEnabled interface{} `field:"optional" json:"fipsEnabled" yaml:"fipsEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster_node_pool#host_group_id KubernetesClusterNodePool#host_group_id}.
	HostGroupId *string `field:"optional" json:"hostGroupId" yaml:"hostGroupId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster_node_pool#id KubernetesClusterNodePool#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// kubelet_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster_node_pool#kubelet_config KubernetesClusterNodePool#kubelet_config}
	KubeletConfig *KubernetesClusterNodePoolKubeletConfig `field:"optional" json:"kubeletConfig" yaml:"kubeletConfig"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster_node_pool#kubelet_disk_type KubernetesClusterNodePool#kubelet_disk_type}.
	KubeletDiskType *string `field:"optional" json:"kubeletDiskType" yaml:"kubeletDiskType"`
	// linux_os_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster_node_pool#linux_os_config KubernetesClusterNodePool#linux_os_config}
	LinuxOsConfig *KubernetesClusterNodePoolLinuxOsConfig `field:"optional" json:"linuxOsConfig" yaml:"linuxOsConfig"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster_node_pool#max_count KubernetesClusterNodePool#max_count}.
	MaxCount *float64 `field:"optional" json:"maxCount" yaml:"maxCount"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster_node_pool#max_pods KubernetesClusterNodePool#max_pods}.
	MaxPods *float64 `field:"optional" json:"maxPods" yaml:"maxPods"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster_node_pool#min_count KubernetesClusterNodePool#min_count}.
	MinCount *float64 `field:"optional" json:"minCount" yaml:"minCount"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster_node_pool#mode KubernetesClusterNodePool#mode}.
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster_node_pool#node_count KubernetesClusterNodePool#node_count}.
	NodeCount *float64 `field:"optional" json:"nodeCount" yaml:"nodeCount"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster_node_pool#node_labels KubernetesClusterNodePool#node_labels}.
	NodeLabels *map[string]*string `field:"optional" json:"nodeLabels" yaml:"nodeLabels"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster_node_pool#node_public_ip_prefix_id KubernetesClusterNodePool#node_public_ip_prefix_id}.
	NodePublicIpPrefixId *string `field:"optional" json:"nodePublicIpPrefixId" yaml:"nodePublicIpPrefixId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster_node_pool#node_taints KubernetesClusterNodePool#node_taints}.
	NodeTaints *[]*string `field:"optional" json:"nodeTaints" yaml:"nodeTaints"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster_node_pool#orchestrator_version KubernetesClusterNodePool#orchestrator_version}.
	OrchestratorVersion *string `field:"optional" json:"orchestratorVersion" yaml:"orchestratorVersion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster_node_pool#os_disk_size_gb KubernetesClusterNodePool#os_disk_size_gb}.
	OsDiskSizeGb *float64 `field:"optional" json:"osDiskSizeGb" yaml:"osDiskSizeGb"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster_node_pool#os_disk_type KubernetesClusterNodePool#os_disk_type}.
	OsDiskType *string `field:"optional" json:"osDiskType" yaml:"osDiskType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster_node_pool#os_sku KubernetesClusterNodePool#os_sku}.
	OsSku *string `field:"optional" json:"osSku" yaml:"osSku"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster_node_pool#os_type KubernetesClusterNodePool#os_type}.
	OsType *string `field:"optional" json:"osType" yaml:"osType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster_node_pool#pod_subnet_id KubernetesClusterNodePool#pod_subnet_id}.
	PodSubnetId *string `field:"optional" json:"podSubnetId" yaml:"podSubnetId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster_node_pool#priority KubernetesClusterNodePool#priority}.
	Priority *string `field:"optional" json:"priority" yaml:"priority"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster_node_pool#proximity_placement_group_id KubernetesClusterNodePool#proximity_placement_group_id}.
	ProximityPlacementGroupId *string `field:"optional" json:"proximityPlacementGroupId" yaml:"proximityPlacementGroupId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster_node_pool#scale_down_mode KubernetesClusterNodePool#scale_down_mode}.
	ScaleDownMode *string `field:"optional" json:"scaleDownMode" yaml:"scaleDownMode"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster_node_pool#spot_max_price KubernetesClusterNodePool#spot_max_price}.
	SpotMaxPrice *float64 `field:"optional" json:"spotMaxPrice" yaml:"spotMaxPrice"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster_node_pool#tags KubernetesClusterNodePool#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster_node_pool#timeouts KubernetesClusterNodePool#timeouts}
	Timeouts *KubernetesClusterNodePoolTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster_node_pool#ultra_ssd_enabled KubernetesClusterNodePool#ultra_ssd_enabled}.
	UltraSsdEnabled interface{} `field:"optional" json:"ultraSsdEnabled" yaml:"ultraSsdEnabled"`
	// upgrade_settings block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster_node_pool#upgrade_settings KubernetesClusterNodePool#upgrade_settings}
	UpgradeSettings *KubernetesClusterNodePoolUpgradeSettings `field:"optional" json:"upgradeSettings" yaml:"upgradeSettings"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster_node_pool#vnet_subnet_id KubernetesClusterNodePool#vnet_subnet_id}.
	VnetSubnetId *string `field:"optional" json:"vnetSubnetId" yaml:"vnetSubnetId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster_node_pool#workload_runtime KubernetesClusterNodePool#workload_runtime}.
	WorkloadRuntime *string `field:"optional" json:"workloadRuntime" yaml:"workloadRuntime"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster_node_pool#zones KubernetesClusterNodePool#zones}.
	Zones *[]*string `field:"optional" json:"zones" yaml:"zones"`
}

type KubernetesClusterNodePoolKubeletConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster_node_pool#allowed_unsafe_sysctls KubernetesClusterNodePool#allowed_unsafe_sysctls}.
	AllowedUnsafeSysctls *[]*string `field:"optional" json:"allowedUnsafeSysctls" yaml:"allowedUnsafeSysctls"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster_node_pool#container_log_max_line KubernetesClusterNodePool#container_log_max_line}.
	ContainerLogMaxLine *float64 `field:"optional" json:"containerLogMaxLine" yaml:"containerLogMaxLine"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster_node_pool#container_log_max_size_mb KubernetesClusterNodePool#container_log_max_size_mb}.
	ContainerLogMaxSizeMb *float64 `field:"optional" json:"containerLogMaxSizeMb" yaml:"containerLogMaxSizeMb"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster_node_pool#cpu_cfs_quota_enabled KubernetesClusterNodePool#cpu_cfs_quota_enabled}.
	CpuCfsQuotaEnabled interface{} `field:"optional" json:"cpuCfsQuotaEnabled" yaml:"cpuCfsQuotaEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster_node_pool#cpu_cfs_quota_period KubernetesClusterNodePool#cpu_cfs_quota_period}.
	CpuCfsQuotaPeriod *string `field:"optional" json:"cpuCfsQuotaPeriod" yaml:"cpuCfsQuotaPeriod"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster_node_pool#cpu_manager_policy KubernetesClusterNodePool#cpu_manager_policy}.
	CpuManagerPolicy *string `field:"optional" json:"cpuManagerPolicy" yaml:"cpuManagerPolicy"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster_node_pool#image_gc_high_threshold KubernetesClusterNodePool#image_gc_high_threshold}.
	ImageGcHighThreshold *float64 `field:"optional" json:"imageGcHighThreshold" yaml:"imageGcHighThreshold"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster_node_pool#image_gc_low_threshold KubernetesClusterNodePool#image_gc_low_threshold}.
	ImageGcLowThreshold *float64 `field:"optional" json:"imageGcLowThreshold" yaml:"imageGcLowThreshold"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster_node_pool#pod_max_pid KubernetesClusterNodePool#pod_max_pid}.
	PodMaxPid *float64 `field:"optional" json:"podMaxPid" yaml:"podMaxPid"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster_node_pool#topology_manager_policy KubernetesClusterNodePool#topology_manager_policy}.
	TopologyManagerPolicy *string `field:"optional" json:"topologyManagerPolicy" yaml:"topologyManagerPolicy"`
}

type KubernetesClusterNodePoolKubeletConfigOutputReference interface {
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
	InternalValue() *KubernetesClusterNodePoolKubeletConfig
	SetInternalValue(val *KubernetesClusterNodePoolKubeletConfig)
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

// The jsii proxy struct for KubernetesClusterNodePoolKubeletConfigOutputReference
type jsiiProxy_KubernetesClusterNodePoolKubeletConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_KubernetesClusterNodePoolKubeletConfigOutputReference) AllowedUnsafeSysctls() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedUnsafeSysctls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolKubeletConfigOutputReference) AllowedUnsafeSysctlsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedUnsafeSysctlsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolKubeletConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolKubeletConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolKubeletConfigOutputReference) ContainerLogMaxLine() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"containerLogMaxLine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolKubeletConfigOutputReference) ContainerLogMaxLineInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"containerLogMaxLineInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolKubeletConfigOutputReference) ContainerLogMaxSizeMb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"containerLogMaxSizeMb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolKubeletConfigOutputReference) ContainerLogMaxSizeMbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"containerLogMaxSizeMbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolKubeletConfigOutputReference) CpuCfsQuotaEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cpuCfsQuotaEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolKubeletConfigOutputReference) CpuCfsQuotaEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cpuCfsQuotaEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolKubeletConfigOutputReference) CpuCfsQuotaPeriod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cpuCfsQuotaPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolKubeletConfigOutputReference) CpuCfsQuotaPeriodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cpuCfsQuotaPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolKubeletConfigOutputReference) CpuManagerPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cpuManagerPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolKubeletConfigOutputReference) CpuManagerPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cpuManagerPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolKubeletConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolKubeletConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolKubeletConfigOutputReference) ImageGcHighThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"imageGcHighThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolKubeletConfigOutputReference) ImageGcHighThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"imageGcHighThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolKubeletConfigOutputReference) ImageGcLowThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"imageGcLowThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolKubeletConfigOutputReference) ImageGcLowThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"imageGcLowThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolKubeletConfigOutputReference) InternalValue() *KubernetesClusterNodePoolKubeletConfig {
	var returns *KubernetesClusterNodePoolKubeletConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolKubeletConfigOutputReference) PodMaxPid() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"podMaxPid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolKubeletConfigOutputReference) PodMaxPidInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"podMaxPidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolKubeletConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolKubeletConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolKubeletConfigOutputReference) TopologyManagerPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"topologyManagerPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolKubeletConfigOutputReference) TopologyManagerPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"topologyManagerPolicyInput",
		&returns,
	)
	return returns
}


func NewKubernetesClusterNodePoolKubeletConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) KubernetesClusterNodePoolKubeletConfigOutputReference {
	_init_.Initialize()

	j := jsiiProxy_KubernetesClusterNodePoolKubeletConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesClusterNodePool.KubernetesClusterNodePoolKubeletConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKubernetesClusterNodePoolKubeletConfigOutputReference_Override(k KubernetesClusterNodePoolKubeletConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesClusterNodePool.KubernetesClusterNodePoolKubeletConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePoolKubeletConfigOutputReference) SetAllowedUnsafeSysctls(val *[]*string) {
	_jsii_.Set(
		j,
		"allowedUnsafeSysctls",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePoolKubeletConfigOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePoolKubeletConfigOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePoolKubeletConfigOutputReference) SetContainerLogMaxLine(val *float64) {
	_jsii_.Set(
		j,
		"containerLogMaxLine",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePoolKubeletConfigOutputReference) SetContainerLogMaxSizeMb(val *float64) {
	_jsii_.Set(
		j,
		"containerLogMaxSizeMb",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePoolKubeletConfigOutputReference) SetCpuCfsQuotaEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"cpuCfsQuotaEnabled",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePoolKubeletConfigOutputReference) SetCpuCfsQuotaPeriod(val *string) {
	_jsii_.Set(
		j,
		"cpuCfsQuotaPeriod",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePoolKubeletConfigOutputReference) SetCpuManagerPolicy(val *string) {
	_jsii_.Set(
		j,
		"cpuManagerPolicy",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePoolKubeletConfigOutputReference) SetImageGcHighThreshold(val *float64) {
	_jsii_.Set(
		j,
		"imageGcHighThreshold",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePoolKubeletConfigOutputReference) SetImageGcLowThreshold(val *float64) {
	_jsii_.Set(
		j,
		"imageGcLowThreshold",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePoolKubeletConfigOutputReference) SetInternalValue(val *KubernetesClusterNodePoolKubeletConfig) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePoolKubeletConfigOutputReference) SetPodMaxPid(val *float64) {
	_jsii_.Set(
		j,
		"podMaxPid",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePoolKubeletConfigOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePoolKubeletConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePoolKubeletConfigOutputReference) SetTopologyManagerPolicy(val *string) {
	_jsii_.Set(
		j,
		"topologyManagerPolicy",
		val,
	)
}

func (k *jsiiProxy_KubernetesClusterNodePoolKubeletConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNodePoolKubeletConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNodePoolKubeletConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNodePoolKubeletConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		k,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNodePoolKubeletConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNodePoolKubeletConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		k,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNodePoolKubeletConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		k,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNodePoolKubeletConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		k,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNodePoolKubeletConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNodePoolKubeletConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		k,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNodePoolKubeletConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNodePoolKubeletConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNodePoolKubeletConfigOutputReference) ResetAllowedUnsafeSysctls() {
	_jsii_.InvokeVoid(
		k,
		"resetAllowedUnsafeSysctls",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNodePoolKubeletConfigOutputReference) ResetContainerLogMaxLine() {
	_jsii_.InvokeVoid(
		k,
		"resetContainerLogMaxLine",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNodePoolKubeletConfigOutputReference) ResetContainerLogMaxSizeMb() {
	_jsii_.InvokeVoid(
		k,
		"resetContainerLogMaxSizeMb",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNodePoolKubeletConfigOutputReference) ResetCpuCfsQuotaEnabled() {
	_jsii_.InvokeVoid(
		k,
		"resetCpuCfsQuotaEnabled",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNodePoolKubeletConfigOutputReference) ResetCpuCfsQuotaPeriod() {
	_jsii_.InvokeVoid(
		k,
		"resetCpuCfsQuotaPeriod",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNodePoolKubeletConfigOutputReference) ResetCpuManagerPolicy() {
	_jsii_.InvokeVoid(
		k,
		"resetCpuManagerPolicy",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNodePoolKubeletConfigOutputReference) ResetImageGcHighThreshold() {
	_jsii_.InvokeVoid(
		k,
		"resetImageGcHighThreshold",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNodePoolKubeletConfigOutputReference) ResetImageGcLowThreshold() {
	_jsii_.InvokeVoid(
		k,
		"resetImageGcLowThreshold",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNodePoolKubeletConfigOutputReference) ResetPodMaxPid() {
	_jsii_.InvokeVoid(
		k,
		"resetPodMaxPid",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNodePoolKubeletConfigOutputReference) ResetTopologyManagerPolicy() {
	_jsii_.InvokeVoid(
		k,
		"resetTopologyManagerPolicy",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNodePoolKubeletConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNodePoolKubeletConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type KubernetesClusterNodePoolLinuxOsConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster_node_pool#swap_file_size_mb KubernetesClusterNodePool#swap_file_size_mb}.
	SwapFileSizeMb *float64 `field:"optional" json:"swapFileSizeMb" yaml:"swapFileSizeMb"`
	// sysctl_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster_node_pool#sysctl_config KubernetesClusterNodePool#sysctl_config}
	SysctlConfig *KubernetesClusterNodePoolLinuxOsConfigSysctlConfig `field:"optional" json:"sysctlConfig" yaml:"sysctlConfig"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster_node_pool#transparent_huge_page_defrag KubernetesClusterNodePool#transparent_huge_page_defrag}.
	TransparentHugePageDefrag *string `field:"optional" json:"transparentHugePageDefrag" yaml:"transparentHugePageDefrag"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster_node_pool#transparent_huge_page_enabled KubernetesClusterNodePool#transparent_huge_page_enabled}.
	TransparentHugePageEnabled *string `field:"optional" json:"transparentHugePageEnabled" yaml:"transparentHugePageEnabled"`
}

type KubernetesClusterNodePoolLinuxOsConfigOutputReference interface {
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
	InternalValue() *KubernetesClusterNodePoolLinuxOsConfig
	SetInternalValue(val *KubernetesClusterNodePoolLinuxOsConfig)
	SwapFileSizeMb() *float64
	SetSwapFileSizeMb(val *float64)
	SwapFileSizeMbInput() *float64
	SysctlConfig() KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference
	SysctlConfigInput() *KubernetesClusterNodePoolLinuxOsConfigSysctlConfig
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
	PutSysctlConfig(value *KubernetesClusterNodePoolLinuxOsConfigSysctlConfig)
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

// The jsii proxy struct for KubernetesClusterNodePoolLinuxOsConfigOutputReference
type jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigOutputReference) InternalValue() *KubernetesClusterNodePoolLinuxOsConfig {
	var returns *KubernetesClusterNodePoolLinuxOsConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigOutputReference) SwapFileSizeMb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"swapFileSizeMb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigOutputReference) SwapFileSizeMbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"swapFileSizeMbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigOutputReference) SysctlConfig() KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference {
	var returns KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference
	_jsii_.Get(
		j,
		"sysctlConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigOutputReference) SysctlConfigInput() *KubernetesClusterNodePoolLinuxOsConfigSysctlConfig {
	var returns *KubernetesClusterNodePoolLinuxOsConfigSysctlConfig
	_jsii_.Get(
		j,
		"sysctlConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigOutputReference) TransparentHugePageDefrag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transparentHugePageDefrag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigOutputReference) TransparentHugePageDefragInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transparentHugePageDefragInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigOutputReference) TransparentHugePageEnabled() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transparentHugePageEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigOutputReference) TransparentHugePageEnabledInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transparentHugePageEnabledInput",
		&returns,
	)
	return returns
}


func NewKubernetesClusterNodePoolLinuxOsConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) KubernetesClusterNodePoolLinuxOsConfigOutputReference {
	_init_.Initialize()

	j := jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesClusterNodePool.KubernetesClusterNodePoolLinuxOsConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKubernetesClusterNodePoolLinuxOsConfigOutputReference_Override(k KubernetesClusterNodePoolLinuxOsConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesClusterNodePool.KubernetesClusterNodePoolLinuxOsConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigOutputReference) SetInternalValue(val *KubernetesClusterNodePoolLinuxOsConfig) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigOutputReference) SetSwapFileSizeMb(val *float64) {
	_jsii_.Set(
		j,
		"swapFileSizeMb",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigOutputReference) SetTransparentHugePageDefrag(val *string) {
	_jsii_.Set(
		j,
		"transparentHugePageDefrag",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigOutputReference) SetTransparentHugePageEnabled(val *string) {
	_jsii_.Set(
		j,
		"transparentHugePageEnabled",
		val,
	)
}

func (k *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		k,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		k,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		k,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		k,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		k,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigOutputReference) PutSysctlConfig(value *KubernetesClusterNodePoolLinuxOsConfigSysctlConfig) {
	_jsii_.InvokeVoid(
		k,
		"putSysctlConfig",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigOutputReference) ResetSwapFileSizeMb() {
	_jsii_.InvokeVoid(
		k,
		"resetSwapFileSizeMb",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigOutputReference) ResetSysctlConfig() {
	_jsii_.InvokeVoid(
		k,
		"resetSysctlConfig",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigOutputReference) ResetTransparentHugePageDefrag() {
	_jsii_.InvokeVoid(
		k,
		"resetTransparentHugePageDefrag",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigOutputReference) ResetTransparentHugePageEnabled() {
	_jsii_.InvokeVoid(
		k,
		"resetTransparentHugePageEnabled",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type KubernetesClusterNodePoolLinuxOsConfigSysctlConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster_node_pool#fs_aio_max_nr KubernetesClusterNodePool#fs_aio_max_nr}.
	FsAioMaxNr *float64 `field:"optional" json:"fsAioMaxNr" yaml:"fsAioMaxNr"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster_node_pool#fs_file_max KubernetesClusterNodePool#fs_file_max}.
	FsFileMax *float64 `field:"optional" json:"fsFileMax" yaml:"fsFileMax"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster_node_pool#fs_inotify_max_user_watches KubernetesClusterNodePool#fs_inotify_max_user_watches}.
	FsInotifyMaxUserWatches *float64 `field:"optional" json:"fsInotifyMaxUserWatches" yaml:"fsInotifyMaxUserWatches"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster_node_pool#fs_nr_open KubernetesClusterNodePool#fs_nr_open}.
	FsNrOpen *float64 `field:"optional" json:"fsNrOpen" yaml:"fsNrOpen"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster_node_pool#kernel_threads_max KubernetesClusterNodePool#kernel_threads_max}.
	KernelThreadsMax *float64 `field:"optional" json:"kernelThreadsMax" yaml:"kernelThreadsMax"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster_node_pool#net_core_netdev_max_backlog KubernetesClusterNodePool#net_core_netdev_max_backlog}.
	NetCoreNetdevMaxBacklog *float64 `field:"optional" json:"netCoreNetdevMaxBacklog" yaml:"netCoreNetdevMaxBacklog"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster_node_pool#net_core_optmem_max KubernetesClusterNodePool#net_core_optmem_max}.
	NetCoreOptmemMax *float64 `field:"optional" json:"netCoreOptmemMax" yaml:"netCoreOptmemMax"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster_node_pool#net_core_rmem_default KubernetesClusterNodePool#net_core_rmem_default}.
	NetCoreRmemDefault *float64 `field:"optional" json:"netCoreRmemDefault" yaml:"netCoreRmemDefault"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster_node_pool#net_core_rmem_max KubernetesClusterNodePool#net_core_rmem_max}.
	NetCoreRmemMax *float64 `field:"optional" json:"netCoreRmemMax" yaml:"netCoreRmemMax"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster_node_pool#net_core_somaxconn KubernetesClusterNodePool#net_core_somaxconn}.
	NetCoreSomaxconn *float64 `field:"optional" json:"netCoreSomaxconn" yaml:"netCoreSomaxconn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster_node_pool#net_core_wmem_default KubernetesClusterNodePool#net_core_wmem_default}.
	NetCoreWmemDefault *float64 `field:"optional" json:"netCoreWmemDefault" yaml:"netCoreWmemDefault"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster_node_pool#net_core_wmem_max KubernetesClusterNodePool#net_core_wmem_max}.
	NetCoreWmemMax *float64 `field:"optional" json:"netCoreWmemMax" yaml:"netCoreWmemMax"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster_node_pool#net_ipv4_ip_local_port_range_max KubernetesClusterNodePool#net_ipv4_ip_local_port_range_max}.
	NetIpv4IpLocalPortRangeMax *float64 `field:"optional" json:"netIpv4IpLocalPortRangeMax" yaml:"netIpv4IpLocalPortRangeMax"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster_node_pool#net_ipv4_ip_local_port_range_min KubernetesClusterNodePool#net_ipv4_ip_local_port_range_min}.
	NetIpv4IpLocalPortRangeMin *float64 `field:"optional" json:"netIpv4IpLocalPortRangeMin" yaml:"netIpv4IpLocalPortRangeMin"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster_node_pool#net_ipv4_neigh_default_gc_thresh1 KubernetesClusterNodePool#net_ipv4_neigh_default_gc_thresh1}.
	NetIpv4NeighDefaultGcThresh1 *float64 `field:"optional" json:"netIpv4NeighDefaultGcThresh1" yaml:"netIpv4NeighDefaultGcThresh1"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster_node_pool#net_ipv4_neigh_default_gc_thresh2 KubernetesClusterNodePool#net_ipv4_neigh_default_gc_thresh2}.
	NetIpv4NeighDefaultGcThresh2 *float64 `field:"optional" json:"netIpv4NeighDefaultGcThresh2" yaml:"netIpv4NeighDefaultGcThresh2"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster_node_pool#net_ipv4_neigh_default_gc_thresh3 KubernetesClusterNodePool#net_ipv4_neigh_default_gc_thresh3}.
	NetIpv4NeighDefaultGcThresh3 *float64 `field:"optional" json:"netIpv4NeighDefaultGcThresh3" yaml:"netIpv4NeighDefaultGcThresh3"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster_node_pool#net_ipv4_tcp_fin_timeout KubernetesClusterNodePool#net_ipv4_tcp_fin_timeout}.
	NetIpv4TcpFinTimeout *float64 `field:"optional" json:"netIpv4TcpFinTimeout" yaml:"netIpv4TcpFinTimeout"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster_node_pool#net_ipv4_tcp_keepalive_intvl KubernetesClusterNodePool#net_ipv4_tcp_keepalive_intvl}.
	NetIpv4TcpKeepaliveIntvl *float64 `field:"optional" json:"netIpv4TcpKeepaliveIntvl" yaml:"netIpv4TcpKeepaliveIntvl"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster_node_pool#net_ipv4_tcp_keepalive_probes KubernetesClusterNodePool#net_ipv4_tcp_keepalive_probes}.
	NetIpv4TcpKeepaliveProbes *float64 `field:"optional" json:"netIpv4TcpKeepaliveProbes" yaml:"netIpv4TcpKeepaliveProbes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster_node_pool#net_ipv4_tcp_keepalive_time KubernetesClusterNodePool#net_ipv4_tcp_keepalive_time}.
	NetIpv4TcpKeepaliveTime *float64 `field:"optional" json:"netIpv4TcpKeepaliveTime" yaml:"netIpv4TcpKeepaliveTime"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster_node_pool#net_ipv4_tcp_max_syn_backlog KubernetesClusterNodePool#net_ipv4_tcp_max_syn_backlog}.
	NetIpv4TcpMaxSynBacklog *float64 `field:"optional" json:"netIpv4TcpMaxSynBacklog" yaml:"netIpv4TcpMaxSynBacklog"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster_node_pool#net_ipv4_tcp_max_tw_buckets KubernetesClusterNodePool#net_ipv4_tcp_max_tw_buckets}.
	NetIpv4TcpMaxTwBuckets *float64 `field:"optional" json:"netIpv4TcpMaxTwBuckets" yaml:"netIpv4TcpMaxTwBuckets"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster_node_pool#net_ipv4_tcp_tw_reuse KubernetesClusterNodePool#net_ipv4_tcp_tw_reuse}.
	NetIpv4TcpTwReuse interface{} `field:"optional" json:"netIpv4TcpTwReuse" yaml:"netIpv4TcpTwReuse"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster_node_pool#net_netfilter_nf_conntrack_buckets KubernetesClusterNodePool#net_netfilter_nf_conntrack_buckets}.
	NetNetfilterNfConntrackBuckets *float64 `field:"optional" json:"netNetfilterNfConntrackBuckets" yaml:"netNetfilterNfConntrackBuckets"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster_node_pool#net_netfilter_nf_conntrack_max KubernetesClusterNodePool#net_netfilter_nf_conntrack_max}.
	NetNetfilterNfConntrackMax *float64 `field:"optional" json:"netNetfilterNfConntrackMax" yaml:"netNetfilterNfConntrackMax"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster_node_pool#vm_max_map_count KubernetesClusterNodePool#vm_max_map_count}.
	VmMaxMapCount *float64 `field:"optional" json:"vmMaxMapCount" yaml:"vmMaxMapCount"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster_node_pool#vm_swappiness KubernetesClusterNodePool#vm_swappiness}.
	VmSwappiness *float64 `field:"optional" json:"vmSwappiness" yaml:"vmSwappiness"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster_node_pool#vm_vfs_cache_pressure KubernetesClusterNodePool#vm_vfs_cache_pressure}.
	VmVfsCachePressure *float64 `field:"optional" json:"vmVfsCachePressure" yaml:"vmVfsCachePressure"`
}

type KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference interface {
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
	InternalValue() *KubernetesClusterNodePoolLinuxOsConfigSysctlConfig
	SetInternalValue(val *KubernetesClusterNodePoolLinuxOsConfigSysctlConfig)
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

// The jsii proxy struct for KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference
type jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) FsAioMaxNr() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fsAioMaxNr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) FsAioMaxNrInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fsAioMaxNrInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) FsFileMax() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fsFileMax",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) FsFileMaxInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fsFileMaxInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) FsInotifyMaxUserWatches() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fsInotifyMaxUserWatches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) FsInotifyMaxUserWatchesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fsInotifyMaxUserWatchesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) FsNrOpen() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fsNrOpen",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) FsNrOpenInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fsNrOpenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) InternalValue() *KubernetesClusterNodePoolLinuxOsConfigSysctlConfig {
	var returns *KubernetesClusterNodePoolLinuxOsConfigSysctlConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) KernelThreadsMax() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"kernelThreadsMax",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) KernelThreadsMaxInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"kernelThreadsMaxInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) NetCoreNetdevMaxBacklog() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netCoreNetdevMaxBacklog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) NetCoreNetdevMaxBacklogInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netCoreNetdevMaxBacklogInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) NetCoreOptmemMax() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netCoreOptmemMax",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) NetCoreOptmemMaxInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netCoreOptmemMaxInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) NetCoreRmemDefault() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netCoreRmemDefault",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) NetCoreRmemDefaultInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netCoreRmemDefaultInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) NetCoreRmemMax() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netCoreRmemMax",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) NetCoreRmemMaxInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netCoreRmemMaxInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) NetCoreSomaxconn() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netCoreSomaxconn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) NetCoreSomaxconnInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netCoreSomaxconnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) NetCoreWmemDefault() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netCoreWmemDefault",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) NetCoreWmemDefaultInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netCoreWmemDefaultInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) NetCoreWmemMax() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netCoreWmemMax",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) NetCoreWmemMaxInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netCoreWmemMaxInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) NetIpv4IpLocalPortRangeMax() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netIpv4IpLocalPortRangeMax",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) NetIpv4IpLocalPortRangeMaxInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netIpv4IpLocalPortRangeMaxInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) NetIpv4IpLocalPortRangeMin() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netIpv4IpLocalPortRangeMin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) NetIpv4IpLocalPortRangeMinInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netIpv4IpLocalPortRangeMinInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) NetIpv4NeighDefaultGcThresh1() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netIpv4NeighDefaultGcThresh1",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) NetIpv4NeighDefaultGcThresh1Input() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netIpv4NeighDefaultGcThresh1Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) NetIpv4NeighDefaultGcThresh2() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netIpv4NeighDefaultGcThresh2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) NetIpv4NeighDefaultGcThresh2Input() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netIpv4NeighDefaultGcThresh2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) NetIpv4NeighDefaultGcThresh3() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netIpv4NeighDefaultGcThresh3",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) NetIpv4NeighDefaultGcThresh3Input() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netIpv4NeighDefaultGcThresh3Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) NetIpv4TcpFinTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netIpv4TcpFinTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) NetIpv4TcpFinTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netIpv4TcpFinTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) NetIpv4TcpKeepaliveIntvl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netIpv4TcpKeepaliveIntvl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) NetIpv4TcpKeepaliveIntvlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netIpv4TcpKeepaliveIntvlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) NetIpv4TcpKeepaliveProbes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netIpv4TcpKeepaliveProbes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) NetIpv4TcpKeepaliveProbesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netIpv4TcpKeepaliveProbesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) NetIpv4TcpKeepaliveTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netIpv4TcpKeepaliveTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) NetIpv4TcpKeepaliveTimeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netIpv4TcpKeepaliveTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) NetIpv4TcpMaxSynBacklog() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netIpv4TcpMaxSynBacklog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) NetIpv4TcpMaxSynBacklogInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netIpv4TcpMaxSynBacklogInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) NetIpv4TcpMaxTwBuckets() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netIpv4TcpMaxTwBuckets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) NetIpv4TcpMaxTwBucketsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netIpv4TcpMaxTwBucketsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) NetIpv4TcpTwReuse() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"netIpv4TcpTwReuse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) NetIpv4TcpTwReuseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"netIpv4TcpTwReuseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) NetNetfilterNfConntrackBuckets() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netNetfilterNfConntrackBuckets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) NetNetfilterNfConntrackBucketsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netNetfilterNfConntrackBucketsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) NetNetfilterNfConntrackMax() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netNetfilterNfConntrackMax",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) NetNetfilterNfConntrackMaxInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netNetfilterNfConntrackMaxInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) VmMaxMapCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"vmMaxMapCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) VmMaxMapCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"vmMaxMapCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) VmSwappiness() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"vmSwappiness",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) VmSwappinessInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"vmSwappinessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) VmVfsCachePressure() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"vmVfsCachePressure",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) VmVfsCachePressureInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"vmVfsCachePressureInput",
		&returns,
	)
	return returns
}


func NewKubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference {
	_init_.Initialize()

	j := jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesClusterNodePool.KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference_Override(k KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesClusterNodePool.KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) SetFsAioMaxNr(val *float64) {
	_jsii_.Set(
		j,
		"fsAioMaxNr",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) SetFsFileMax(val *float64) {
	_jsii_.Set(
		j,
		"fsFileMax",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) SetFsInotifyMaxUserWatches(val *float64) {
	_jsii_.Set(
		j,
		"fsInotifyMaxUserWatches",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) SetFsNrOpen(val *float64) {
	_jsii_.Set(
		j,
		"fsNrOpen",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) SetInternalValue(val *KubernetesClusterNodePoolLinuxOsConfigSysctlConfig) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) SetKernelThreadsMax(val *float64) {
	_jsii_.Set(
		j,
		"kernelThreadsMax",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) SetNetCoreNetdevMaxBacklog(val *float64) {
	_jsii_.Set(
		j,
		"netCoreNetdevMaxBacklog",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) SetNetCoreOptmemMax(val *float64) {
	_jsii_.Set(
		j,
		"netCoreOptmemMax",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) SetNetCoreRmemDefault(val *float64) {
	_jsii_.Set(
		j,
		"netCoreRmemDefault",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) SetNetCoreRmemMax(val *float64) {
	_jsii_.Set(
		j,
		"netCoreRmemMax",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) SetNetCoreSomaxconn(val *float64) {
	_jsii_.Set(
		j,
		"netCoreSomaxconn",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) SetNetCoreWmemDefault(val *float64) {
	_jsii_.Set(
		j,
		"netCoreWmemDefault",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) SetNetCoreWmemMax(val *float64) {
	_jsii_.Set(
		j,
		"netCoreWmemMax",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) SetNetIpv4IpLocalPortRangeMax(val *float64) {
	_jsii_.Set(
		j,
		"netIpv4IpLocalPortRangeMax",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) SetNetIpv4IpLocalPortRangeMin(val *float64) {
	_jsii_.Set(
		j,
		"netIpv4IpLocalPortRangeMin",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) SetNetIpv4NeighDefaultGcThresh1(val *float64) {
	_jsii_.Set(
		j,
		"netIpv4NeighDefaultGcThresh1",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) SetNetIpv4NeighDefaultGcThresh2(val *float64) {
	_jsii_.Set(
		j,
		"netIpv4NeighDefaultGcThresh2",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) SetNetIpv4NeighDefaultGcThresh3(val *float64) {
	_jsii_.Set(
		j,
		"netIpv4NeighDefaultGcThresh3",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) SetNetIpv4TcpFinTimeout(val *float64) {
	_jsii_.Set(
		j,
		"netIpv4TcpFinTimeout",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) SetNetIpv4TcpKeepaliveIntvl(val *float64) {
	_jsii_.Set(
		j,
		"netIpv4TcpKeepaliveIntvl",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) SetNetIpv4TcpKeepaliveProbes(val *float64) {
	_jsii_.Set(
		j,
		"netIpv4TcpKeepaliveProbes",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) SetNetIpv4TcpKeepaliveTime(val *float64) {
	_jsii_.Set(
		j,
		"netIpv4TcpKeepaliveTime",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) SetNetIpv4TcpMaxSynBacklog(val *float64) {
	_jsii_.Set(
		j,
		"netIpv4TcpMaxSynBacklog",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) SetNetIpv4TcpMaxTwBuckets(val *float64) {
	_jsii_.Set(
		j,
		"netIpv4TcpMaxTwBuckets",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) SetNetIpv4TcpTwReuse(val interface{}) {
	_jsii_.Set(
		j,
		"netIpv4TcpTwReuse",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) SetNetNetfilterNfConntrackBuckets(val *float64) {
	_jsii_.Set(
		j,
		"netNetfilterNfConntrackBuckets",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) SetNetNetfilterNfConntrackMax(val *float64) {
	_jsii_.Set(
		j,
		"netNetfilterNfConntrackMax",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) SetVmMaxMapCount(val *float64) {
	_jsii_.Set(
		j,
		"vmMaxMapCount",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) SetVmSwappiness(val *float64) {
	_jsii_.Set(
		j,
		"vmSwappiness",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) SetVmVfsCachePressure(val *float64) {
	_jsii_.Set(
		j,
		"vmVfsCachePressure",
		val,
	)
}

func (k *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		k,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		k,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		k,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		k,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		k,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) ResetFsAioMaxNr() {
	_jsii_.InvokeVoid(
		k,
		"resetFsAioMaxNr",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) ResetFsFileMax() {
	_jsii_.InvokeVoid(
		k,
		"resetFsFileMax",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) ResetFsInotifyMaxUserWatches() {
	_jsii_.InvokeVoid(
		k,
		"resetFsInotifyMaxUserWatches",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) ResetFsNrOpen() {
	_jsii_.InvokeVoid(
		k,
		"resetFsNrOpen",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) ResetKernelThreadsMax() {
	_jsii_.InvokeVoid(
		k,
		"resetKernelThreadsMax",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) ResetNetCoreNetdevMaxBacklog() {
	_jsii_.InvokeVoid(
		k,
		"resetNetCoreNetdevMaxBacklog",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) ResetNetCoreOptmemMax() {
	_jsii_.InvokeVoid(
		k,
		"resetNetCoreOptmemMax",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) ResetNetCoreRmemDefault() {
	_jsii_.InvokeVoid(
		k,
		"resetNetCoreRmemDefault",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) ResetNetCoreRmemMax() {
	_jsii_.InvokeVoid(
		k,
		"resetNetCoreRmemMax",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) ResetNetCoreSomaxconn() {
	_jsii_.InvokeVoid(
		k,
		"resetNetCoreSomaxconn",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) ResetNetCoreWmemDefault() {
	_jsii_.InvokeVoid(
		k,
		"resetNetCoreWmemDefault",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) ResetNetCoreWmemMax() {
	_jsii_.InvokeVoid(
		k,
		"resetNetCoreWmemMax",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) ResetNetIpv4IpLocalPortRangeMax() {
	_jsii_.InvokeVoid(
		k,
		"resetNetIpv4IpLocalPortRangeMax",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) ResetNetIpv4IpLocalPortRangeMin() {
	_jsii_.InvokeVoid(
		k,
		"resetNetIpv4IpLocalPortRangeMin",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) ResetNetIpv4NeighDefaultGcThresh1() {
	_jsii_.InvokeVoid(
		k,
		"resetNetIpv4NeighDefaultGcThresh1",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) ResetNetIpv4NeighDefaultGcThresh2() {
	_jsii_.InvokeVoid(
		k,
		"resetNetIpv4NeighDefaultGcThresh2",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) ResetNetIpv4NeighDefaultGcThresh3() {
	_jsii_.InvokeVoid(
		k,
		"resetNetIpv4NeighDefaultGcThresh3",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) ResetNetIpv4TcpFinTimeout() {
	_jsii_.InvokeVoid(
		k,
		"resetNetIpv4TcpFinTimeout",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) ResetNetIpv4TcpKeepaliveIntvl() {
	_jsii_.InvokeVoid(
		k,
		"resetNetIpv4TcpKeepaliveIntvl",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) ResetNetIpv4TcpKeepaliveProbes() {
	_jsii_.InvokeVoid(
		k,
		"resetNetIpv4TcpKeepaliveProbes",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) ResetNetIpv4TcpKeepaliveTime() {
	_jsii_.InvokeVoid(
		k,
		"resetNetIpv4TcpKeepaliveTime",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) ResetNetIpv4TcpMaxSynBacklog() {
	_jsii_.InvokeVoid(
		k,
		"resetNetIpv4TcpMaxSynBacklog",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) ResetNetIpv4TcpMaxTwBuckets() {
	_jsii_.InvokeVoid(
		k,
		"resetNetIpv4TcpMaxTwBuckets",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) ResetNetIpv4TcpTwReuse() {
	_jsii_.InvokeVoid(
		k,
		"resetNetIpv4TcpTwReuse",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) ResetNetNetfilterNfConntrackBuckets() {
	_jsii_.InvokeVoid(
		k,
		"resetNetNetfilterNfConntrackBuckets",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) ResetNetNetfilterNfConntrackMax() {
	_jsii_.InvokeVoid(
		k,
		"resetNetNetfilterNfConntrackMax",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) ResetVmMaxMapCount() {
	_jsii_.InvokeVoid(
		k,
		"resetVmMaxMapCount",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) ResetVmSwappiness() {
	_jsii_.InvokeVoid(
		k,
		"resetVmSwappiness",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) ResetVmVfsCachePressure() {
	_jsii_.InvokeVoid(
		k,
		"resetVmVfsCachePressure",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNodePoolLinuxOsConfigSysctlConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type KubernetesClusterNodePoolTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster_node_pool#create KubernetesClusterNodePool#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster_node_pool#delete KubernetesClusterNodePool#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster_node_pool#read KubernetesClusterNodePool#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster_node_pool#update KubernetesClusterNodePool#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

type KubernetesClusterNodePoolTimeoutsOutputReference interface {
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

// The jsii proxy struct for KubernetesClusterNodePoolTimeoutsOutputReference
type jsiiProxy_KubernetesClusterNodePoolTimeoutsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_KubernetesClusterNodePoolTimeoutsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolTimeoutsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolTimeoutsOutputReference) Create() *string {
	var returns *string
	_jsii_.Get(
		j,
		"create",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolTimeoutsOutputReference) CreateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolTimeoutsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolTimeoutsOutputReference) Delete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolTimeoutsOutputReference) DeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolTimeoutsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolTimeoutsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolTimeoutsOutputReference) Read() *string {
	var returns *string
	_jsii_.Get(
		j,
		"read",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolTimeoutsOutputReference) ReadInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolTimeoutsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolTimeoutsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolTimeoutsOutputReference) Update() *string {
	var returns *string
	_jsii_.Get(
		j,
		"update",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolTimeoutsOutputReference) UpdateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateInput",
		&returns,
	)
	return returns
}


func NewKubernetesClusterNodePoolTimeoutsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) KubernetesClusterNodePoolTimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_KubernetesClusterNodePoolTimeoutsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesClusterNodePool.KubernetesClusterNodePoolTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKubernetesClusterNodePoolTimeoutsOutputReference_Override(k KubernetesClusterNodePoolTimeoutsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesClusterNodePool.KubernetesClusterNodePoolTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePoolTimeoutsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePoolTimeoutsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePoolTimeoutsOutputReference) SetCreate(val *string) {
	_jsii_.Set(
		j,
		"create",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePoolTimeoutsOutputReference) SetDelete(val *string) {
	_jsii_.Set(
		j,
		"delete",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePoolTimeoutsOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePoolTimeoutsOutputReference) SetRead(val *string) {
	_jsii_.Set(
		j,
		"read",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePoolTimeoutsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePoolTimeoutsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePoolTimeoutsOutputReference) SetUpdate(val *string) {
	_jsii_.Set(
		j,
		"update",
		val,
	)
}

func (k *jsiiProxy_KubernetesClusterNodePoolTimeoutsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNodePoolTimeoutsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNodePoolTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNodePoolTimeoutsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		k,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNodePoolTimeoutsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNodePoolTimeoutsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		k,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNodePoolTimeoutsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		k,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNodePoolTimeoutsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		k,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNodePoolTimeoutsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNodePoolTimeoutsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		k,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNodePoolTimeoutsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNodePoolTimeoutsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNodePoolTimeoutsOutputReference) ResetCreate() {
	_jsii_.InvokeVoid(
		k,
		"resetCreate",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNodePoolTimeoutsOutputReference) ResetDelete() {
	_jsii_.InvokeVoid(
		k,
		"resetDelete",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNodePoolTimeoutsOutputReference) ResetRead() {
	_jsii_.InvokeVoid(
		k,
		"resetRead",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNodePoolTimeoutsOutputReference) ResetUpdate() {
	_jsii_.InvokeVoid(
		k,
		"resetUpdate",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterNodePoolTimeoutsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNodePoolTimeoutsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type KubernetesClusterNodePoolUpgradeSettings struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster_node_pool#max_surge KubernetesClusterNodePool#max_surge}.
	MaxSurge *string `field:"required" json:"maxSurge" yaml:"maxSurge"`
}

type KubernetesClusterNodePoolUpgradeSettingsOutputReference interface {
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
	InternalValue() *KubernetesClusterNodePoolUpgradeSettings
	SetInternalValue(val *KubernetesClusterNodePoolUpgradeSettings)
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

// The jsii proxy struct for KubernetesClusterNodePoolUpgradeSettingsOutputReference
type jsiiProxy_KubernetesClusterNodePoolUpgradeSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_KubernetesClusterNodePoolUpgradeSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolUpgradeSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolUpgradeSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolUpgradeSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolUpgradeSettingsOutputReference) InternalValue() *KubernetesClusterNodePoolUpgradeSettings {
	var returns *KubernetesClusterNodePoolUpgradeSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolUpgradeSettingsOutputReference) MaxSurge() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxSurge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolUpgradeSettingsOutputReference) MaxSurgeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxSurgeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolUpgradeSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterNodePoolUpgradeSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewKubernetesClusterNodePoolUpgradeSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) KubernetesClusterNodePoolUpgradeSettingsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_KubernetesClusterNodePoolUpgradeSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesClusterNodePool.KubernetesClusterNodePoolUpgradeSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKubernetesClusterNodePoolUpgradeSettingsOutputReference_Override(k KubernetesClusterNodePoolUpgradeSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesClusterNodePool.KubernetesClusterNodePoolUpgradeSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePoolUpgradeSettingsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePoolUpgradeSettingsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePoolUpgradeSettingsOutputReference) SetInternalValue(val *KubernetesClusterNodePoolUpgradeSettings) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePoolUpgradeSettingsOutputReference) SetMaxSurge(val *string) {
	_jsii_.Set(
		j,
		"maxSurge",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePoolUpgradeSettingsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterNodePoolUpgradeSettingsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (k *jsiiProxy_KubernetesClusterNodePoolUpgradeSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNodePoolUpgradeSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNodePoolUpgradeSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNodePoolUpgradeSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		k,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNodePoolUpgradeSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNodePoolUpgradeSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		k,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNodePoolUpgradeSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		k,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNodePoolUpgradeSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		k,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNodePoolUpgradeSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNodePoolUpgradeSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		k,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNodePoolUpgradeSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNodePoolUpgradeSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNodePoolUpgradeSettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterNodePoolUpgradeSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

