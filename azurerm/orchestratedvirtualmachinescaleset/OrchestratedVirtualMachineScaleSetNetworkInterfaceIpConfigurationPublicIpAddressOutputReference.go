// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package orchestratedvirtualmachinescaleset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v10/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v10/orchestratedvirtualmachinescaleset/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OrchestratedVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference interface {
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
	DomainNameLabel() *string
	SetDomainNameLabel(val *string)
	DomainNameLabelInput() *string
	// Experimental.
	Fqn() *string
	IdleTimeoutInMinutes() *float64
	SetIdleTimeoutInMinutes(val *float64)
	IdleTimeoutInMinutesInput() *float64
	InternalValue() interface{}
	SetInternalValue(val interface{})
	IpTag() OrchestratedVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressIpTagList
	IpTagInput() interface{}
	Name() *string
	SetName(val *string)
	NameInput() *string
	PublicIpPrefixId() *string
	SetPublicIpPrefixId(val *string)
	PublicIpPrefixIdInput() *string
	SkuName() *string
	SetSkuName(val *string)
	SkuNameInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Version() *string
	SetVersion(val *string)
	VersionInput() *string
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
	PutIpTag(value interface{})
	ResetDomainNameLabel()
	ResetIdleTimeoutInMinutes()
	ResetIpTag()
	ResetPublicIpPrefixId()
	ResetSkuName()
	ResetVersion()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OrchestratedVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference
type jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference) DomainNameLabel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainNameLabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference) DomainNameLabelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainNameLabelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference) IdleTimeoutInMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"idleTimeoutInMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference) IdleTimeoutInMinutesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"idleTimeoutInMinutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference) IpTag() OrchestratedVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressIpTagList {
	var returns OrchestratedVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressIpTagList
	_jsii_.Get(
		j,
		"ipTag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference) IpTagInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipTagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference) PublicIpPrefixId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicIpPrefixId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference) PublicIpPrefixIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicIpPrefixIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference) SkuName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"skuName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference) SkuNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"skuNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference) VersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}


func NewOrchestratedVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) OrchestratedVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference {
	_init_.Initialize()

	if err := validateNewOrchestratedVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.orchestratedVirtualMachineScaleSet.OrchestratedVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewOrchestratedVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference_Override(o OrchestratedVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.orchestratedVirtualMachineScaleSet.OrchestratedVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		o,
	)
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference)SetDomainNameLabel(val *string) {
	if err := j.validateSetDomainNameLabelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainNameLabel",
		val,
	)
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference)SetIdleTimeoutInMinutes(val *float64) {
	if err := j.validateSetIdleTimeoutInMinutesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"idleTimeoutInMinutes",
		val,
	)
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference)SetPublicIpPrefixId(val *string) {
	if err := j.validateSetPublicIpPrefixIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publicIpPrefixId",
		val,
	)
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference)SetSkuName(val *string) {
	if err := j.validateSetSkuNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skuName",
		val,
	)
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference)SetVersion(val *string) {
	if err := j.validateSetVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := o.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := o.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := o.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		o,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := o.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := o.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		o,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := o.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		o,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := o.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		o,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := o.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		o,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := o.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		o,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := o.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference) PutIpTag(value interface{}) {
	if err := o.validatePutIpTagParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putIpTag",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference) ResetDomainNameLabel() {
	_jsii_.InvokeVoid(
		o,
		"resetDomainNameLabel",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference) ResetIdleTimeoutInMinutes() {
	_jsii_.InvokeVoid(
		o,
		"resetIdleTimeoutInMinutes",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference) ResetIpTag() {
	_jsii_.InvokeVoid(
		o,
		"resetIpTag",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference) ResetPublicIpPrefixId() {
	_jsii_.InvokeVoid(
		o,
		"resetPublicIpPrefixId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference) ResetSkuName() {
	_jsii_.InvokeVoid(
		o,
		"resetSkuName",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference) ResetVersion() {
	_jsii_.InvokeVoid(
		o,
		"resetVersion",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := o.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		o,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

