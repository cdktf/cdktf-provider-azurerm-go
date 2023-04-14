package eventgrideventsubscription

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v6/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v6/eventgrideventsubscription/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EventgridEventSubscriptionWebhookEndpointOutputReference interface {
	cdktf.ComplexObject
	ActiveDirectoryAppIdOrUri() *string
	SetActiveDirectoryAppIdOrUri(val *string)
	ActiveDirectoryAppIdOrUriInput() *string
	ActiveDirectoryTenantId() *string
	SetActiveDirectoryTenantId(val *string)
	ActiveDirectoryTenantIdInput() *string
	BaseUrl() *string
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
	InternalValue() *EventgridEventSubscriptionWebhookEndpoint
	SetInternalValue(val *EventgridEventSubscriptionWebhookEndpoint)
	MaxEventsPerBatch() *float64
	SetMaxEventsPerBatch(val *float64)
	MaxEventsPerBatchInput() *float64
	PreferredBatchSizeInKilobytes() *float64
	SetPreferredBatchSizeInKilobytes(val *float64)
	PreferredBatchSizeInKilobytesInput() *float64
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
	ResetActiveDirectoryAppIdOrUri()
	ResetActiveDirectoryTenantId()
	ResetMaxEventsPerBatch()
	ResetPreferredBatchSizeInKilobytes()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EventgridEventSubscriptionWebhookEndpointOutputReference
type jsiiProxy_EventgridEventSubscriptionWebhookEndpointOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EventgridEventSubscriptionWebhookEndpointOutputReference) ActiveDirectoryAppIdOrUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"activeDirectoryAppIdOrUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionWebhookEndpointOutputReference) ActiveDirectoryAppIdOrUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"activeDirectoryAppIdOrUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionWebhookEndpointOutputReference) ActiveDirectoryTenantId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"activeDirectoryTenantId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionWebhookEndpointOutputReference) ActiveDirectoryTenantIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"activeDirectoryTenantIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionWebhookEndpointOutputReference) BaseUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baseUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionWebhookEndpointOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionWebhookEndpointOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionWebhookEndpointOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionWebhookEndpointOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionWebhookEndpointOutputReference) InternalValue() *EventgridEventSubscriptionWebhookEndpoint {
	var returns *EventgridEventSubscriptionWebhookEndpoint
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionWebhookEndpointOutputReference) MaxEventsPerBatch() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxEventsPerBatch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionWebhookEndpointOutputReference) MaxEventsPerBatchInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxEventsPerBatchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionWebhookEndpointOutputReference) PreferredBatchSizeInKilobytes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"preferredBatchSizeInKilobytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionWebhookEndpointOutputReference) PreferredBatchSizeInKilobytesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"preferredBatchSizeInKilobytesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionWebhookEndpointOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionWebhookEndpointOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionWebhookEndpointOutputReference) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionWebhookEndpointOutputReference) UrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"urlInput",
		&returns,
	)
	return returns
}


func NewEventgridEventSubscriptionWebhookEndpointOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) EventgridEventSubscriptionWebhookEndpointOutputReference {
	_init_.Initialize()

	if err := validateNewEventgridEventSubscriptionWebhookEndpointOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_EventgridEventSubscriptionWebhookEndpointOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridEventSubscription.EventgridEventSubscriptionWebhookEndpointOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEventgridEventSubscriptionWebhookEndpointOutputReference_Override(e EventgridEventSubscriptionWebhookEndpointOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridEventSubscription.EventgridEventSubscriptionWebhookEndpointOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionWebhookEndpointOutputReference)SetActiveDirectoryAppIdOrUri(val *string) {
	if err := j.validateSetActiveDirectoryAppIdOrUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"activeDirectoryAppIdOrUri",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionWebhookEndpointOutputReference)SetActiveDirectoryTenantId(val *string) {
	if err := j.validateSetActiveDirectoryTenantIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"activeDirectoryTenantId",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionWebhookEndpointOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionWebhookEndpointOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionWebhookEndpointOutputReference)SetInternalValue(val *EventgridEventSubscriptionWebhookEndpoint) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionWebhookEndpointOutputReference)SetMaxEventsPerBatch(val *float64) {
	if err := j.validateSetMaxEventsPerBatchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxEventsPerBatch",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionWebhookEndpointOutputReference)SetPreferredBatchSizeInKilobytes(val *float64) {
	if err := j.validateSetPreferredBatchSizeInKilobytesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preferredBatchSizeInKilobytes",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionWebhookEndpointOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionWebhookEndpointOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionWebhookEndpointOutputReference)SetUrl(val *string) {
	if err := j.validateSetUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"url",
		val,
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionWebhookEndpointOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionWebhookEndpointOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := e.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionWebhookEndpointOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionWebhookEndpointOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := e.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionWebhookEndpointOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := e.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionWebhookEndpointOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := e.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionWebhookEndpointOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := e.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionWebhookEndpointOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := e.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionWebhookEndpointOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := e.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionWebhookEndpointOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := e.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionWebhookEndpointOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionWebhookEndpointOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionWebhookEndpointOutputReference) ResetActiveDirectoryAppIdOrUri() {
	_jsii_.InvokeVoid(
		e,
		"resetActiveDirectoryAppIdOrUri",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionWebhookEndpointOutputReference) ResetActiveDirectoryTenantId() {
	_jsii_.InvokeVoid(
		e,
		"resetActiveDirectoryTenantId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionWebhookEndpointOutputReference) ResetMaxEventsPerBatch() {
	_jsii_.InvokeVoid(
		e,
		"resetMaxEventsPerBatch",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionWebhookEndpointOutputReference) ResetPreferredBatchSizeInKilobytes() {
	_jsii_.InvokeVoid(
		e,
		"resetPreferredBatchSizeInKilobytes",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionWebhookEndpointOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := e.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionWebhookEndpointOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

