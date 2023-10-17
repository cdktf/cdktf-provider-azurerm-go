// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package applicationgateway

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v11/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v11/applicationgateway/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApplicationGatewayHttpListenerOutputReference interface {
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
	CustomErrorConfiguration() ApplicationGatewayHttpListenerCustomErrorConfigurationList
	CustomErrorConfigurationInput() interface{}
	FirewallPolicyId() *string
	SetFirewallPolicyId(val *string)
	FirewallPolicyIdInput() *string
	// Experimental.
	Fqn() *string
	FrontendIpConfigurationId() *string
	FrontendIpConfigurationName() *string
	SetFrontendIpConfigurationName(val *string)
	FrontendIpConfigurationNameInput() *string
	FrontendPortId() *string
	FrontendPortName() *string
	SetFrontendPortName(val *string)
	FrontendPortNameInput() *string
	HostName() *string
	SetHostName(val *string)
	HostNameInput() *string
	HostNames() *[]*string
	SetHostNames(val *[]*string)
	HostNamesInput() *[]*string
	Id() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	SetName(val *string)
	NameInput() *string
	Protocol() *string
	SetProtocol(val *string)
	ProtocolInput() *string
	RequireSni() interface{}
	SetRequireSni(val interface{})
	RequireSniInput() interface{}
	SslCertificateId() *string
	SslCertificateName() *string
	SetSslCertificateName(val *string)
	SslCertificateNameInput() *string
	SslProfileId() *string
	SslProfileName() *string
	SetSslProfileName(val *string)
	SslProfileNameInput() *string
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
	PutCustomErrorConfiguration(value interface{})
	ResetCustomErrorConfiguration()
	ResetFirewallPolicyId()
	ResetHostName()
	ResetHostNames()
	ResetRequireSni()
	ResetSslCertificateName()
	ResetSslProfileName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ApplicationGatewayHttpListenerOutputReference
type jsiiProxy_ApplicationGatewayHttpListenerOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ApplicationGatewayHttpListenerOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayHttpListenerOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayHttpListenerOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayHttpListenerOutputReference) CustomErrorConfiguration() ApplicationGatewayHttpListenerCustomErrorConfigurationList {
	var returns ApplicationGatewayHttpListenerCustomErrorConfigurationList
	_jsii_.Get(
		j,
		"customErrorConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayHttpListenerOutputReference) CustomErrorConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customErrorConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayHttpListenerOutputReference) FirewallPolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallPolicyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayHttpListenerOutputReference) FirewallPolicyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallPolicyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayHttpListenerOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayHttpListenerOutputReference) FrontendIpConfigurationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"frontendIpConfigurationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayHttpListenerOutputReference) FrontendIpConfigurationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"frontendIpConfigurationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayHttpListenerOutputReference) FrontendIpConfigurationNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"frontendIpConfigurationNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayHttpListenerOutputReference) FrontendPortId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"frontendPortId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayHttpListenerOutputReference) FrontendPortName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"frontendPortName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayHttpListenerOutputReference) FrontendPortNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"frontendPortNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayHttpListenerOutputReference) HostName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayHttpListenerOutputReference) HostNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayHttpListenerOutputReference) HostNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"hostNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayHttpListenerOutputReference) HostNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"hostNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayHttpListenerOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayHttpListenerOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayHttpListenerOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayHttpListenerOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayHttpListenerOutputReference) Protocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayHttpListenerOutputReference) ProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayHttpListenerOutputReference) RequireSni() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireSni",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayHttpListenerOutputReference) RequireSniInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireSniInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayHttpListenerOutputReference) SslCertificateId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslCertificateId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayHttpListenerOutputReference) SslCertificateName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslCertificateName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayHttpListenerOutputReference) SslCertificateNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslCertificateNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayHttpListenerOutputReference) SslProfileId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslProfileId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayHttpListenerOutputReference) SslProfileName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslProfileName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayHttpListenerOutputReference) SslProfileNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslProfileNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayHttpListenerOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayHttpListenerOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewApplicationGatewayHttpListenerOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ApplicationGatewayHttpListenerOutputReference {
	_init_.Initialize()

	if err := validateNewApplicationGatewayHttpListenerOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApplicationGatewayHttpListenerOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.applicationGateway.ApplicationGatewayHttpListenerOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewApplicationGatewayHttpListenerOutputReference_Override(a ApplicationGatewayHttpListenerOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.applicationGateway.ApplicationGatewayHttpListenerOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_ApplicationGatewayHttpListenerOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewayHttpListenerOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewayHttpListenerOutputReference)SetFirewallPolicyId(val *string) {
	if err := j.validateSetFirewallPolicyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"firewallPolicyId",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewayHttpListenerOutputReference)SetFrontendIpConfigurationName(val *string) {
	if err := j.validateSetFrontendIpConfigurationNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"frontendIpConfigurationName",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewayHttpListenerOutputReference)SetFrontendPortName(val *string) {
	if err := j.validateSetFrontendPortNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"frontendPortName",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewayHttpListenerOutputReference)SetHostName(val *string) {
	if err := j.validateSetHostNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostName",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewayHttpListenerOutputReference)SetHostNames(val *[]*string) {
	if err := j.validateSetHostNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostNames",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewayHttpListenerOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewayHttpListenerOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewayHttpListenerOutputReference)SetProtocol(val *string) {
	if err := j.validateSetProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protocol",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewayHttpListenerOutputReference)SetRequireSni(val interface{}) {
	if err := j.validateSetRequireSniParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requireSni",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewayHttpListenerOutputReference)SetSslCertificateName(val *string) {
	if err := j.validateSetSslCertificateNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sslCertificateName",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewayHttpListenerOutputReference)SetSslProfileName(val *string) {
	if err := j.validateSetSslProfileNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sslProfileName",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewayHttpListenerOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewayHttpListenerOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_ApplicationGatewayHttpListenerOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationGatewayHttpListenerOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationGatewayHttpListenerOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationGatewayHttpListenerOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationGatewayHttpListenerOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationGatewayHttpListenerOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationGatewayHttpListenerOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationGatewayHttpListenerOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationGatewayHttpListenerOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationGatewayHttpListenerOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationGatewayHttpListenerOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationGatewayHttpListenerOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationGatewayHttpListenerOutputReference) PutCustomErrorConfiguration(value interface{}) {
	if err := a.validatePutCustomErrorConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCustomErrorConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApplicationGatewayHttpListenerOutputReference) ResetCustomErrorConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomErrorConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationGatewayHttpListenerOutputReference) ResetFirewallPolicyId() {
	_jsii_.InvokeVoid(
		a,
		"resetFirewallPolicyId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationGatewayHttpListenerOutputReference) ResetHostName() {
	_jsii_.InvokeVoid(
		a,
		"resetHostName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationGatewayHttpListenerOutputReference) ResetHostNames() {
	_jsii_.InvokeVoid(
		a,
		"resetHostNames",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationGatewayHttpListenerOutputReference) ResetRequireSni() {
	_jsii_.InvokeVoid(
		a,
		"resetRequireSni",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationGatewayHttpListenerOutputReference) ResetSslCertificateName() {
	_jsii_.InvokeVoid(
		a,
		"resetSslCertificateName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationGatewayHttpListenerOutputReference) ResetSslProfileName() {
	_jsii_.InvokeVoid(
		a,
		"resetSslProfileName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationGatewayHttpListenerOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := a.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationGatewayHttpListenerOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

