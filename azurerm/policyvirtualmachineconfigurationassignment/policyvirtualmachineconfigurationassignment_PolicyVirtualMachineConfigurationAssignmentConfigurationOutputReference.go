package policyvirtualmachineconfigurationassignment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v4/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v4/policyvirtualmachineconfigurationassignment/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PolicyVirtualMachineConfigurationAssignmentConfigurationOutputReference interface {
	cdktf.ComplexObject
	AssignmentType() *string
	SetAssignmentType(val *string)
	AssignmentTypeInput() *string
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
	ContentHash() *string
	SetContentHash(val *string)
	ContentHashInput() *string
	ContentUri() *string
	SetContentUri(val *string)
	ContentUriInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *PolicyVirtualMachineConfigurationAssignmentConfiguration
	SetInternalValue(val *PolicyVirtualMachineConfigurationAssignmentConfiguration)
	Parameter() PolicyVirtualMachineConfigurationAssignmentConfigurationParameterList
	ParameterInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Version() *string
	SetVersion(val *string)
	VersionInput() *string
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
	PutParameter(value interface{})
	ResetAssignmentType()
	ResetContentHash()
	ResetContentUri()
	ResetParameter()
	ResetVersion()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PolicyVirtualMachineConfigurationAssignmentConfigurationOutputReference
type jsiiProxy_PolicyVirtualMachineConfigurationAssignmentConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PolicyVirtualMachineConfigurationAssignmentConfigurationOutputReference) AssignmentType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assignmentType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyVirtualMachineConfigurationAssignmentConfigurationOutputReference) AssignmentTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assignmentTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyVirtualMachineConfigurationAssignmentConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyVirtualMachineConfigurationAssignmentConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyVirtualMachineConfigurationAssignmentConfigurationOutputReference) ContentHash() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentHash",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyVirtualMachineConfigurationAssignmentConfigurationOutputReference) ContentHashInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentHashInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyVirtualMachineConfigurationAssignmentConfigurationOutputReference) ContentUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyVirtualMachineConfigurationAssignmentConfigurationOutputReference) ContentUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyVirtualMachineConfigurationAssignmentConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyVirtualMachineConfigurationAssignmentConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyVirtualMachineConfigurationAssignmentConfigurationOutputReference) InternalValue() *PolicyVirtualMachineConfigurationAssignmentConfiguration {
	var returns *PolicyVirtualMachineConfigurationAssignmentConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyVirtualMachineConfigurationAssignmentConfigurationOutputReference) Parameter() PolicyVirtualMachineConfigurationAssignmentConfigurationParameterList {
	var returns PolicyVirtualMachineConfigurationAssignmentConfigurationParameterList
	_jsii_.Get(
		j,
		"parameter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyVirtualMachineConfigurationAssignmentConfigurationOutputReference) ParameterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"parameterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyVirtualMachineConfigurationAssignmentConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyVirtualMachineConfigurationAssignmentConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyVirtualMachineConfigurationAssignmentConfigurationOutputReference) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyVirtualMachineConfigurationAssignmentConfigurationOutputReference) VersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}


func NewPolicyVirtualMachineConfigurationAssignmentConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) PolicyVirtualMachineConfigurationAssignmentConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewPolicyVirtualMachineConfigurationAssignmentConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PolicyVirtualMachineConfigurationAssignmentConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.policyVirtualMachineConfigurationAssignment.PolicyVirtualMachineConfigurationAssignmentConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPolicyVirtualMachineConfigurationAssignmentConfigurationOutputReference_Override(p PolicyVirtualMachineConfigurationAssignmentConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.policyVirtualMachineConfigurationAssignment.PolicyVirtualMachineConfigurationAssignmentConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PolicyVirtualMachineConfigurationAssignmentConfigurationOutputReference)SetAssignmentType(val *string) {
	if err := j.validateSetAssignmentTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"assignmentType",
		val,
	)
}

func (j *jsiiProxy_PolicyVirtualMachineConfigurationAssignmentConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PolicyVirtualMachineConfigurationAssignmentConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PolicyVirtualMachineConfigurationAssignmentConfigurationOutputReference)SetContentHash(val *string) {
	if err := j.validateSetContentHashParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contentHash",
		val,
	)
}

func (j *jsiiProxy_PolicyVirtualMachineConfigurationAssignmentConfigurationOutputReference)SetContentUri(val *string) {
	if err := j.validateSetContentUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contentUri",
		val,
	)
}

func (j *jsiiProxy_PolicyVirtualMachineConfigurationAssignmentConfigurationOutputReference)SetInternalValue(val *PolicyVirtualMachineConfigurationAssignmentConfiguration) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PolicyVirtualMachineConfigurationAssignmentConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PolicyVirtualMachineConfigurationAssignmentConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_PolicyVirtualMachineConfigurationAssignmentConfigurationOutputReference)SetVersion(val *string) {
	if err := j.validateSetVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

func (p *jsiiProxy_PolicyVirtualMachineConfigurationAssignmentConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyVirtualMachineConfigurationAssignmentConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := p.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyVirtualMachineConfigurationAssignmentConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := p.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyVirtualMachineConfigurationAssignmentConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := p.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyVirtualMachineConfigurationAssignmentConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := p.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyVirtualMachineConfigurationAssignmentConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := p.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyVirtualMachineConfigurationAssignmentConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := p.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyVirtualMachineConfigurationAssignmentConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := p.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyVirtualMachineConfigurationAssignmentConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := p.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyVirtualMachineConfigurationAssignmentConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := p.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyVirtualMachineConfigurationAssignmentConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyVirtualMachineConfigurationAssignmentConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := p.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyVirtualMachineConfigurationAssignmentConfigurationOutputReference) PutParameter(value interface{}) {
	if err := p.validatePutParameterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putParameter",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PolicyVirtualMachineConfigurationAssignmentConfigurationOutputReference) ResetAssignmentType() {
	_jsii_.InvokeVoid(
		p,
		"resetAssignmentType",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyVirtualMachineConfigurationAssignmentConfigurationOutputReference) ResetContentHash() {
	_jsii_.InvokeVoid(
		p,
		"resetContentHash",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyVirtualMachineConfigurationAssignmentConfigurationOutputReference) ResetContentUri() {
	_jsii_.InvokeVoid(
		p,
		"resetContentUri",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyVirtualMachineConfigurationAssignmentConfigurationOutputReference) ResetParameter() {
	_jsii_.InvokeVoid(
		p,
		"resetParameter",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyVirtualMachineConfigurationAssignmentConfigurationOutputReference) ResetVersion() {
	_jsii_.InvokeVoid(
		p,
		"resetVersion",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyVirtualMachineConfigurationAssignmentConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := p.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		p,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyVirtualMachineConfigurationAssignmentConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

