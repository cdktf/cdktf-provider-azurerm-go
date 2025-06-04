// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dashboardgrafana

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/dashboardgrafana/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DashboardGrafanaSmtpOutputReference interface {
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
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	// Experimental.
	Fqn() *string
	FromAddress() *string
	SetFromAddress(val *string)
	FromAddressInput() *string
	FromName() *string
	SetFromName(val *string)
	FromNameInput() *string
	Host() *string
	SetHost(val *string)
	HostInput() *string
	InternalValue() *DashboardGrafanaSmtp
	SetInternalValue(val *DashboardGrafanaSmtp)
	Password() *string
	SetPassword(val *string)
	PasswordInput() *string
	StartTlsPolicy() *string
	SetStartTlsPolicy(val *string)
	StartTlsPolicyInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	User() *string
	SetUser(val *string)
	UserInput() *string
	VerificationSkipEnabled() interface{}
	SetVerificationSkipEnabled(val interface{})
	VerificationSkipEnabledInput() interface{}
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
	ResetEnabled()
	ResetFromName()
	ResetVerificationSkipEnabled()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DashboardGrafanaSmtpOutputReference
type jsiiProxy_DashboardGrafanaSmtpOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DashboardGrafanaSmtpOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardGrafanaSmtpOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardGrafanaSmtpOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardGrafanaSmtpOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardGrafanaSmtpOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardGrafanaSmtpOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardGrafanaSmtpOutputReference) FromAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fromAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardGrafanaSmtpOutputReference) FromAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fromAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardGrafanaSmtpOutputReference) FromName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fromName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardGrafanaSmtpOutputReference) FromNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fromNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardGrafanaSmtpOutputReference) Host() *string {
	var returns *string
	_jsii_.Get(
		j,
		"host",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardGrafanaSmtpOutputReference) HostInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardGrafanaSmtpOutputReference) InternalValue() *DashboardGrafanaSmtp {
	var returns *DashboardGrafanaSmtp
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardGrafanaSmtpOutputReference) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardGrafanaSmtpOutputReference) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardGrafanaSmtpOutputReference) StartTlsPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startTlsPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardGrafanaSmtpOutputReference) StartTlsPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startTlsPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardGrafanaSmtpOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardGrafanaSmtpOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardGrafanaSmtpOutputReference) User() *string {
	var returns *string
	_jsii_.Get(
		j,
		"user",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardGrafanaSmtpOutputReference) UserInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardGrafanaSmtpOutputReference) VerificationSkipEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"verificationSkipEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardGrafanaSmtpOutputReference) VerificationSkipEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"verificationSkipEnabledInput",
		&returns,
	)
	return returns
}


func NewDashboardGrafanaSmtpOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DashboardGrafanaSmtpOutputReference {
	_init_.Initialize()

	if err := validateNewDashboardGrafanaSmtpOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DashboardGrafanaSmtpOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.dashboardGrafana.DashboardGrafanaSmtpOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDashboardGrafanaSmtpOutputReference_Override(d DashboardGrafanaSmtpOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.dashboardGrafana.DashboardGrafanaSmtpOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DashboardGrafanaSmtpOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DashboardGrafanaSmtpOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DashboardGrafanaSmtpOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_DashboardGrafanaSmtpOutputReference)SetFromAddress(val *string) {
	if err := j.validateSetFromAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fromAddress",
		val,
	)
}

func (j *jsiiProxy_DashboardGrafanaSmtpOutputReference)SetFromName(val *string) {
	if err := j.validateSetFromNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fromName",
		val,
	)
}

func (j *jsiiProxy_DashboardGrafanaSmtpOutputReference)SetHost(val *string) {
	if err := j.validateSetHostParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"host",
		val,
	)
}

func (j *jsiiProxy_DashboardGrafanaSmtpOutputReference)SetInternalValue(val *DashboardGrafanaSmtp) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DashboardGrafanaSmtpOutputReference)SetPassword(val *string) {
	if err := j.validateSetPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_DashboardGrafanaSmtpOutputReference)SetStartTlsPolicy(val *string) {
	if err := j.validateSetStartTlsPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startTlsPolicy",
		val,
	)
}

func (j *jsiiProxy_DashboardGrafanaSmtpOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DashboardGrafanaSmtpOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DashboardGrafanaSmtpOutputReference)SetUser(val *string) {
	if err := j.validateSetUserParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"user",
		val,
	)
}

func (j *jsiiProxy_DashboardGrafanaSmtpOutputReference)SetVerificationSkipEnabled(val interface{}) {
	if err := j.validateSetVerificationSkipEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"verificationSkipEnabled",
		val,
	)
}

func (d *jsiiProxy_DashboardGrafanaSmtpOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardGrafanaSmtpOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardGrafanaSmtpOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardGrafanaSmtpOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardGrafanaSmtpOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardGrafanaSmtpOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardGrafanaSmtpOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardGrafanaSmtpOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardGrafanaSmtpOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardGrafanaSmtpOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardGrafanaSmtpOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardGrafanaSmtpOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardGrafanaSmtpOutputReference) ResetEnabled() {
	_jsii_.InvokeVoid(
		d,
		"resetEnabled",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardGrafanaSmtpOutputReference) ResetFromName() {
	_jsii_.InvokeVoid(
		d,
		"resetFromName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardGrafanaSmtpOutputReference) ResetVerificationSkipEnabled() {
	_jsii_.InvokeVoid(
		d,
		"resetVerificationSkipEnabled",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardGrafanaSmtpOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardGrafanaSmtpOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

