// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataazurermrediscache

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/dataazurermrediscache/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataAzurermRedisCacheRedisConfigurationOutputReference interface {
	cdktf.ComplexObject
	ActiveDirectoryAuthenticationEnabled() cdktf.IResolvable
	AofBackupEnabled() cdktf.IResolvable
	AofStorageConnectionString0() *string
	AofStorageConnectionString1() *string
	AuthenticationEnabled() cdktf.IResolvable
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
	DataPersistenceAuthenticationMethod() *string
	EnableAuthentication() cdktf.IResolvable
	// Experimental.
	Fqn() *string
	InternalValue() *DataAzurermRedisCacheRedisConfiguration
	SetInternalValue(val *DataAzurermRedisCacheRedisConfiguration)
	Maxclients() *float64
	MaxfragmentationmemoryReserved() *float64
	MaxmemoryDelta() *float64
	MaxmemoryPolicy() *string
	MaxmemoryReserved() *float64
	NotifyKeyspaceEvents() *string
	RdbBackupEnabled() cdktf.IResolvable
	RdbBackupFrequency() *float64
	RdbBackupMaxSnapshotCount() *float64
	RdbStorageConnectionString() *string
	StorageAccountSubscriptionId() *string
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

// The jsii proxy struct for DataAzurermRedisCacheRedisConfigurationOutputReference
type jsiiProxy_DataAzurermRedisCacheRedisConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataAzurermRedisCacheRedisConfigurationOutputReference) ActiveDirectoryAuthenticationEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"activeDirectoryAuthenticationEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermRedisCacheRedisConfigurationOutputReference) AofBackupEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"aofBackupEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermRedisCacheRedisConfigurationOutputReference) AofStorageConnectionString0() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aofStorageConnectionString0",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermRedisCacheRedisConfigurationOutputReference) AofStorageConnectionString1() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aofStorageConnectionString1",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermRedisCacheRedisConfigurationOutputReference) AuthenticationEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"authenticationEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermRedisCacheRedisConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermRedisCacheRedisConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermRedisCacheRedisConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermRedisCacheRedisConfigurationOutputReference) DataPersistenceAuthenticationMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataPersistenceAuthenticationMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermRedisCacheRedisConfigurationOutputReference) EnableAuthentication() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"enableAuthentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermRedisCacheRedisConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermRedisCacheRedisConfigurationOutputReference) InternalValue() *DataAzurermRedisCacheRedisConfiguration {
	var returns *DataAzurermRedisCacheRedisConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermRedisCacheRedisConfigurationOutputReference) Maxclients() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxclients",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermRedisCacheRedisConfigurationOutputReference) MaxfragmentationmemoryReserved() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxfragmentationmemoryReserved",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermRedisCacheRedisConfigurationOutputReference) MaxmemoryDelta() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxmemoryDelta",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermRedisCacheRedisConfigurationOutputReference) MaxmemoryPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxmemoryPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermRedisCacheRedisConfigurationOutputReference) MaxmemoryReserved() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxmemoryReserved",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermRedisCacheRedisConfigurationOutputReference) NotifyKeyspaceEvents() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notifyKeyspaceEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermRedisCacheRedisConfigurationOutputReference) RdbBackupEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"rdbBackupEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermRedisCacheRedisConfigurationOutputReference) RdbBackupFrequency() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rdbBackupFrequency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermRedisCacheRedisConfigurationOutputReference) RdbBackupMaxSnapshotCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rdbBackupMaxSnapshotCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermRedisCacheRedisConfigurationOutputReference) RdbStorageConnectionString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rdbStorageConnectionString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermRedisCacheRedisConfigurationOutputReference) StorageAccountSubscriptionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccountSubscriptionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermRedisCacheRedisConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermRedisCacheRedisConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataAzurermRedisCacheRedisConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataAzurermRedisCacheRedisConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewDataAzurermRedisCacheRedisConfigurationOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAzurermRedisCacheRedisConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataAzurermRedisCache.DataAzurermRedisCacheRedisConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataAzurermRedisCacheRedisConfigurationOutputReference_Override(d DataAzurermRedisCacheRedisConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataAzurermRedisCache.DataAzurermRedisCacheRedisConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataAzurermRedisCacheRedisConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAzurermRedisCacheRedisConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAzurermRedisCacheRedisConfigurationOutputReference)SetInternalValue(val *DataAzurermRedisCacheRedisConfiguration) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAzurermRedisCacheRedisConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAzurermRedisCacheRedisConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAzurermRedisCacheRedisConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermRedisCacheRedisConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAzurermRedisCacheRedisConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAzurermRedisCacheRedisConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAzurermRedisCacheRedisConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAzurermRedisCacheRedisConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAzurermRedisCacheRedisConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAzurermRedisCacheRedisConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAzurermRedisCacheRedisConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAzurermRedisCacheRedisConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAzurermRedisCacheRedisConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermRedisCacheRedisConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAzurermRedisCacheRedisConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAzurermRedisCacheRedisConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

