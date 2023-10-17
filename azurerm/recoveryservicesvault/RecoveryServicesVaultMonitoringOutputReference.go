// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package recoveryservicesvault

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v11/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v11/recoveryservicesvault/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RecoveryServicesVaultMonitoringOutputReference interface {
	cdktf.ComplexObject
	AlertsForAllJobFailuresEnabled() interface{}
	SetAlertsForAllJobFailuresEnabled(val interface{})
	AlertsForAllJobFailuresEnabledInput() interface{}
	AlertsForCriticalOperationFailuresEnabled() interface{}
	SetAlertsForCriticalOperationFailuresEnabled(val interface{})
	AlertsForCriticalOperationFailuresEnabledInput() interface{}
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
	InternalValue() *RecoveryServicesVaultMonitoring
	SetInternalValue(val *RecoveryServicesVaultMonitoring)
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
	ResetAlertsForAllJobFailuresEnabled()
	ResetAlertsForCriticalOperationFailuresEnabled()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for RecoveryServicesVaultMonitoringOutputReference
type jsiiProxy_RecoveryServicesVaultMonitoringOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_RecoveryServicesVaultMonitoringOutputReference) AlertsForAllJobFailuresEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"alertsForAllJobFailuresEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecoveryServicesVaultMonitoringOutputReference) AlertsForAllJobFailuresEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"alertsForAllJobFailuresEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecoveryServicesVaultMonitoringOutputReference) AlertsForCriticalOperationFailuresEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"alertsForCriticalOperationFailuresEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecoveryServicesVaultMonitoringOutputReference) AlertsForCriticalOperationFailuresEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"alertsForCriticalOperationFailuresEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecoveryServicesVaultMonitoringOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecoveryServicesVaultMonitoringOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecoveryServicesVaultMonitoringOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecoveryServicesVaultMonitoringOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecoveryServicesVaultMonitoringOutputReference) InternalValue() *RecoveryServicesVaultMonitoring {
	var returns *RecoveryServicesVaultMonitoring
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecoveryServicesVaultMonitoringOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecoveryServicesVaultMonitoringOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewRecoveryServicesVaultMonitoringOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) RecoveryServicesVaultMonitoringOutputReference {
	_init_.Initialize()

	if err := validateNewRecoveryServicesVaultMonitoringOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_RecoveryServicesVaultMonitoringOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.recoveryServicesVault.RecoveryServicesVaultMonitoringOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewRecoveryServicesVaultMonitoringOutputReference_Override(r RecoveryServicesVaultMonitoringOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.recoveryServicesVault.RecoveryServicesVaultMonitoringOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		r,
	)
}

func (j *jsiiProxy_RecoveryServicesVaultMonitoringOutputReference)SetAlertsForAllJobFailuresEnabled(val interface{}) {
	if err := j.validateSetAlertsForAllJobFailuresEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"alertsForAllJobFailuresEnabled",
		val,
	)
}

func (j *jsiiProxy_RecoveryServicesVaultMonitoringOutputReference)SetAlertsForCriticalOperationFailuresEnabled(val interface{}) {
	if err := j.validateSetAlertsForCriticalOperationFailuresEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"alertsForCriticalOperationFailuresEnabled",
		val,
	)
}

func (j *jsiiProxy_RecoveryServicesVaultMonitoringOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_RecoveryServicesVaultMonitoringOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_RecoveryServicesVaultMonitoringOutputReference)SetInternalValue(val *RecoveryServicesVaultMonitoring) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_RecoveryServicesVaultMonitoringOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_RecoveryServicesVaultMonitoringOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (r *jsiiProxy_RecoveryServicesVaultMonitoringOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RecoveryServicesVaultMonitoringOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (r *jsiiProxy_RecoveryServicesVaultMonitoringOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (r *jsiiProxy_RecoveryServicesVaultMonitoringOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (r *jsiiProxy_RecoveryServicesVaultMonitoringOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (r *jsiiProxy_RecoveryServicesVaultMonitoringOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (r *jsiiProxy_RecoveryServicesVaultMonitoringOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (r *jsiiProxy_RecoveryServicesVaultMonitoringOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (r *jsiiProxy_RecoveryServicesVaultMonitoringOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (r *jsiiProxy_RecoveryServicesVaultMonitoringOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (r *jsiiProxy_RecoveryServicesVaultMonitoringOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RecoveryServicesVaultMonitoringOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (r *jsiiProxy_RecoveryServicesVaultMonitoringOutputReference) ResetAlertsForAllJobFailuresEnabled() {
	_jsii_.InvokeVoid(
		r,
		"resetAlertsForAllJobFailuresEnabled",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RecoveryServicesVaultMonitoringOutputReference) ResetAlertsForCriticalOperationFailuresEnabled() {
	_jsii_.InvokeVoid(
		r,
		"resetAlertsForCriticalOperationFailuresEnabled",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RecoveryServicesVaultMonitoringOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (r *jsiiProxy_RecoveryServicesVaultMonitoringOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

