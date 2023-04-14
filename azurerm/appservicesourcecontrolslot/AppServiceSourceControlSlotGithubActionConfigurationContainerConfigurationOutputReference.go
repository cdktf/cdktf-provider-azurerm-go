package appservicesourcecontrolslot

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v6/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v6/appservicesourcecontrolslot/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AppServiceSourceControlSlotGithubActionConfigurationContainerConfigurationOutputReference interface {
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
	ImageName() *string
	SetImageName(val *string)
	ImageNameInput() *string
	InternalValue() *AppServiceSourceControlSlotGithubActionConfigurationContainerConfiguration
	SetInternalValue(val *AppServiceSourceControlSlotGithubActionConfigurationContainerConfiguration)
	RegistryPassword() *string
	SetRegistryPassword(val *string)
	RegistryPasswordInput() *string
	RegistryUrl() *string
	SetRegistryUrl(val *string)
	RegistryUrlInput() *string
	RegistryUsername() *string
	SetRegistryUsername(val *string)
	RegistryUsernameInput() *string
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
	ResetRegistryPassword()
	ResetRegistryUsername()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AppServiceSourceControlSlotGithubActionConfigurationContainerConfigurationOutputReference
type jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationContainerConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationContainerConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationContainerConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationContainerConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationContainerConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationContainerConfigurationOutputReference) ImageName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationContainerConfigurationOutputReference) ImageNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationContainerConfigurationOutputReference) InternalValue() *AppServiceSourceControlSlotGithubActionConfigurationContainerConfiguration {
	var returns *AppServiceSourceControlSlotGithubActionConfigurationContainerConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationContainerConfigurationOutputReference) RegistryPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"registryPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationContainerConfigurationOutputReference) RegistryPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"registryPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationContainerConfigurationOutputReference) RegistryUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"registryUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationContainerConfigurationOutputReference) RegistryUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"registryUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationContainerConfigurationOutputReference) RegistryUsername() *string {
	var returns *string
	_jsii_.Get(
		j,
		"registryUsername",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationContainerConfigurationOutputReference) RegistryUsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"registryUsernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationContainerConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationContainerConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAppServiceSourceControlSlotGithubActionConfigurationContainerConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AppServiceSourceControlSlotGithubActionConfigurationContainerConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewAppServiceSourceControlSlotGithubActionConfigurationContainerConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationContainerConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.appServiceSourceControlSlot.AppServiceSourceControlSlotGithubActionConfigurationContainerConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAppServiceSourceControlSlotGithubActionConfigurationContainerConfigurationOutputReference_Override(a AppServiceSourceControlSlotGithubActionConfigurationContainerConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.appServiceSourceControlSlot.AppServiceSourceControlSlotGithubActionConfigurationContainerConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationContainerConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationContainerConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationContainerConfigurationOutputReference)SetImageName(val *string) {
	if err := j.validateSetImageNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imageName",
		val,
	)
}

func (j *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationContainerConfigurationOutputReference)SetInternalValue(val *AppServiceSourceControlSlotGithubActionConfigurationContainerConfiguration) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationContainerConfigurationOutputReference)SetRegistryPassword(val *string) {
	if err := j.validateSetRegistryPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"registryPassword",
		val,
	)
}

func (j *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationContainerConfigurationOutputReference)SetRegistryUrl(val *string) {
	if err := j.validateSetRegistryUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"registryUrl",
		val,
	)
}

func (j *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationContainerConfigurationOutputReference)SetRegistryUsername(val *string) {
	if err := j.validateSetRegistryUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"registryUsername",
		val,
	)
}

func (j *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationContainerConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationContainerConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationContainerConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationContainerConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationContainerConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationContainerConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationContainerConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationContainerConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationContainerConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationContainerConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationContainerConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationContainerConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationContainerConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationContainerConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationContainerConfigurationOutputReference) ResetRegistryPassword() {
	_jsii_.InvokeVoid(
		a,
		"resetRegistryPassword",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationContainerConfigurationOutputReference) ResetRegistryUsername() {
	_jsii_.InvokeVoid(
		a,
		"resetRegistryUsername",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationContainerConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_AppServiceSourceControlSlotGithubActionConfigurationContainerConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

