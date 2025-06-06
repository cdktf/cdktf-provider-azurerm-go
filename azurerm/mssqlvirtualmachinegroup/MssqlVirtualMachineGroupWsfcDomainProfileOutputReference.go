// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mssqlvirtualmachinegroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/mssqlvirtualmachinegroup/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MssqlVirtualMachineGroupWsfcDomainProfileOutputReference interface {
	cdktf.ComplexObject
	ClusterBootstrapAccountName() *string
	SetClusterBootstrapAccountName(val *string)
	ClusterBootstrapAccountNameInput() *string
	ClusterOperatorAccountName() *string
	SetClusterOperatorAccountName(val *string)
	ClusterOperatorAccountNameInput() *string
	ClusterSubnetType() *string
	SetClusterSubnetType(val *string)
	ClusterSubnetTypeInput() *string
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
	Fqdn() *string
	SetFqdn(val *string)
	FqdnInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *MssqlVirtualMachineGroupWsfcDomainProfile
	SetInternalValue(val *MssqlVirtualMachineGroupWsfcDomainProfile)
	OrganizationalUnitPath() *string
	SetOrganizationalUnitPath(val *string)
	OrganizationalUnitPathInput() *string
	SqlServiceAccountName() *string
	SetSqlServiceAccountName(val *string)
	SqlServiceAccountNameInput() *string
	StorageAccountPrimaryKey() *string
	SetStorageAccountPrimaryKey(val *string)
	StorageAccountPrimaryKeyInput() *string
	StorageAccountUrl() *string
	SetStorageAccountUrl(val *string)
	StorageAccountUrlInput() *string
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
	ResetClusterBootstrapAccountName()
	ResetClusterOperatorAccountName()
	ResetOrganizationalUnitPath()
	ResetSqlServiceAccountName()
	ResetStorageAccountPrimaryKey()
	ResetStorageAccountUrl()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MssqlVirtualMachineGroupWsfcDomainProfileOutputReference
type jsiiProxy_MssqlVirtualMachineGroupWsfcDomainProfileOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MssqlVirtualMachineGroupWsfcDomainProfileOutputReference) ClusterBootstrapAccountName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterBootstrapAccountName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineGroupWsfcDomainProfileOutputReference) ClusterBootstrapAccountNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterBootstrapAccountNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineGroupWsfcDomainProfileOutputReference) ClusterOperatorAccountName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterOperatorAccountName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineGroupWsfcDomainProfileOutputReference) ClusterOperatorAccountNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterOperatorAccountNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineGroupWsfcDomainProfileOutputReference) ClusterSubnetType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterSubnetType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineGroupWsfcDomainProfileOutputReference) ClusterSubnetTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterSubnetTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineGroupWsfcDomainProfileOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineGroupWsfcDomainProfileOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineGroupWsfcDomainProfileOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineGroupWsfcDomainProfileOutputReference) Fqdn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqdn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineGroupWsfcDomainProfileOutputReference) FqdnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqdnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineGroupWsfcDomainProfileOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineGroupWsfcDomainProfileOutputReference) InternalValue() *MssqlVirtualMachineGroupWsfcDomainProfile {
	var returns *MssqlVirtualMachineGroupWsfcDomainProfile
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineGroupWsfcDomainProfileOutputReference) OrganizationalUnitPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organizationalUnitPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineGroupWsfcDomainProfileOutputReference) OrganizationalUnitPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organizationalUnitPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineGroupWsfcDomainProfileOutputReference) SqlServiceAccountName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqlServiceAccountName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineGroupWsfcDomainProfileOutputReference) SqlServiceAccountNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqlServiceAccountNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineGroupWsfcDomainProfileOutputReference) StorageAccountPrimaryKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccountPrimaryKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineGroupWsfcDomainProfileOutputReference) StorageAccountPrimaryKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccountPrimaryKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineGroupWsfcDomainProfileOutputReference) StorageAccountUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccountUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineGroupWsfcDomainProfileOutputReference) StorageAccountUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccountUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineGroupWsfcDomainProfileOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineGroupWsfcDomainProfileOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMssqlVirtualMachineGroupWsfcDomainProfileOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MssqlVirtualMachineGroupWsfcDomainProfileOutputReference {
	_init_.Initialize()

	if err := validateNewMssqlVirtualMachineGroupWsfcDomainProfileOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MssqlVirtualMachineGroupWsfcDomainProfileOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.mssqlVirtualMachineGroup.MssqlVirtualMachineGroupWsfcDomainProfileOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMssqlVirtualMachineGroupWsfcDomainProfileOutputReference_Override(m MssqlVirtualMachineGroupWsfcDomainProfileOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.mssqlVirtualMachineGroup.MssqlVirtualMachineGroupWsfcDomainProfileOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineGroupWsfcDomainProfileOutputReference)SetClusterBootstrapAccountName(val *string) {
	if err := j.validateSetClusterBootstrapAccountNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterBootstrapAccountName",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineGroupWsfcDomainProfileOutputReference)SetClusterOperatorAccountName(val *string) {
	if err := j.validateSetClusterOperatorAccountNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterOperatorAccountName",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineGroupWsfcDomainProfileOutputReference)SetClusterSubnetType(val *string) {
	if err := j.validateSetClusterSubnetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterSubnetType",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineGroupWsfcDomainProfileOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineGroupWsfcDomainProfileOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineGroupWsfcDomainProfileOutputReference)SetFqdn(val *string) {
	if err := j.validateSetFqdnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fqdn",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineGroupWsfcDomainProfileOutputReference)SetInternalValue(val *MssqlVirtualMachineGroupWsfcDomainProfile) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineGroupWsfcDomainProfileOutputReference)SetOrganizationalUnitPath(val *string) {
	if err := j.validateSetOrganizationalUnitPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"organizationalUnitPath",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineGroupWsfcDomainProfileOutputReference)SetSqlServiceAccountName(val *string) {
	if err := j.validateSetSqlServiceAccountNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sqlServiceAccountName",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineGroupWsfcDomainProfileOutputReference)SetStorageAccountPrimaryKey(val *string) {
	if err := j.validateSetStorageAccountPrimaryKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageAccountPrimaryKey",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineGroupWsfcDomainProfileOutputReference)SetStorageAccountUrl(val *string) {
	if err := j.validateSetStorageAccountUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageAccountUrl",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineGroupWsfcDomainProfileOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineGroupWsfcDomainProfileOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MssqlVirtualMachineGroupWsfcDomainProfileOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineGroupWsfcDomainProfileOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MssqlVirtualMachineGroupWsfcDomainProfileOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MssqlVirtualMachineGroupWsfcDomainProfileOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MssqlVirtualMachineGroupWsfcDomainProfileOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MssqlVirtualMachineGroupWsfcDomainProfileOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MssqlVirtualMachineGroupWsfcDomainProfileOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MssqlVirtualMachineGroupWsfcDomainProfileOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MssqlVirtualMachineGroupWsfcDomainProfileOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MssqlVirtualMachineGroupWsfcDomainProfileOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MssqlVirtualMachineGroupWsfcDomainProfileOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineGroupWsfcDomainProfileOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MssqlVirtualMachineGroupWsfcDomainProfileOutputReference) ResetClusterBootstrapAccountName() {
	_jsii_.InvokeVoid(
		m,
		"resetClusterBootstrapAccountName",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlVirtualMachineGroupWsfcDomainProfileOutputReference) ResetClusterOperatorAccountName() {
	_jsii_.InvokeVoid(
		m,
		"resetClusterOperatorAccountName",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlVirtualMachineGroupWsfcDomainProfileOutputReference) ResetOrganizationalUnitPath() {
	_jsii_.InvokeVoid(
		m,
		"resetOrganizationalUnitPath",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlVirtualMachineGroupWsfcDomainProfileOutputReference) ResetSqlServiceAccountName() {
	_jsii_.InvokeVoid(
		m,
		"resetSqlServiceAccountName",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlVirtualMachineGroupWsfcDomainProfileOutputReference) ResetStorageAccountPrimaryKey() {
	_jsii_.InvokeVoid(
		m,
		"resetStorageAccountPrimaryKey",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlVirtualMachineGroupWsfcDomainProfileOutputReference) ResetStorageAccountUrl() {
	_jsii_.InvokeVoid(
		m,
		"resetStorageAccountUrl",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlVirtualMachineGroupWsfcDomainProfileOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_MssqlVirtualMachineGroupWsfcDomainProfileOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

