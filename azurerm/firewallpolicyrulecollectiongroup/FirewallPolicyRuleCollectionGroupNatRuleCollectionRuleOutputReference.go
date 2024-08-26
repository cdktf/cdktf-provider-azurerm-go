// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package firewallpolicyrulecollectiongroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/firewallpolicyrulecollectiongroup/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference interface {
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
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DestinationAddress() *string
	SetDestinationAddress(val *string)
	DestinationAddressInput() *string
	DestinationPorts() *[]*string
	SetDestinationPorts(val *[]*string)
	DestinationPortsInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	SetName(val *string)
	NameInput() *string
	Protocols() *[]*string
	SetProtocols(val *[]*string)
	ProtocolsInput() *[]*string
	SourceAddresses() *[]*string
	SetSourceAddresses(val *[]*string)
	SourceAddressesInput() *[]*string
	SourceIpGroups() *[]*string
	SetSourceIpGroups(val *[]*string)
	SourceIpGroupsInput() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TranslatedAddress() *string
	SetTranslatedAddress(val *string)
	TranslatedAddressInput() *string
	TranslatedFqdn() *string
	SetTranslatedFqdn(val *string)
	TranslatedFqdnInput() *string
	TranslatedPort() *float64
	SetTranslatedPort(val *float64)
	TranslatedPortInput() *float64
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
	ResetDescription()
	ResetDestinationAddress()
	ResetDestinationPorts()
	ResetSourceAddresses()
	ResetSourceIpGroups()
	ResetTranslatedAddress()
	ResetTranslatedFqdn()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference
type jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference) DestinationAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference) DestinationAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference) DestinationPorts() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destinationPorts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference) DestinationPortsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destinationPortsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference) Protocols() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"protocols",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference) ProtocolsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"protocolsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference) SourceAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourceAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference) SourceAddressesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourceAddressesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference) SourceIpGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourceIpGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference) SourceIpGroupsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourceIpGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference) TranslatedAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"translatedAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference) TranslatedAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"translatedAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference) TranslatedFqdn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"translatedFqdn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference) TranslatedFqdnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"translatedFqdnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference) TranslatedPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"translatedPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference) TranslatedPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"translatedPortInput",
		&returns,
	)
	return returns
}


func NewFirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference {
	_init_.Initialize()

	if err := validateNewFirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.firewallPolicyRuleCollectionGroup.FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewFirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference_Override(f FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.firewallPolicyRuleCollectionGroup.FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		f,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference)SetDestinationAddress(val *string) {
	if err := j.validateSetDestinationAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationAddress",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference)SetDestinationPorts(val *[]*string) {
	if err := j.validateSetDestinationPortsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationPorts",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference)SetProtocols(val *[]*string) {
	if err := j.validateSetProtocolsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protocols",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference)SetSourceAddresses(val *[]*string) {
	if err := j.validateSetSourceAddressesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceAddresses",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference)SetSourceIpGroups(val *[]*string) {
	if err := j.validateSetSourceIpGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceIpGroups",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference)SetTranslatedAddress(val *string) {
	if err := j.validateSetTranslatedAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"translatedAddress",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference)SetTranslatedFqdn(val *string) {
	if err := j.validateSetTranslatedFqdnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"translatedFqdn",
		val,
	)
}

func (j *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference)SetTranslatedPort(val *float64) {
	if err := j.validateSetTranslatedPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"translatedPort",
		val,
	)
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := f.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := f.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := f.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := f.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := f.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := f.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := f.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := f.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := f.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := f.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		f,
		"resetDescription",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference) ResetDestinationAddress() {
	_jsii_.InvokeVoid(
		f,
		"resetDestinationAddress",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference) ResetDestinationPorts() {
	_jsii_.InvokeVoid(
		f,
		"resetDestinationPorts",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference) ResetSourceAddresses() {
	_jsii_.InvokeVoid(
		f,
		"resetSourceAddresses",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference) ResetSourceIpGroups() {
	_jsii_.InvokeVoid(
		f,
		"resetSourceIpGroups",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference) ResetTranslatedAddress() {
	_jsii_.InvokeVoid(
		f,
		"resetTranslatedAddress",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference) ResetTranslatedFqdn() {
	_jsii_.InvokeVoid(
		f,
		"resetTranslatedFqdn",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := f.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		f,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallPolicyRuleCollectionGroupNatRuleCollectionRuleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

