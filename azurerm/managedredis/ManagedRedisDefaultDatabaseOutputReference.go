// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package managedredis

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/managedredis/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ManagedRedisDefaultDatabaseOutputReference interface {
	cdktf.ComplexObject
	AccessKeysAuthenticationEnabled() interface{}
	SetAccessKeysAuthenticationEnabled(val interface{})
	AccessKeysAuthenticationEnabledInput() interface{}
	ClientProtocol() *string
	SetClientProtocol(val *string)
	ClientProtocolInput() *string
	ClusteringPolicy() *string
	SetClusteringPolicy(val *string)
	ClusteringPolicyInput() *string
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
	EvictionPolicy() *string
	SetEvictionPolicy(val *string)
	EvictionPolicyInput() *string
	// Experimental.
	Fqn() *string
	GeoReplicationGroupName() *string
	SetGeoReplicationGroupName(val *string)
	GeoReplicationGroupNameInput() *string
	InternalValue() *ManagedRedisDefaultDatabase
	SetInternalValue(val *ManagedRedisDefaultDatabase)
	Module() ManagedRedisDefaultDatabaseModuleList
	ModuleInput() interface{}
	Port() *float64
	PrimaryAccessKey() *string
	SecondaryAccessKey() *string
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
	PutModule(value interface{})
	ResetAccessKeysAuthenticationEnabled()
	ResetClientProtocol()
	ResetClusteringPolicy()
	ResetEvictionPolicy()
	ResetGeoReplicationGroupName()
	ResetModule()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ManagedRedisDefaultDatabaseOutputReference
type jsiiProxy_ManagedRedisDefaultDatabaseOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ManagedRedisDefaultDatabaseOutputReference) AccessKeysAuthenticationEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accessKeysAuthenticationEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedRedisDefaultDatabaseOutputReference) AccessKeysAuthenticationEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accessKeysAuthenticationEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedRedisDefaultDatabaseOutputReference) ClientProtocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedRedisDefaultDatabaseOutputReference) ClientProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientProtocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedRedisDefaultDatabaseOutputReference) ClusteringPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusteringPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedRedisDefaultDatabaseOutputReference) ClusteringPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusteringPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedRedisDefaultDatabaseOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedRedisDefaultDatabaseOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedRedisDefaultDatabaseOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedRedisDefaultDatabaseOutputReference) EvictionPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"evictionPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedRedisDefaultDatabaseOutputReference) EvictionPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"evictionPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedRedisDefaultDatabaseOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedRedisDefaultDatabaseOutputReference) GeoReplicationGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"geoReplicationGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedRedisDefaultDatabaseOutputReference) GeoReplicationGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"geoReplicationGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedRedisDefaultDatabaseOutputReference) InternalValue() *ManagedRedisDefaultDatabase {
	var returns *ManagedRedisDefaultDatabase
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedRedisDefaultDatabaseOutputReference) Module() ManagedRedisDefaultDatabaseModuleList {
	var returns ManagedRedisDefaultDatabaseModuleList
	_jsii_.Get(
		j,
		"module",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedRedisDefaultDatabaseOutputReference) ModuleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"moduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedRedisDefaultDatabaseOutputReference) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedRedisDefaultDatabaseOutputReference) PrimaryAccessKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryAccessKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedRedisDefaultDatabaseOutputReference) SecondaryAccessKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryAccessKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedRedisDefaultDatabaseOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedRedisDefaultDatabaseOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewManagedRedisDefaultDatabaseOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ManagedRedisDefaultDatabaseOutputReference {
	_init_.Initialize()

	if err := validateNewManagedRedisDefaultDatabaseOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ManagedRedisDefaultDatabaseOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.managedRedis.ManagedRedisDefaultDatabaseOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewManagedRedisDefaultDatabaseOutputReference_Override(m ManagedRedisDefaultDatabaseOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.managedRedis.ManagedRedisDefaultDatabaseOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_ManagedRedisDefaultDatabaseOutputReference)SetAccessKeysAuthenticationEnabled(val interface{}) {
	if err := j.validateSetAccessKeysAuthenticationEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessKeysAuthenticationEnabled",
		val,
	)
}

func (j *jsiiProxy_ManagedRedisDefaultDatabaseOutputReference)SetClientProtocol(val *string) {
	if err := j.validateSetClientProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientProtocol",
		val,
	)
}

func (j *jsiiProxy_ManagedRedisDefaultDatabaseOutputReference)SetClusteringPolicy(val *string) {
	if err := j.validateSetClusteringPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusteringPolicy",
		val,
	)
}

func (j *jsiiProxy_ManagedRedisDefaultDatabaseOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ManagedRedisDefaultDatabaseOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ManagedRedisDefaultDatabaseOutputReference)SetEvictionPolicy(val *string) {
	if err := j.validateSetEvictionPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"evictionPolicy",
		val,
	)
}

func (j *jsiiProxy_ManagedRedisDefaultDatabaseOutputReference)SetGeoReplicationGroupName(val *string) {
	if err := j.validateSetGeoReplicationGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"geoReplicationGroupName",
		val,
	)
}

func (j *jsiiProxy_ManagedRedisDefaultDatabaseOutputReference)SetInternalValue(val *ManagedRedisDefaultDatabase) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ManagedRedisDefaultDatabaseOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ManagedRedisDefaultDatabaseOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_ManagedRedisDefaultDatabaseOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedRedisDefaultDatabaseOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := m.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedRedisDefaultDatabaseOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedRedisDefaultDatabaseOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := m.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedRedisDefaultDatabaseOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := m.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedRedisDefaultDatabaseOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := m.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedRedisDefaultDatabaseOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := m.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedRedisDefaultDatabaseOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := m.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedRedisDefaultDatabaseOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := m.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedRedisDefaultDatabaseOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := m.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedRedisDefaultDatabaseOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedRedisDefaultDatabaseOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedRedisDefaultDatabaseOutputReference) PutModule(value interface{}) {
	if err := m.validatePutModuleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putModule",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_ManagedRedisDefaultDatabaseOutputReference) ResetAccessKeysAuthenticationEnabled() {
	_jsii_.InvokeVoid(
		m,
		"resetAccessKeysAuthenticationEnabled",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedRedisDefaultDatabaseOutputReference) ResetClientProtocol() {
	_jsii_.InvokeVoid(
		m,
		"resetClientProtocol",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedRedisDefaultDatabaseOutputReference) ResetClusteringPolicy() {
	_jsii_.InvokeVoid(
		m,
		"resetClusteringPolicy",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedRedisDefaultDatabaseOutputReference) ResetEvictionPolicy() {
	_jsii_.InvokeVoid(
		m,
		"resetEvictionPolicy",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedRedisDefaultDatabaseOutputReference) ResetGeoReplicationGroupName() {
	_jsii_.InvokeVoid(
		m,
		"resetGeoReplicationGroupName",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedRedisDefaultDatabaseOutputReference) ResetModule() {
	_jsii_.InvokeVoid(
		m,
		"resetModule",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedRedisDefaultDatabaseOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := m.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedRedisDefaultDatabaseOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

