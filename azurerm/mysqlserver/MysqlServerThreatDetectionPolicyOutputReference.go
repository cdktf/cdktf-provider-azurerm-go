// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mysqlserver

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v11/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v11/mysqlserver/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MysqlServerThreatDetectionPolicyOutputReference interface {
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
	DisabledAlerts() *[]*string
	SetDisabledAlerts(val *[]*string)
	DisabledAlertsInput() *[]*string
	EmailAccountAdmins() interface{}
	SetEmailAccountAdmins(val interface{})
	EmailAccountAdminsInput() interface{}
	EmailAddresses() *[]*string
	SetEmailAddresses(val *[]*string)
	EmailAddressesInput() *[]*string
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *MysqlServerThreatDetectionPolicy
	SetInternalValue(val *MysqlServerThreatDetectionPolicy)
	RetentionDays() *float64
	SetRetentionDays(val *float64)
	RetentionDaysInput() *float64
	StorageAccountAccessKey() *string
	SetStorageAccountAccessKey(val *string)
	StorageAccountAccessKeyInput() *string
	StorageEndpoint() *string
	SetStorageEndpoint(val *string)
	StorageEndpointInput() *string
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
	ResetDisabledAlerts()
	ResetEmailAccountAdmins()
	ResetEmailAddresses()
	ResetEnabled()
	ResetRetentionDays()
	ResetStorageAccountAccessKey()
	ResetStorageEndpoint()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MysqlServerThreatDetectionPolicyOutputReference
type jsiiProxy_MysqlServerThreatDetectionPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MysqlServerThreatDetectionPolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServerThreatDetectionPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServerThreatDetectionPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServerThreatDetectionPolicyOutputReference) DisabledAlerts() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"disabledAlerts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServerThreatDetectionPolicyOutputReference) DisabledAlertsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"disabledAlertsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServerThreatDetectionPolicyOutputReference) EmailAccountAdmins() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"emailAccountAdmins",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServerThreatDetectionPolicyOutputReference) EmailAccountAdminsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"emailAccountAdminsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServerThreatDetectionPolicyOutputReference) EmailAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"emailAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServerThreatDetectionPolicyOutputReference) EmailAddressesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"emailAddressesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServerThreatDetectionPolicyOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServerThreatDetectionPolicyOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServerThreatDetectionPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServerThreatDetectionPolicyOutputReference) InternalValue() *MysqlServerThreatDetectionPolicy {
	var returns *MysqlServerThreatDetectionPolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServerThreatDetectionPolicyOutputReference) RetentionDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retentionDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServerThreatDetectionPolicyOutputReference) RetentionDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retentionDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServerThreatDetectionPolicyOutputReference) StorageAccountAccessKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccountAccessKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServerThreatDetectionPolicyOutputReference) StorageAccountAccessKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccountAccessKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServerThreatDetectionPolicyOutputReference) StorageEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServerThreatDetectionPolicyOutputReference) StorageEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServerThreatDetectionPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlServerThreatDetectionPolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMysqlServerThreatDetectionPolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MysqlServerThreatDetectionPolicyOutputReference {
	_init_.Initialize()

	if err := validateNewMysqlServerThreatDetectionPolicyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MysqlServerThreatDetectionPolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.mysqlServer.MysqlServerThreatDetectionPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMysqlServerThreatDetectionPolicyOutputReference_Override(m MysqlServerThreatDetectionPolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.mysqlServer.MysqlServerThreatDetectionPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MysqlServerThreatDetectionPolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MysqlServerThreatDetectionPolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MysqlServerThreatDetectionPolicyOutputReference)SetDisabledAlerts(val *[]*string) {
	if err := j.validateSetDisabledAlertsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disabledAlerts",
		val,
	)
}

func (j *jsiiProxy_MysqlServerThreatDetectionPolicyOutputReference)SetEmailAccountAdmins(val interface{}) {
	if err := j.validateSetEmailAccountAdminsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"emailAccountAdmins",
		val,
	)
}

func (j *jsiiProxy_MysqlServerThreatDetectionPolicyOutputReference)SetEmailAddresses(val *[]*string) {
	if err := j.validateSetEmailAddressesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"emailAddresses",
		val,
	)
}

func (j *jsiiProxy_MysqlServerThreatDetectionPolicyOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_MysqlServerThreatDetectionPolicyOutputReference)SetInternalValue(val *MysqlServerThreatDetectionPolicy) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MysqlServerThreatDetectionPolicyOutputReference)SetRetentionDays(val *float64) {
	if err := j.validateSetRetentionDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retentionDays",
		val,
	)
}

func (j *jsiiProxy_MysqlServerThreatDetectionPolicyOutputReference)SetStorageAccountAccessKey(val *string) {
	if err := j.validateSetStorageAccountAccessKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageAccountAccessKey",
		val,
	)
}

func (j *jsiiProxy_MysqlServerThreatDetectionPolicyOutputReference)SetStorageEndpoint(val *string) {
	if err := j.validateSetStorageEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageEndpoint",
		val,
	)
}

func (j *jsiiProxy_MysqlServerThreatDetectionPolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MysqlServerThreatDetectionPolicyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MysqlServerThreatDetectionPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MysqlServerThreatDetectionPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MysqlServerThreatDetectionPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MysqlServerThreatDetectionPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MysqlServerThreatDetectionPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MysqlServerThreatDetectionPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MysqlServerThreatDetectionPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MysqlServerThreatDetectionPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MysqlServerThreatDetectionPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MysqlServerThreatDetectionPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MysqlServerThreatDetectionPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MysqlServerThreatDetectionPolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MysqlServerThreatDetectionPolicyOutputReference) ResetDisabledAlerts() {
	_jsii_.InvokeVoid(
		m,
		"resetDisabledAlerts",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MysqlServerThreatDetectionPolicyOutputReference) ResetEmailAccountAdmins() {
	_jsii_.InvokeVoid(
		m,
		"resetEmailAccountAdmins",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MysqlServerThreatDetectionPolicyOutputReference) ResetEmailAddresses() {
	_jsii_.InvokeVoid(
		m,
		"resetEmailAddresses",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MysqlServerThreatDetectionPolicyOutputReference) ResetEnabled() {
	_jsii_.InvokeVoid(
		m,
		"resetEnabled",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MysqlServerThreatDetectionPolicyOutputReference) ResetRetentionDays() {
	_jsii_.InvokeVoid(
		m,
		"resetRetentionDays",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MysqlServerThreatDetectionPolicyOutputReference) ResetStorageAccountAccessKey() {
	_jsii_.InvokeVoid(
		m,
		"resetStorageAccountAccessKey",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MysqlServerThreatDetectionPolicyOutputReference) ResetStorageEndpoint() {
	_jsii_.InvokeVoid(
		m,
		"resetStorageEndpoint",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MysqlServerThreatDetectionPolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_MysqlServerThreatDetectionPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

