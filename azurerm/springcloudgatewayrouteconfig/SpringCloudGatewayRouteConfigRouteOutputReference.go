package springcloudgatewayrouteconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v8/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v8/springcloudgatewayrouteconfig/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SpringCloudGatewayRouteConfigRouteOutputReference interface {
	cdktf.ComplexObject
	ClassificationTags() *[]*string
	SetClassificationTags(val *[]*string)
	ClassificationTagsInput() *[]*string
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
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	Filters() *[]*string
	SetFilters(val *[]*string)
	FiltersInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Order() *float64
	SetOrder(val *float64)
	OrderInput() *float64
	Predicates() *[]*string
	SetPredicates(val *[]*string)
	PredicatesInput() *[]*string
	SsoValidationEnabled() interface{}
	SetSsoValidationEnabled(val interface{})
	SsoValidationEnabledInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Title() *string
	SetTitle(val *string)
	TitleInput() *string
	TokenRelay() interface{}
	SetTokenRelay(val interface{})
	TokenRelayInput() interface{}
	Uri() *string
	SetUri(val *string)
	UriInput() *string
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
	ResetClassificationTags()
	ResetDescription()
	ResetFilters()
	ResetPredicates()
	ResetSsoValidationEnabled()
	ResetTitle()
	ResetTokenRelay()
	ResetUri()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SpringCloudGatewayRouteConfigRouteOutputReference
type jsiiProxy_SpringCloudGatewayRouteConfigRouteOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SpringCloudGatewayRouteConfigRouteOutputReference) ClassificationTags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"classificationTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudGatewayRouteConfigRouteOutputReference) ClassificationTagsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"classificationTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudGatewayRouteConfigRouteOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudGatewayRouteConfigRouteOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudGatewayRouteConfigRouteOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudGatewayRouteConfigRouteOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudGatewayRouteConfigRouteOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudGatewayRouteConfigRouteOutputReference) Filters() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"filters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudGatewayRouteConfigRouteOutputReference) FiltersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"filtersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudGatewayRouteConfigRouteOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudGatewayRouteConfigRouteOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudGatewayRouteConfigRouteOutputReference) Order() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"order",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudGatewayRouteConfigRouteOutputReference) OrderInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"orderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudGatewayRouteConfigRouteOutputReference) Predicates() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"predicates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudGatewayRouteConfigRouteOutputReference) PredicatesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"predicatesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudGatewayRouteConfigRouteOutputReference) SsoValidationEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ssoValidationEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudGatewayRouteConfigRouteOutputReference) SsoValidationEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ssoValidationEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudGatewayRouteConfigRouteOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudGatewayRouteConfigRouteOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudGatewayRouteConfigRouteOutputReference) Title() *string {
	var returns *string
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudGatewayRouteConfigRouteOutputReference) TitleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"titleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudGatewayRouteConfigRouteOutputReference) TokenRelay() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tokenRelay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudGatewayRouteConfigRouteOutputReference) TokenRelayInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tokenRelayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudGatewayRouteConfigRouteOutputReference) Uri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudGatewayRouteConfigRouteOutputReference) UriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uriInput",
		&returns,
	)
	return returns
}


func NewSpringCloudGatewayRouteConfigRouteOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) SpringCloudGatewayRouteConfigRouteOutputReference {
	_init_.Initialize()

	if err := validateNewSpringCloudGatewayRouteConfigRouteOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_SpringCloudGatewayRouteConfigRouteOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.springCloudGatewayRouteConfig.SpringCloudGatewayRouteConfigRouteOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewSpringCloudGatewayRouteConfigRouteOutputReference_Override(s SpringCloudGatewayRouteConfigRouteOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.springCloudGatewayRouteConfig.SpringCloudGatewayRouteConfigRouteOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_SpringCloudGatewayRouteConfigRouteOutputReference)SetClassificationTags(val *[]*string) {
	if err := j.validateSetClassificationTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"classificationTags",
		val,
	)
}

func (j *jsiiProxy_SpringCloudGatewayRouteConfigRouteOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SpringCloudGatewayRouteConfigRouteOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SpringCloudGatewayRouteConfigRouteOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_SpringCloudGatewayRouteConfigRouteOutputReference)SetFilters(val *[]*string) {
	if err := j.validateSetFiltersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"filters",
		val,
	)
}

func (j *jsiiProxy_SpringCloudGatewayRouteConfigRouteOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SpringCloudGatewayRouteConfigRouteOutputReference)SetOrder(val *float64) {
	if err := j.validateSetOrderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"order",
		val,
	)
}

func (j *jsiiProxy_SpringCloudGatewayRouteConfigRouteOutputReference)SetPredicates(val *[]*string) {
	if err := j.validateSetPredicatesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"predicates",
		val,
	)
}

func (j *jsiiProxy_SpringCloudGatewayRouteConfigRouteOutputReference)SetSsoValidationEnabled(val interface{}) {
	if err := j.validateSetSsoValidationEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ssoValidationEnabled",
		val,
	)
}

func (j *jsiiProxy_SpringCloudGatewayRouteConfigRouteOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SpringCloudGatewayRouteConfigRouteOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_SpringCloudGatewayRouteConfigRouteOutputReference)SetTitle(val *string) {
	if err := j.validateSetTitleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"title",
		val,
	)
}

func (j *jsiiProxy_SpringCloudGatewayRouteConfigRouteOutputReference)SetTokenRelay(val interface{}) {
	if err := j.validateSetTokenRelayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenRelay",
		val,
	)
}

func (j *jsiiProxy_SpringCloudGatewayRouteConfigRouteOutputReference)SetUri(val *string) {
	if err := j.validateSetUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"uri",
		val,
	)
}

func (s *jsiiProxy_SpringCloudGatewayRouteConfigRouteOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudGatewayRouteConfigRouteOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SpringCloudGatewayRouteConfigRouteOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SpringCloudGatewayRouteConfigRouteOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SpringCloudGatewayRouteConfigRouteOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SpringCloudGatewayRouteConfigRouteOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SpringCloudGatewayRouteConfigRouteOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SpringCloudGatewayRouteConfigRouteOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SpringCloudGatewayRouteConfigRouteOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SpringCloudGatewayRouteConfigRouteOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SpringCloudGatewayRouteConfigRouteOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudGatewayRouteConfigRouteOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SpringCloudGatewayRouteConfigRouteOutputReference) ResetClassificationTags() {
	_jsii_.InvokeVoid(
		s,
		"resetClassificationTags",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpringCloudGatewayRouteConfigRouteOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		s,
		"resetDescription",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpringCloudGatewayRouteConfigRouteOutputReference) ResetFilters() {
	_jsii_.InvokeVoid(
		s,
		"resetFilters",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpringCloudGatewayRouteConfigRouteOutputReference) ResetPredicates() {
	_jsii_.InvokeVoid(
		s,
		"resetPredicates",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpringCloudGatewayRouteConfigRouteOutputReference) ResetSsoValidationEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetSsoValidationEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpringCloudGatewayRouteConfigRouteOutputReference) ResetTitle() {
	_jsii_.InvokeVoid(
		s,
		"resetTitle",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpringCloudGatewayRouteConfigRouteOutputReference) ResetTokenRelay() {
	_jsii_.InvokeVoid(
		s,
		"resetTokenRelay",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpringCloudGatewayRouteConfigRouteOutputReference) ResetUri() {
	_jsii_.InvokeVoid(
		s,
		"resetUri",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpringCloudGatewayRouteConfigRouteOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_SpringCloudGatewayRouteConfigRouteOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

