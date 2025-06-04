// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package springcloudservice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/springcloudservice/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SpringCloudServiceConfigServerGitSettingRepositorySshAuthOutputReference interface {
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
	// Experimental.
	Fqn() *string
	HostKey() *string
	SetHostKey(val *string)
	HostKeyAlgorithm() *string
	SetHostKeyAlgorithm(val *string)
	HostKeyAlgorithmInput() *string
	HostKeyInput() *string
	InternalValue() *SpringCloudServiceConfigServerGitSettingRepositorySshAuth
	SetInternalValue(val *SpringCloudServiceConfigServerGitSettingRepositorySshAuth)
	PrivateKey() *string
	SetPrivateKey(val *string)
	PrivateKeyInput() *string
	StrictHostKeyCheckingEnabled() interface{}
	SetStrictHostKeyCheckingEnabled(val interface{})
	StrictHostKeyCheckingEnabledInput() interface{}
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
	ResetHostKey()
	ResetHostKeyAlgorithm()
	ResetStrictHostKeyCheckingEnabled()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SpringCloudServiceConfigServerGitSettingRepositorySshAuthOutputReference
type jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositorySshAuthOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositorySshAuthOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositorySshAuthOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositorySshAuthOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositorySshAuthOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositorySshAuthOutputReference) HostKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositorySshAuthOutputReference) HostKeyAlgorithm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostKeyAlgorithm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositorySshAuthOutputReference) HostKeyAlgorithmInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostKeyAlgorithmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositorySshAuthOutputReference) HostKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositorySshAuthOutputReference) InternalValue() *SpringCloudServiceConfigServerGitSettingRepositorySshAuth {
	var returns *SpringCloudServiceConfigServerGitSettingRepositorySshAuth
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositorySshAuthOutputReference) PrivateKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositorySshAuthOutputReference) PrivateKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositorySshAuthOutputReference) StrictHostKeyCheckingEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"strictHostKeyCheckingEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositorySshAuthOutputReference) StrictHostKeyCheckingEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"strictHostKeyCheckingEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositorySshAuthOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositorySshAuthOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSpringCloudServiceConfigServerGitSettingRepositorySshAuthOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) SpringCloudServiceConfigServerGitSettingRepositorySshAuthOutputReference {
	_init_.Initialize()

	if err := validateNewSpringCloudServiceConfigServerGitSettingRepositorySshAuthOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositorySshAuthOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.springCloudService.SpringCloudServiceConfigServerGitSettingRepositorySshAuthOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSpringCloudServiceConfigServerGitSettingRepositorySshAuthOutputReference_Override(s SpringCloudServiceConfigServerGitSettingRepositorySshAuthOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.springCloudService.SpringCloudServiceConfigServerGitSettingRepositorySshAuthOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositorySshAuthOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositorySshAuthOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositorySshAuthOutputReference)SetHostKey(val *string) {
	if err := j.validateSetHostKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostKey",
		val,
	)
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositorySshAuthOutputReference)SetHostKeyAlgorithm(val *string) {
	if err := j.validateSetHostKeyAlgorithmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostKeyAlgorithm",
		val,
	)
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositorySshAuthOutputReference)SetInternalValue(val *SpringCloudServiceConfigServerGitSettingRepositorySshAuth) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositorySshAuthOutputReference)SetPrivateKey(val *string) {
	if err := j.validateSetPrivateKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateKey",
		val,
	)
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositorySshAuthOutputReference)SetStrictHostKeyCheckingEnabled(val interface{}) {
	if err := j.validateSetStrictHostKeyCheckingEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"strictHostKeyCheckingEnabled",
		val,
	)
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositorySshAuthOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositorySshAuthOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositorySshAuthOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositorySshAuthOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositorySshAuthOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositorySshAuthOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositorySshAuthOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositorySshAuthOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositorySshAuthOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositorySshAuthOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositorySshAuthOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositorySshAuthOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositorySshAuthOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositorySshAuthOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositorySshAuthOutputReference) ResetHostKey() {
	_jsii_.InvokeVoid(
		s,
		"resetHostKey",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositorySshAuthOutputReference) ResetHostKeyAlgorithm() {
	_jsii_.InvokeVoid(
		s,
		"resetHostKeyAlgorithm",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositorySshAuthOutputReference) ResetStrictHostKeyCheckingEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetStrictHostKeyCheckingEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositorySshAuthOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositorySshAuthOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

