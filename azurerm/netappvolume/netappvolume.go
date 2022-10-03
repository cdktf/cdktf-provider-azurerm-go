package netappvolume

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/netappvolume/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/azurerm/r/netapp_volume azurerm_netapp_volume}.
type NetappVolume interface {
	cdktf.TerraformResource
	AccountName() *string
	SetAccountName(val *string)
	AccountNameInput() *string
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
	CreateFromSnapshotResourceId() *string
	SetCreateFromSnapshotResourceId(val *string)
	CreateFromSnapshotResourceIdInput() *string
	DataProtectionReplication() NetappVolumeDataProtectionReplicationOutputReference
	DataProtectionReplicationInput() *NetappVolumeDataProtectionReplication
	DataProtectionSnapshotPolicy() NetappVolumeDataProtectionSnapshotPolicyOutputReference
	DataProtectionSnapshotPolicyInput() *NetappVolumeDataProtectionSnapshotPolicy
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	ExportPolicyRule() NetappVolumeExportPolicyRuleList
	ExportPolicyRuleInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	MountIpAddresses() *[]*string
	Name() *string
	SetName(val *string)
	NameInput() *string
	NetworkFeatures() *string
	SetNetworkFeatures(val *string)
	NetworkFeaturesInput() *string
	// The tree node.
	Node() constructs.Node
	PoolName() *string
	SetPoolName(val *string)
	PoolNameInput() *string
	Protocols() *[]*string
	SetProtocols(val *[]*string)
	ProtocolsInput() *[]*string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	ResourceGroupName() *string
	SetResourceGroupName(val *string)
	ResourceGroupNameInput() *string
	SecurityStyle() *string
	SetSecurityStyle(val *string)
	SecurityStyleInput() *string
	ServiceLevel() *string
	SetServiceLevel(val *string)
	ServiceLevelInput() *string
	SnapshotDirectoryVisible() interface{}
	SetSnapshotDirectoryVisible(val interface{})
	SnapshotDirectoryVisibleInput() interface{}
	StorageQuotaInGb() *float64
	SetStorageQuotaInGb(val *float64)
	StorageQuotaInGbInput() *float64
	SubnetId() *string
	SetSubnetId(val *string)
	SubnetIdInput() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	ThroughputInMibps() *float64
	SetThroughputInMibps(val *float64)
	ThroughputInMibpsInput() *float64
	Timeouts() NetappVolumeTimeoutsOutputReference
	TimeoutsInput() interface{}
	VolumePath() *string
	SetVolumePath(val *string)
	VolumePathInput() *string
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
	PutDataProtectionReplication(value *NetappVolumeDataProtectionReplication)
	PutDataProtectionSnapshotPolicy(value *NetappVolumeDataProtectionSnapshotPolicy)
	PutExportPolicyRule(value interface{})
	PutTimeouts(value *NetappVolumeTimeouts)
	ResetCreateFromSnapshotResourceId()
	ResetDataProtectionReplication()
	ResetDataProtectionSnapshotPolicy()
	ResetExportPolicyRule()
	ResetId()
	ResetNetworkFeatures()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProtocols()
	ResetSecurityStyle()
	ResetSnapshotDirectoryVisible()
	ResetTags()
	ResetThroughputInMibps()
	ResetTimeouts()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for NetappVolume
type jsiiProxy_NetappVolume struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_NetappVolume) AccountName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) AccountNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) CreateFromSnapshotResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createFromSnapshotResourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) CreateFromSnapshotResourceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createFromSnapshotResourceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) DataProtectionReplication() NetappVolumeDataProtectionReplicationOutputReference {
	var returns NetappVolumeDataProtectionReplicationOutputReference
	_jsii_.Get(
		j,
		"dataProtectionReplication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) DataProtectionReplicationInput() *NetappVolumeDataProtectionReplication {
	var returns *NetappVolumeDataProtectionReplication
	_jsii_.Get(
		j,
		"dataProtectionReplicationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) DataProtectionSnapshotPolicy() NetappVolumeDataProtectionSnapshotPolicyOutputReference {
	var returns NetappVolumeDataProtectionSnapshotPolicyOutputReference
	_jsii_.Get(
		j,
		"dataProtectionSnapshotPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) DataProtectionSnapshotPolicyInput() *NetappVolumeDataProtectionSnapshotPolicy {
	var returns *NetappVolumeDataProtectionSnapshotPolicy
	_jsii_.Get(
		j,
		"dataProtectionSnapshotPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) ExportPolicyRule() NetappVolumeExportPolicyRuleList {
	var returns NetappVolumeExportPolicyRuleList
	_jsii_.Get(
		j,
		"exportPolicyRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) ExportPolicyRuleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"exportPolicyRuleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) MountIpAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"mountIpAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) NetworkFeatures() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkFeatures",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) NetworkFeaturesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkFeaturesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) PoolName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"poolName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) PoolNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"poolNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) Protocols() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"protocols",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) ProtocolsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"protocolsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) SecurityStyle() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityStyle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) SecurityStyleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityStyleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) ServiceLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) ServiceLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) SnapshotDirectoryVisible() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"snapshotDirectoryVisible",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) SnapshotDirectoryVisibleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"snapshotDirectoryVisibleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) StorageQuotaInGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"storageQuotaInGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) StorageQuotaInGbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"storageQuotaInGbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) SubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) SubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) ThroughputInMibps() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"throughputInMibps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) ThroughputInMibpsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"throughputInMibpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) Timeouts() NetappVolumeTimeoutsOutputReference {
	var returns NetappVolumeTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) VolumePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) VolumePathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumePathInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/netapp_volume azurerm_netapp_volume} Resource.
func NewNetappVolume(scope constructs.Construct, id *string, config *NetappVolumeConfig) NetappVolume {
	_init_.Initialize()

	j := jsiiProxy_NetappVolume{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.netappVolume.NetappVolume",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/netapp_volume azurerm_netapp_volume} Resource.
func NewNetappVolume_Override(n NetappVolume, scope constructs.Construct, id *string, config *NetappVolumeConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.netappVolume.NetappVolume",
		[]interface{}{scope, id, config},
		n,
	)
}

func (j *jsiiProxy_NetappVolume) SetAccountName(val *string) {
	_jsii_.Set(
		j,
		"accountName",
		val,
	)
}

func (j *jsiiProxy_NetappVolume) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_NetappVolume) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_NetappVolume) SetCreateFromSnapshotResourceId(val *string) {
	_jsii_.Set(
		j,
		"createFromSnapshotResourceId",
		val,
	)
}

func (j *jsiiProxy_NetappVolume) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_NetappVolume) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_NetappVolume) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_NetappVolume) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_NetappVolume) SetLocation(val *string) {
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_NetappVolume) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_NetappVolume) SetNetworkFeatures(val *string) {
	_jsii_.Set(
		j,
		"networkFeatures",
		val,
	)
}

func (j *jsiiProxy_NetappVolume) SetPoolName(val *string) {
	_jsii_.Set(
		j,
		"poolName",
		val,
	)
}

func (j *jsiiProxy_NetappVolume) SetProtocols(val *[]*string) {
	_jsii_.Set(
		j,
		"protocols",
		val,
	)
}

func (j *jsiiProxy_NetappVolume) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_NetappVolume) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_NetappVolume) SetResourceGroupName(val *string) {
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_NetappVolume) SetSecurityStyle(val *string) {
	_jsii_.Set(
		j,
		"securityStyle",
		val,
	)
}

func (j *jsiiProxy_NetappVolume) SetServiceLevel(val *string) {
	_jsii_.Set(
		j,
		"serviceLevel",
		val,
	)
}

func (j *jsiiProxy_NetappVolume) SetSnapshotDirectoryVisible(val interface{}) {
	_jsii_.Set(
		j,
		"snapshotDirectoryVisible",
		val,
	)
}

func (j *jsiiProxy_NetappVolume) SetStorageQuotaInGb(val *float64) {
	_jsii_.Set(
		j,
		"storageQuotaInGb",
		val,
	)
}

func (j *jsiiProxy_NetappVolume) SetSubnetId(val *string) {
	_jsii_.Set(
		j,
		"subnetId",
		val,
	)
}

func (j *jsiiProxy_NetappVolume) SetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_NetappVolume) SetThroughputInMibps(val *float64) {
	_jsii_.Set(
		j,
		"throughputInMibps",
		val,
	)
}

func (j *jsiiProxy_NetappVolume) SetVolumePath(val *string) {
	_jsii_.Set(
		j,
		"volumePath",
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
func NetappVolume_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.netappVolume.NetappVolume",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func NetappVolume_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.netappVolume.NetappVolume",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (n *jsiiProxy_NetappVolume) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		n,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (n *jsiiProxy_NetappVolume) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolume) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolume) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		n,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolume) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		n,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolume) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		n,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolume) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		n,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolume) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		n,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolume) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolume) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		n,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolume) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolume) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		n,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (n *jsiiProxy_NetappVolume) PutDataProtectionReplication(value *NetappVolumeDataProtectionReplication) {
	_jsii_.InvokeVoid(
		n,
		"putDataProtectionReplication",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetappVolume) PutDataProtectionSnapshotPolicy(value *NetappVolumeDataProtectionSnapshotPolicy) {
	_jsii_.InvokeVoid(
		n,
		"putDataProtectionSnapshotPolicy",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetappVolume) PutExportPolicyRule(value interface{}) {
	_jsii_.InvokeVoid(
		n,
		"putExportPolicyRule",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetappVolume) PutTimeouts(value *NetappVolumeTimeouts) {
	_jsii_.InvokeVoid(
		n,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetappVolume) ResetCreateFromSnapshotResourceId() {
	_jsii_.InvokeVoid(
		n,
		"resetCreateFromSnapshotResourceId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolume) ResetDataProtectionReplication() {
	_jsii_.InvokeVoid(
		n,
		"resetDataProtectionReplication",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolume) ResetDataProtectionSnapshotPolicy() {
	_jsii_.InvokeVoid(
		n,
		"resetDataProtectionSnapshotPolicy",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolume) ResetExportPolicyRule() {
	_jsii_.InvokeVoid(
		n,
		"resetExportPolicyRule",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolume) ResetId() {
	_jsii_.InvokeVoid(
		n,
		"resetId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolume) ResetNetworkFeatures() {
	_jsii_.InvokeVoid(
		n,
		"resetNetworkFeatures",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolume) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		n,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolume) ResetProtocols() {
	_jsii_.InvokeVoid(
		n,
		"resetProtocols",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolume) ResetSecurityStyle() {
	_jsii_.InvokeVoid(
		n,
		"resetSecurityStyle",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolume) ResetSnapshotDirectoryVisible() {
	_jsii_.InvokeVoid(
		n,
		"resetSnapshotDirectoryVisible",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolume) ResetTags() {
	_jsii_.InvokeVoid(
		n,
		"resetTags",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolume) ResetThroughputInMibps() {
	_jsii_.InvokeVoid(
		n,
		"resetThroughputInMibps",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolume) ResetTimeouts() {
	_jsii_.InvokeVoid(
		n,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolume) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolume) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolume) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolume) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type NetappVolumeConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/netapp_volume#account_name NetappVolume#account_name}.
	AccountName *string `field:"required" json:"accountName" yaml:"accountName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/netapp_volume#location NetappVolume#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/netapp_volume#name NetappVolume#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/netapp_volume#pool_name NetappVolume#pool_name}.
	PoolName *string `field:"required" json:"poolName" yaml:"poolName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/netapp_volume#resource_group_name NetappVolume#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/netapp_volume#service_level NetappVolume#service_level}.
	ServiceLevel *string `field:"required" json:"serviceLevel" yaml:"serviceLevel"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/netapp_volume#storage_quota_in_gb NetappVolume#storage_quota_in_gb}.
	StorageQuotaInGb *float64 `field:"required" json:"storageQuotaInGb" yaml:"storageQuotaInGb"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/netapp_volume#subnet_id NetappVolume#subnet_id}.
	SubnetId *string `field:"required" json:"subnetId" yaml:"subnetId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/netapp_volume#volume_path NetappVolume#volume_path}.
	VolumePath *string `field:"required" json:"volumePath" yaml:"volumePath"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/netapp_volume#create_from_snapshot_resource_id NetappVolume#create_from_snapshot_resource_id}.
	CreateFromSnapshotResourceId *string `field:"optional" json:"createFromSnapshotResourceId" yaml:"createFromSnapshotResourceId"`
	// data_protection_replication block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/netapp_volume#data_protection_replication NetappVolume#data_protection_replication}
	DataProtectionReplication *NetappVolumeDataProtectionReplication `field:"optional" json:"dataProtectionReplication" yaml:"dataProtectionReplication"`
	// data_protection_snapshot_policy block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/netapp_volume#data_protection_snapshot_policy NetappVolume#data_protection_snapshot_policy}
	DataProtectionSnapshotPolicy *NetappVolumeDataProtectionSnapshotPolicy `field:"optional" json:"dataProtectionSnapshotPolicy" yaml:"dataProtectionSnapshotPolicy"`
	// export_policy_rule block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/netapp_volume#export_policy_rule NetappVolume#export_policy_rule}
	ExportPolicyRule interface{} `field:"optional" json:"exportPolicyRule" yaml:"exportPolicyRule"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/netapp_volume#id NetappVolume#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/netapp_volume#network_features NetappVolume#network_features}.
	NetworkFeatures *string `field:"optional" json:"networkFeatures" yaml:"networkFeatures"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/netapp_volume#protocols NetappVolume#protocols}.
	Protocols *[]*string `field:"optional" json:"protocols" yaml:"protocols"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/netapp_volume#security_style NetappVolume#security_style}.
	SecurityStyle *string `field:"optional" json:"securityStyle" yaml:"securityStyle"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/netapp_volume#snapshot_directory_visible NetappVolume#snapshot_directory_visible}.
	SnapshotDirectoryVisible interface{} `field:"optional" json:"snapshotDirectoryVisible" yaml:"snapshotDirectoryVisible"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/netapp_volume#tags NetappVolume#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/netapp_volume#throughput_in_mibps NetappVolume#throughput_in_mibps}.
	ThroughputInMibps *float64 `field:"optional" json:"throughputInMibps" yaml:"throughputInMibps"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/netapp_volume#timeouts NetappVolume#timeouts}
	Timeouts *NetappVolumeTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

type NetappVolumeDataProtectionReplication struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/netapp_volume#remote_volume_location NetappVolume#remote_volume_location}.
	RemoteVolumeLocation *string `field:"required" json:"remoteVolumeLocation" yaml:"remoteVolumeLocation"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/netapp_volume#remote_volume_resource_id NetappVolume#remote_volume_resource_id}.
	RemoteVolumeResourceId *string `field:"required" json:"remoteVolumeResourceId" yaml:"remoteVolumeResourceId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/netapp_volume#replication_frequency NetappVolume#replication_frequency}.
	ReplicationFrequency *string `field:"required" json:"replicationFrequency" yaml:"replicationFrequency"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/netapp_volume#endpoint_type NetappVolume#endpoint_type}.
	EndpointType *string `field:"optional" json:"endpointType" yaml:"endpointType"`
}

type NetappVolumeDataProtectionReplicationOutputReference interface {
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
	EndpointType() *string
	SetEndpointType(val *string)
	EndpointTypeInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *NetappVolumeDataProtectionReplication
	SetInternalValue(val *NetappVolumeDataProtectionReplication)
	RemoteVolumeLocation() *string
	SetRemoteVolumeLocation(val *string)
	RemoteVolumeLocationInput() *string
	RemoteVolumeResourceId() *string
	SetRemoteVolumeResourceId(val *string)
	RemoteVolumeResourceIdInput() *string
	ReplicationFrequency() *string
	SetReplicationFrequency(val *string)
	ReplicationFrequencyInput() *string
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
	ResetEndpointType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for NetappVolumeDataProtectionReplicationOutputReference
type jsiiProxy_NetappVolumeDataProtectionReplicationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_NetappVolumeDataProtectionReplicationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeDataProtectionReplicationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeDataProtectionReplicationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeDataProtectionReplicationOutputReference) EndpointType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeDataProtectionReplicationOutputReference) EndpointTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeDataProtectionReplicationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeDataProtectionReplicationOutputReference) InternalValue() *NetappVolumeDataProtectionReplication {
	var returns *NetappVolumeDataProtectionReplication
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeDataProtectionReplicationOutputReference) RemoteVolumeLocation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"remoteVolumeLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeDataProtectionReplicationOutputReference) RemoteVolumeLocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"remoteVolumeLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeDataProtectionReplicationOutputReference) RemoteVolumeResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"remoteVolumeResourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeDataProtectionReplicationOutputReference) RemoteVolumeResourceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"remoteVolumeResourceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeDataProtectionReplicationOutputReference) ReplicationFrequency() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicationFrequency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeDataProtectionReplicationOutputReference) ReplicationFrequencyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicationFrequencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeDataProtectionReplicationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeDataProtectionReplicationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewNetappVolumeDataProtectionReplicationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) NetappVolumeDataProtectionReplicationOutputReference {
	_init_.Initialize()

	j := jsiiProxy_NetappVolumeDataProtectionReplicationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.netappVolume.NetappVolumeDataProtectionReplicationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewNetappVolumeDataProtectionReplicationOutputReference_Override(n NetappVolumeDataProtectionReplicationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.netappVolume.NetappVolumeDataProtectionReplicationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		n,
	)
}

func (j *jsiiProxy_NetappVolumeDataProtectionReplicationOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeDataProtectionReplicationOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeDataProtectionReplicationOutputReference) SetEndpointType(val *string) {
	_jsii_.Set(
		j,
		"endpointType",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeDataProtectionReplicationOutputReference) SetInternalValue(val *NetappVolumeDataProtectionReplication) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeDataProtectionReplicationOutputReference) SetRemoteVolumeLocation(val *string) {
	_jsii_.Set(
		j,
		"remoteVolumeLocation",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeDataProtectionReplicationOutputReference) SetRemoteVolumeResourceId(val *string) {
	_jsii_.Set(
		j,
		"remoteVolumeResourceId",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeDataProtectionReplicationOutputReference) SetReplicationFrequency(val *string) {
	_jsii_.Set(
		j,
		"replicationFrequency",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeDataProtectionReplicationOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeDataProtectionReplicationOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (n *jsiiProxy_NetappVolumeDataProtectionReplicationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeDataProtectionReplicationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeDataProtectionReplicationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeDataProtectionReplicationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		n,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeDataProtectionReplicationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		n,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeDataProtectionReplicationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		n,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeDataProtectionReplicationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		n,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeDataProtectionReplicationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		n,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeDataProtectionReplicationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeDataProtectionReplicationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		n,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeDataProtectionReplicationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeDataProtectionReplicationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeDataProtectionReplicationOutputReference) ResetEndpointType() {
	_jsii_.InvokeVoid(
		n,
		"resetEndpointType",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolumeDataProtectionReplicationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeDataProtectionReplicationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type NetappVolumeDataProtectionSnapshotPolicy struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/netapp_volume#snapshot_policy_id NetappVolume#snapshot_policy_id}.
	SnapshotPolicyId *string `field:"required" json:"snapshotPolicyId" yaml:"snapshotPolicyId"`
}

type NetappVolumeDataProtectionSnapshotPolicyOutputReference interface {
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
	InternalValue() *NetappVolumeDataProtectionSnapshotPolicy
	SetInternalValue(val *NetappVolumeDataProtectionSnapshotPolicy)
	SnapshotPolicyId() *string
	SetSnapshotPolicyId(val *string)
	SnapshotPolicyIdInput() *string
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

// The jsii proxy struct for NetappVolumeDataProtectionSnapshotPolicyOutputReference
type jsiiProxy_NetappVolumeDataProtectionSnapshotPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_NetappVolumeDataProtectionSnapshotPolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeDataProtectionSnapshotPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeDataProtectionSnapshotPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeDataProtectionSnapshotPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeDataProtectionSnapshotPolicyOutputReference) InternalValue() *NetappVolumeDataProtectionSnapshotPolicy {
	var returns *NetappVolumeDataProtectionSnapshotPolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeDataProtectionSnapshotPolicyOutputReference) SnapshotPolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotPolicyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeDataProtectionSnapshotPolicyOutputReference) SnapshotPolicyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotPolicyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeDataProtectionSnapshotPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeDataProtectionSnapshotPolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewNetappVolumeDataProtectionSnapshotPolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) NetappVolumeDataProtectionSnapshotPolicyOutputReference {
	_init_.Initialize()

	j := jsiiProxy_NetappVolumeDataProtectionSnapshotPolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.netappVolume.NetappVolumeDataProtectionSnapshotPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewNetappVolumeDataProtectionSnapshotPolicyOutputReference_Override(n NetappVolumeDataProtectionSnapshotPolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.netappVolume.NetappVolumeDataProtectionSnapshotPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		n,
	)
}

func (j *jsiiProxy_NetappVolumeDataProtectionSnapshotPolicyOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeDataProtectionSnapshotPolicyOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeDataProtectionSnapshotPolicyOutputReference) SetInternalValue(val *NetappVolumeDataProtectionSnapshotPolicy) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeDataProtectionSnapshotPolicyOutputReference) SetSnapshotPolicyId(val *string) {
	_jsii_.Set(
		j,
		"snapshotPolicyId",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeDataProtectionSnapshotPolicyOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeDataProtectionSnapshotPolicyOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (n *jsiiProxy_NetappVolumeDataProtectionSnapshotPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeDataProtectionSnapshotPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeDataProtectionSnapshotPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeDataProtectionSnapshotPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		n,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeDataProtectionSnapshotPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		n,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeDataProtectionSnapshotPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		n,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeDataProtectionSnapshotPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		n,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeDataProtectionSnapshotPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		n,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeDataProtectionSnapshotPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeDataProtectionSnapshotPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		n,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeDataProtectionSnapshotPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeDataProtectionSnapshotPolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeDataProtectionSnapshotPolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeDataProtectionSnapshotPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type NetappVolumeExportPolicyRule struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/netapp_volume#allowed_clients NetappVolume#allowed_clients}.
	AllowedClients *[]*string `field:"required" json:"allowedClients" yaml:"allowedClients"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/netapp_volume#rule_index NetappVolume#rule_index}.
	RuleIndex *float64 `field:"required" json:"ruleIndex" yaml:"ruleIndex"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/netapp_volume#protocols_enabled NetappVolume#protocols_enabled}.
	ProtocolsEnabled *[]*string `field:"optional" json:"protocolsEnabled" yaml:"protocolsEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/netapp_volume#root_access_enabled NetappVolume#root_access_enabled}.
	RootAccessEnabled interface{} `field:"optional" json:"rootAccessEnabled" yaml:"rootAccessEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/netapp_volume#unix_read_only NetappVolume#unix_read_only}.
	UnixReadOnly interface{} `field:"optional" json:"unixReadOnly" yaml:"unixReadOnly"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/netapp_volume#unix_read_write NetappVolume#unix_read_write}.
	UnixReadWrite interface{} `field:"optional" json:"unixReadWrite" yaml:"unixReadWrite"`
}

type NetappVolumeExportPolicyRuleList interface {
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
	Get(index *float64) NetappVolumeExportPolicyRuleOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for NetappVolumeExportPolicyRuleList
type jsiiProxy_NetappVolumeExportPolicyRuleList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_NetappVolumeExportPolicyRuleList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeExportPolicyRuleList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeExportPolicyRuleList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeExportPolicyRuleList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeExportPolicyRuleList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeExportPolicyRuleList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewNetappVolumeExportPolicyRuleList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) NetappVolumeExportPolicyRuleList {
	_init_.Initialize()

	j := jsiiProxy_NetappVolumeExportPolicyRuleList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.netappVolume.NetappVolumeExportPolicyRuleList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewNetappVolumeExportPolicyRuleList_Override(n NetappVolumeExportPolicyRuleList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.netappVolume.NetappVolumeExportPolicyRuleList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		n,
	)
}

func (j *jsiiProxy_NetappVolumeExportPolicyRuleList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeExportPolicyRuleList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeExportPolicyRuleList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeExportPolicyRuleList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (n *jsiiProxy_NetappVolumeExportPolicyRuleList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeExportPolicyRuleList) Get(index *float64) NetappVolumeExportPolicyRuleOutputReference {
	var returns NetappVolumeExportPolicyRuleOutputReference

	_jsii_.Invoke(
		n,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeExportPolicyRuleList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeExportPolicyRuleList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type NetappVolumeExportPolicyRuleOutputReference interface {
	cdktf.ComplexObject
	AllowedClients() *[]*string
	SetAllowedClients(val *[]*string)
	AllowedClientsInput() *[]*string
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	ProtocolsEnabled() *[]*string
	SetProtocolsEnabled(val *[]*string)
	ProtocolsEnabledInput() *[]*string
	RootAccessEnabled() interface{}
	SetRootAccessEnabled(val interface{})
	RootAccessEnabledInput() interface{}
	RuleIndex() *float64
	SetRuleIndex(val *float64)
	RuleIndexInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UnixReadOnly() interface{}
	SetUnixReadOnly(val interface{})
	UnixReadOnlyInput() interface{}
	UnixReadWrite() interface{}
	SetUnixReadWrite(val interface{})
	UnixReadWriteInput() interface{}
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
	ResetProtocolsEnabled()
	ResetRootAccessEnabled()
	ResetUnixReadOnly()
	ResetUnixReadWrite()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for NetappVolumeExportPolicyRuleOutputReference
type jsiiProxy_NetappVolumeExportPolicyRuleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_NetappVolumeExportPolicyRuleOutputReference) AllowedClients() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedClients",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeExportPolicyRuleOutputReference) AllowedClientsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedClientsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeExportPolicyRuleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeExportPolicyRuleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeExportPolicyRuleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeExportPolicyRuleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeExportPolicyRuleOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeExportPolicyRuleOutputReference) ProtocolsEnabled() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"protocolsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeExportPolicyRuleOutputReference) ProtocolsEnabledInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"protocolsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeExportPolicyRuleOutputReference) RootAccessEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rootAccessEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeExportPolicyRuleOutputReference) RootAccessEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rootAccessEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeExportPolicyRuleOutputReference) RuleIndex() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ruleIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeExportPolicyRuleOutputReference) RuleIndexInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ruleIndexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeExportPolicyRuleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeExportPolicyRuleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeExportPolicyRuleOutputReference) UnixReadOnly() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"unixReadOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeExportPolicyRuleOutputReference) UnixReadOnlyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"unixReadOnlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeExportPolicyRuleOutputReference) UnixReadWrite() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"unixReadWrite",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeExportPolicyRuleOutputReference) UnixReadWriteInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"unixReadWriteInput",
		&returns,
	)
	return returns
}


func NewNetappVolumeExportPolicyRuleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) NetappVolumeExportPolicyRuleOutputReference {
	_init_.Initialize()

	j := jsiiProxy_NetappVolumeExportPolicyRuleOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.netappVolume.NetappVolumeExportPolicyRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewNetappVolumeExportPolicyRuleOutputReference_Override(n NetappVolumeExportPolicyRuleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.netappVolume.NetappVolumeExportPolicyRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		n,
	)
}

func (j *jsiiProxy_NetappVolumeExportPolicyRuleOutputReference) SetAllowedClients(val *[]*string) {
	_jsii_.Set(
		j,
		"allowedClients",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeExportPolicyRuleOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeExportPolicyRuleOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeExportPolicyRuleOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeExportPolicyRuleOutputReference) SetProtocolsEnabled(val *[]*string) {
	_jsii_.Set(
		j,
		"protocolsEnabled",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeExportPolicyRuleOutputReference) SetRootAccessEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"rootAccessEnabled",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeExportPolicyRuleOutputReference) SetRuleIndex(val *float64) {
	_jsii_.Set(
		j,
		"ruleIndex",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeExportPolicyRuleOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeExportPolicyRuleOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeExportPolicyRuleOutputReference) SetUnixReadOnly(val interface{}) {
	_jsii_.Set(
		j,
		"unixReadOnly",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeExportPolicyRuleOutputReference) SetUnixReadWrite(val interface{}) {
	_jsii_.Set(
		j,
		"unixReadWrite",
		val,
	)
}

func (n *jsiiProxy_NetappVolumeExportPolicyRuleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeExportPolicyRuleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeExportPolicyRuleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeExportPolicyRuleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		n,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeExportPolicyRuleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		n,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeExportPolicyRuleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		n,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeExportPolicyRuleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		n,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeExportPolicyRuleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		n,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeExportPolicyRuleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeExportPolicyRuleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		n,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeExportPolicyRuleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeExportPolicyRuleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeExportPolicyRuleOutputReference) ResetProtocolsEnabled() {
	_jsii_.InvokeVoid(
		n,
		"resetProtocolsEnabled",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolumeExportPolicyRuleOutputReference) ResetRootAccessEnabled() {
	_jsii_.InvokeVoid(
		n,
		"resetRootAccessEnabled",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolumeExportPolicyRuleOutputReference) ResetUnixReadOnly() {
	_jsii_.InvokeVoid(
		n,
		"resetUnixReadOnly",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolumeExportPolicyRuleOutputReference) ResetUnixReadWrite() {
	_jsii_.InvokeVoid(
		n,
		"resetUnixReadWrite",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolumeExportPolicyRuleOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeExportPolicyRuleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type NetappVolumeTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/netapp_volume#create NetappVolume#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/netapp_volume#delete NetappVolume#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/netapp_volume#read NetappVolume#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/netapp_volume#update NetappVolume#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

type NetappVolumeTimeoutsOutputReference interface {
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

// The jsii proxy struct for NetappVolumeTimeoutsOutputReference
type jsiiProxy_NetappVolumeTimeoutsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_NetappVolumeTimeoutsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeTimeoutsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeTimeoutsOutputReference) Create() *string {
	var returns *string
	_jsii_.Get(
		j,
		"create",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeTimeoutsOutputReference) CreateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeTimeoutsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeTimeoutsOutputReference) Delete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeTimeoutsOutputReference) DeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeTimeoutsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeTimeoutsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeTimeoutsOutputReference) Read() *string {
	var returns *string
	_jsii_.Get(
		j,
		"read",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeTimeoutsOutputReference) ReadInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeTimeoutsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeTimeoutsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeTimeoutsOutputReference) Update() *string {
	var returns *string
	_jsii_.Get(
		j,
		"update",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeTimeoutsOutputReference) UpdateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateInput",
		&returns,
	)
	return returns
}


func NewNetappVolumeTimeoutsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) NetappVolumeTimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_NetappVolumeTimeoutsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.netappVolume.NetappVolumeTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewNetappVolumeTimeoutsOutputReference_Override(n NetappVolumeTimeoutsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.netappVolume.NetappVolumeTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		n,
	)
}

func (j *jsiiProxy_NetappVolumeTimeoutsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeTimeoutsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeTimeoutsOutputReference) SetCreate(val *string) {
	_jsii_.Set(
		j,
		"create",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeTimeoutsOutputReference) SetDelete(val *string) {
	_jsii_.Set(
		j,
		"delete",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeTimeoutsOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeTimeoutsOutputReference) SetRead(val *string) {
	_jsii_.Set(
		j,
		"read",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeTimeoutsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeTimeoutsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeTimeoutsOutputReference) SetUpdate(val *string) {
	_jsii_.Set(
		j,
		"update",
		val,
	)
}

func (n *jsiiProxy_NetappVolumeTimeoutsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeTimeoutsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeTimeoutsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		n,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeTimeoutsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		n,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeTimeoutsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		n,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeTimeoutsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		n,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeTimeoutsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		n,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeTimeoutsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeTimeoutsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		n,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeTimeoutsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeTimeoutsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeTimeoutsOutputReference) ResetCreate() {
	_jsii_.InvokeVoid(
		n,
		"resetCreate",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolumeTimeoutsOutputReference) ResetDelete() {
	_jsii_.InvokeVoid(
		n,
		"resetDelete",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolumeTimeoutsOutputReference) ResetRead() {
	_jsii_.InvokeVoid(
		n,
		"resetRead",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolumeTimeoutsOutputReference) ResetUpdate() {
	_jsii_.InvokeVoid(
		n,
		"resetUpdate",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolumeTimeoutsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeTimeoutsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

