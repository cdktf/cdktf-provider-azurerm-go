// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sqldatabase

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v10/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v10/sqldatabase/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SqlDatabaseThreatDetectionPolicyOutputReference interface {
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
	EmailAccountAdmins() *string
	SetEmailAccountAdmins(val *string)
	EmailAccountAdminsInput() *string
	EmailAddresses() *[]*string
	SetEmailAddresses(val *[]*string)
	EmailAddressesInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *SqlDatabaseThreatDetectionPolicy
	SetInternalValue(val *SqlDatabaseThreatDetectionPolicy)
	RetentionDays() *float64
	SetRetentionDays(val *float64)
	RetentionDaysInput() *float64
	State() *string
	SetState(val *string)
	StateInput() *string
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
	ResetRetentionDays()
	ResetState()
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

// The jsii proxy struct for SqlDatabaseThreatDetectionPolicyOutputReference
type jsiiProxy_SqlDatabaseThreatDetectionPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SqlDatabaseThreatDetectionPolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseThreatDetectionPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseThreatDetectionPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseThreatDetectionPolicyOutputReference) DisabledAlerts() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"disabledAlerts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseThreatDetectionPolicyOutputReference) DisabledAlertsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"disabledAlertsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseThreatDetectionPolicyOutputReference) EmailAccountAdmins() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailAccountAdmins",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseThreatDetectionPolicyOutputReference) EmailAccountAdminsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailAccountAdminsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseThreatDetectionPolicyOutputReference) EmailAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"emailAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseThreatDetectionPolicyOutputReference) EmailAddressesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"emailAddressesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseThreatDetectionPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseThreatDetectionPolicyOutputReference) InternalValue() *SqlDatabaseThreatDetectionPolicy {
	var returns *SqlDatabaseThreatDetectionPolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseThreatDetectionPolicyOutputReference) RetentionDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retentionDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseThreatDetectionPolicyOutputReference) RetentionDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retentionDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseThreatDetectionPolicyOutputReference) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseThreatDetectionPolicyOutputReference) StateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseThreatDetectionPolicyOutputReference) StorageAccountAccessKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccountAccessKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseThreatDetectionPolicyOutputReference) StorageAccountAccessKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccountAccessKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseThreatDetectionPolicyOutputReference) StorageEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseThreatDetectionPolicyOutputReference) StorageEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseThreatDetectionPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqlDatabaseThreatDetectionPolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSqlDatabaseThreatDetectionPolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) SqlDatabaseThreatDetectionPolicyOutputReference {
	_init_.Initialize()

	if err := validateNewSqlDatabaseThreatDetectionPolicyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SqlDatabaseThreatDetectionPolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.sqlDatabase.SqlDatabaseThreatDetectionPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSqlDatabaseThreatDetectionPolicyOutputReference_Override(s SqlDatabaseThreatDetectionPolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.sqlDatabase.SqlDatabaseThreatDetectionPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SqlDatabaseThreatDetectionPolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SqlDatabaseThreatDetectionPolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SqlDatabaseThreatDetectionPolicyOutputReference)SetDisabledAlerts(val *[]*string) {
	if err := j.validateSetDisabledAlertsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disabledAlerts",
		val,
	)
}

func (j *jsiiProxy_SqlDatabaseThreatDetectionPolicyOutputReference)SetEmailAccountAdmins(val *string) {
	if err := j.validateSetEmailAccountAdminsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"emailAccountAdmins",
		val,
	)
}

func (j *jsiiProxy_SqlDatabaseThreatDetectionPolicyOutputReference)SetEmailAddresses(val *[]*string) {
	if err := j.validateSetEmailAddressesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"emailAddresses",
		val,
	)
}

func (j *jsiiProxy_SqlDatabaseThreatDetectionPolicyOutputReference)SetInternalValue(val *SqlDatabaseThreatDetectionPolicy) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SqlDatabaseThreatDetectionPolicyOutputReference)SetRetentionDays(val *float64) {
	if err := j.validateSetRetentionDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retentionDays",
		val,
	)
}

func (j *jsiiProxy_SqlDatabaseThreatDetectionPolicyOutputReference)SetState(val *string) {
	if err := j.validateSetStateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"state",
		val,
	)
}

func (j *jsiiProxy_SqlDatabaseThreatDetectionPolicyOutputReference)SetStorageAccountAccessKey(val *string) {
	if err := j.validateSetStorageAccountAccessKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageAccountAccessKey",
		val,
	)
}

func (j *jsiiProxy_SqlDatabaseThreatDetectionPolicyOutputReference)SetStorageEndpoint(val *string) {
	if err := j.validateSetStorageEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageEndpoint",
		val,
	)
}

func (j *jsiiProxy_SqlDatabaseThreatDetectionPolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SqlDatabaseThreatDetectionPolicyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SqlDatabaseThreatDetectionPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlDatabaseThreatDetectionPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlDatabaseThreatDetectionPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlDatabaseThreatDetectionPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlDatabaseThreatDetectionPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlDatabaseThreatDetectionPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlDatabaseThreatDetectionPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlDatabaseThreatDetectionPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlDatabaseThreatDetectionPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlDatabaseThreatDetectionPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlDatabaseThreatDetectionPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlDatabaseThreatDetectionPolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlDatabaseThreatDetectionPolicyOutputReference) ResetDisabledAlerts() {
	_jsii_.InvokeVoid(
		s,
		"resetDisabledAlerts",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqlDatabaseThreatDetectionPolicyOutputReference) ResetEmailAccountAdmins() {
	_jsii_.InvokeVoid(
		s,
		"resetEmailAccountAdmins",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqlDatabaseThreatDetectionPolicyOutputReference) ResetEmailAddresses() {
	_jsii_.InvokeVoid(
		s,
		"resetEmailAddresses",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqlDatabaseThreatDetectionPolicyOutputReference) ResetRetentionDays() {
	_jsii_.InvokeVoid(
		s,
		"resetRetentionDays",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqlDatabaseThreatDetectionPolicyOutputReference) ResetState() {
	_jsii_.InvokeVoid(
		s,
		"resetState",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqlDatabaseThreatDetectionPolicyOutputReference) ResetStorageAccountAccessKey() {
	_jsii_.InvokeVoid(
		s,
		"resetStorageAccountAccessKey",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqlDatabaseThreatDetectionPolicyOutputReference) ResetStorageEndpoint() {
	_jsii_.InvokeVoid(
		s,
		"resetStorageEndpoint",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SqlDatabaseThreatDetectionPolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := s.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqlDatabaseThreatDetectionPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

