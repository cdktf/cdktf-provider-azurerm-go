// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package stackhcideploymentsetting

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/stackhcideploymentsetting/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type StackHciDeploymentSettingScaleUnitInfrastructureNetworkOutputReference interface {
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
	DhcpEnabled() interface{}
	SetDhcpEnabled(val interface{})
	DhcpEnabledInput() interface{}
	DnsServer() *[]*string
	SetDnsServer(val *[]*string)
	DnsServerInput() *[]*string
	// Experimental.
	Fqn() *string
	Gateway() *string
	SetGateway(val *string)
	GatewayInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	IpPool() StackHciDeploymentSettingScaleUnitInfrastructureNetworkIpPoolList
	IpPoolInput() interface{}
	SubnetMask() *string
	SetSubnetMask(val *string)
	SubnetMaskInput() *string
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
	PutIpPool(value interface{})
	ResetDhcpEnabled()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for StackHciDeploymentSettingScaleUnitInfrastructureNetworkOutputReference
type jsiiProxy_StackHciDeploymentSettingScaleUnitInfrastructureNetworkOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitInfrastructureNetworkOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitInfrastructureNetworkOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitInfrastructureNetworkOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitInfrastructureNetworkOutputReference) DhcpEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dhcpEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitInfrastructureNetworkOutputReference) DhcpEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dhcpEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitInfrastructureNetworkOutputReference) DnsServer() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dnsServer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitInfrastructureNetworkOutputReference) DnsServerInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dnsServerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitInfrastructureNetworkOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitInfrastructureNetworkOutputReference) Gateway() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitInfrastructureNetworkOutputReference) GatewayInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitInfrastructureNetworkOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitInfrastructureNetworkOutputReference) IpPool() StackHciDeploymentSettingScaleUnitInfrastructureNetworkIpPoolList {
	var returns StackHciDeploymentSettingScaleUnitInfrastructureNetworkIpPoolList
	_jsii_.Get(
		j,
		"ipPool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitInfrastructureNetworkOutputReference) IpPoolInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipPoolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitInfrastructureNetworkOutputReference) SubnetMask() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetMask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitInfrastructureNetworkOutputReference) SubnetMaskInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetMaskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitInfrastructureNetworkOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitInfrastructureNetworkOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewStackHciDeploymentSettingScaleUnitInfrastructureNetworkOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) StackHciDeploymentSettingScaleUnitInfrastructureNetworkOutputReference {
	_init_.Initialize()

	if err := validateNewStackHciDeploymentSettingScaleUnitInfrastructureNetworkOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_StackHciDeploymentSettingScaleUnitInfrastructureNetworkOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.stackHciDeploymentSetting.StackHciDeploymentSettingScaleUnitInfrastructureNetworkOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewStackHciDeploymentSettingScaleUnitInfrastructureNetworkOutputReference_Override(s StackHciDeploymentSettingScaleUnitInfrastructureNetworkOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.stackHciDeploymentSetting.StackHciDeploymentSettingScaleUnitInfrastructureNetworkOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitInfrastructureNetworkOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitInfrastructureNetworkOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitInfrastructureNetworkOutputReference)SetDhcpEnabled(val interface{}) {
	if err := j.validateSetDhcpEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dhcpEnabled",
		val,
	)
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitInfrastructureNetworkOutputReference)SetDnsServer(val *[]*string) {
	if err := j.validateSetDnsServerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dnsServer",
		val,
	)
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitInfrastructureNetworkOutputReference)SetGateway(val *string) {
	if err := j.validateSetGatewayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gateway",
		val,
	)
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitInfrastructureNetworkOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitInfrastructureNetworkOutputReference)SetSubnetMask(val *string) {
	if err := j.validateSetSubnetMaskParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetMask",
		val,
	)
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitInfrastructureNetworkOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitInfrastructureNetworkOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_StackHciDeploymentSettingScaleUnitInfrastructureNetworkOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StackHciDeploymentSettingScaleUnitInfrastructureNetworkOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_StackHciDeploymentSettingScaleUnitInfrastructureNetworkOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_StackHciDeploymentSettingScaleUnitInfrastructureNetworkOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_StackHciDeploymentSettingScaleUnitInfrastructureNetworkOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_StackHciDeploymentSettingScaleUnitInfrastructureNetworkOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_StackHciDeploymentSettingScaleUnitInfrastructureNetworkOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_StackHciDeploymentSettingScaleUnitInfrastructureNetworkOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_StackHciDeploymentSettingScaleUnitInfrastructureNetworkOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_StackHciDeploymentSettingScaleUnitInfrastructureNetworkOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_StackHciDeploymentSettingScaleUnitInfrastructureNetworkOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StackHciDeploymentSettingScaleUnitInfrastructureNetworkOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (s *jsiiProxy_StackHciDeploymentSettingScaleUnitInfrastructureNetworkOutputReference) PutIpPool(value interface{}) {
	if err := s.validatePutIpPoolParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putIpPool",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StackHciDeploymentSettingScaleUnitInfrastructureNetworkOutputReference) ResetDhcpEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetDhcpEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StackHciDeploymentSettingScaleUnitInfrastructureNetworkOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_StackHciDeploymentSettingScaleUnitInfrastructureNetworkOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

