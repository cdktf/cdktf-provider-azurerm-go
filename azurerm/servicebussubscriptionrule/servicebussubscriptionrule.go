package servicebussubscriptionrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/servicebussubscriptionrule/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/azurerm/r/servicebus_subscription_rule azurerm_servicebus_subscription_rule}.
type ServicebusSubscriptionRule interface {
	cdktf.TerraformResource
	Action() *string
	SetAction(val *string)
	ActionInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	CorrelationFilter() ServicebusSubscriptionRuleCorrelationFilterOutputReference
	CorrelationFilterInput() *ServicebusSubscriptionRuleCorrelationFilter
	// Experimental.
	Count() *float64
	// Experimental.
	SetCount(val *float64)
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	FilterType() *string
	SetFilterType(val *string)
	FilterTypeInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	SqlFilter() *string
	SetSqlFilter(val *string)
	SqlFilterCompatibilityLevel() *float64
	SqlFilterInput() *string
	SubscriptionId() *string
	SetSubscriptionId(val *string)
	SubscriptionIdInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() ServicebusSubscriptionRuleTimeoutsOutputReference
	TimeoutsInput() interface{}
	// Experimental.
	AddOverride(path *string, value interface{})
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutCorrelationFilter(value *ServicebusSubscriptionRuleCorrelationFilter)
	PutTimeouts(value *ServicebusSubscriptionRuleTimeouts)
	ResetAction()
	ResetCorrelationFilter()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetSqlFilter()
	ResetTimeouts()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for ServicebusSubscriptionRule
type jsiiProxy_ServicebusSubscriptionRule struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ServicebusSubscriptionRule) Action() *string {
	var returns *string
	_jsii_.Get(
		j,
		"action",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscriptionRule) ActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscriptionRule) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscriptionRule) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscriptionRule) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscriptionRule) CorrelationFilter() ServicebusSubscriptionRuleCorrelationFilterOutputReference {
	var returns ServicebusSubscriptionRuleCorrelationFilterOutputReference
	_jsii_.Get(
		j,
		"correlationFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscriptionRule) CorrelationFilterInput() *ServicebusSubscriptionRuleCorrelationFilter {
	var returns *ServicebusSubscriptionRuleCorrelationFilter
	_jsii_.Get(
		j,
		"correlationFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscriptionRule) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscriptionRule) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscriptionRule) FilterType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filterType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscriptionRule) FilterTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filterTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscriptionRule) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscriptionRule) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscriptionRule) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscriptionRule) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscriptionRule) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscriptionRule) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscriptionRule) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscriptionRule) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscriptionRule) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscriptionRule) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscriptionRule) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscriptionRule) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscriptionRule) SqlFilter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqlFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscriptionRule) SqlFilterCompatibilityLevel() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sqlFilterCompatibilityLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscriptionRule) SqlFilterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqlFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscriptionRule) SubscriptionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subscriptionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscriptionRule) SubscriptionIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subscriptionIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscriptionRule) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscriptionRule) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscriptionRule) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscriptionRule) Timeouts() ServicebusSubscriptionRuleTimeoutsOutputReference {
	var returns ServicebusSubscriptionRuleTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscriptionRule) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/servicebus_subscription_rule azurerm_servicebus_subscription_rule} Resource.
func NewServicebusSubscriptionRule(scope constructs.Construct, id *string, config *ServicebusSubscriptionRuleConfig) ServicebusSubscriptionRule {
	_init_.Initialize()

	j := jsiiProxy_ServicebusSubscriptionRule{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.servicebusSubscriptionRule.ServicebusSubscriptionRule",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/servicebus_subscription_rule azurerm_servicebus_subscription_rule} Resource.
func NewServicebusSubscriptionRule_Override(s ServicebusSubscriptionRule, scope constructs.Construct, id *string, config *ServicebusSubscriptionRuleConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.servicebusSubscriptionRule.ServicebusSubscriptionRule",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_ServicebusSubscriptionRule) SetAction(val *string) {
	_jsii_.Set(
		j,
		"action",
		val,
	)
}

func (j *jsiiProxy_ServicebusSubscriptionRule) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ServicebusSubscriptionRule) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ServicebusSubscriptionRule) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ServicebusSubscriptionRule) SetFilterType(val *string) {
	_jsii_.Set(
		j,
		"filterType",
		val,
	)
}

func (j *jsiiProxy_ServicebusSubscriptionRule) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ServicebusSubscriptionRule) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ServicebusSubscriptionRule) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ServicebusSubscriptionRule) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ServicebusSubscriptionRule) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ServicebusSubscriptionRule) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ServicebusSubscriptionRule) SetSqlFilter(val *string) {
	_jsii_.Set(
		j,
		"sqlFilter",
		val,
	)
}

func (j *jsiiProxy_ServicebusSubscriptionRule) SetSubscriptionId(val *string) {
	_jsii_.Set(
		j,
		"subscriptionId",
		val,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func ServicebusSubscriptionRule_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.servicebusSubscriptionRule.ServicebusSubscriptionRule",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ServicebusSubscriptionRule_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.servicebusSubscriptionRule.ServicebusSubscriptionRule",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_ServicebusSubscriptionRule) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_ServicebusSubscriptionRule) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicebusSubscriptionRule) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicebusSubscriptionRule) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicebusSubscriptionRule) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicebusSubscriptionRule) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicebusSubscriptionRule) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicebusSubscriptionRule) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicebusSubscriptionRule) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicebusSubscriptionRule) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicebusSubscriptionRule) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicebusSubscriptionRule) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_ServicebusSubscriptionRule) PutCorrelationFilter(value *ServicebusSubscriptionRuleCorrelationFilter) {
	_jsii_.InvokeVoid(
		s,
		"putCorrelationFilter",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_ServicebusSubscriptionRule) PutTimeouts(value *ServicebusSubscriptionRuleTimeouts) {
	_jsii_.InvokeVoid(
		s,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_ServicebusSubscriptionRule) ResetAction() {
	_jsii_.InvokeVoid(
		s,
		"resetAction",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicebusSubscriptionRule) ResetCorrelationFilter() {
	_jsii_.InvokeVoid(
		s,
		"resetCorrelationFilter",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicebusSubscriptionRule) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicebusSubscriptionRule) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicebusSubscriptionRule) ResetSqlFilter() {
	_jsii_.InvokeVoid(
		s,
		"resetSqlFilter",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicebusSubscriptionRule) ResetTimeouts() {
	_jsii_.InvokeVoid(
		s,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicebusSubscriptionRule) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicebusSubscriptionRule) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicebusSubscriptionRule) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicebusSubscriptionRule) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ServicebusSubscriptionRuleConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/servicebus_subscription_rule#filter_type ServicebusSubscriptionRule#filter_type}.
	FilterType *string `field:"required" json:"filterType" yaml:"filterType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/servicebus_subscription_rule#name ServicebusSubscriptionRule#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/servicebus_subscription_rule#subscription_id ServicebusSubscriptionRule#subscription_id}.
	SubscriptionId *string `field:"required" json:"subscriptionId" yaml:"subscriptionId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/servicebus_subscription_rule#action ServicebusSubscriptionRule#action}.
	Action *string `field:"optional" json:"action" yaml:"action"`
	// correlation_filter block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/servicebus_subscription_rule#correlation_filter ServicebusSubscriptionRule#correlation_filter}
	CorrelationFilter *ServicebusSubscriptionRuleCorrelationFilter `field:"optional" json:"correlationFilter" yaml:"correlationFilter"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/servicebus_subscription_rule#id ServicebusSubscriptionRule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/servicebus_subscription_rule#sql_filter ServicebusSubscriptionRule#sql_filter}.
	SqlFilter *string `field:"optional" json:"sqlFilter" yaml:"sqlFilter"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/servicebus_subscription_rule#timeouts ServicebusSubscriptionRule#timeouts}
	Timeouts *ServicebusSubscriptionRuleTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

type ServicebusSubscriptionRuleCorrelationFilter struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/servicebus_subscription_rule#content_type ServicebusSubscriptionRule#content_type}.
	ContentType *string `field:"optional" json:"contentType" yaml:"contentType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/servicebus_subscription_rule#correlation_id ServicebusSubscriptionRule#correlation_id}.
	CorrelationId *string `field:"optional" json:"correlationId" yaml:"correlationId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/servicebus_subscription_rule#label ServicebusSubscriptionRule#label}.
	Label *string `field:"optional" json:"label" yaml:"label"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/servicebus_subscription_rule#message_id ServicebusSubscriptionRule#message_id}.
	MessageId *string `field:"optional" json:"messageId" yaml:"messageId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/servicebus_subscription_rule#properties ServicebusSubscriptionRule#properties}.
	Properties *map[string]*string `field:"optional" json:"properties" yaml:"properties"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/servicebus_subscription_rule#reply_to ServicebusSubscriptionRule#reply_to}.
	ReplyTo *string `field:"optional" json:"replyTo" yaml:"replyTo"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/servicebus_subscription_rule#reply_to_session_id ServicebusSubscriptionRule#reply_to_session_id}.
	ReplyToSessionId *string `field:"optional" json:"replyToSessionId" yaml:"replyToSessionId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/servicebus_subscription_rule#session_id ServicebusSubscriptionRule#session_id}.
	SessionId *string `field:"optional" json:"sessionId" yaml:"sessionId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/servicebus_subscription_rule#to ServicebusSubscriptionRule#to}.
	To *string `field:"optional" json:"to" yaml:"to"`
}

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

func (j *jsiiProxy_ServicebusSubscriptionRuleCorrelationFilterOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ServicebusSubscriptionRuleCorrelationFilterOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ServicebusSubscriptionRuleCorrelationFilterOutputReference) SetContentType(val *string) {
	_jsii_.Set(
		j,
		"contentType",
		val,
	)
}

func (j *jsiiProxy_ServicebusSubscriptionRuleCorrelationFilterOutputReference) SetCorrelationId(val *string) {
	_jsii_.Set(
		j,
		"correlationId",
		val,
	)
}

func (j *jsiiProxy_ServicebusSubscriptionRuleCorrelationFilterOutputReference) SetInternalValue(val *ServicebusSubscriptionRuleCorrelationFilter) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ServicebusSubscriptionRuleCorrelationFilterOutputReference) SetLabel(val *string) {
	_jsii_.Set(
		j,
		"label",
		val,
	)
}

func (j *jsiiProxy_ServicebusSubscriptionRuleCorrelationFilterOutputReference) SetMessageId(val *string) {
	_jsii_.Set(
		j,
		"messageId",
		val,
	)
}

func (j *jsiiProxy_ServicebusSubscriptionRuleCorrelationFilterOutputReference) SetProperties(val *map[string]*string) {
	_jsii_.Set(
		j,
		"properties",
		val,
	)
}

func (j *jsiiProxy_ServicebusSubscriptionRuleCorrelationFilterOutputReference) SetReplyTo(val *string) {
	_jsii_.Set(
		j,
		"replyTo",
		val,
	)
}

func (j *jsiiProxy_ServicebusSubscriptionRuleCorrelationFilterOutputReference) SetReplyToSessionId(val *string) {
	_jsii_.Set(
		j,
		"replyToSessionId",
		val,
	)
}

func (j *jsiiProxy_ServicebusSubscriptionRuleCorrelationFilterOutputReference) SetSessionId(val *string) {
	_jsii_.Set(
		j,
		"sessionId",
		val,
	)
}

func (j *jsiiProxy_ServicebusSubscriptionRuleCorrelationFilterOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ServicebusSubscriptionRuleCorrelationFilterOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ServicebusSubscriptionRuleCorrelationFilterOutputReference) SetTo(val *string) {
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

type ServicebusSubscriptionRuleTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/servicebus_subscription_rule#create ServicebusSubscriptionRule#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/servicebus_subscription_rule#delete ServicebusSubscriptionRule#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/servicebus_subscription_rule#read ServicebusSubscriptionRule#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/servicebus_subscription_rule#update ServicebusSubscriptionRule#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

type ServicebusSubscriptionRuleTimeoutsOutputReference interface {
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
	Create() *string
	SetCreate(val *string)
	CreateInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Delete() *string
	SetDelete(val *string)
	DeleteInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Read() *string
	SetRead(val *string)
	ReadInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Update() *string
	SetUpdate(val *string)
	UpdateInput() *string
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
	ResetCreate()
	ResetDelete()
	ResetRead()
	ResetUpdate()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ServicebusSubscriptionRuleTimeoutsOutputReference
type jsiiProxy_ServicebusSubscriptionRuleTimeoutsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ServicebusSubscriptionRuleTimeoutsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscriptionRuleTimeoutsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscriptionRuleTimeoutsOutputReference) Create() *string {
	var returns *string
	_jsii_.Get(
		j,
		"create",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscriptionRuleTimeoutsOutputReference) CreateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscriptionRuleTimeoutsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscriptionRuleTimeoutsOutputReference) Delete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscriptionRuleTimeoutsOutputReference) DeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscriptionRuleTimeoutsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscriptionRuleTimeoutsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscriptionRuleTimeoutsOutputReference) Read() *string {
	var returns *string
	_jsii_.Get(
		j,
		"read",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscriptionRuleTimeoutsOutputReference) ReadInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscriptionRuleTimeoutsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscriptionRuleTimeoutsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscriptionRuleTimeoutsOutputReference) Update() *string {
	var returns *string
	_jsii_.Get(
		j,
		"update",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscriptionRuleTimeoutsOutputReference) UpdateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateInput",
		&returns,
	)
	return returns
}


func NewServicebusSubscriptionRuleTimeoutsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ServicebusSubscriptionRuleTimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ServicebusSubscriptionRuleTimeoutsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.servicebusSubscriptionRule.ServicebusSubscriptionRuleTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewServicebusSubscriptionRuleTimeoutsOutputReference_Override(s ServicebusSubscriptionRuleTimeoutsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.servicebusSubscriptionRule.ServicebusSubscriptionRuleTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_ServicebusSubscriptionRuleTimeoutsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ServicebusSubscriptionRuleTimeoutsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ServicebusSubscriptionRuleTimeoutsOutputReference) SetCreate(val *string) {
	_jsii_.Set(
		j,
		"create",
		val,
	)
}

func (j *jsiiProxy_ServicebusSubscriptionRuleTimeoutsOutputReference) SetDelete(val *string) {
	_jsii_.Set(
		j,
		"delete",
		val,
	)
}

func (j *jsiiProxy_ServicebusSubscriptionRuleTimeoutsOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ServicebusSubscriptionRuleTimeoutsOutputReference) SetRead(val *string) {
	_jsii_.Set(
		j,
		"read",
		val,
	)
}

func (j *jsiiProxy_ServicebusSubscriptionRuleTimeoutsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ServicebusSubscriptionRuleTimeoutsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ServicebusSubscriptionRuleTimeoutsOutputReference) SetUpdate(val *string) {
	_jsii_.Set(
		j,
		"update",
		val,
	)
}

func (s *jsiiProxy_ServicebusSubscriptionRuleTimeoutsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicebusSubscriptionRuleTimeoutsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicebusSubscriptionRuleTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicebusSubscriptionRuleTimeoutsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicebusSubscriptionRuleTimeoutsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicebusSubscriptionRuleTimeoutsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicebusSubscriptionRuleTimeoutsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicebusSubscriptionRuleTimeoutsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicebusSubscriptionRuleTimeoutsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicebusSubscriptionRuleTimeoutsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicebusSubscriptionRuleTimeoutsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicebusSubscriptionRuleTimeoutsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicebusSubscriptionRuleTimeoutsOutputReference) ResetCreate() {
	_jsii_.InvokeVoid(
		s,
		"resetCreate",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicebusSubscriptionRuleTimeoutsOutputReference) ResetDelete() {
	_jsii_.InvokeVoid(
		s,
		"resetDelete",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicebusSubscriptionRuleTimeoutsOutputReference) ResetRead() {
	_jsii_.InvokeVoid(
		s,
		"resetRead",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicebusSubscriptionRuleTimeoutsOutputReference) ResetUpdate() {
	_jsii_.InvokeVoid(
		s,
		"resetUpdate",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicebusSubscriptionRuleTimeoutsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicebusSubscriptionRuleTimeoutsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

