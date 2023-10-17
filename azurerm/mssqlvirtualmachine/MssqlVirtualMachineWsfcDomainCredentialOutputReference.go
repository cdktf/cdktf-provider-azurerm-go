// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mssqlvirtualmachine

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v11/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v11/mssqlvirtualmachine/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MssqlVirtualMachineWsfcDomainCredentialOutputReference interface {
	cdktf.ComplexObject
	ClusterBootstrapAccountPassword() *string
	SetClusterBootstrapAccountPassword(val *string)
	ClusterBootstrapAccountPasswordInput() *string
	ClusterOperatorAccountPassword() *string
	SetClusterOperatorAccountPassword(val *string)
	ClusterOperatorAccountPasswordInput() *string
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
	InternalValue() *MssqlVirtualMachineWsfcDomainCredential
	SetInternalValue(val *MssqlVirtualMachineWsfcDomainCredential)
	SqlServiceAccountPassword() *string
	SetSqlServiceAccountPassword(val *string)
	SqlServiceAccountPasswordInput() *string
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

// The jsii proxy struct for MssqlVirtualMachineWsfcDomainCredentialOutputReference
type jsiiProxy_MssqlVirtualMachineWsfcDomainCredentialOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MssqlVirtualMachineWsfcDomainCredentialOutputReference) ClusterBootstrapAccountPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterBootstrapAccountPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineWsfcDomainCredentialOutputReference) ClusterBootstrapAccountPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterBootstrapAccountPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineWsfcDomainCredentialOutputReference) ClusterOperatorAccountPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterOperatorAccountPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineWsfcDomainCredentialOutputReference) ClusterOperatorAccountPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterOperatorAccountPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineWsfcDomainCredentialOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineWsfcDomainCredentialOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineWsfcDomainCredentialOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineWsfcDomainCredentialOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineWsfcDomainCredentialOutputReference) InternalValue() *MssqlVirtualMachineWsfcDomainCredential {
	var returns *MssqlVirtualMachineWsfcDomainCredential
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineWsfcDomainCredentialOutputReference) SqlServiceAccountPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqlServiceAccountPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineWsfcDomainCredentialOutputReference) SqlServiceAccountPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqlServiceAccountPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineWsfcDomainCredentialOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineWsfcDomainCredentialOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMssqlVirtualMachineWsfcDomainCredentialOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MssqlVirtualMachineWsfcDomainCredentialOutputReference {
	_init_.Initialize()

	if err := validateNewMssqlVirtualMachineWsfcDomainCredentialOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MssqlVirtualMachineWsfcDomainCredentialOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.mssqlVirtualMachine.MssqlVirtualMachineWsfcDomainCredentialOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMssqlVirtualMachineWsfcDomainCredentialOutputReference_Override(m MssqlVirtualMachineWsfcDomainCredentialOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.mssqlVirtualMachine.MssqlVirtualMachineWsfcDomainCredentialOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineWsfcDomainCredentialOutputReference)SetClusterBootstrapAccountPassword(val *string) {
	if err := j.validateSetClusterBootstrapAccountPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterBootstrapAccountPassword",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineWsfcDomainCredentialOutputReference)SetClusterOperatorAccountPassword(val *string) {
	if err := j.validateSetClusterOperatorAccountPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterOperatorAccountPassword",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineWsfcDomainCredentialOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineWsfcDomainCredentialOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineWsfcDomainCredentialOutputReference)SetInternalValue(val *MssqlVirtualMachineWsfcDomainCredential) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineWsfcDomainCredentialOutputReference)SetSqlServiceAccountPassword(val *string) {
	if err := j.validateSetSqlServiceAccountPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sqlServiceAccountPassword",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineWsfcDomainCredentialOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineWsfcDomainCredentialOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MssqlVirtualMachineWsfcDomainCredentialOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineWsfcDomainCredentialOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MssqlVirtualMachineWsfcDomainCredentialOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MssqlVirtualMachineWsfcDomainCredentialOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MssqlVirtualMachineWsfcDomainCredentialOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MssqlVirtualMachineWsfcDomainCredentialOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MssqlVirtualMachineWsfcDomainCredentialOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MssqlVirtualMachineWsfcDomainCredentialOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MssqlVirtualMachineWsfcDomainCredentialOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MssqlVirtualMachineWsfcDomainCredentialOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MssqlVirtualMachineWsfcDomainCredentialOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineWsfcDomainCredentialOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MssqlVirtualMachineWsfcDomainCredentialOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_MssqlVirtualMachineWsfcDomainCredentialOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

