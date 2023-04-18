package automationsoftwareupdateconfiguration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v7/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v7/automationsoftwareupdateconfiguration/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AutomationSoftwareUpdateConfigurationWindowsOutputReference interface {
	cdktf.ComplexObject
	ClassificationIncluded() *string
	SetClassificationIncluded(val *string)
	ClassificationIncludedInput() *string
	ClassificationsIncluded() *[]*string
	SetClassificationsIncluded(val *[]*string)
	ClassificationsIncludedInput() *[]*string
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
	ExcludedKnowledgeBaseNumbers() *[]*string
	SetExcludedKnowledgeBaseNumbers(val *[]*string)
	ExcludedKnowledgeBaseNumbersInput() *[]*string
	// Experimental.
	Fqn() *string
	IncludedKnowledgeBaseNumbers() *[]*string
	SetIncludedKnowledgeBaseNumbers(val *[]*string)
	IncludedKnowledgeBaseNumbersInput() *[]*string
	InternalValue() *AutomationSoftwareUpdateConfigurationWindows
	SetInternalValue(val *AutomationSoftwareUpdateConfigurationWindows)
	Reboot() *string
	SetReboot(val *string)
	RebootInput() *string
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
	ResetClassificationIncluded()
	ResetClassificationsIncluded()
	ResetExcludedKnowledgeBaseNumbers()
	ResetIncludedKnowledgeBaseNumbers()
	ResetReboot()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AutomationSoftwareUpdateConfigurationWindowsOutputReference
type jsiiProxy_AutomationSoftwareUpdateConfigurationWindowsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationWindowsOutputReference) ClassificationIncluded() *string {
	var returns *string
	_jsii_.Get(
		j,
		"classificationIncluded",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationWindowsOutputReference) ClassificationIncludedInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"classificationIncludedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationWindowsOutputReference) ClassificationsIncluded() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"classificationsIncluded",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationWindowsOutputReference) ClassificationsIncludedInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"classificationsIncludedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationWindowsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationWindowsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationWindowsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationWindowsOutputReference) ExcludedKnowledgeBaseNumbers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedKnowledgeBaseNumbers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationWindowsOutputReference) ExcludedKnowledgeBaseNumbersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedKnowledgeBaseNumbersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationWindowsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationWindowsOutputReference) IncludedKnowledgeBaseNumbers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includedKnowledgeBaseNumbers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationWindowsOutputReference) IncludedKnowledgeBaseNumbersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includedKnowledgeBaseNumbersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationWindowsOutputReference) InternalValue() *AutomationSoftwareUpdateConfigurationWindows {
	var returns *AutomationSoftwareUpdateConfigurationWindows
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationWindowsOutputReference) Reboot() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reboot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationWindowsOutputReference) RebootInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rebootInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationWindowsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationWindowsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAutomationSoftwareUpdateConfigurationWindowsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AutomationSoftwareUpdateConfigurationWindowsOutputReference {
	_init_.Initialize()

	if err := validateNewAutomationSoftwareUpdateConfigurationWindowsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AutomationSoftwareUpdateConfigurationWindowsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.automationSoftwareUpdateConfiguration.AutomationSoftwareUpdateConfigurationWindowsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAutomationSoftwareUpdateConfigurationWindowsOutputReference_Override(a AutomationSoftwareUpdateConfigurationWindowsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.automationSoftwareUpdateConfiguration.AutomationSoftwareUpdateConfigurationWindowsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationWindowsOutputReference)SetClassificationIncluded(val *string) {
	if err := j.validateSetClassificationIncludedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"classificationIncluded",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationWindowsOutputReference)SetClassificationsIncluded(val *[]*string) {
	if err := j.validateSetClassificationsIncludedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"classificationsIncluded",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationWindowsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationWindowsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationWindowsOutputReference)SetExcludedKnowledgeBaseNumbers(val *[]*string) {
	if err := j.validateSetExcludedKnowledgeBaseNumbersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludedKnowledgeBaseNumbers",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationWindowsOutputReference)SetIncludedKnowledgeBaseNumbers(val *[]*string) {
	if err := j.validateSetIncludedKnowledgeBaseNumbersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includedKnowledgeBaseNumbers",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationWindowsOutputReference)SetInternalValue(val *AutomationSoftwareUpdateConfigurationWindows) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationWindowsOutputReference)SetReboot(val *string) {
	if err := j.validateSetRebootParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"reboot",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationWindowsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationWindowsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationWindowsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationWindowsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationWindowsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationWindowsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationWindowsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationWindowsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationWindowsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationWindowsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationWindowsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationWindowsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationWindowsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationWindowsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationWindowsOutputReference) ResetClassificationIncluded() {
	_jsii_.InvokeVoid(
		a,
		"resetClassificationIncluded",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationWindowsOutputReference) ResetClassificationsIncluded() {
	_jsii_.InvokeVoid(
		a,
		"resetClassificationsIncluded",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationWindowsOutputReference) ResetExcludedKnowledgeBaseNumbers() {
	_jsii_.InvokeVoid(
		a,
		"resetExcludedKnowledgeBaseNumbers",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationWindowsOutputReference) ResetIncludedKnowledgeBaseNumbers() {
	_jsii_.InvokeVoid(
		a,
		"resetIncludedKnowledgeBaseNumbers",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationWindowsOutputReference) ResetReboot() {
	_jsii_.InvokeVoid(
		a,
		"resetReboot",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationWindowsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationWindowsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

