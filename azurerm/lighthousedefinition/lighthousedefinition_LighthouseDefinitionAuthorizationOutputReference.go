package lighthousedefinition

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/jsii"

	"github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/lighthousedefinition/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LighthouseDefinitionAuthorizationOutputReference interface {
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
	DelegatedRoleDefinitionIds() *[]*string
	SetDelegatedRoleDefinitionIds(val *[]*string)
	DelegatedRoleDefinitionIdsInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	PrincipalDisplayName() *string
	SetPrincipalDisplayName(val *string)
	PrincipalDisplayNameInput() *string
	PrincipalId() *string
	SetPrincipalId(val *string)
	PrincipalIdInput() *string
	RoleDefinitionId() *string
	SetRoleDefinitionId(val *string)
	RoleDefinitionIdInput() *string
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
	ResetDelegatedRoleDefinitionIds()
	ResetPrincipalDisplayName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for LighthouseDefinitionAuthorizationOutputReference
type jsiiProxy_LighthouseDefinitionAuthorizationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LighthouseDefinitionAuthorizationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LighthouseDefinitionAuthorizationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LighthouseDefinitionAuthorizationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LighthouseDefinitionAuthorizationOutputReference) DelegatedRoleDefinitionIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"delegatedRoleDefinitionIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LighthouseDefinitionAuthorizationOutputReference) DelegatedRoleDefinitionIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"delegatedRoleDefinitionIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LighthouseDefinitionAuthorizationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LighthouseDefinitionAuthorizationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LighthouseDefinitionAuthorizationOutputReference) PrincipalDisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"principalDisplayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LighthouseDefinitionAuthorizationOutputReference) PrincipalDisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"principalDisplayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LighthouseDefinitionAuthorizationOutputReference) PrincipalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"principalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LighthouseDefinitionAuthorizationOutputReference) PrincipalIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"principalIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LighthouseDefinitionAuthorizationOutputReference) RoleDefinitionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleDefinitionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LighthouseDefinitionAuthorizationOutputReference) RoleDefinitionIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleDefinitionIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LighthouseDefinitionAuthorizationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LighthouseDefinitionAuthorizationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewLighthouseDefinitionAuthorizationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) LighthouseDefinitionAuthorizationOutputReference {
	_init_.Initialize()

	if err := validateNewLighthouseDefinitionAuthorizationOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_LighthouseDefinitionAuthorizationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.lighthouseDefinition.LighthouseDefinitionAuthorizationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewLighthouseDefinitionAuthorizationOutputReference_Override(l LighthouseDefinitionAuthorizationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.lighthouseDefinition.LighthouseDefinitionAuthorizationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		l,
	)
}

func (j *jsiiProxy_LighthouseDefinitionAuthorizationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LighthouseDefinitionAuthorizationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LighthouseDefinitionAuthorizationOutputReference)SetDelegatedRoleDefinitionIds(val *[]*string) {
	if err := j.validateSetDelegatedRoleDefinitionIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"delegatedRoleDefinitionIds",
		val,
	)
}

func (j *jsiiProxy_LighthouseDefinitionAuthorizationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LighthouseDefinitionAuthorizationOutputReference)SetPrincipalDisplayName(val *string) {
	if err := j.validateSetPrincipalDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"principalDisplayName",
		val,
	)
}

func (j *jsiiProxy_LighthouseDefinitionAuthorizationOutputReference)SetPrincipalId(val *string) {
	if err := j.validateSetPrincipalIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"principalId",
		val,
	)
}

func (j *jsiiProxy_LighthouseDefinitionAuthorizationOutputReference)SetRoleDefinitionId(val *string) {
	if err := j.validateSetRoleDefinitionIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleDefinitionId",
		val,
	)
}

func (j *jsiiProxy_LighthouseDefinitionAuthorizationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LighthouseDefinitionAuthorizationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (l *jsiiProxy_LighthouseDefinitionAuthorizationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LighthouseDefinitionAuthorizationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := l.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LighthouseDefinitionAuthorizationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := l.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LighthouseDefinitionAuthorizationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := l.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LighthouseDefinitionAuthorizationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := l.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LighthouseDefinitionAuthorizationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := l.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LighthouseDefinitionAuthorizationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := l.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LighthouseDefinitionAuthorizationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := l.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LighthouseDefinitionAuthorizationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := l.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LighthouseDefinitionAuthorizationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := l.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LighthouseDefinitionAuthorizationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LighthouseDefinitionAuthorizationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := l.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LighthouseDefinitionAuthorizationOutputReference) ResetDelegatedRoleDefinitionIds() {
	_jsii_.InvokeVoid(
		l,
		"resetDelegatedRoleDefinitionIds",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LighthouseDefinitionAuthorizationOutputReference) ResetPrincipalDisplayName() {
	_jsii_.InvokeVoid(
		l,
		"resetPrincipalDisplayName",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LighthouseDefinitionAuthorizationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := l.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		l,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LighthouseDefinitionAuthorizationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

