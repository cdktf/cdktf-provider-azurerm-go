// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package springcloudservice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/springcloudservice/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SpringCloudServiceConfigServerGitSettingRepositoryOutputReference interface {
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
	HttpBasicAuth() SpringCloudServiceConfigServerGitSettingRepositoryHttpBasicAuthOutputReference
	HttpBasicAuthInput() *SpringCloudServiceConfigServerGitSettingRepositoryHttpBasicAuth
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Label() *string
	SetLabel(val *string)
	LabelInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	Pattern() *[]*string
	SetPattern(val *[]*string)
	PatternInput() *[]*string
	SearchPaths() *[]*string
	SetSearchPaths(val *[]*string)
	SearchPathsInput() *[]*string
	SshAuth() SpringCloudServiceConfigServerGitSettingRepositorySshAuthOutputReference
	SshAuthInput() *SpringCloudServiceConfigServerGitSettingRepositorySshAuth
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Uri() *string
	SetUri(val *string)
	UriInput() *string
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	PutHttpBasicAuth(value *SpringCloudServiceConfigServerGitSettingRepositoryHttpBasicAuth)
	PutSshAuth(value *SpringCloudServiceConfigServerGitSettingRepositorySshAuth)
	ResetHttpBasicAuth()
	ResetLabel()
	ResetPattern()
	ResetSearchPaths()
	ResetSshAuth()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SpringCloudServiceConfigServerGitSettingRepositoryOutputReference
type jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryOutputReference) HttpBasicAuth() SpringCloudServiceConfigServerGitSettingRepositoryHttpBasicAuthOutputReference {
	var returns SpringCloudServiceConfigServerGitSettingRepositoryHttpBasicAuthOutputReference
	_jsii_.Get(
		j,
		"httpBasicAuth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryOutputReference) HttpBasicAuthInput() *SpringCloudServiceConfigServerGitSettingRepositoryHttpBasicAuth {
	var returns *SpringCloudServiceConfigServerGitSettingRepositoryHttpBasicAuth
	_jsii_.Get(
		j,
		"httpBasicAuthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryOutputReference) Label() *string {
	var returns *string
	_jsii_.Get(
		j,
		"label",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryOutputReference) LabelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"labelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryOutputReference) Pattern() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"pattern",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryOutputReference) PatternInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"patternInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryOutputReference) SearchPaths() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"searchPaths",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryOutputReference) SearchPathsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"searchPathsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryOutputReference) SshAuth() SpringCloudServiceConfigServerGitSettingRepositorySshAuthOutputReference {
	var returns SpringCloudServiceConfigServerGitSettingRepositorySshAuthOutputReference
	_jsii_.Get(
		j,
		"sshAuth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryOutputReference) SshAuthInput() *SpringCloudServiceConfigServerGitSettingRepositorySshAuth {
	var returns *SpringCloudServiceConfigServerGitSettingRepositorySshAuth
	_jsii_.Get(
		j,
		"sshAuthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryOutputReference) Uri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryOutputReference) UriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uriInput",
		&returns,
	)
	return returns
}


func NewSpringCloudServiceConfigServerGitSettingRepositoryOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) SpringCloudServiceConfigServerGitSettingRepositoryOutputReference {
	_init_.Initialize()

	if err := validateNewSpringCloudServiceConfigServerGitSettingRepositoryOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.springCloudService.SpringCloudServiceConfigServerGitSettingRepositoryOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewSpringCloudServiceConfigServerGitSettingRepositoryOutputReference_Override(s SpringCloudServiceConfigServerGitSettingRepositoryOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.springCloudService.SpringCloudServiceConfigServerGitSettingRepositoryOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryOutputReference)SetLabel(val *string) {
	if err := j.validateSetLabelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"label",
		val,
	)
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryOutputReference)SetPattern(val *[]*string) {
	if err := j.validateSetPatternParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pattern",
		val,
	)
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryOutputReference)SetSearchPaths(val *[]*string) {
	if err := j.validateSetSearchPathsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"searchPaths",
		val,
	)
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryOutputReference)SetUri(val *string) {
	if err := j.validateSetUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"uri",
		val,
	)
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryOutputReference) PutHttpBasicAuth(value *SpringCloudServiceConfigServerGitSettingRepositoryHttpBasicAuth) {
	if err := s.validatePutHttpBasicAuthParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putHttpBasicAuth",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryOutputReference) PutSshAuth(value *SpringCloudServiceConfigServerGitSettingRepositorySshAuth) {
	if err := s.validatePutSshAuthParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putSshAuth",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryOutputReference) ResetHttpBasicAuth() {
	_jsii_.InvokeVoid(
		s,
		"resetHttpBasicAuth",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryOutputReference) ResetLabel() {
	_jsii_.InvokeVoid(
		s,
		"resetLabel",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryOutputReference) ResetPattern() {
	_jsii_.InvokeVoid(
		s,
		"resetPattern",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryOutputReference) ResetSearchPaths() {
	_jsii_.InvokeVoid(
		s,
		"resetSearchPaths",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryOutputReference) ResetSshAuth() {
	_jsii_.InvokeVoid(
		s,
		"resetSshAuth",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := s.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingRepositoryOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

