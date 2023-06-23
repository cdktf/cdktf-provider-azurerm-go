package springcloudcustomizedaccelerator

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v9/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v9/springcloudcustomizedaccelerator/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SpringCloudCustomizedAcceleratorGitRepositoryOutputReference interface {
	cdktf.ComplexObject
	BasicAuth() SpringCloudCustomizedAcceleratorGitRepositoryBasicAuthOutputReference
	BasicAuthInput() *SpringCloudCustomizedAcceleratorGitRepositoryBasicAuth
	Branch() *string
	SetBranch(val *string)
	BranchInput() *string
	CaCertificateId() *string
	SetCaCertificateId(val *string)
	CaCertificateIdInput() *string
	Commit() *string
	SetCommit(val *string)
	CommitInput() *string
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
	GitTag() *string
	SetGitTag(val *string)
	GitTagInput() *string
	InternalValue() *SpringCloudCustomizedAcceleratorGitRepository
	SetInternalValue(val *SpringCloudCustomizedAcceleratorGitRepository)
	IntervalInSeconds() *float64
	SetIntervalInSeconds(val *float64)
	IntervalInSecondsInput() *float64
	SshAuth() SpringCloudCustomizedAcceleratorGitRepositorySshAuthOutputReference
	SshAuthInput() *SpringCloudCustomizedAcceleratorGitRepositorySshAuth
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Url() *string
	SetUrl(val *string)
	UrlInput() *string
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
	PutBasicAuth(value *SpringCloudCustomizedAcceleratorGitRepositoryBasicAuth)
	PutSshAuth(value *SpringCloudCustomizedAcceleratorGitRepositorySshAuth)
	ResetBasicAuth()
	ResetBranch()
	ResetCaCertificateId()
	ResetCommit()
	ResetGitTag()
	ResetIntervalInSeconds()
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

// The jsii proxy struct for SpringCloudCustomizedAcceleratorGitRepositoryOutputReference
type jsiiProxy_SpringCloudCustomizedAcceleratorGitRepositoryOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SpringCloudCustomizedAcceleratorGitRepositoryOutputReference) BasicAuth() SpringCloudCustomizedAcceleratorGitRepositoryBasicAuthOutputReference {
	var returns SpringCloudCustomizedAcceleratorGitRepositoryBasicAuthOutputReference
	_jsii_.Get(
		j,
		"basicAuth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudCustomizedAcceleratorGitRepositoryOutputReference) BasicAuthInput() *SpringCloudCustomizedAcceleratorGitRepositoryBasicAuth {
	var returns *SpringCloudCustomizedAcceleratorGitRepositoryBasicAuth
	_jsii_.Get(
		j,
		"basicAuthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudCustomizedAcceleratorGitRepositoryOutputReference) Branch() *string {
	var returns *string
	_jsii_.Get(
		j,
		"branch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudCustomizedAcceleratorGitRepositoryOutputReference) BranchInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"branchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudCustomizedAcceleratorGitRepositoryOutputReference) CaCertificateId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caCertificateId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudCustomizedAcceleratorGitRepositoryOutputReference) CaCertificateIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caCertificateIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudCustomizedAcceleratorGitRepositoryOutputReference) Commit() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudCustomizedAcceleratorGitRepositoryOutputReference) CommitInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudCustomizedAcceleratorGitRepositoryOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudCustomizedAcceleratorGitRepositoryOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudCustomizedAcceleratorGitRepositoryOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudCustomizedAcceleratorGitRepositoryOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudCustomizedAcceleratorGitRepositoryOutputReference) GitTag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gitTag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudCustomizedAcceleratorGitRepositoryOutputReference) GitTagInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gitTagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudCustomizedAcceleratorGitRepositoryOutputReference) InternalValue() *SpringCloudCustomizedAcceleratorGitRepository {
	var returns *SpringCloudCustomizedAcceleratorGitRepository
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudCustomizedAcceleratorGitRepositoryOutputReference) IntervalInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"intervalInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudCustomizedAcceleratorGitRepositoryOutputReference) IntervalInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"intervalInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudCustomizedAcceleratorGitRepositoryOutputReference) SshAuth() SpringCloudCustomizedAcceleratorGitRepositorySshAuthOutputReference {
	var returns SpringCloudCustomizedAcceleratorGitRepositorySshAuthOutputReference
	_jsii_.Get(
		j,
		"sshAuth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudCustomizedAcceleratorGitRepositoryOutputReference) SshAuthInput() *SpringCloudCustomizedAcceleratorGitRepositorySshAuth {
	var returns *SpringCloudCustomizedAcceleratorGitRepositorySshAuth
	_jsii_.Get(
		j,
		"sshAuthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudCustomizedAcceleratorGitRepositoryOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudCustomizedAcceleratorGitRepositoryOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudCustomizedAcceleratorGitRepositoryOutputReference) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudCustomizedAcceleratorGitRepositoryOutputReference) UrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"urlInput",
		&returns,
	)
	return returns
}


func NewSpringCloudCustomizedAcceleratorGitRepositoryOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) SpringCloudCustomizedAcceleratorGitRepositoryOutputReference {
	_init_.Initialize()

	if err := validateNewSpringCloudCustomizedAcceleratorGitRepositoryOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SpringCloudCustomizedAcceleratorGitRepositoryOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.springCloudCustomizedAccelerator.SpringCloudCustomizedAcceleratorGitRepositoryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSpringCloudCustomizedAcceleratorGitRepositoryOutputReference_Override(s SpringCloudCustomizedAcceleratorGitRepositoryOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.springCloudCustomizedAccelerator.SpringCloudCustomizedAcceleratorGitRepositoryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SpringCloudCustomizedAcceleratorGitRepositoryOutputReference)SetBranch(val *string) {
	if err := j.validateSetBranchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"branch",
		val,
	)
}

func (j *jsiiProxy_SpringCloudCustomizedAcceleratorGitRepositoryOutputReference)SetCaCertificateId(val *string) {
	if err := j.validateSetCaCertificateIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"caCertificateId",
		val,
	)
}

func (j *jsiiProxy_SpringCloudCustomizedAcceleratorGitRepositoryOutputReference)SetCommit(val *string) {
	if err := j.validateSetCommitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"commit",
		val,
	)
}

func (j *jsiiProxy_SpringCloudCustomizedAcceleratorGitRepositoryOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SpringCloudCustomizedAcceleratorGitRepositoryOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SpringCloudCustomizedAcceleratorGitRepositoryOutputReference)SetGitTag(val *string) {
	if err := j.validateSetGitTagParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gitTag",
		val,
	)
}

func (j *jsiiProxy_SpringCloudCustomizedAcceleratorGitRepositoryOutputReference)SetInternalValue(val *SpringCloudCustomizedAcceleratorGitRepository) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SpringCloudCustomizedAcceleratorGitRepositoryOutputReference)SetIntervalInSeconds(val *float64) {
	if err := j.validateSetIntervalInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"intervalInSeconds",
		val,
	)
}

func (j *jsiiProxy_SpringCloudCustomizedAcceleratorGitRepositoryOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SpringCloudCustomizedAcceleratorGitRepositoryOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_SpringCloudCustomizedAcceleratorGitRepositoryOutputReference)SetUrl(val *string) {
	if err := j.validateSetUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"url",
		val,
	)
}

func (s *jsiiProxy_SpringCloudCustomizedAcceleratorGitRepositoryOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudCustomizedAcceleratorGitRepositoryOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SpringCloudCustomizedAcceleratorGitRepositoryOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SpringCloudCustomizedAcceleratorGitRepositoryOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SpringCloudCustomizedAcceleratorGitRepositoryOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SpringCloudCustomizedAcceleratorGitRepositoryOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SpringCloudCustomizedAcceleratorGitRepositoryOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SpringCloudCustomizedAcceleratorGitRepositoryOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SpringCloudCustomizedAcceleratorGitRepositoryOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SpringCloudCustomizedAcceleratorGitRepositoryOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SpringCloudCustomizedAcceleratorGitRepositoryOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudCustomizedAcceleratorGitRepositoryOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SpringCloudCustomizedAcceleratorGitRepositoryOutputReference) PutBasicAuth(value *SpringCloudCustomizedAcceleratorGitRepositoryBasicAuth) {
	if err := s.validatePutBasicAuthParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putBasicAuth",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SpringCloudCustomizedAcceleratorGitRepositoryOutputReference) PutSshAuth(value *SpringCloudCustomizedAcceleratorGitRepositorySshAuth) {
	if err := s.validatePutSshAuthParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putSshAuth",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SpringCloudCustomizedAcceleratorGitRepositoryOutputReference) ResetBasicAuth() {
	_jsii_.InvokeVoid(
		s,
		"resetBasicAuth",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpringCloudCustomizedAcceleratorGitRepositoryOutputReference) ResetBranch() {
	_jsii_.InvokeVoid(
		s,
		"resetBranch",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpringCloudCustomizedAcceleratorGitRepositoryOutputReference) ResetCaCertificateId() {
	_jsii_.InvokeVoid(
		s,
		"resetCaCertificateId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpringCloudCustomizedAcceleratorGitRepositoryOutputReference) ResetCommit() {
	_jsii_.InvokeVoid(
		s,
		"resetCommit",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpringCloudCustomizedAcceleratorGitRepositoryOutputReference) ResetGitTag() {
	_jsii_.InvokeVoid(
		s,
		"resetGitTag",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpringCloudCustomizedAcceleratorGitRepositoryOutputReference) ResetIntervalInSeconds() {
	_jsii_.InvokeVoid(
		s,
		"resetIntervalInSeconds",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpringCloudCustomizedAcceleratorGitRepositoryOutputReference) ResetSshAuth() {
	_jsii_.InvokeVoid(
		s,
		"resetSshAuth",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpringCloudCustomizedAcceleratorGitRepositoryOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_SpringCloudCustomizedAcceleratorGitRepositoryOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

