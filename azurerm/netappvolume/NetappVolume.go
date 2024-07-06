// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package netappvolume

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/netappvolume/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.111.0/docs/resources/netapp_volume azurerm_netapp_volume}.
type NetappVolume interface {
	cdktf.TerraformResource
	AccountName() *string
	SetAccountName(val *string)
	AccountNameInput() *string
	AzureVmwareDataStoreEnabled() interface{}
	SetAzureVmwareDataStoreEnabled(val interface{})
	AzureVmwareDataStoreEnabledInput() interface{}
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
	EncryptionKeySource() *string
	SetEncryptionKeySource(val *string)
	EncryptionKeySourceInput() *string
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
	KerberosEnabled() interface{}
	SetKerberosEnabled(val interface{})
	KerberosEnabledInput() interface{}
	KeyVaultPrivateEndpointId() *string
	SetKeyVaultPrivateEndpointId(val *string)
	KeyVaultPrivateEndpointIdInput() *string
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
	SmbAccessBasedEnumerationEnabled() interface{}
	SetSmbAccessBasedEnumerationEnabled(val interface{})
	SmbAccessBasedEnumerationEnabledInput() interface{}
	SmbContinuousAvailabilityEnabled() interface{}
	SetSmbContinuousAvailabilityEnabled(val interface{})
	SmbContinuousAvailabilityEnabledInput() interface{}
	SmbNonBrowsableEnabled() interface{}
	SetSmbNonBrowsableEnabled(val interface{})
	SmbNonBrowsableEnabledInput() interface{}
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
	Zone() *string
	SetZone(val *string)
	ZoneInput() *string
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
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
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutDataProtectionReplication(value *NetappVolumeDataProtectionReplication)
	PutDataProtectionSnapshotPolicy(value *NetappVolumeDataProtectionSnapshotPolicy)
	PutExportPolicyRule(value interface{})
	PutTimeouts(value *NetappVolumeTimeouts)
	ResetAzureVmwareDataStoreEnabled()
	ResetCreateFromSnapshotResourceId()
	ResetDataProtectionReplication()
	ResetDataProtectionSnapshotPolicy()
	ResetEncryptionKeySource()
	ResetExportPolicyRule()
	ResetId()
	ResetKerberosEnabled()
	ResetKeyVaultPrivateEndpointId()
	ResetNetworkFeatures()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProtocols()
	ResetSecurityStyle()
	ResetSmbAccessBasedEnumerationEnabled()
	ResetSmbContinuousAvailabilityEnabled()
	ResetSmbNonBrowsableEnabled()
	ResetSnapshotDirectoryVisible()
	ResetTags()
	ResetThroughputInMibps()
	ResetTimeouts()
	ResetZone()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
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

func (j *jsiiProxy_NetappVolume) AzureVmwareDataStoreEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"azureVmwareDataStoreEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) AzureVmwareDataStoreEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"azureVmwareDataStoreEnabledInput",
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

func (j *jsiiProxy_NetappVolume) Count() interface{} {
	var returns interface{}
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

func (j *jsiiProxy_NetappVolume) EncryptionKeySource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionKeySource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) EncryptionKeySourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionKeySourceInput",
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

func (j *jsiiProxy_NetappVolume) KerberosEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kerberosEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) KerberosEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kerberosEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) KeyVaultPrivateEndpointId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyVaultPrivateEndpointId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) KeyVaultPrivateEndpointIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyVaultPrivateEndpointIdInput",
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

func (j *jsiiProxy_NetappVolume) SmbAccessBasedEnumerationEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"smbAccessBasedEnumerationEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) SmbAccessBasedEnumerationEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"smbAccessBasedEnumerationEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) SmbContinuousAvailabilityEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"smbContinuousAvailabilityEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) SmbContinuousAvailabilityEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"smbContinuousAvailabilityEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) SmbNonBrowsableEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"smbNonBrowsableEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) SmbNonBrowsableEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"smbNonBrowsableEnabledInput",
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

func (j *jsiiProxy_NetappVolume) Zone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolume) ZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zoneInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.111.0/docs/resources/netapp_volume azurerm_netapp_volume} Resource.
func NewNetappVolume(scope constructs.Construct, id *string, config *NetappVolumeConfig) NetappVolume {
	_init_.Initialize()

	if err := validateNewNetappVolumeParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_NetappVolume{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.netappVolume.NetappVolume",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.111.0/docs/resources/netapp_volume azurerm_netapp_volume} Resource.
func NewNetappVolume_Override(n NetappVolume, scope constructs.Construct, id *string, config *NetappVolumeConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.netappVolume.NetappVolume",
		[]interface{}{scope, id, config},
		n,
	)
}

func (j *jsiiProxy_NetappVolume)SetAccountName(val *string) {
	if err := j.validateSetAccountNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountName",
		val,
	)
}

func (j *jsiiProxy_NetappVolume)SetAzureVmwareDataStoreEnabled(val interface{}) {
	if err := j.validateSetAzureVmwareDataStoreEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"azureVmwareDataStoreEnabled",
		val,
	)
}

func (j *jsiiProxy_NetappVolume)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_NetappVolume)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_NetappVolume)SetCreateFromSnapshotResourceId(val *string) {
	if err := j.validateSetCreateFromSnapshotResourceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createFromSnapshotResourceId",
		val,
	)
}

func (j *jsiiProxy_NetappVolume)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_NetappVolume)SetEncryptionKeySource(val *string) {
	if err := j.validateSetEncryptionKeySourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"encryptionKeySource",
		val,
	)
}

func (j *jsiiProxy_NetappVolume)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_NetappVolume)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_NetappVolume)SetKerberosEnabled(val interface{}) {
	if err := j.validateSetKerberosEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kerberosEnabled",
		val,
	)
}

func (j *jsiiProxy_NetappVolume)SetKeyVaultPrivateEndpointId(val *string) {
	if err := j.validateSetKeyVaultPrivateEndpointIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyVaultPrivateEndpointId",
		val,
	)
}

func (j *jsiiProxy_NetappVolume)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_NetappVolume)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_NetappVolume)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_NetappVolume)SetNetworkFeatures(val *string) {
	if err := j.validateSetNetworkFeaturesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkFeatures",
		val,
	)
}

func (j *jsiiProxy_NetappVolume)SetPoolName(val *string) {
	if err := j.validateSetPoolNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"poolName",
		val,
	)
}

func (j *jsiiProxy_NetappVolume)SetProtocols(val *[]*string) {
	if err := j.validateSetProtocolsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protocols",
		val,
	)
}

func (j *jsiiProxy_NetappVolume)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_NetappVolume)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_NetappVolume)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_NetappVolume)SetSecurityStyle(val *string) {
	if err := j.validateSetSecurityStyleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityStyle",
		val,
	)
}

func (j *jsiiProxy_NetappVolume)SetServiceLevel(val *string) {
	if err := j.validateSetServiceLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceLevel",
		val,
	)
}

func (j *jsiiProxy_NetappVolume)SetSmbAccessBasedEnumerationEnabled(val interface{}) {
	if err := j.validateSetSmbAccessBasedEnumerationEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"smbAccessBasedEnumerationEnabled",
		val,
	)
}

func (j *jsiiProxy_NetappVolume)SetSmbContinuousAvailabilityEnabled(val interface{}) {
	if err := j.validateSetSmbContinuousAvailabilityEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"smbContinuousAvailabilityEnabled",
		val,
	)
}

func (j *jsiiProxy_NetappVolume)SetSmbNonBrowsableEnabled(val interface{}) {
	if err := j.validateSetSmbNonBrowsableEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"smbNonBrowsableEnabled",
		val,
	)
}

func (j *jsiiProxy_NetappVolume)SetSnapshotDirectoryVisible(val interface{}) {
	if err := j.validateSetSnapshotDirectoryVisibleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"snapshotDirectoryVisible",
		val,
	)
}

func (j *jsiiProxy_NetappVolume)SetStorageQuotaInGb(val *float64) {
	if err := j.validateSetStorageQuotaInGbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageQuotaInGb",
		val,
	)
}

func (j *jsiiProxy_NetappVolume)SetSubnetId(val *string) {
	if err := j.validateSetSubnetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetId",
		val,
	)
}

func (j *jsiiProxy_NetappVolume)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_NetappVolume)SetThroughputInMibps(val *float64) {
	if err := j.validateSetThroughputInMibpsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"throughputInMibps",
		val,
	)
}

func (j *jsiiProxy_NetappVolume)SetVolumePath(val *string) {
	if err := j.validateSetVolumePathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"volumePath",
		val,
	)
}

func (j *jsiiProxy_NetappVolume)SetZone(val *string) {
	if err := j.validateSetZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zone",
		val,
	)
}

// Generates CDKTF code for importing a NetappVolume resource upon running "cdktf plan <stack-name>".
func NetappVolume_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateNetappVolume_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.netappVolume.NetappVolume",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
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

	if err := validateNetappVolume_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.netappVolume.NetappVolume",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func NetappVolume_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNetappVolume_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.netappVolume.NetappVolume",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func NetappVolume_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNetappVolume_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.netappVolume.NetappVolume",
		"isTerraformResource",
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

func (n *jsiiProxy_NetappVolume) AddMoveTarget(moveTarget *string) {
	if err := n.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (n *jsiiProxy_NetappVolume) AddOverride(path *string, value interface{}) {
	if err := n.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (n *jsiiProxy_NetappVolume) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := n.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
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
	if err := n.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
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
	if err := n.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
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
	if err := n.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
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
	if err := n.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
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
	if err := n.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
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
	if err := n.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
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
	if err := n.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
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
	if err := n.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		n,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolume) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolume) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := n.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (n *jsiiProxy_NetappVolume) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := n.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolume) MoveFromId(id *string) {
	if err := n.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"moveFromId",
		[]interface{}{id},
	)
}

func (n *jsiiProxy_NetappVolume) MoveTo(moveTarget *string, index interface{}) {
	if err := n.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (n *jsiiProxy_NetappVolume) MoveToId(id *string) {
	if err := n.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"moveToId",
		[]interface{}{id},
	)
}

func (n *jsiiProxy_NetappVolume) OverrideLogicalId(newLogicalId *string) {
	if err := n.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (n *jsiiProxy_NetappVolume) PutDataProtectionReplication(value *NetappVolumeDataProtectionReplication) {
	if err := n.validatePutDataProtectionReplicationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putDataProtectionReplication",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetappVolume) PutDataProtectionSnapshotPolicy(value *NetappVolumeDataProtectionSnapshotPolicy) {
	if err := n.validatePutDataProtectionSnapshotPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putDataProtectionSnapshotPolicy",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetappVolume) PutExportPolicyRule(value interface{}) {
	if err := n.validatePutExportPolicyRuleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putExportPolicyRule",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetappVolume) PutTimeouts(value *NetappVolumeTimeouts) {
	if err := n.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetappVolume) ResetAzureVmwareDataStoreEnabled() {
	_jsii_.InvokeVoid(
		n,
		"resetAzureVmwareDataStoreEnabled",
		nil, // no parameters
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

func (n *jsiiProxy_NetappVolume) ResetEncryptionKeySource() {
	_jsii_.InvokeVoid(
		n,
		"resetEncryptionKeySource",
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

func (n *jsiiProxy_NetappVolume) ResetKerberosEnabled() {
	_jsii_.InvokeVoid(
		n,
		"resetKerberosEnabled",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolume) ResetKeyVaultPrivateEndpointId() {
	_jsii_.InvokeVoid(
		n,
		"resetKeyVaultPrivateEndpointId",
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

func (n *jsiiProxy_NetappVolume) ResetSmbAccessBasedEnumerationEnabled() {
	_jsii_.InvokeVoid(
		n,
		"resetSmbAccessBasedEnumerationEnabled",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolume) ResetSmbContinuousAvailabilityEnabled() {
	_jsii_.InvokeVoid(
		n,
		"resetSmbContinuousAvailabilityEnabled",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolume) ResetSmbNonBrowsableEnabled() {
	_jsii_.InvokeVoid(
		n,
		"resetSmbNonBrowsableEnabled",
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

func (n *jsiiProxy_NetappVolume) ResetZone() {
	_jsii_.InvokeVoid(
		n,
		"resetZone",
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

func (n *jsiiProxy_NetappVolume) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolume) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toHclTerraform",
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

