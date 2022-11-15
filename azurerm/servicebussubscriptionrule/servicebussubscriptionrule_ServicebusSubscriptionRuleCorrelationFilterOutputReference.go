package servicebussubscriptionrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v4/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v4/servicebussubscriptionrule/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ServicebusSubscriptionRuleCorrelationFilterOutputReference interface {
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
	ContentType() *string
	SetContentType(val *string)
	ContentTypeInput() *string
	CorrelationId() *string
	SetCorrelationId(val *string)
	CorrelationIdInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *ServicebusSubscriptionRuleCorrelationFilter
	SetInternalValue(val *ServicebusSubscriptionRuleCorrelationFilter)
	Label() *string
	SetLabel(val *string)
	LabelInput() *string
	MessageId() *string
	SetMessageId(val *string)
	MessageIdInput() *string
	Properties() *map[string]*string
	SetProperties(val *map[string]*string)
	PropertiesInput() *map[string]*string
	ReplyTo() *string
	SetReplyTo(val *string)
	ReplyToInput() *string
	ReplyToSessionId() *string
	SetReplyToSessionId(val *string)
	ReplyToSessionIdInput() *string
	SessionId() *string
	SetSessionId(val *string)
	SessionIdInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	To() *string
	SetTo(val *string)
	ToInput() *string
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
	ResetContentType()
	ResetCorrelationId()
	ResetLabel()
	ResetMessageId()
	ResetProperties()
	ResetReplyTo()
	ResetReplyToSessionId()
	ResetSessionId()
	ResetTo()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ServicebusSubscriptionRuleCorrelationFilterOutputReference
type jsiiProxy_ServicebusSubscriptionRuleCorrelationFilterOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ServicebusSubscriptionRuleCorrelationFilterOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscriptionRuleCorrelationFilterOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscriptionRuleCorrelationFilterOutputReference) ContentType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscriptionRuleCorrelationFilterOutputReference) ContentTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscriptionRuleCorrelationFilterOutputReference) CorrelationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"correlationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscriptionRuleCorrelationFilterOutputReference) CorrelationIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"correlationIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscriptionRuleCorrelationFilterOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscriptionRuleCorrelationFilterOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscriptionRuleCorrelationFilterOutputReference) InternalValue() *ServicebusSubscriptionRuleCorrelationFilter {
	var returns *ServicebusSubscriptionRuleCorrelationFilter
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscriptionRuleCorrelationFilterOutputReference) Label() *string {
	var returns *string
	_jsii_.Get(
		j,
		"label",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscriptionRuleCorrelationFilterOutputReference) LabelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"labelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscriptionRuleCorrelationFilterOutputReference) MessageId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"messageId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscriptionRuleCorrelationFilterOutputReference) MessageIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"messageIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscriptionRuleCorrelationFilterOutputReference) Properties() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"properties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscriptionRuleCorrelationFilterOutputReference) PropertiesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"propertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscriptionRuleCorrelationFilterOutputReference) ReplyTo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replyTo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscriptionRuleCorrelationFilterOutputReference) ReplyToInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replyToInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscriptionRuleCorrelationFilterOutputReference) ReplyToSessionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replyToSessionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscriptionRuleCorrelationFilterOutputReference) ReplyToSessionIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replyToSessionIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscriptionRuleCorrelationFilterOutputReference) SessionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sessionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscriptionRuleCorrelationFilterOutputReference) SessionIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sessionIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscriptionRuleCorrelationFilterOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscriptionRuleCorrelationFilterOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscriptionRuleCorrelationFilterOutputReference) To() *string {
	var returns *string
	_jsii_.Get(
		j,
		"to",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscriptionRuleCorrelationFilterOutputReference) ToInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"toInput",
		&returns,
	)
	return returns
}


func NewServicebusSubscriptionRuleCorrelationFilterOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ServicebusSubscriptionRuleCorrelationFilterOutputReference {
	_init_.Initialize()

	if err := validateNewServicebusSubscriptionRuleCorrelationFilterOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ServicebusSubscriptionRuleCorrelationFilterOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.servicebusSubscriptionRule.ServicebusSubscriptionRuleCorrelationFilterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewServicebusSubscriptionRuleCorrelationFilterOutputReference_Override(s ServicebusSubscriptionRuleCorrelationFilterOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.servicebusSubscriptionRule.ServicebusSubscriptionRuleCorrelationFilterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_ServicebusSubscriptionRuleCorrelationFilterOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ServicebusSubscriptionRuleCorrelationFilterOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ServicebusSubscriptionRuleCorrelationFilterOutputReference)SetContentType(val *string) {
	if err := j.validateSetContentTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contentType",
		val,
	)
}

func (j *jsiiProxy_ServicebusSubscriptionRuleCorrelationFilterOutputReference)SetCorrelationId(val *string) {
	if err := j.validateSetCorrelationIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"correlationId",
		val,
	)
}

func (j *jsiiProxy_ServicebusSubscriptionRuleCorrelationFilterOutputReference)SetInternalValue(val *ServicebusSubscriptionRuleCorrelationFilter) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ServicebusSubscriptionRuleCorrelationFilterOutputReference)SetLabel(val *string) {
	if err := j.validateSetLabelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"label",
		val,
	)
}

func (j *jsiiProxy_ServicebusSubscriptionRuleCorrelationFilterOutputReference)SetMessageId(val *string) {
	if err := j.validateSetMessageIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"messageId",
		val,
	)
}

func (j *jsiiProxy_ServicebusSubscriptionRuleCorrelationFilterOutputReference)SetProperties(val *map[string]*string) {
	if err := j.validateSetPropertiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"properties",
		val,
	)
}

func (j *jsiiProxy_ServicebusSubscriptionRuleCorrelationFilterOutputReference)SetReplyTo(val *string) {
	if err := j.validateSetReplyToParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replyTo",
		val,
	)
}

func (j *jsiiProxy_ServicebusSubscriptionRuleCorrelationFilterOutputReference)SetReplyToSessionId(val *string) {
	if err := j.validateSetReplyToSessionIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replyToSessionId",
		val,
	)
}

func (j *jsiiProxy_ServicebusSubscriptionRuleCorrelationFilterOutputReference)SetSessionId(val *string) {
	if err := j.validateSetSessionIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sessionId",
		val,
	)
}

func (j *jsiiProxy_ServicebusSubscriptionRuleCorrelationFilterOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ServicebusSubscriptionRuleCorrelationFilterOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ServicebusSubscriptionRuleCorrelationFilterOutputReference)SetTo(val *string) {
	if err := j.validateSetToParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"to",
		val,
	)
}

func (s *jsiiProxy_ServicebusSubscriptionRuleCorrelationFilterOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicebusSubscriptionRuleCorrelationFilterOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_ServicebusSubscriptionRuleCorrelationFilterOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_ServicebusSubscriptionRuleCorrelationFilterOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_ServicebusSubscriptionRuleCorrelationFilterOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_ServicebusSubscriptionRuleCorrelationFilterOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_ServicebusSubscriptionRuleCorrelationFilterOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_ServicebusSubscriptionRuleCorrelationFilterOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_ServicebusSubscriptionRuleCorrelationFilterOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_ServicebusSubscriptionRuleCorrelationFilterOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_ServicebusSubscriptionRuleCorrelationFilterOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicebusSubscriptionRuleCorrelationFilterOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (s *jsiiProxy_ServicebusSubscriptionRuleCorrelationFilterOutputReference) ResetContentType() {
	_jsii_.InvokeVoid(
		s,
		"resetContentType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicebusSubscriptionRuleCorrelationFilterOutputReference) ResetCorrelationId() {
	_jsii_.InvokeVoid(
		s,
		"resetCorrelationId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicebusSubscriptionRuleCorrelationFilterOutputReference) ResetLabel() {
	_jsii_.InvokeVoid(
		s,
		"resetLabel",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicebusSubscriptionRuleCorrelationFilterOutputReference) ResetMessageId() {
	_jsii_.InvokeVoid(
		s,
		"resetMessageId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicebusSubscriptionRuleCorrelationFilterOutputReference) ResetProperties() {
	_jsii_.InvokeVoid(
		s,
		"resetProperties",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicebusSubscriptionRuleCorrelationFilterOutputReference) ResetReplyTo() {
	_jsii_.InvokeVoid(
		s,
		"resetReplyTo",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicebusSubscriptionRuleCorrelationFilterOutputReference) ResetReplyToSessionId() {
	_jsii_.InvokeVoid(
		s,
		"resetReplyToSessionId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicebusSubscriptionRuleCorrelationFilterOutputReference) ResetSessionId() {
	_jsii_.InvokeVoid(
		s,
		"resetSessionId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicebusSubscriptionRuleCorrelationFilterOutputReference) ResetTo() {
	_jsii_.InvokeVoid(
		s,
		"resetTo",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicebusSubscriptionRuleCorrelationFilterOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_ServicebusSubscriptionRuleCorrelationFilterOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

