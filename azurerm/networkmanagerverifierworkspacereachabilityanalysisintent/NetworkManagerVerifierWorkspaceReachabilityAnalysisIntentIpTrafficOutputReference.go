// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkmanagerverifierworkspacereachabilityanalysisintent

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/networkmanagerverifierworkspacereachabilityanalysisintent/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetworkManagerVerifierWorkspaceReachabilityAnalysisIntentIpTrafficOutputReference interface {
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
	DestinationIps() *[]*string
	SetDestinationIps(val *[]*string)
	DestinationIpsInput() *[]*string
	DestinationPorts() *[]*string
	SetDestinationPorts(val *[]*string)
	DestinationPortsInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *NetworkManagerVerifierWorkspaceReachabilityAnalysisIntentIpTraffic
	SetInternalValue(val *NetworkManagerVerifierWorkspaceReachabilityAnalysisIntentIpTraffic)
	Protocols() *[]*string
	SetProtocols(val *[]*string)
	ProtocolsInput() *[]*string
	SourceIps() *[]*string
	SetSourceIps(val *[]*string)
	SourceIpsInput() *[]*string
	SourcePorts() *[]*string
	SetSourcePorts(val *[]*string)
	SourcePortsInput() *[]*string
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

// The jsii proxy struct for NetworkManagerVerifierWorkspaceReachabilityAnalysisIntentIpTrafficOutputReference
type jsiiProxy_NetworkManagerVerifierWorkspaceReachabilityAnalysisIntentIpTrafficOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_NetworkManagerVerifierWorkspaceReachabilityAnalysisIntentIpTrafficOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagerVerifierWorkspaceReachabilityAnalysisIntentIpTrafficOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagerVerifierWorkspaceReachabilityAnalysisIntentIpTrafficOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagerVerifierWorkspaceReachabilityAnalysisIntentIpTrafficOutputReference) DestinationIps() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destinationIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagerVerifierWorkspaceReachabilityAnalysisIntentIpTrafficOutputReference) DestinationIpsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destinationIpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagerVerifierWorkspaceReachabilityAnalysisIntentIpTrafficOutputReference) DestinationPorts() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destinationPorts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagerVerifierWorkspaceReachabilityAnalysisIntentIpTrafficOutputReference) DestinationPortsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destinationPortsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagerVerifierWorkspaceReachabilityAnalysisIntentIpTrafficOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagerVerifierWorkspaceReachabilityAnalysisIntentIpTrafficOutputReference) InternalValue() *NetworkManagerVerifierWorkspaceReachabilityAnalysisIntentIpTraffic {
	var returns *NetworkManagerVerifierWorkspaceReachabilityAnalysisIntentIpTraffic
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagerVerifierWorkspaceReachabilityAnalysisIntentIpTrafficOutputReference) Protocols() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"protocols",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagerVerifierWorkspaceReachabilityAnalysisIntentIpTrafficOutputReference) ProtocolsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"protocolsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagerVerifierWorkspaceReachabilityAnalysisIntentIpTrafficOutputReference) SourceIps() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourceIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagerVerifierWorkspaceReachabilityAnalysisIntentIpTrafficOutputReference) SourceIpsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourceIpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagerVerifierWorkspaceReachabilityAnalysisIntentIpTrafficOutputReference) SourcePorts() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourcePorts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagerVerifierWorkspaceReachabilityAnalysisIntentIpTrafficOutputReference) SourcePortsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourcePortsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagerVerifierWorkspaceReachabilityAnalysisIntentIpTrafficOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkManagerVerifierWorkspaceReachabilityAnalysisIntentIpTrafficOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewNetworkManagerVerifierWorkspaceReachabilityAnalysisIntentIpTrafficOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) NetworkManagerVerifierWorkspaceReachabilityAnalysisIntentIpTrafficOutputReference {
	_init_.Initialize()

	if err := validateNewNetworkManagerVerifierWorkspaceReachabilityAnalysisIntentIpTrafficOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_NetworkManagerVerifierWorkspaceReachabilityAnalysisIntentIpTrafficOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.networkManagerVerifierWorkspaceReachabilityAnalysisIntent.NetworkManagerVerifierWorkspaceReachabilityAnalysisIntentIpTrafficOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewNetworkManagerVerifierWorkspaceReachabilityAnalysisIntentIpTrafficOutputReference_Override(n NetworkManagerVerifierWorkspaceReachabilityAnalysisIntentIpTrafficOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.networkManagerVerifierWorkspaceReachabilityAnalysisIntent.NetworkManagerVerifierWorkspaceReachabilityAnalysisIntentIpTrafficOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		n,
	)
}

func (j *jsiiProxy_NetworkManagerVerifierWorkspaceReachabilityAnalysisIntentIpTrafficOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_NetworkManagerVerifierWorkspaceReachabilityAnalysisIntentIpTrafficOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_NetworkManagerVerifierWorkspaceReachabilityAnalysisIntentIpTrafficOutputReference)SetDestinationIps(val *[]*string) {
	if err := j.validateSetDestinationIpsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationIps",
		val,
	)
}

func (j *jsiiProxy_NetworkManagerVerifierWorkspaceReachabilityAnalysisIntentIpTrafficOutputReference)SetDestinationPorts(val *[]*string) {
	if err := j.validateSetDestinationPortsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationPorts",
		val,
	)
}

func (j *jsiiProxy_NetworkManagerVerifierWorkspaceReachabilityAnalysisIntentIpTrafficOutputReference)SetInternalValue(val *NetworkManagerVerifierWorkspaceReachabilityAnalysisIntentIpTraffic) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_NetworkManagerVerifierWorkspaceReachabilityAnalysisIntentIpTrafficOutputReference)SetProtocols(val *[]*string) {
	if err := j.validateSetProtocolsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protocols",
		val,
	)
}

func (j *jsiiProxy_NetworkManagerVerifierWorkspaceReachabilityAnalysisIntentIpTrafficOutputReference)SetSourceIps(val *[]*string) {
	if err := j.validateSetSourceIpsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceIps",
		val,
	)
}

func (j *jsiiProxy_NetworkManagerVerifierWorkspaceReachabilityAnalysisIntentIpTrafficOutputReference)SetSourcePorts(val *[]*string) {
	if err := j.validateSetSourcePortsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourcePorts",
		val,
	)
}

func (j *jsiiProxy_NetworkManagerVerifierWorkspaceReachabilityAnalysisIntentIpTrafficOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_NetworkManagerVerifierWorkspaceReachabilityAnalysisIntentIpTrafficOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (n *jsiiProxy_NetworkManagerVerifierWorkspaceReachabilityAnalysisIntentIpTrafficOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkManagerVerifierWorkspaceReachabilityAnalysisIntentIpTrafficOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := n.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkManagerVerifierWorkspaceReachabilityAnalysisIntentIpTrafficOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := n.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkManagerVerifierWorkspaceReachabilityAnalysisIntentIpTrafficOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := n.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		n,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkManagerVerifierWorkspaceReachabilityAnalysisIntentIpTrafficOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := n.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		n,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkManagerVerifierWorkspaceReachabilityAnalysisIntentIpTrafficOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := n.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		n,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkManagerVerifierWorkspaceReachabilityAnalysisIntentIpTrafficOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := n.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		n,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkManagerVerifierWorkspaceReachabilityAnalysisIntentIpTrafficOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := n.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		n,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkManagerVerifierWorkspaceReachabilityAnalysisIntentIpTrafficOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := n.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		n,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkManagerVerifierWorkspaceReachabilityAnalysisIntentIpTrafficOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := n.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		n,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkManagerVerifierWorkspaceReachabilityAnalysisIntentIpTrafficOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkManagerVerifierWorkspaceReachabilityAnalysisIntentIpTrafficOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := n.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkManagerVerifierWorkspaceReachabilityAnalysisIntentIpTrafficOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := n.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		n,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkManagerVerifierWorkspaceReachabilityAnalysisIntentIpTrafficOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

