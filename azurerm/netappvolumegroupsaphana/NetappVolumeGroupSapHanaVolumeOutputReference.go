package netappvolumegroupsaphana

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v9/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v9/netappvolumegroupsaphana/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetappVolumeGroupSapHanaVolumeOutputReference interface {
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
	DataProtectionReplication() NetappVolumeGroupSapHanaVolumeDataProtectionReplicationOutputReference
	DataProtectionReplicationInput() *NetappVolumeGroupSapHanaVolumeDataProtectionReplication
	DataProtectionSnapshotPolicy() NetappVolumeGroupSapHanaVolumeDataProtectionSnapshotPolicyOutputReference
	DataProtectionSnapshotPolicyInput() *NetappVolumeGroupSapHanaVolumeDataProtectionSnapshotPolicy
	ExportPolicyRule() NetappVolumeGroupSapHanaVolumeExportPolicyRuleList
	ExportPolicyRuleInput() interface{}
	// Experimental.
	Fqn() *string
	Id() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MountIpAddresses() *[]*string
	Name() *string
	SetName(val *string)
	NameInput() *string
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
	PutDataProtectionReplication(value *NetappVolumeGroupSapHanaVolumeDataProtectionReplication)
	PutDataProtectionSnapshotPolicy(value *NetappVolumeGroupSapHanaVolumeDataProtectionSnapshotPolicy)
	PutExportPolicyRule(value interface{})
	ResetDataProtectionReplication()
	ResetDataProtectionSnapshotPolicy()
	ResetProximityPlacementGroupId()
	ResetTags()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for NetappVolumeGroupSapHanaVolumeOutputReference
type jsiiProxy_NetappVolumeGroupSapHanaVolumeOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_NetappVolumeGroupSapHanaVolumeOutputReference) CapacityPoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"capacityPoolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeGroupSapHanaVolumeOutputReference) CapacityPoolIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"capacityPoolIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeGroupSapHanaVolumeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeGroupSapHanaVolumeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeGroupSapHanaVolumeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeGroupSapHanaVolumeOutputReference) DataProtectionReplication() NetappVolumeGroupSapHanaVolumeDataProtectionReplicationOutputReference {
	var returns NetappVolumeGroupSapHanaVolumeDataProtectionReplicationOutputReference
	_jsii_.Get(
		j,
		"dataProtectionReplication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeGroupSapHanaVolumeOutputReference) DataProtectionReplicationInput() *NetappVolumeGroupSapHanaVolumeDataProtectionReplication {
	var returns *NetappVolumeGroupSapHanaVolumeDataProtectionReplication
	_jsii_.Get(
		j,
		"dataProtectionReplicationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeGroupSapHanaVolumeOutputReference) DataProtectionSnapshotPolicy() NetappVolumeGroupSapHanaVolumeDataProtectionSnapshotPolicyOutputReference {
	var returns NetappVolumeGroupSapHanaVolumeDataProtectionSnapshotPolicyOutputReference
	_jsii_.Get(
		j,
		"dataProtectionSnapshotPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeGroupSapHanaVolumeOutputReference) DataProtectionSnapshotPolicyInput() *NetappVolumeGroupSapHanaVolumeDataProtectionSnapshotPolicy {
	var returns *NetappVolumeGroupSapHanaVolumeDataProtectionSnapshotPolicy
	_jsii_.Get(
		j,
		"dataProtectionSnapshotPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeGroupSapHanaVolumeOutputReference) ExportPolicyRule() NetappVolumeGroupSapHanaVolumeExportPolicyRuleList {
	var returns NetappVolumeGroupSapHanaVolumeExportPolicyRuleList
	_jsii_.Get(
		j,
		"exportPolicyRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeGroupSapHanaVolumeOutputReference) ExportPolicyRuleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"exportPolicyRuleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeGroupSapHanaVolumeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeGroupSapHanaVolumeOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeGroupSapHanaVolumeOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeGroupSapHanaVolumeOutputReference) MountIpAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"mountIpAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeGroupSapHanaVolumeOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeGroupSapHanaVolumeOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeGroupSapHanaVolumeOutputReference) Protocols() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"protocols",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeGroupSapHanaVolumeOutputReference) ProtocolsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"protocolsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeGroupSapHanaVolumeOutputReference) ProximityPlacementGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"proximityPlacementGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeGroupSapHanaVolumeOutputReference) ProximityPlacementGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"proximityPlacementGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeGroupSapHanaVolumeOutputReference) SecurityStyle() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityStyle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeGroupSapHanaVolumeOutputReference) SecurityStyleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityStyleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeGroupSapHanaVolumeOutputReference) ServiceLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeGroupSapHanaVolumeOutputReference) ServiceLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeGroupSapHanaVolumeOutputReference) SnapshotDirectoryVisible() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"snapshotDirectoryVisible",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeGroupSapHanaVolumeOutputReference) SnapshotDirectoryVisibleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"snapshotDirectoryVisibleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeGroupSapHanaVolumeOutputReference) StorageQuotaInGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"storageQuotaInGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeGroupSapHanaVolumeOutputReference) StorageQuotaInGbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"storageQuotaInGbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeGroupSapHanaVolumeOutputReference) SubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeGroupSapHanaVolumeOutputReference) SubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeGroupSapHanaVolumeOutputReference) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeGroupSapHanaVolumeOutputReference) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeGroupSapHanaVolumeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeGroupSapHanaVolumeOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeGroupSapHanaVolumeOutputReference) ThroughputInMibps() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"throughputInMibps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeGroupSapHanaVolumeOutputReference) ThroughputInMibpsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"throughputInMibpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeGroupSapHanaVolumeOutputReference) VolumePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeGroupSapHanaVolumeOutputReference) VolumePathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumePathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeGroupSapHanaVolumeOutputReference) VolumeSpecName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumeSpecName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeGroupSapHanaVolumeOutputReference) VolumeSpecNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumeSpecNameInput",
		&returns,
	)
	return returns
}


func NewNetappVolumeGroupSapHanaVolumeOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) NetappVolumeGroupSapHanaVolumeOutputReference {
	_init_.Initialize()

	if err := validateNewNetappVolumeGroupSapHanaVolumeOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_NetappVolumeGroupSapHanaVolumeOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.netappVolumeGroupSapHana.NetappVolumeGroupSapHanaVolumeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewNetappVolumeGroupSapHanaVolumeOutputReference_Override(n NetappVolumeGroupSapHanaVolumeOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.netappVolumeGroupSapHana.NetappVolumeGroupSapHanaVolumeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		n,
	)
}

func (j *jsiiProxy_NetappVolumeGroupSapHanaVolumeOutputReference)SetCapacityPoolId(val *string) {
	if err := j.validateSetCapacityPoolIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"capacityPoolId",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeGroupSapHanaVolumeOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeGroupSapHanaVolumeOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeGroupSapHanaVolumeOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeGroupSapHanaVolumeOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeGroupSapHanaVolumeOutputReference)SetProtocols(val *[]*string) {
	if err := j.validateSetProtocolsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protocols",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeGroupSapHanaVolumeOutputReference)SetProximityPlacementGroupId(val *string) {
	if err := j.validateSetProximityPlacementGroupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"proximityPlacementGroupId",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeGroupSapHanaVolumeOutputReference)SetSecurityStyle(val *string) {
	if err := j.validateSetSecurityStyleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityStyle",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeGroupSapHanaVolumeOutputReference)SetServiceLevel(val *string) {
	if err := j.validateSetServiceLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceLevel",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeGroupSapHanaVolumeOutputReference)SetSnapshotDirectoryVisible(val interface{}) {
	if err := j.validateSetSnapshotDirectoryVisibleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"snapshotDirectoryVisible",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeGroupSapHanaVolumeOutputReference)SetStorageQuotaInGb(val *float64) {
	if err := j.validateSetStorageQuotaInGbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageQuotaInGb",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeGroupSapHanaVolumeOutputReference)SetSubnetId(val *string) {
	if err := j.validateSetSubnetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetId",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeGroupSapHanaVolumeOutputReference)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeGroupSapHanaVolumeOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeGroupSapHanaVolumeOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeGroupSapHanaVolumeOutputReference)SetThroughputInMibps(val *float64) {
	if err := j.validateSetThroughputInMibpsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"throughputInMibps",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeGroupSapHanaVolumeOutputReference)SetVolumePath(val *string) {
	if err := j.validateSetVolumePathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"volumePath",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeGroupSapHanaVolumeOutputReference)SetVolumeSpecName(val *string) {
	if err := j.validateSetVolumeSpecNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"volumeSpecName",
		val,
	)
}

func (n *jsiiProxy_NetappVolumeGroupSapHanaVolumeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeGroupSapHanaVolumeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (n *jsiiProxy_NetappVolumeGroupSapHanaVolumeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (n *jsiiProxy_NetappVolumeGroupSapHanaVolumeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (n *jsiiProxy_NetappVolumeGroupSapHanaVolumeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (n *jsiiProxy_NetappVolumeGroupSapHanaVolumeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (n *jsiiProxy_NetappVolumeGroupSapHanaVolumeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (n *jsiiProxy_NetappVolumeGroupSapHanaVolumeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (n *jsiiProxy_NetappVolumeGroupSapHanaVolumeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (n *jsiiProxy_NetappVolumeGroupSapHanaVolumeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (n *jsiiProxy_NetappVolumeGroupSapHanaVolumeOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeGroupSapHanaVolumeOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (n *jsiiProxy_NetappVolumeGroupSapHanaVolumeOutputReference) PutDataProtectionReplication(value *NetappVolumeGroupSapHanaVolumeDataProtectionReplication) {
	if err := n.validatePutDataProtectionReplicationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putDataProtectionReplication",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetappVolumeGroupSapHanaVolumeOutputReference) PutDataProtectionSnapshotPolicy(value *NetappVolumeGroupSapHanaVolumeDataProtectionSnapshotPolicy) {
	if err := n.validatePutDataProtectionSnapshotPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putDataProtectionSnapshotPolicy",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetappVolumeGroupSapHanaVolumeOutputReference) PutExportPolicyRule(value interface{}) {
	if err := n.validatePutExportPolicyRuleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putExportPolicyRule",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetappVolumeGroupSapHanaVolumeOutputReference) ResetDataProtectionReplication() {
	_jsii_.InvokeVoid(
		n,
		"resetDataProtectionReplication",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolumeGroupSapHanaVolumeOutputReference) ResetDataProtectionSnapshotPolicy() {
	_jsii_.InvokeVoid(
		n,
		"resetDataProtectionSnapshotPolicy",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolumeGroupSapHanaVolumeOutputReference) ResetProximityPlacementGroupId() {
	_jsii_.InvokeVoid(
		n,
		"resetProximityPlacementGroupId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolumeGroupSapHanaVolumeOutputReference) ResetTags() {
	_jsii_.InvokeVoid(
		n,
		"resetTags",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolumeGroupSapHanaVolumeOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (n *jsiiProxy_NetappVolumeGroupSapHanaVolumeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

