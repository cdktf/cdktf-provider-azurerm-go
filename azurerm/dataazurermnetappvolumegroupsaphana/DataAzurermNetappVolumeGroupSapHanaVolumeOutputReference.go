// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataazurermnetappvolumegroupsaphana

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v10/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v10/dataazurermnetappvolumegroupsaphana/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataAzurermNetappVolumeGroupSapHanaVolumeOutputReference interface {
	cdktf.ComplexObject
	CapacityPoolId() *string
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
	DataProtectionReplication() DataAzurermNetappVolumeGroupSapHanaVolumeDataProtectionReplicationList
	DataProtectionSnapshotPolicy() DataAzurermNetappVolumeGroupSapHanaVolumeDataProtectionSnapshotPolicyList
	ExportPolicyRule() DataAzurermNetappVolumeGroupSapHanaVolumeExportPolicyRuleList
	// Experimental.
	Fqn() *string
	Id() *string
	InternalValue() *DataAzurermNetappVolumeGroupSapHanaVolume
	SetInternalValue(val *DataAzurermNetappVolumeGroupSapHanaVolume)
	MountIpAddresses() *[]*string
	Name() *string
	Protocols() *[]*string
	ProximityPlacementGroupId() *string
	SecurityStyle() *string
	ServiceLevel() *string
	SnapshotDirectoryVisible() cdktf.IResolvable
	StorageQuotaInGb() *float64
	SubnetId() *string
	Tags() cdktf.StringMap
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	ThroughputInMibps() *float64
	VolumePath() *string
	VolumeSpecName() *string
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

// The jsii proxy struct for DataAzurermNetappVolumeGroupSapHanaVolumeOutputReference
type jsiiProxy_DataAzurermNetappVolumeGroupSapHanaVolumeOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataAzurermNetappVolumeGroupSapHanaVolumeOutputReference) CapacityPoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"capacityPoolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermNetappVolumeGroupSapHanaVolumeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermNetappVolumeGroupSapHanaVolumeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermNetappVolumeGroupSapHanaVolumeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermNetappVolumeGroupSapHanaVolumeOutputReference) DataProtectionReplication() DataAzurermNetappVolumeGroupSapHanaVolumeDataProtectionReplicationList {
	var returns DataAzurermNetappVolumeGroupSapHanaVolumeDataProtectionReplicationList
	_jsii_.Get(
		j,
		"dataProtectionReplication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermNetappVolumeGroupSapHanaVolumeOutputReference) DataProtectionSnapshotPolicy() DataAzurermNetappVolumeGroupSapHanaVolumeDataProtectionSnapshotPolicyList {
	var returns DataAzurermNetappVolumeGroupSapHanaVolumeDataProtectionSnapshotPolicyList
	_jsii_.Get(
		j,
		"dataProtectionSnapshotPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermNetappVolumeGroupSapHanaVolumeOutputReference) ExportPolicyRule() DataAzurermNetappVolumeGroupSapHanaVolumeExportPolicyRuleList {
	var returns DataAzurermNetappVolumeGroupSapHanaVolumeExportPolicyRuleList
	_jsii_.Get(
		j,
		"exportPolicyRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermNetappVolumeGroupSapHanaVolumeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermNetappVolumeGroupSapHanaVolumeOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermNetappVolumeGroupSapHanaVolumeOutputReference) InternalValue() *DataAzurermNetappVolumeGroupSapHanaVolume {
	var returns *DataAzurermNetappVolumeGroupSapHanaVolume
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermNetappVolumeGroupSapHanaVolumeOutputReference) MountIpAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"mountIpAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermNetappVolumeGroupSapHanaVolumeOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermNetappVolumeGroupSapHanaVolumeOutputReference) Protocols() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"protocols",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermNetappVolumeGroupSapHanaVolumeOutputReference) ProximityPlacementGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"proximityPlacementGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermNetappVolumeGroupSapHanaVolumeOutputReference) SecurityStyle() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityStyle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermNetappVolumeGroupSapHanaVolumeOutputReference) ServiceLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermNetappVolumeGroupSapHanaVolumeOutputReference) SnapshotDirectoryVisible() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"snapshotDirectoryVisible",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermNetappVolumeGroupSapHanaVolumeOutputReference) StorageQuotaInGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"storageQuotaInGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermNetappVolumeGroupSapHanaVolumeOutputReference) SubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermNetappVolumeGroupSapHanaVolumeOutputReference) Tags() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermNetappVolumeGroupSapHanaVolumeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermNetappVolumeGroupSapHanaVolumeOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermNetappVolumeGroupSapHanaVolumeOutputReference) ThroughputInMibps() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"throughputInMibps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermNetappVolumeGroupSapHanaVolumeOutputReference) VolumePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermNetappVolumeGroupSapHanaVolumeOutputReference) VolumeSpecName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumeSpecName",
		&returns,
	)
	return returns
}


func NewDataAzurermNetappVolumeGroupSapHanaVolumeOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataAzurermNetappVolumeGroupSapHanaVolumeOutputReference {
	_init_.Initialize()

	if err := validateNewDataAzurermNetappVolumeGroupSapHanaVolumeOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAzurermNetappVolumeGroupSapHanaVolumeOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataAzurermNetappVolumeGroupSapHana.DataAzurermNetappVolumeGroupSapHanaVolumeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataAzurermNetappVolumeGroupSapHanaVolumeOutputReference_Override(d DataAzurermNetappVolumeGroupSapHanaVolumeOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataAzurermNetappVolumeGroupSapHana.DataAzurermNetappVolumeGroupSapHanaVolumeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataAzurermNetappVolumeGroupSapHanaVolumeOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAzurermNetappVolumeGroupSapHanaVolumeOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAzurermNetappVolumeGroupSapHanaVolumeOutputReference)SetInternalValue(val *DataAzurermNetappVolumeGroupSapHanaVolume) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAzurermNetappVolumeGroupSapHanaVolumeOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAzurermNetappVolumeGroupSapHanaVolumeOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAzurermNetappVolumeGroupSapHanaVolumeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermNetappVolumeGroupSapHanaVolumeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermNetappVolumeGroupSapHanaVolumeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermNetappVolumeGroupSapHanaVolumeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermNetappVolumeGroupSapHanaVolumeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermNetappVolumeGroupSapHanaVolumeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermNetappVolumeGroupSapHanaVolumeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermNetappVolumeGroupSapHanaVolumeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermNetappVolumeGroupSapHanaVolumeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermNetappVolumeGroupSapHanaVolumeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermNetappVolumeGroupSapHanaVolumeOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermNetappVolumeGroupSapHanaVolumeOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermNetappVolumeGroupSapHanaVolumeOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermNetappVolumeGroupSapHanaVolumeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

