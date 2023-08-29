// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package siterecoveryreplicatedvm

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v10/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v10/siterecoveryreplicatedvm/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SiteRecoveryReplicatedVmNetworkInterfaceOutputReference interface {
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
	FailoverTestPublicIpAddressId() *string
	SetFailoverTestPublicIpAddressId(val *string)
	FailoverTestPublicIpAddressIdInput() *string
	FailoverTestStaticIp() *string
	SetFailoverTestStaticIp(val *string)
	FailoverTestStaticIpInput() *string
	FailoverTestSubnetName() *string
	SetFailoverTestSubnetName(val *string)
	FailoverTestSubnetNameInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	IsPrimary() interface{}
	SetIsPrimary(val interface{})
	IsPrimaryInput() interface{}
	RecoveryPublicIpAddressId() *string
	SetRecoveryPublicIpAddressId(val *string)
	RecoveryPublicIpAddressIdInput() *string
	SourceNetworkInterfaceId() *string
	SetSourceNetworkInterfaceId(val *string)
	SourceNetworkInterfaceIdInput() *string
	TargetStaticIp() *string
	SetTargetStaticIp(val *string)
	TargetStaticIpInput() *string
	TargetSubnetName() *string
	SetTargetSubnetName(val *string)
	TargetSubnetNameInput() *string
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
	ResetFailoverTestPublicIpAddressId()
	ResetFailoverTestStaticIp()
	ResetFailoverTestSubnetName()
	ResetIsPrimary()
	ResetRecoveryPublicIpAddressId()
	ResetSourceNetworkInterfaceId()
	ResetTargetStaticIp()
	ResetTargetSubnetName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SiteRecoveryReplicatedVmNetworkInterfaceOutputReference
type jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceOutputReference) FailoverTestPublicIpAddressId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"failoverTestPublicIpAddressId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceOutputReference) FailoverTestPublicIpAddressIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"failoverTestPublicIpAddressIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceOutputReference) FailoverTestStaticIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"failoverTestStaticIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceOutputReference) FailoverTestStaticIpInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"failoverTestStaticIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceOutputReference) FailoverTestSubnetName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"failoverTestSubnetName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceOutputReference) FailoverTestSubnetNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"failoverTestSubnetNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceOutputReference) IsPrimary() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isPrimary",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceOutputReference) IsPrimaryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isPrimaryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceOutputReference) RecoveryPublicIpAddressId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recoveryPublicIpAddressId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceOutputReference) RecoveryPublicIpAddressIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recoveryPublicIpAddressIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceOutputReference) SourceNetworkInterfaceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceNetworkInterfaceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceOutputReference) SourceNetworkInterfaceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceNetworkInterfaceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceOutputReference) TargetStaticIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetStaticIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceOutputReference) TargetStaticIpInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetStaticIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceOutputReference) TargetSubnetName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetSubnetName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceOutputReference) TargetSubnetNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetSubnetNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSiteRecoveryReplicatedVmNetworkInterfaceOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) SiteRecoveryReplicatedVmNetworkInterfaceOutputReference {
	_init_.Initialize()

	if err := validateNewSiteRecoveryReplicatedVmNetworkInterfaceOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.siteRecoveryReplicatedVm.SiteRecoveryReplicatedVmNetworkInterfaceOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewSiteRecoveryReplicatedVmNetworkInterfaceOutputReference_Override(s SiteRecoveryReplicatedVmNetworkInterfaceOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.siteRecoveryReplicatedVm.SiteRecoveryReplicatedVmNetworkInterfaceOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceOutputReference)SetFailoverTestPublicIpAddressId(val *string) {
	if err := j.validateSetFailoverTestPublicIpAddressIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"failoverTestPublicIpAddressId",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceOutputReference)SetFailoverTestStaticIp(val *string) {
	if err := j.validateSetFailoverTestStaticIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"failoverTestStaticIp",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceOutputReference)SetFailoverTestSubnetName(val *string) {
	if err := j.validateSetFailoverTestSubnetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"failoverTestSubnetName",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceOutputReference)SetIsPrimary(val interface{}) {
	if err := j.validateSetIsPrimaryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isPrimary",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceOutputReference)SetRecoveryPublicIpAddressId(val *string) {
	if err := j.validateSetRecoveryPublicIpAddressIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"recoveryPublicIpAddressId",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceOutputReference)SetSourceNetworkInterfaceId(val *string) {
	if err := j.validateSetSourceNetworkInterfaceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceNetworkInterfaceId",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceOutputReference)SetTargetStaticIp(val *string) {
	if err := j.validateSetTargetStaticIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetStaticIp",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceOutputReference)SetTargetSubnetName(val *string) {
	if err := j.validateSetTargetSubnetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetSubnetName",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceOutputReference) ResetFailoverTestPublicIpAddressId() {
	_jsii_.InvokeVoid(
		s,
		"resetFailoverTestPublicIpAddressId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceOutputReference) ResetFailoverTestStaticIp() {
	_jsii_.InvokeVoid(
		s,
		"resetFailoverTestStaticIp",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceOutputReference) ResetFailoverTestSubnetName() {
	_jsii_.InvokeVoid(
		s,
		"resetFailoverTestSubnetName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceOutputReference) ResetIsPrimary() {
	_jsii_.InvokeVoid(
		s,
		"resetIsPrimary",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceOutputReference) ResetRecoveryPublicIpAddressId() {
	_jsii_.InvokeVoid(
		s,
		"resetRecoveryPublicIpAddressId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceOutputReference) ResetSourceNetworkInterfaceId() {
	_jsii_.InvokeVoid(
		s,
		"resetSourceNetworkInterfaceId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceOutputReference) ResetTargetStaticIp() {
	_jsii_.InvokeVoid(
		s,
		"resetTargetStaticIp",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceOutputReference) ResetTargetSubnetName() {
	_jsii_.InvokeVoid(
		s,
		"resetTargetSubnetName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

