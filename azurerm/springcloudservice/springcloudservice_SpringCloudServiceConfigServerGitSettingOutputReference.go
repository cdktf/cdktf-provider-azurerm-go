package springcloudservice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v3/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v3/springcloudservice/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SpringCloudServiceConfigServerGitSettingOutputReference interface {
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
	HttpBasicAuth() SpringCloudServiceConfigServerGitSettingHttpBasicAuthOutputReference
	HttpBasicAuthInput() *SpringCloudServiceConfigServerGitSettingHttpBasicAuth
	InternalValue() *SpringCloudServiceConfigServerGitSetting
	SetInternalValue(val *SpringCloudServiceConfigServerGitSetting)
	Label() *string
	SetLabel(val *string)
	LabelInput() *string
	Repository() SpringCloudServiceConfigServerGitSettingRepositoryList
	RepositoryInput() interface{}
	SearchPaths() *[]*string
	SetSearchPaths(val *[]*string)
	SearchPathsInput() *[]*string
	SshAuth() SpringCloudServiceConfigServerGitSettingSshAuthOutputReference
	SshAuthInput() *SpringCloudServiceConfigServerGitSettingSshAuth
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
	InterpolationForAttribute(property *string) cdktf.IResolvable
	PutHttpBasicAuth(value *SpringCloudServiceConfigServerGitSettingHttpBasicAuth)
	PutRepository(value interface{})
	PutSshAuth(value *SpringCloudServiceConfigServerGitSettingSshAuth)
	ResetHttpBasicAuth()
	ResetLabel()
	ResetRepository()
	ResetSearchPaths()
	ResetSshAuth()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SpringCloudServiceConfigServerGitSettingOutputReference
type jsiiProxy_SpringCloudServiceConfigServerGitSettingOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingOutputReference) HttpBasicAuth() SpringCloudServiceConfigServerGitSettingHttpBasicAuthOutputReference {
	var returns SpringCloudServiceConfigServerGitSettingHttpBasicAuthOutputReference
	_jsii_.Get(
		j,
		"httpBasicAuth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingOutputReference) HttpBasicAuthInput() *SpringCloudServiceConfigServerGitSettingHttpBasicAuth {
	var returns *SpringCloudServiceConfigServerGitSettingHttpBasicAuth
	_jsii_.Get(
		j,
		"httpBasicAuthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingOutputReference) InternalValue() *SpringCloudServiceConfigServerGitSetting {
	var returns *SpringCloudServiceConfigServerGitSetting
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingOutputReference) Label() *string {
	var returns *string
	_jsii_.Get(
		j,
		"label",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingOutputReference) LabelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"labelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingOutputReference) Repository() SpringCloudServiceConfigServerGitSettingRepositoryList {
	var returns SpringCloudServiceConfigServerGitSettingRepositoryList
	_jsii_.Get(
		j,
		"repository",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingOutputReference) RepositoryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"repositoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingOutputReference) SearchPaths() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"searchPaths",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingOutputReference) SearchPathsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"searchPathsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingOutputReference) SshAuth() SpringCloudServiceConfigServerGitSettingSshAuthOutputReference {
	var returns SpringCloudServiceConfigServerGitSettingSshAuthOutputReference
	_jsii_.Get(
		j,
		"sshAuth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingOutputReference) SshAuthInput() *SpringCloudServiceConfigServerGitSettingSshAuth {
	var returns *SpringCloudServiceConfigServerGitSettingSshAuth
	_jsii_.Get(
		j,
		"sshAuthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingOutputReference) Uri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingOutputReference) UriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uriInput",
		&returns,
	)
	return returns
}


func NewSpringCloudServiceConfigServerGitSettingOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) SpringCloudServiceConfigServerGitSettingOutputReference {
	_init_.Initialize()

	if err := validateNewSpringCloudServiceConfigServerGitSettingOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SpringCloudServiceConfigServerGitSettingOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.springCloudService.SpringCloudServiceConfigServerGitSettingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSpringCloudServiceConfigServerGitSettingOutputReference_Override(s SpringCloudServiceConfigServerGitSettingOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.springCloudService.SpringCloudServiceConfigServerGitSettingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingOutputReference)SetInternalValue(val *SpringCloudServiceConfigServerGitSetting) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingOutputReference)SetLabel(val *string) {
	if err := j.validateSetLabelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"label",
		val,
	)
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingOutputReference)SetSearchPaths(val *[]*string) {
	if err := j.validateSetSearchPathsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"searchPaths",
		val,
	)
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_SpringCloudServiceConfigServerGitSettingOutputReference)SetUri(val *string) {
	if err := j.validateSetUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"uri",
		val,
	)
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingOutputReference) PutHttpBasicAuth(value *SpringCloudServiceConfigServerGitSettingHttpBasicAuth) {
	if err := s.validatePutHttpBasicAuthParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putHttpBasicAuth",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingOutputReference) PutRepository(value interface{}) {
	if err := s.validatePutRepositoryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putRepository",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingOutputReference) PutSshAuth(value *SpringCloudServiceConfigServerGitSettingSshAuth) {
	if err := s.validatePutSshAuthParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putSshAuth",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingOutputReference) ResetHttpBasicAuth() {
	_jsii_.InvokeVoid(
		s,
		"resetHttpBasicAuth",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingOutputReference) ResetLabel() {
	_jsii_.InvokeVoid(
		s,
		"resetLabel",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingOutputReference) ResetRepository() {
	_jsii_.InvokeVoid(
		s,
		"resetRepository",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingOutputReference) ResetSearchPaths() {
	_jsii_.InvokeVoid(
		s,
		"resetSearchPaths",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingOutputReference) ResetSshAuth() {
	_jsii_.InvokeVoid(
		s,
		"resetSshAuth",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_SpringCloudServiceConfigServerGitSettingOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

