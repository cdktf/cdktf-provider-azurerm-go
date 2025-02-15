// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package netappvolumegrouporacle

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/netappvolumegrouporacle/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetappVolumeGroupOracleVolumeOutputReference interface {
	cdktf.ComplexObject
	CapacityPoolId() *string
	SetCapacityPoolId(val *string)
	CapacityPoolIdInput() *string
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
	DataProtectionSnapshotPolicy() NetappVolumeGroupOracleVolumeDataProtectionSnapshotPolicyOutputReference
	DataProtectionSnapshotPolicyInput() *NetappVolumeGroupOracleVolumeDataProtectionSnapshotPolicy
	EncryptionKeySource() *string
	SetEncryptionKeySource(val *string)
	EncryptionKeySourceInput() *string
	ExportPolicyRule() NetappVolumeGroupOracleVolumeExportPolicyRuleList
	ExportPolicyRuleInput() interface{}
	// Experimental.
	Fqn() *string
	Id() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	KeyVaultPrivateEndpointId() *string
	SetKeyVaultPrivateEndpointId(val *string)
	KeyVaultPrivateEndpointIdInput() *string
	MountIpAddresses() *[]*string
	Name() *string
	SetName(val *string)
	NameInput() *string
	NetworkFeatures() *string
	SetNetworkFeatures(val *string)
	NetworkFeaturesInput() *string
	Protocols() *[]*string
	SetProtocols(val *[]*string)
	ProtocolsInput() *[]*string
	ProximityPlacementGroupId() *string
	SetProximityPlacementGroupId(val *string)
	ProximityPlacementGroupIdInput() *string
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
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	ThroughputInMibps() *float64
	SetThroughputInMibps(val *float64)
	ThroughputInMibpsInput() *float64
	VolumePath() *string
	SetVolumePath(val *string)
	VolumePathInput() *string
	VolumeSpecName() *string
	SetVolumeSpecName(val *string)
	VolumeSpecNameInput() *string
	Zone() *string
	SetZone(val *string)
	ZoneInput() *string
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
	PutDataProtectionSnapshotPolicy(value *NetappVolumeGroupOracleVolumeDataProtectionSnapshotPolicy)
	PutExportPolicyRule(value interface{})
	ResetDataProtectionSnapshotPolicy()
	ResetEncryptionKeySource()
	ResetKeyVaultPrivateEndpointId()
	ResetNetworkFeatures()
	ResetProximityPlacementGroupId()
	ResetTags()
	ResetZone()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for NetappVolumeGroupOracleVolumeOutputReference
type jsiiProxy_NetappVolumeGroupOracleVolumeOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_NetappVolumeGroupOracleVolumeOutputReference) CapacityPoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"capacityPoolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeGroupOracleVolumeOutputReference) CapacityPoolIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"capacityPoolIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeGroupOracleVolumeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeGroupOracleVolumeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeGroupOracleVolumeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeGroupOracleVolumeOutputReference) DataProtectionSnapshotPolicy() NetappVolumeGroupOracleVolumeDataProtectionSnapshotPolicyOutputReference {
	var returns NetappVolumeGroupOracleVolumeDataProtectionSnapshotPolicyOutputReference
	_jsii_.Get(
		j,
		"dataProtectionSnapshotPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeGroupOracleVolumeOutputReference) DataProtectionSnapshotPolicyInput() *NetappVolumeGroupOracleVolumeDataProtectionSnapshotPolicy {
	var returns *NetappVolumeGroupOracleVolumeDataProtectionSnapshotPolicy
	_jsii_.Get(
		j,
		"dataProtectionSnapshotPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeGroupOracleVolumeOutputReference) EncryptionKeySource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionKeySource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeGroupOracleVolumeOutputReference) EncryptionKeySourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionKeySourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeGroupOracleVolumeOutputReference) ExportPolicyRule() NetappVolumeGroupOracleVolumeExportPolicyRuleList {
	var returns NetappVolumeGroupOracleVolumeExportPolicyRuleList
	_jsii_.Get(
		j,
		"exportPolicyRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeGroupOracleVolumeOutputReference) ExportPolicyRuleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"exportPolicyRuleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeGroupOracleVolumeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeGroupOracleVolumeOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeGroupOracleVolumeOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeGroupOracleVolumeOutputReference) KeyVaultPrivateEndpointId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyVaultPrivateEndpointId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeGroupOracleVolumeOutputReference) KeyVaultPrivateEndpointIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyVaultPrivateEndpointIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeGroupOracleVolumeOutputReference) MountIpAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"mountIpAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeGroupOracleVolumeOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeGroupOracleVolumeOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeGroupOracleVolumeOutputReference) NetworkFeatures() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkFeatures",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeGroupOracleVolumeOutputReference) NetworkFeaturesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkFeaturesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeGroupOracleVolumeOutputReference) Protocols() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"protocols",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeGroupOracleVolumeOutputReference) ProtocolsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"protocolsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeGroupOracleVolumeOutputReference) ProximityPlacementGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"proximityPlacementGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeGroupOracleVolumeOutputReference) ProximityPlacementGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"proximityPlacementGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeGroupOracleVolumeOutputReference) SecurityStyle() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityStyle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeGroupOracleVolumeOutputReference) SecurityStyleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityStyleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeGroupOracleVolumeOutputReference) ServiceLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeGroupOracleVolumeOutputReference) ServiceLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeGroupOracleVolumeOutputReference) SnapshotDirectoryVisible() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"snapshotDirectoryVisible",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeGroupOracleVolumeOutputReference) SnapshotDirectoryVisibleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"snapshotDirectoryVisibleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeGroupOracleVolumeOutputReference) StorageQuotaInGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"storageQuotaInGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeGroupOracleVolumeOutputReference) StorageQuotaInGbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"storageQuotaInGbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeGroupOracleVolumeOutputReference) SubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeGroupOracleVolumeOutputReference) SubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeGroupOracleVolumeOutputReference) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeGroupOracleVolumeOutputReference) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeGroupOracleVolumeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeGroupOracleVolumeOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeGroupOracleVolumeOutputReference) ThroughputInMibps() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"throughputInMibps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeGroupOracleVolumeOutputReference) ThroughputInMibpsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"throughputInMibpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeGroupOracleVolumeOutputReference) VolumePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeGroupOracleVolumeOutputReference) VolumePathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumePathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeGroupOracleVolumeOutputReference) VolumeSpecName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumeSpecName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeGroupOracleVolumeOutputReference) VolumeSpecNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumeSpecNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeGroupOracleVolumeOutputReference) Zone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeGroupOracleVolumeOutputReference) ZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zoneInput",
		&returns,
	)
	return returns
}


func NewNetappVolumeGroupOracleVolumeOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) NetappVolumeGroupOracleVolumeOutputReference {
	_init_.Initialize()

	if err := validateNewNetappVolumeGroupOracleVolumeOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_NetappVolumeGroupOracleVolumeOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.netappVolumeGroupOracle.NetappVolumeGroupOracleVolumeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewNetappVolumeGroupOracleVolumeOutputReference_Override(n NetappVolumeGroupOracleVolumeOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.netappVolumeGroupOracle.NetappVolumeGroupOracleVolumeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		n,
	)
}

func (j *jsiiProxy_NetappVolumeGroupOracleVolumeOutputReference)SetCapacityPoolId(val *string) {
	if err := j.validateSetCapacityPoolIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"capacityPoolId",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeGroupOracleVolumeOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeGroupOracleVolumeOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeGroupOracleVolumeOutputReference)SetEncryptionKeySource(val *string) {
	if err := j.validateSetEncryptionKeySourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"encryptionKeySource",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeGroupOracleVolumeOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeGroupOracleVolumeOutputReference)SetKeyVaultPrivateEndpointId(val *string) {
	if err := j.validateSetKeyVaultPrivateEndpointIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyVaultPrivateEndpointId",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeGroupOracleVolumeOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeGroupOracleVolumeOutputReference)SetNetworkFeatures(val *string) {
	if err := j.validateSetNetworkFeaturesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkFeatures",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeGroupOracleVolumeOutputReference)SetProtocols(val *[]*string) {
	if err := j.validateSetProtocolsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protocols",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeGroupOracleVolumeOutputReference)SetProximityPlacementGroupId(val *string) {
	if err := j.validateSetProximityPlacementGroupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"proximityPlacementGroupId",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeGroupOracleVolumeOutputReference)SetSecurityStyle(val *string) {
	if err := j.validateSetSecurityStyleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityStyle",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeGroupOracleVolumeOutputReference)SetServiceLevel(val *string) {
	if err := j.validateSetServiceLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceLevel",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeGroupOracleVolumeOutputReference)SetSnapshotDirectoryVisible(val interface{}) {
	if err := j.validateSetSnapshotDirectoryVisibleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"snapshotDirectoryVisible",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeGroupOracleVolumeOutputReference)SetStorageQuotaInGb(val *float64) {
	if err := j.validateSetStorageQuotaInGbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageQuotaInGb",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeGroupOracleVolumeOutputReference)SetSubnetId(val *string) {
	if err := j.validateSetSubnetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetId",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeGroupOracleVolumeOutputReference)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeGroupOracleVolumeOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeGroupOracleVolumeOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeGroupOracleVolumeOutputReference)SetThroughputInMibps(val *float64) {
	if err := j.validateSetThroughputInMibpsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"throughputInMibps",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeGroupOracleVolumeOutputReference)SetVolumePath(val *string) {
	if err := j.validateSetVolumePathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"volumePath",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeGroupOracleVolumeOutputReference)SetVolumeSpecName(val *string) {
	if err := j.validateSetVolumeSpecNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"volumeSpecName",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeGroupOracleVolumeOutputReference)SetZone(val *string) {
	if err := j.validateSetZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zone",
		val,
	)
}

func (n *jsiiProxy_NetappVolumeGroupOracleVolumeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeGroupOracleVolumeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (n *jsiiProxy_NetappVolumeGroupOracleVolumeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (n *jsiiProxy_NetappVolumeGroupOracleVolumeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (n *jsiiProxy_NetappVolumeGroupOracleVolumeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (n *jsiiProxy_NetappVolumeGroupOracleVolumeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (n *jsiiProxy_NetappVolumeGroupOracleVolumeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (n *jsiiProxy_NetappVolumeGroupOracleVolumeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (n *jsiiProxy_NetappVolumeGroupOracleVolumeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (n *jsiiProxy_NetappVolumeGroupOracleVolumeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (n *jsiiProxy_NetappVolumeGroupOracleVolumeOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeGroupOracleVolumeOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := n.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeGroupOracleVolumeOutputReference) PutDataProtectionSnapshotPolicy(value *NetappVolumeGroupOracleVolumeDataProtectionSnapshotPolicy) {
	if err := n.validatePutDataProtectionSnapshotPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putDataProtectionSnapshotPolicy",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetappVolumeGroupOracleVolumeOutputReference) PutExportPolicyRule(value interface{}) {
	if err := n.validatePutExportPolicyRuleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putExportPolicyRule",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetappVolumeGroupOracleVolumeOutputReference) ResetDataProtectionSnapshotPolicy() {
	_jsii_.InvokeVoid(
		n,
		"resetDataProtectionSnapshotPolicy",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolumeGroupOracleVolumeOutputReference) ResetEncryptionKeySource() {
	_jsii_.InvokeVoid(
		n,
		"resetEncryptionKeySource",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolumeGroupOracleVolumeOutputReference) ResetKeyVaultPrivateEndpointId() {
	_jsii_.InvokeVoid(
		n,
		"resetKeyVaultPrivateEndpointId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolumeGroupOracleVolumeOutputReference) ResetNetworkFeatures() {
	_jsii_.InvokeVoid(
		n,
		"resetNetworkFeatures",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolumeGroupOracleVolumeOutputReference) ResetProximityPlacementGroupId() {
	_jsii_.InvokeVoid(
		n,
		"resetProximityPlacementGroupId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolumeGroupOracleVolumeOutputReference) ResetTags() {
	_jsii_.InvokeVoid(
		n,
		"resetTags",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolumeGroupOracleVolumeOutputReference) ResetZone() {
	_jsii_.InvokeVoid(
		n,
		"resetZone",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolumeGroupOracleVolumeOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := n.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		n,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeGroupOracleVolumeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

