package synapseworkspace

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v7/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v7/synapseworkspace/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SynapseWorkspaceGithubRepoOutputReference interface {
	cdktf.ComplexObject
	AccountName() *string
	SetAccountName(val *string)
	AccountNameInput() *string
	BranchName() *string
	SetBranchName(val *string)
	BranchNameInput() *string
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
	GitUrl() *string
	SetGitUrl(val *string)
	GitUrlInput() *string
	InternalValue() *SynapseWorkspaceGithubRepo
	SetInternalValue(val *SynapseWorkspaceGithubRepo)
	LastCommitId() *string
	SetLastCommitId(val *string)
	LastCommitIdInput() *string
	RepositoryName() *string
	SetRepositoryName(val *string)
	RepositoryNameInput() *string
	RootFolder() *string
	SetRootFolder(val *string)
	RootFolderInput() *string
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
	ResetGitUrl()
	ResetLastCommitId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SynapseWorkspaceGithubRepoOutputReference
type jsiiProxy_SynapseWorkspaceGithubRepoOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SynapseWorkspaceGithubRepoOutputReference) AccountName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceGithubRepoOutputReference) AccountNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceGithubRepoOutputReference) BranchName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"branchName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceGithubRepoOutputReference) BranchNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"branchNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceGithubRepoOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceGithubRepoOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceGithubRepoOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceGithubRepoOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceGithubRepoOutputReference) GitUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gitUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceGithubRepoOutputReference) GitUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gitUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceGithubRepoOutputReference) InternalValue() *SynapseWorkspaceGithubRepo {
	var returns *SynapseWorkspaceGithubRepo
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceGithubRepoOutputReference) LastCommitId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastCommitId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceGithubRepoOutputReference) LastCommitIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastCommitIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceGithubRepoOutputReference) RepositoryName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceGithubRepoOutputReference) RepositoryNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceGithubRepoOutputReference) RootFolder() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rootFolder",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceGithubRepoOutputReference) RootFolderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rootFolderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceGithubRepoOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseWorkspaceGithubRepoOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSynapseWorkspaceGithubRepoOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) SynapseWorkspaceGithubRepoOutputReference {
	_init_.Initialize()

	if err := validateNewSynapseWorkspaceGithubRepoOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SynapseWorkspaceGithubRepoOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.synapseWorkspace.SynapseWorkspaceGithubRepoOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSynapseWorkspaceGithubRepoOutputReference_Override(s SynapseWorkspaceGithubRepoOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.synapseWorkspace.SynapseWorkspaceGithubRepoOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SynapseWorkspaceGithubRepoOutputReference)SetAccountName(val *string) {
	if err := j.validateSetAccountNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountName",
		val,
	)
}

func (j *jsiiProxy_SynapseWorkspaceGithubRepoOutputReference)SetBranchName(val *string) {
	if err := j.validateSetBranchNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"branchName",
		val,
	)
}

func (j *jsiiProxy_SynapseWorkspaceGithubRepoOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SynapseWorkspaceGithubRepoOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SynapseWorkspaceGithubRepoOutputReference)SetGitUrl(val *string) {
	if err := j.validateSetGitUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gitUrl",
		val,
	)
}

func (j *jsiiProxy_SynapseWorkspaceGithubRepoOutputReference)SetInternalValue(val *SynapseWorkspaceGithubRepo) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SynapseWorkspaceGithubRepoOutputReference)SetLastCommitId(val *string) {
	if err := j.validateSetLastCommitIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lastCommitId",
		val,
	)
}

func (j *jsiiProxy_SynapseWorkspaceGithubRepoOutputReference)SetRepositoryName(val *string) {
	if err := j.validateSetRepositoryNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"repositoryName",
		val,
	)
}

func (j *jsiiProxy_SynapseWorkspaceGithubRepoOutputReference)SetRootFolder(val *string) {
	if err := j.validateSetRootFolderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rootFolder",
		val,
	)
}

func (j *jsiiProxy_SynapseWorkspaceGithubRepoOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SynapseWorkspaceGithubRepoOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SynapseWorkspaceGithubRepoOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseWorkspaceGithubRepoOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SynapseWorkspaceGithubRepoOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SynapseWorkspaceGithubRepoOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SynapseWorkspaceGithubRepoOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SynapseWorkspaceGithubRepoOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SynapseWorkspaceGithubRepoOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SynapseWorkspaceGithubRepoOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SynapseWorkspaceGithubRepoOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SynapseWorkspaceGithubRepoOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SynapseWorkspaceGithubRepoOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseWorkspaceGithubRepoOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SynapseWorkspaceGithubRepoOutputReference) ResetGitUrl() {
	_jsii_.InvokeVoid(
		s,
		"resetGitUrl",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SynapseWorkspaceGithubRepoOutputReference) ResetLastCommitId() {
	_jsii_.InvokeVoid(
		s,
		"resetLastCommitId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SynapseWorkspaceGithubRepoOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_SynapseWorkspaceGithubRepoOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

