package logzsubaccounttagrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v5/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v5/logzsubaccounttagrule/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/azurerm/r/logz_sub_account_tag_rule azurerm_logz_sub_account_tag_rule}.
type LogzSubAccountTagRule interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() *float64
	// Experimental.
	SetCount(val *float64)
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
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
	LogzSubAccountId() *string
	SetLogzSubAccountId(val *string)
	LogzSubAccountIdInput() *string
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
	SendAadLogs() interface{}
	SetSendAadLogs(val interface{})
	SendAadLogsInput() interface{}
	SendActivityLogs() interface{}
	SetSendActivityLogs(val interface{})
	SendActivityLogsInput() interface{}
	SendSubscriptionLogs() interface{}
	SetSendSubscriptionLogs(val interface{})
	SendSubscriptionLogsInput() interface{}
	TagFilter() LogzSubAccountTagRuleTagFilterList
	TagFilterInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() LogzSubAccountTagRuleTimeoutsOutputReference
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
	PutTagFilter(value interface{})
	PutTimeouts(value *LogzSubAccountTagRuleTimeouts)
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetSendAadLogs()
	ResetSendActivityLogs()
	ResetSendSubscriptionLogs()
	ResetTagFilter()
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

// The jsii proxy struct for LogzSubAccountTagRule
type jsiiProxy_LogzSubAccountTagRule struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_LogzSubAccountTagRule) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogzSubAccountTagRule) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogzSubAccountTagRule) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogzSubAccountTagRule) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogzSubAccountTagRule) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogzSubAccountTagRule) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogzSubAccountTagRule) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogzSubAccountTagRule) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogzSubAccountTagRule) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogzSubAccountTagRule) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogzSubAccountTagRule) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogzSubAccountTagRule) LogzSubAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logzSubAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogzSubAccountTagRule) LogzSubAccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logzSubAccountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogzSubAccountTagRule) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogzSubAccountTagRule) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogzSubAccountTagRule) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogzSubAccountTagRule) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogzSubAccountTagRule) SendAadLogs() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sendAadLogs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogzSubAccountTagRule) SendAadLogsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sendAadLogsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogzSubAccountTagRule) SendActivityLogs() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sendActivityLogs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogzSubAccountTagRule) SendActivityLogsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sendActivityLogsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogzSubAccountTagRule) SendSubscriptionLogs() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sendSubscriptionLogs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogzSubAccountTagRule) SendSubscriptionLogsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sendSubscriptionLogsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogzSubAccountTagRule) TagFilter() LogzSubAccountTagRuleTagFilterList {
	var returns LogzSubAccountTagRuleTagFilterList
	_jsii_.Get(
		j,
		"tagFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogzSubAccountTagRule) TagFilterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogzSubAccountTagRule) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogzSubAccountTagRule) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogzSubAccountTagRule) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogzSubAccountTagRule) Timeouts() LogzSubAccountTagRuleTimeoutsOutputReference {
	var returns LogzSubAccountTagRuleTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogzSubAccountTagRule) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/logz_sub_account_tag_rule azurerm_logz_sub_account_tag_rule} Resource.
func NewLogzSubAccountTagRule(scope constructs.Construct, id *string, config *LogzSubAccountTagRuleConfig) LogzSubAccountTagRule {
	_init_.Initialize()

	if err := validateNewLogzSubAccountTagRuleParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_LogzSubAccountTagRule{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.logzSubAccountTagRule.LogzSubAccountTagRule",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/logz_sub_account_tag_rule azurerm_logz_sub_account_tag_rule} Resource.
func NewLogzSubAccountTagRule_Override(l LogzSubAccountTagRule, scope constructs.Construct, id *string, config *LogzSubAccountTagRuleConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.logzSubAccountTagRule.LogzSubAccountTagRule",
		[]interface{}{scope, id, config},
		l,
	)
}

func (j *jsiiProxy_LogzSubAccountTagRule)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_LogzSubAccountTagRule)SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_LogzSubAccountTagRule)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_LogzSubAccountTagRule)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_LogzSubAccountTagRule)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_LogzSubAccountTagRule)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_LogzSubAccountTagRule)SetLogzSubAccountId(val *string) {
	if err := j.validateSetLogzSubAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logzSubAccountId",
		val,
	)
}

func (j *jsiiProxy_LogzSubAccountTagRule)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_LogzSubAccountTagRule)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_LogzSubAccountTagRule)SetSendAadLogs(val interface{}) {
	if err := j.validateSetSendAadLogsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sendAadLogs",
		val,
	)
}

func (j *jsiiProxy_LogzSubAccountTagRule)SetSendActivityLogs(val interface{}) {
	if err := j.validateSetSendActivityLogsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sendActivityLogs",
		val,
	)
}

func (j *jsiiProxy_LogzSubAccountTagRule)SetSendSubscriptionLogs(val interface{}) {
	if err := j.validateSetSendSubscriptionLogsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sendSubscriptionLogs",
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
func LogzSubAccountTagRule_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLogzSubAccountTagRule_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.logzSubAccountTagRule.LogzSubAccountTagRule",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func LogzSubAccountTagRule_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLogzSubAccountTagRule_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.logzSubAccountTagRule.LogzSubAccountTagRule",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func LogzSubAccountTagRule_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLogzSubAccountTagRule_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.logzSubAccountTagRule.LogzSubAccountTagRule",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func LogzSubAccountTagRule_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.logzSubAccountTagRule.LogzSubAccountTagRule",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (l *jsiiProxy_LogzSubAccountTagRule) AddOverride(path *string, value interface{}) {
	if err := l.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (l *jsiiProxy_LogzSubAccountTagRule) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (l *jsiiProxy_LogzSubAccountTagRule) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (l *jsiiProxy_LogzSubAccountTagRule) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (l *jsiiProxy_LogzSubAccountTagRule) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (l *jsiiProxy_LogzSubAccountTagRule) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (l *jsiiProxy_LogzSubAccountTagRule) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (l *jsiiProxy_LogzSubAccountTagRule) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (l *jsiiProxy_LogzSubAccountTagRule) GetStringAttribute(terraformAttribute *string) *string {
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

func (l *jsiiProxy_LogzSubAccountTagRule) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (l *jsiiProxy_LogzSubAccountTagRule) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := l.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogzSubAccountTagRule) OverrideLogicalId(newLogicalId *string) {
	if err := l.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (l *jsiiProxy_LogzSubAccountTagRule) PutTagFilter(value interface{}) {
	if err := l.validatePutTagFilterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putTagFilter",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LogzSubAccountTagRule) PutTimeouts(value *LogzSubAccountTagRuleTimeouts) {
	if err := l.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LogzSubAccountTagRule) ResetId() {
	_jsii_.InvokeVoid(
		l,
		"resetId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogzSubAccountTagRule) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		l,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogzSubAccountTagRule) ResetSendAadLogs() {
	_jsii_.InvokeVoid(
		l,
		"resetSendAadLogs",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogzSubAccountTagRule) ResetSendActivityLogs() {
	_jsii_.InvokeVoid(
		l,
		"resetSendActivityLogs",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogzSubAccountTagRule) ResetSendSubscriptionLogs() {
	_jsii_.InvokeVoid(
		l,
		"resetSendSubscriptionLogs",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogzSubAccountTagRule) ResetTagFilter() {
	_jsii_.InvokeVoid(
		l,
		"resetTagFilter",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogzSubAccountTagRule) ResetTimeouts() {
	_jsii_.InvokeVoid(
		l,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogzSubAccountTagRule) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogzSubAccountTagRule) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogzSubAccountTagRule) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogzSubAccountTagRule) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}
