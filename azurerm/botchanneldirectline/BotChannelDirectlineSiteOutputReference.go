package botchanneldirectline

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v7/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v7/botchanneldirectline/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BotChannelDirectlineSiteOutputReference interface {
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
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	EnhancedAuthenticationEnabled() interface{}
	SetEnhancedAuthenticationEnabled(val interface{})
	EnhancedAuthenticationEnabledInput() interface{}
	// Experimental.
	Fqn() *string
	Id() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Key() *string
	Key2() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TrustedOrigins() *[]*string
	SetTrustedOrigins(val *[]*string)
	TrustedOriginsInput() *[]*string
	V1Allowed() interface{}
	SetV1Allowed(val interface{})
	V1AllowedInput() interface{}
	V3Allowed() interface{}
	SetV3Allowed(val interface{})
	V3AllowedInput() interface{}
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
	ResetEnabled()
	ResetEnhancedAuthenticationEnabled()
	ResetTrustedOrigins()
	ResetV1Allowed()
	ResetV3Allowed()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BotChannelDirectlineSiteOutputReference
type jsiiProxy_BotChannelDirectlineSiteOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BotChannelDirectlineSiteOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotChannelDirectlineSiteOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotChannelDirectlineSiteOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotChannelDirectlineSiteOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotChannelDirectlineSiteOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotChannelDirectlineSiteOutputReference) EnhancedAuthenticationEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enhancedAuthenticationEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotChannelDirectlineSiteOutputReference) EnhancedAuthenticationEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enhancedAuthenticationEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotChannelDirectlineSiteOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotChannelDirectlineSiteOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotChannelDirectlineSiteOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotChannelDirectlineSiteOutputReference) Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotChannelDirectlineSiteOutputReference) Key2() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotChannelDirectlineSiteOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotChannelDirectlineSiteOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotChannelDirectlineSiteOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotChannelDirectlineSiteOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotChannelDirectlineSiteOutputReference) TrustedOrigins() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"trustedOrigins",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotChannelDirectlineSiteOutputReference) TrustedOriginsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"trustedOriginsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotChannelDirectlineSiteOutputReference) V1Allowed() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"v1Allowed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotChannelDirectlineSiteOutputReference) V1AllowedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"v1AllowedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotChannelDirectlineSiteOutputReference) V3Allowed() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"v3Allowed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotChannelDirectlineSiteOutputReference) V3AllowedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"v3AllowedInput",
		&returns,
	)
	return returns
}


func NewBotChannelDirectlineSiteOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) BotChannelDirectlineSiteOutputReference {
	_init_.Initialize()

	if err := validateNewBotChannelDirectlineSiteOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_BotChannelDirectlineSiteOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.botChannelDirectline.BotChannelDirectlineSiteOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewBotChannelDirectlineSiteOutputReference_Override(b BotChannelDirectlineSiteOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.botChannelDirectline.BotChannelDirectlineSiteOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		b,
	)
}

func (j *jsiiProxy_BotChannelDirectlineSiteOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BotChannelDirectlineSiteOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BotChannelDirectlineSiteOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_BotChannelDirectlineSiteOutputReference)SetEnhancedAuthenticationEnabled(val interface{}) {
	if err := j.validateSetEnhancedAuthenticationEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enhancedAuthenticationEnabled",
		val,
	)
}

func (j *jsiiProxy_BotChannelDirectlineSiteOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BotChannelDirectlineSiteOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_BotChannelDirectlineSiteOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BotChannelDirectlineSiteOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_BotChannelDirectlineSiteOutputReference)SetTrustedOrigins(val *[]*string) {
	if err := j.validateSetTrustedOriginsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trustedOrigins",
		val,
	)
}

func (j *jsiiProxy_BotChannelDirectlineSiteOutputReference)SetV1Allowed(val interface{}) {
	if err := j.validateSetV1AllowedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"v1Allowed",
		val,
	)
}

func (j *jsiiProxy_BotChannelDirectlineSiteOutputReference)SetV3Allowed(val interface{}) {
	if err := j.validateSetV3AllowedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"v3Allowed",
		val,
	)
}

func (b *jsiiProxy_BotChannelDirectlineSiteOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BotChannelDirectlineSiteOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := b.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BotChannelDirectlineSiteOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := b.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BotChannelDirectlineSiteOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := b.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BotChannelDirectlineSiteOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := b.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BotChannelDirectlineSiteOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := b.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BotChannelDirectlineSiteOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := b.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BotChannelDirectlineSiteOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := b.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BotChannelDirectlineSiteOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := b.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BotChannelDirectlineSiteOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := b.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BotChannelDirectlineSiteOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BotChannelDirectlineSiteOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := b.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BotChannelDirectlineSiteOutputReference) ResetEnabled() {
	_jsii_.InvokeVoid(
		b,
		"resetEnabled",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BotChannelDirectlineSiteOutputReference) ResetEnhancedAuthenticationEnabled() {
	_jsii_.InvokeVoid(
		b,
		"resetEnhancedAuthenticationEnabled",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BotChannelDirectlineSiteOutputReference) ResetTrustedOrigins() {
	_jsii_.InvokeVoid(
		b,
		"resetTrustedOrigins",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BotChannelDirectlineSiteOutputReference) ResetV1Allowed() {
	_jsii_.InvokeVoid(
		b,
		"resetV1Allowed",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BotChannelDirectlineSiteOutputReference) ResetV3Allowed() {
	_jsii_.InvokeVoid(
		b,
		"resetV3Allowed",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BotChannelDirectlineSiteOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := b.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BotChannelDirectlineSiteOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

