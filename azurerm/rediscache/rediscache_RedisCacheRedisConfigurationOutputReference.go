package rediscache

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v4/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v4/rediscache/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RedisCacheRedisConfigurationOutputReference interface {
	cdktf.ComplexObject
	AofBackupEnabled() interface{}
	SetAofBackupEnabled(val interface{})
	AofBackupEnabledInput() interface{}
	AofStorageConnectionString0() *string
	SetAofStorageConnectionString0(val *string)
	AofStorageConnectionString0Input() *string
	AofStorageConnectionString1() *string
	SetAofStorageConnectionString1(val *string)
	AofStorageConnectionString1Input() *string
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
	EnableAuthentication() interface{}
	SetEnableAuthentication(val interface{})
	EnableAuthenticationInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *RedisCacheRedisConfiguration
	SetInternalValue(val *RedisCacheRedisConfiguration)
	Maxclients() *float64
	MaxfragmentationmemoryReserved() *float64
	SetMaxfragmentationmemoryReserved(val *float64)
	MaxfragmentationmemoryReservedInput() *float64
	MaxmemoryDelta() *float64
	SetMaxmemoryDelta(val *float64)
	MaxmemoryDeltaInput() *float64
	MaxmemoryPolicy() *string
	SetMaxmemoryPolicy(val *string)
	MaxmemoryPolicyInput() *string
	MaxmemoryReserved() *float64
	SetMaxmemoryReserved(val *float64)
	MaxmemoryReservedInput() *float64
	NotifyKeyspaceEvents() *string
	SetNotifyKeyspaceEvents(val *string)
	NotifyKeyspaceEventsInput() *string
	RdbBackupEnabled() interface{}
	SetRdbBackupEnabled(val interface{})
	RdbBackupEnabledInput() interface{}
	RdbBackupFrequency() *float64
	SetRdbBackupFrequency(val *float64)
	RdbBackupFrequencyInput() *float64
	RdbBackupMaxSnapshotCount() *float64
	SetRdbBackupMaxSnapshotCount(val *float64)
	RdbBackupMaxSnapshotCountInput() *float64
	RdbStorageConnectionString() *string
	SetRdbStorageConnectionString(val *string)
	RdbStorageConnectionStringInput() *string
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
	ResetAofBackupEnabled()
	ResetAofStorageConnectionString0()
	ResetAofStorageConnectionString1()
	ResetEnableAuthentication()
	ResetMaxfragmentationmemoryReserved()
	ResetMaxmemoryDelta()
	ResetMaxmemoryPolicy()
	ResetMaxmemoryReserved()
	ResetNotifyKeyspaceEvents()
	ResetRdbBackupEnabled()
	ResetRdbBackupFrequency()
	ResetRdbBackupMaxSnapshotCount()
	ResetRdbStorageConnectionString()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for RedisCacheRedisConfigurationOutputReference
type jsiiProxy_RedisCacheRedisConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_RedisCacheRedisConfigurationOutputReference) AofBackupEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"aofBackupEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCacheRedisConfigurationOutputReference) AofBackupEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"aofBackupEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCacheRedisConfigurationOutputReference) AofStorageConnectionString0() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aofStorageConnectionString0",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCacheRedisConfigurationOutputReference) AofStorageConnectionString0Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aofStorageConnectionString0Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCacheRedisConfigurationOutputReference) AofStorageConnectionString1() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aofStorageConnectionString1",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCacheRedisConfigurationOutputReference) AofStorageConnectionString1Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aofStorageConnectionString1Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCacheRedisConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCacheRedisConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCacheRedisConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCacheRedisConfigurationOutputReference) EnableAuthentication() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableAuthentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCacheRedisConfigurationOutputReference) EnableAuthenticationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableAuthenticationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCacheRedisConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCacheRedisConfigurationOutputReference) InternalValue() *RedisCacheRedisConfiguration {
	var returns *RedisCacheRedisConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCacheRedisConfigurationOutputReference) Maxclients() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxclients",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCacheRedisConfigurationOutputReference) MaxfragmentationmemoryReserved() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxfragmentationmemoryReserved",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCacheRedisConfigurationOutputReference) MaxfragmentationmemoryReservedInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxfragmentationmemoryReservedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCacheRedisConfigurationOutputReference) MaxmemoryDelta() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxmemoryDelta",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCacheRedisConfigurationOutputReference) MaxmemoryDeltaInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxmemoryDeltaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCacheRedisConfigurationOutputReference) MaxmemoryPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxmemoryPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCacheRedisConfigurationOutputReference) MaxmemoryPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxmemoryPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCacheRedisConfigurationOutputReference) MaxmemoryReserved() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxmemoryReserved",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCacheRedisConfigurationOutputReference) MaxmemoryReservedInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxmemoryReservedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCacheRedisConfigurationOutputReference) NotifyKeyspaceEvents() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notifyKeyspaceEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCacheRedisConfigurationOutputReference) NotifyKeyspaceEventsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notifyKeyspaceEventsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCacheRedisConfigurationOutputReference) RdbBackupEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rdbBackupEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCacheRedisConfigurationOutputReference) RdbBackupEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rdbBackupEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCacheRedisConfigurationOutputReference) RdbBackupFrequency() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rdbBackupFrequency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCacheRedisConfigurationOutputReference) RdbBackupFrequencyInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rdbBackupFrequencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCacheRedisConfigurationOutputReference) RdbBackupMaxSnapshotCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rdbBackupMaxSnapshotCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCacheRedisConfigurationOutputReference) RdbBackupMaxSnapshotCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rdbBackupMaxSnapshotCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCacheRedisConfigurationOutputReference) RdbStorageConnectionString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rdbStorageConnectionString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCacheRedisConfigurationOutputReference) RdbStorageConnectionStringInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rdbStorageConnectionStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCacheRedisConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCacheRedisConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewRedisCacheRedisConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) RedisCacheRedisConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewRedisCacheRedisConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_RedisCacheRedisConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.redisCache.RedisCacheRedisConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewRedisCacheRedisConfigurationOutputReference_Override(r RedisCacheRedisConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.redisCache.RedisCacheRedisConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		r,
	)
}

func (j *jsiiProxy_RedisCacheRedisConfigurationOutputReference)SetAofBackupEnabled(val interface{}) {
	if err := j.validateSetAofBackupEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"aofBackupEnabled",
		val,
	)
}

func (j *jsiiProxy_RedisCacheRedisConfigurationOutputReference)SetAofStorageConnectionString0(val *string) {
	if err := j.validateSetAofStorageConnectionString0Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"aofStorageConnectionString0",
		val,
	)
}

func (j *jsiiProxy_RedisCacheRedisConfigurationOutputReference)SetAofStorageConnectionString1(val *string) {
	if err := j.validateSetAofStorageConnectionString1Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"aofStorageConnectionString1",
		val,
	)
}

func (j *jsiiProxy_RedisCacheRedisConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_RedisCacheRedisConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_RedisCacheRedisConfigurationOutputReference)SetEnableAuthentication(val interface{}) {
	if err := j.validateSetEnableAuthenticationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableAuthentication",
		val,
	)
}

func (j *jsiiProxy_RedisCacheRedisConfigurationOutputReference)SetInternalValue(val *RedisCacheRedisConfiguration) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_RedisCacheRedisConfigurationOutputReference)SetMaxfragmentationmemoryReserved(val *float64) {
	if err := j.validateSetMaxfragmentationmemoryReservedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxfragmentationmemoryReserved",
		val,
	)
}

func (j *jsiiProxy_RedisCacheRedisConfigurationOutputReference)SetMaxmemoryDelta(val *float64) {
	if err := j.validateSetMaxmemoryDeltaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxmemoryDelta",
		val,
	)
}

func (j *jsiiProxy_RedisCacheRedisConfigurationOutputReference)SetMaxmemoryPolicy(val *string) {
	if err := j.validateSetMaxmemoryPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxmemoryPolicy",
		val,
	)
}

func (j *jsiiProxy_RedisCacheRedisConfigurationOutputReference)SetMaxmemoryReserved(val *float64) {
	if err := j.validateSetMaxmemoryReservedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxmemoryReserved",
		val,
	)
}

func (j *jsiiProxy_RedisCacheRedisConfigurationOutputReference)SetNotifyKeyspaceEvents(val *string) {
	if err := j.validateSetNotifyKeyspaceEventsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notifyKeyspaceEvents",
		val,
	)
}

func (j *jsiiProxy_RedisCacheRedisConfigurationOutputReference)SetRdbBackupEnabled(val interface{}) {
	if err := j.validateSetRdbBackupEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rdbBackupEnabled",
		val,
	)
}

func (j *jsiiProxy_RedisCacheRedisConfigurationOutputReference)SetRdbBackupFrequency(val *float64) {
	if err := j.validateSetRdbBackupFrequencyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rdbBackupFrequency",
		val,
	)
}

func (j *jsiiProxy_RedisCacheRedisConfigurationOutputReference)SetRdbBackupMaxSnapshotCount(val *float64) {
	if err := j.validateSetRdbBackupMaxSnapshotCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rdbBackupMaxSnapshotCount",
		val,
	)
}

func (j *jsiiProxy_RedisCacheRedisConfigurationOutputReference)SetRdbStorageConnectionString(val *string) {
	if err := j.validateSetRdbStorageConnectionStringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rdbStorageConnectionString",
		val,
	)
}

func (j *jsiiProxy_RedisCacheRedisConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_RedisCacheRedisConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (r *jsiiProxy_RedisCacheRedisConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisCacheRedisConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := r.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisCacheRedisConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := r.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisCacheRedisConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := r.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		r,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisCacheRedisConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := r.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisCacheRedisConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := r.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		r,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisCacheRedisConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := r.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		r,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisCacheRedisConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := r.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		r,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisCacheRedisConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := r.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		r,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisCacheRedisConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := r.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		r,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisCacheRedisConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisCacheRedisConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := r.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisCacheRedisConfigurationOutputReference) ResetAofBackupEnabled() {
	_jsii_.InvokeVoid(
		r,
		"resetAofBackupEnabled",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedisCacheRedisConfigurationOutputReference) ResetAofStorageConnectionString0() {
	_jsii_.InvokeVoid(
		r,
		"resetAofStorageConnectionString0",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedisCacheRedisConfigurationOutputReference) ResetAofStorageConnectionString1() {
	_jsii_.InvokeVoid(
		r,
		"resetAofStorageConnectionString1",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedisCacheRedisConfigurationOutputReference) ResetEnableAuthentication() {
	_jsii_.InvokeVoid(
		r,
		"resetEnableAuthentication",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedisCacheRedisConfigurationOutputReference) ResetMaxfragmentationmemoryReserved() {
	_jsii_.InvokeVoid(
		r,
		"resetMaxfragmentationmemoryReserved",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedisCacheRedisConfigurationOutputReference) ResetMaxmemoryDelta() {
	_jsii_.InvokeVoid(
		r,
		"resetMaxmemoryDelta",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedisCacheRedisConfigurationOutputReference) ResetMaxmemoryPolicy() {
	_jsii_.InvokeVoid(
		r,
		"resetMaxmemoryPolicy",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedisCacheRedisConfigurationOutputReference) ResetMaxmemoryReserved() {
	_jsii_.InvokeVoid(
		r,
		"resetMaxmemoryReserved",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedisCacheRedisConfigurationOutputReference) ResetNotifyKeyspaceEvents() {
	_jsii_.InvokeVoid(
		r,
		"resetNotifyKeyspaceEvents",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedisCacheRedisConfigurationOutputReference) ResetRdbBackupEnabled() {
	_jsii_.InvokeVoid(
		r,
		"resetRdbBackupEnabled",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedisCacheRedisConfigurationOutputReference) ResetRdbBackupFrequency() {
	_jsii_.InvokeVoid(
		r,
		"resetRdbBackupFrequency",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedisCacheRedisConfigurationOutputReference) ResetRdbBackupMaxSnapshotCount() {
	_jsii_.InvokeVoid(
		r,
		"resetRdbBackupMaxSnapshotCount",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedisCacheRedisConfigurationOutputReference) ResetRdbStorageConnectionString() {
	_jsii_.InvokeVoid(
		r,
		"resetRdbStorageConnectionString",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedisCacheRedisConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := r.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		r,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisCacheRedisConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

