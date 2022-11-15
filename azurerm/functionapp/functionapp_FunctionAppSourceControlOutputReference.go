package functionapp

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v4/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v4/functionapp/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type FunctionAppSourceControlOutputReference interface {
	cdktf.ComplexObject
	Branch() *string
	SetBranch(val *string)
	BranchInput() *string
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
	InternalValue() *FunctionAppSourceControl
	SetInternalValue(val *FunctionAppSourceControl)
	ManualIntegration() interface{}
	SetManualIntegration(val interface{})
	ManualIntegrationInput() interface{}
	RepoUrl() *string
	SetRepoUrl(val *string)
	RepoUrlInput() *string
	RollbackEnabled() interface{}
	SetRollbackEnabled(val interface{})
	RollbackEnabledInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UseMercurial() interface{}
	SetUseMercurial(val interface{})
	UseMercurialInput() interface{}
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
	ResetBranch()
	ResetManualIntegration()
	ResetRepoUrl()
	ResetRollbackEnabled()
	ResetUseMercurial()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for FunctionAppSourceControlOutputReference
type jsiiProxy_FunctionAppSourceControlOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_FunctionAppSourceControlOutputReference) Branch() *string {
	var returns *string
	_jsii_.Get(
		j,
		"branch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppSourceControlOutputReference) BranchInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"branchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppSourceControlOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppSourceControlOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppSourceControlOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppSourceControlOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppSourceControlOutputReference) InternalValue() *FunctionAppSourceControl {
	var returns *FunctionAppSourceControl
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppSourceControlOutputReference) ManualIntegration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"manualIntegration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppSourceControlOutputReference) ManualIntegrationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"manualIntegrationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppSourceControlOutputReference) RepoUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repoUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppSourceControlOutputReference) RepoUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repoUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppSourceControlOutputReference) RollbackEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rollbackEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppSourceControlOutputReference) RollbackEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rollbackEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppSourceControlOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppSourceControlOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppSourceControlOutputReference) UseMercurial() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useMercurial",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppSourceControlOutputReference) UseMercurialInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useMercurialInput",
		&returns,
	)
	return returns
}


func NewFunctionAppSourceControlOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) FunctionAppSourceControlOutputReference {
	_init_.Initialize()

	if err := validateNewFunctionAppSourceControlOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_FunctionAppSourceControlOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.functionApp.FunctionAppSourceControlOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewFunctionAppSourceControlOutputReference_Override(f FunctionAppSourceControlOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.functionApp.FunctionAppSourceControlOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		f,
	)
}

func (j *jsiiProxy_FunctionAppSourceControlOutputReference)SetBranch(val *string) {
	if err := j.validateSetBranchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"branch",
		val,
	)
}

func (j *jsiiProxy_FunctionAppSourceControlOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_FunctionAppSourceControlOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_FunctionAppSourceControlOutputReference)SetInternalValue(val *FunctionAppSourceControl) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FunctionAppSourceControlOutputReference)SetManualIntegration(val interface{}) {
	if err := j.validateSetManualIntegrationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"manualIntegration",
		val,
	)
}

func (j *jsiiProxy_FunctionAppSourceControlOutputReference)SetRepoUrl(val *string) {
	if err := j.validateSetRepoUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"repoUrl",
		val,
	)
}

func (j *jsiiProxy_FunctionAppSourceControlOutputReference)SetRollbackEnabled(val interface{}) {
	if err := j.validateSetRollbackEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rollbackEnabled",
		val,
	)
}

func (j *jsiiProxy_FunctionAppSourceControlOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FunctionAppSourceControlOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_FunctionAppSourceControlOutputReference)SetUseMercurial(val interface{}) {
	if err := j.validateSetUseMercurialParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useMercurial",
		val,
	)
}

func (f *jsiiProxy_FunctionAppSourceControlOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionAppSourceControlOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (f *jsiiProxy_FunctionAppSourceControlOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (f *jsiiProxy_FunctionAppSourceControlOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (f *jsiiProxy_FunctionAppSourceControlOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (f *jsiiProxy_FunctionAppSourceControlOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (f *jsiiProxy_FunctionAppSourceControlOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (f *jsiiProxy_FunctionAppSourceControlOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (f *jsiiProxy_FunctionAppSourceControlOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (f *jsiiProxy_FunctionAppSourceControlOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (f *jsiiProxy_FunctionAppSourceControlOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionAppSourceControlOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (f *jsiiProxy_FunctionAppSourceControlOutputReference) ResetBranch() {
	_jsii_.InvokeVoid(
		f,
		"resetBranch",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppSourceControlOutputReference) ResetManualIntegration() {
	_jsii_.InvokeVoid(
		f,
		"resetManualIntegration",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppSourceControlOutputReference) ResetRepoUrl() {
	_jsii_.InvokeVoid(
		f,
		"resetRepoUrl",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppSourceControlOutputReference) ResetRollbackEnabled() {
	_jsii_.InvokeVoid(
		f,
		"resetRollbackEnabled",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppSourceControlOutputReference) ResetUseMercurial() {
	_jsii_.InvokeVoid(
		f,
		"resetUseMercurial",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppSourceControlOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (f *jsiiProxy_FunctionAppSourceControlOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

