// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sentinelthreatintelligenceindicator

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/sentinelthreatintelligenceindicator/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.51.0/docs/resources/sentinel_threat_intelligence_indicator azurerm_sentinel_threat_intelligence_indicator}.
type SentinelThreatIntelligenceIndicator interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	Confidence() *float64
	SetConfidence(val *float64)
	ConfidenceInput() *float64
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CreatedBy() *string
	SetCreatedBy(val *string)
	CreatedByInput() *string
	CreatedOn() *string
	Defanged() cdktf.IResolvable
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	Extension() *string
	SetExtension(val *string)
	ExtensionInput() *string
	ExternalId() *string
	ExternalLastUpdatedTimeUtc() *string
	ExternalReference() SentinelThreatIntelligenceIndicatorExternalReferenceList
	ExternalReferenceInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GranularMarking() SentinelThreatIntelligenceIndicatorGranularMarkingList
	GranularMarkingInput() interface{}
	Guid() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	IndicatorType() *[]*string
	KillChainPhase() SentinelThreatIntelligenceIndicatorKillChainPhaseList
	KillChainPhaseInput() interface{}
	Language() *string
	SetLanguage(val *string)
	LanguageInput() *string
	LastUpdatedTimeUtc() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	ObjectMarkingRefs() *[]*string
	SetObjectMarkingRefs(val *[]*string)
	ObjectMarkingRefsInput() *[]*string
	ParsedPattern() SentinelThreatIntelligenceIndicatorParsedPatternList
	Pattern() *string
	SetPattern(val *string)
	PatternInput() *string
	PatternType() *string
	SetPatternType(val *string)
	PatternTypeInput() *string
	PatternVersion() *string
	SetPatternVersion(val *string)
	PatternVersionInput() *string
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
	Revoked() interface{}
	SetRevoked(val interface{})
	RevokedInput() interface{}
	Source() *string
	SetSource(val *string)
	SourceInput() *string
	Tags() *[]*string
	SetTags(val *[]*string)
	TagsInput() *[]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	ThreatTypes() *[]*string
	SetThreatTypes(val *[]*string)
	ThreatTypesInput() *[]*string
	Timeouts() SentinelThreatIntelligenceIndicatorTimeoutsOutputReference
	TimeoutsInput() interface{}
	ValidateFromUtc() *string
	SetValidateFromUtc(val *string)
	ValidateFromUtcInput() *string
	ValidateUntilUtc() *string
	SetValidateUntilUtc(val *string)
	ValidateUntilUtcInput() *string
	WorkspaceId() *string
	SetWorkspaceId(val *string)
	WorkspaceIdInput() *string
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
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
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutExternalReference(value interface{})
	PutGranularMarking(value interface{})
	PutKillChainPhase(value interface{})
	PutTimeouts(value *SentinelThreatIntelligenceIndicatorTimeouts)
	ResetConfidence()
	ResetCreatedBy()
	ResetDescription()
	ResetExtension()
	ResetExternalReference()
	ResetGranularMarking()
	ResetId()
	ResetKillChainPhase()
	ResetLanguage()
	ResetObjectMarkingRefs()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPatternVersion()
	ResetRevoked()
	ResetTags()
	ResetThreatTypes()
	ResetTimeouts()
	ResetValidateUntilUtc()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for SentinelThreatIntelligenceIndicator
type jsiiProxy_SentinelThreatIntelligenceIndicator struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_SentinelThreatIntelligenceIndicator) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelThreatIntelligenceIndicator) Confidence() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"confidence",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelThreatIntelligenceIndicator) ConfidenceInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"confidenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelThreatIntelligenceIndicator) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelThreatIntelligenceIndicator) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelThreatIntelligenceIndicator) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelThreatIntelligenceIndicator) CreatedBy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelThreatIntelligenceIndicator) CreatedByInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdByInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelThreatIntelligenceIndicator) CreatedOn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelThreatIntelligenceIndicator) Defanged() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"defanged",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelThreatIntelligenceIndicator) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelThreatIntelligenceIndicator) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelThreatIntelligenceIndicator) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelThreatIntelligenceIndicator) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelThreatIntelligenceIndicator) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelThreatIntelligenceIndicator) Extension() *string {
	var returns *string
	_jsii_.Get(
		j,
		"extension",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelThreatIntelligenceIndicator) ExtensionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"extensionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelThreatIntelligenceIndicator) ExternalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelThreatIntelligenceIndicator) ExternalLastUpdatedTimeUtc() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalLastUpdatedTimeUtc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelThreatIntelligenceIndicator) ExternalReference() SentinelThreatIntelligenceIndicatorExternalReferenceList {
	var returns SentinelThreatIntelligenceIndicatorExternalReferenceList
	_jsii_.Get(
		j,
		"externalReference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelThreatIntelligenceIndicator) ExternalReferenceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"externalReferenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelThreatIntelligenceIndicator) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelThreatIntelligenceIndicator) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelThreatIntelligenceIndicator) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelThreatIntelligenceIndicator) GranularMarking() SentinelThreatIntelligenceIndicatorGranularMarkingList {
	var returns SentinelThreatIntelligenceIndicatorGranularMarkingList
	_jsii_.Get(
		j,
		"granularMarking",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelThreatIntelligenceIndicator) GranularMarkingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"granularMarkingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelThreatIntelligenceIndicator) Guid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"guid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelThreatIntelligenceIndicator) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelThreatIntelligenceIndicator) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelThreatIntelligenceIndicator) IndicatorType() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"indicatorType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelThreatIntelligenceIndicator) KillChainPhase() SentinelThreatIntelligenceIndicatorKillChainPhaseList {
	var returns SentinelThreatIntelligenceIndicatorKillChainPhaseList
	_jsii_.Get(
		j,
		"killChainPhase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelThreatIntelligenceIndicator) KillChainPhaseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"killChainPhaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelThreatIntelligenceIndicator) Language() *string {
	var returns *string
	_jsii_.Get(
		j,
		"language",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelThreatIntelligenceIndicator) LanguageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"languageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelThreatIntelligenceIndicator) LastUpdatedTimeUtc() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastUpdatedTimeUtc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelThreatIntelligenceIndicator) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelThreatIntelligenceIndicator) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelThreatIntelligenceIndicator) ObjectMarkingRefs() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"objectMarkingRefs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelThreatIntelligenceIndicator) ObjectMarkingRefsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"objectMarkingRefsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelThreatIntelligenceIndicator) ParsedPattern() SentinelThreatIntelligenceIndicatorParsedPatternList {
	var returns SentinelThreatIntelligenceIndicatorParsedPatternList
	_jsii_.Get(
		j,
		"parsedPattern",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelThreatIntelligenceIndicator) Pattern() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pattern",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelThreatIntelligenceIndicator) PatternInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"patternInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelThreatIntelligenceIndicator) PatternType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"patternType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelThreatIntelligenceIndicator) PatternTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"patternTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelThreatIntelligenceIndicator) PatternVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"patternVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelThreatIntelligenceIndicator) PatternVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"patternVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelThreatIntelligenceIndicator) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelThreatIntelligenceIndicator) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelThreatIntelligenceIndicator) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelThreatIntelligenceIndicator) Revoked() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"revoked",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelThreatIntelligenceIndicator) RevokedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"revokedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelThreatIntelligenceIndicator) Source() *string {
	var returns *string
	_jsii_.Get(
		j,
		"source",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelThreatIntelligenceIndicator) SourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelThreatIntelligenceIndicator) Tags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelThreatIntelligenceIndicator) TagsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelThreatIntelligenceIndicator) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelThreatIntelligenceIndicator) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelThreatIntelligenceIndicator) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelThreatIntelligenceIndicator) ThreatTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"threatTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelThreatIntelligenceIndicator) ThreatTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"threatTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelThreatIntelligenceIndicator) Timeouts() SentinelThreatIntelligenceIndicatorTimeoutsOutputReference {
	var returns SentinelThreatIntelligenceIndicatorTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelThreatIntelligenceIndicator) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelThreatIntelligenceIndicator) ValidateFromUtc() *string {
	var returns *string
	_jsii_.Get(
		j,
		"validateFromUtc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelThreatIntelligenceIndicator) ValidateFromUtcInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"validateFromUtcInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelThreatIntelligenceIndicator) ValidateUntilUtc() *string {
	var returns *string
	_jsii_.Get(
		j,
		"validateUntilUtc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelThreatIntelligenceIndicator) ValidateUntilUtcInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"validateUntilUtcInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelThreatIntelligenceIndicator) WorkspaceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workspaceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelThreatIntelligenceIndicator) WorkspaceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workspaceIdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.51.0/docs/resources/sentinel_threat_intelligence_indicator azurerm_sentinel_threat_intelligence_indicator} Resource.
func NewSentinelThreatIntelligenceIndicator(scope constructs.Construct, id *string, config *SentinelThreatIntelligenceIndicatorConfig) SentinelThreatIntelligenceIndicator {
	_init_.Initialize()

	if err := validateNewSentinelThreatIntelligenceIndicatorParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_SentinelThreatIntelligenceIndicator{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.sentinelThreatIntelligenceIndicator.SentinelThreatIntelligenceIndicator",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.51.0/docs/resources/sentinel_threat_intelligence_indicator azurerm_sentinel_threat_intelligence_indicator} Resource.
func NewSentinelThreatIntelligenceIndicator_Override(s SentinelThreatIntelligenceIndicator, scope constructs.Construct, id *string, config *SentinelThreatIntelligenceIndicatorConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.sentinelThreatIntelligenceIndicator.SentinelThreatIntelligenceIndicator",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SentinelThreatIntelligenceIndicator)SetConfidence(val *float64) {
	if err := j.validateSetConfidenceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"confidence",
		val,
	)
}

func (j *jsiiProxy_SentinelThreatIntelligenceIndicator)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_SentinelThreatIntelligenceIndicator)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SentinelThreatIntelligenceIndicator)SetCreatedBy(val *string) {
	if err := j.validateSetCreatedByParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createdBy",
		val,
	)
}

func (j *jsiiProxy_SentinelThreatIntelligenceIndicator)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SentinelThreatIntelligenceIndicator)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_SentinelThreatIntelligenceIndicator)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_SentinelThreatIntelligenceIndicator)SetExtension(val *string) {
	if err := j.validateSetExtensionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"extension",
		val,
	)
}

func (j *jsiiProxy_SentinelThreatIntelligenceIndicator)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_SentinelThreatIntelligenceIndicator)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_SentinelThreatIntelligenceIndicator)SetLanguage(val *string) {
	if err := j.validateSetLanguageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"language",
		val,
	)
}

func (j *jsiiProxy_SentinelThreatIntelligenceIndicator)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SentinelThreatIntelligenceIndicator)SetObjectMarkingRefs(val *[]*string) {
	if err := j.validateSetObjectMarkingRefsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"objectMarkingRefs",
		val,
	)
}

func (j *jsiiProxy_SentinelThreatIntelligenceIndicator)SetPattern(val *string) {
	if err := j.validateSetPatternParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pattern",
		val,
	)
}

func (j *jsiiProxy_SentinelThreatIntelligenceIndicator)SetPatternType(val *string) {
	if err := j.validateSetPatternTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"patternType",
		val,
	)
}

func (j *jsiiProxy_SentinelThreatIntelligenceIndicator)SetPatternVersion(val *string) {
	if err := j.validateSetPatternVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"patternVersion",
		val,
	)
}

func (j *jsiiProxy_SentinelThreatIntelligenceIndicator)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_SentinelThreatIntelligenceIndicator)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_SentinelThreatIntelligenceIndicator)SetRevoked(val interface{}) {
	if err := j.validateSetRevokedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"revoked",
		val,
	)
}

func (j *jsiiProxy_SentinelThreatIntelligenceIndicator)SetSource(val *string) {
	if err := j.validateSetSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"source",
		val,
	)
}

func (j *jsiiProxy_SentinelThreatIntelligenceIndicator)SetTags(val *[]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_SentinelThreatIntelligenceIndicator)SetThreatTypes(val *[]*string) {
	if err := j.validateSetThreatTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"threatTypes",
		val,
	)
}

func (j *jsiiProxy_SentinelThreatIntelligenceIndicator)SetValidateFromUtc(val *string) {
	if err := j.validateSetValidateFromUtcParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"validateFromUtc",
		val,
	)
}

func (j *jsiiProxy_SentinelThreatIntelligenceIndicator)SetValidateUntilUtc(val *string) {
	if err := j.validateSetValidateUntilUtcParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"validateUntilUtc",
		val,
	)
}

func (j *jsiiProxy_SentinelThreatIntelligenceIndicator)SetWorkspaceId(val *string) {
	if err := j.validateSetWorkspaceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workspaceId",
		val,
	)
}

// Generates CDKTF code for importing a SentinelThreatIntelligenceIndicator resource upon running "cdktf plan <stack-name>".
func SentinelThreatIntelligenceIndicator_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateSentinelThreatIntelligenceIndicator_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.sentinelThreatIntelligenceIndicator.SentinelThreatIntelligenceIndicator",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
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
func SentinelThreatIntelligenceIndicator_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSentinelThreatIntelligenceIndicator_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.sentinelThreatIntelligenceIndicator.SentinelThreatIntelligenceIndicator",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SentinelThreatIntelligenceIndicator_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSentinelThreatIntelligenceIndicator_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.sentinelThreatIntelligenceIndicator.SentinelThreatIntelligenceIndicator",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SentinelThreatIntelligenceIndicator_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSentinelThreatIntelligenceIndicator_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.sentinelThreatIntelligenceIndicator.SentinelThreatIntelligenceIndicator",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SentinelThreatIntelligenceIndicator_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.sentinelThreatIntelligenceIndicator.SentinelThreatIntelligenceIndicator",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_SentinelThreatIntelligenceIndicator) AddMoveTarget(moveTarget *string) {
	if err := s.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (s *jsiiProxy_SentinelThreatIntelligenceIndicator) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_SentinelThreatIntelligenceIndicator) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SentinelThreatIntelligenceIndicator) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SentinelThreatIntelligenceIndicator) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SentinelThreatIntelligenceIndicator) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SentinelThreatIntelligenceIndicator) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SentinelThreatIntelligenceIndicator) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SentinelThreatIntelligenceIndicator) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SentinelThreatIntelligenceIndicator) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SentinelThreatIntelligenceIndicator) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SentinelThreatIntelligenceIndicator) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelThreatIntelligenceIndicator) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := s.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (s *jsiiProxy_SentinelThreatIntelligenceIndicator) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelThreatIntelligenceIndicator) MoveFromId(id *string) {
	if err := s.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveFromId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SentinelThreatIntelligenceIndicator) MoveTo(moveTarget *string, index interface{}) {
	if err := s.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (s *jsiiProxy_SentinelThreatIntelligenceIndicator) MoveToId(id *string) {
	if err := s.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveToId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SentinelThreatIntelligenceIndicator) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_SentinelThreatIntelligenceIndicator) PutExternalReference(value interface{}) {
	if err := s.validatePutExternalReferenceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putExternalReference",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SentinelThreatIntelligenceIndicator) PutGranularMarking(value interface{}) {
	if err := s.validatePutGranularMarkingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putGranularMarking",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SentinelThreatIntelligenceIndicator) PutKillChainPhase(value interface{}) {
	if err := s.validatePutKillChainPhaseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putKillChainPhase",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SentinelThreatIntelligenceIndicator) PutTimeouts(value *SentinelThreatIntelligenceIndicatorTimeouts) {
	if err := s.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SentinelThreatIntelligenceIndicator) ResetConfidence() {
	_jsii_.InvokeVoid(
		s,
		"resetConfidence",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelThreatIntelligenceIndicator) ResetCreatedBy() {
	_jsii_.InvokeVoid(
		s,
		"resetCreatedBy",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelThreatIntelligenceIndicator) ResetDescription() {
	_jsii_.InvokeVoid(
		s,
		"resetDescription",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelThreatIntelligenceIndicator) ResetExtension() {
	_jsii_.InvokeVoid(
		s,
		"resetExtension",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelThreatIntelligenceIndicator) ResetExternalReference() {
	_jsii_.InvokeVoid(
		s,
		"resetExternalReference",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelThreatIntelligenceIndicator) ResetGranularMarking() {
	_jsii_.InvokeVoid(
		s,
		"resetGranularMarking",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelThreatIntelligenceIndicator) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelThreatIntelligenceIndicator) ResetKillChainPhase() {
	_jsii_.InvokeVoid(
		s,
		"resetKillChainPhase",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelThreatIntelligenceIndicator) ResetLanguage() {
	_jsii_.InvokeVoid(
		s,
		"resetLanguage",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelThreatIntelligenceIndicator) ResetObjectMarkingRefs() {
	_jsii_.InvokeVoid(
		s,
		"resetObjectMarkingRefs",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelThreatIntelligenceIndicator) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelThreatIntelligenceIndicator) ResetPatternVersion() {
	_jsii_.InvokeVoid(
		s,
		"resetPatternVersion",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelThreatIntelligenceIndicator) ResetRevoked() {
	_jsii_.InvokeVoid(
		s,
		"resetRevoked",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelThreatIntelligenceIndicator) ResetTags() {
	_jsii_.InvokeVoid(
		s,
		"resetTags",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelThreatIntelligenceIndicator) ResetThreatTypes() {
	_jsii_.InvokeVoid(
		s,
		"resetThreatTypes",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelThreatIntelligenceIndicator) ResetTimeouts() {
	_jsii_.InvokeVoid(
		s,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelThreatIntelligenceIndicator) ResetValidateUntilUtc() {
	_jsii_.InvokeVoid(
		s,
		"resetValidateUntilUtc",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelThreatIntelligenceIndicator) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelThreatIntelligenceIndicator) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelThreatIntelligenceIndicator) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelThreatIntelligenceIndicator) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelThreatIntelligenceIndicator) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelThreatIntelligenceIndicator) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

