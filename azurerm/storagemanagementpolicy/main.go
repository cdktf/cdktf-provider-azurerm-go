// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storagemanagementpolicy

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-azurerm.storageManagementPolicy.StorageManagementPolicy",
		reflect.TypeOf((*StorageManagementPolicy)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "forEach", GoGetter: "ForEach"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberMethod{JsiiMethod: "putRule", GoMethod: "PutRule"},
			_jsii_.MemberMethod{JsiiMethod: "putTimeouts", GoMethod: "PutTimeouts"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetRule", GoMethod: "ResetRule"},
			_jsii_.MemberMethod{JsiiMethod: "resetTimeouts", GoMethod: "ResetTimeouts"},
			_jsii_.MemberProperty{JsiiProperty: "rule", GoGetter: "Rule"},
			_jsii_.MemberProperty{JsiiProperty: "ruleInput", GoGetter: "RuleInput"},
			_jsii_.MemberProperty{JsiiProperty: "storageAccountId", GoGetter: "StorageAccountId"},
			_jsii_.MemberProperty{JsiiProperty: "storageAccountIdInput", GoGetter: "StorageAccountIdInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "timeouts", GoGetter: "Timeouts"},
			_jsii_.MemberProperty{JsiiProperty: "timeoutsInput", GoGetter: "TimeoutsInput"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_StorageManagementPolicy{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-azurerm.storageManagementPolicy.StorageManagementPolicyConfig",
		reflect.TypeOf((*StorageManagementPolicyConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-azurerm.storageManagementPolicy.StorageManagementPolicyRule",
		reflect.TypeOf((*StorageManagementPolicyRule)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-azurerm.storageManagementPolicy.StorageManagementPolicyRuleActions",
		reflect.TypeOf((*StorageManagementPolicyRuleActions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-azurerm.storageManagementPolicy.StorageManagementPolicyRuleActionsBaseBlob",
		reflect.TypeOf((*StorageManagementPolicyRuleActionsBaseBlob)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-azurerm.storageManagementPolicy.StorageManagementPolicyRuleActionsBaseBlobOutputReference",
		reflect.TypeOf((*StorageManagementPolicyRuleActionsBaseBlobOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "autoTierToHotFromCoolEnabled", GoGetter: "AutoTierToHotFromCoolEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "autoTierToHotFromCoolEnabledInput", GoGetter: "AutoTierToHotFromCoolEnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "deleteAfterDaysSinceCreationGreaterThan", GoGetter: "DeleteAfterDaysSinceCreationGreaterThan"},
			_jsii_.MemberProperty{JsiiProperty: "deleteAfterDaysSinceCreationGreaterThanInput", GoGetter: "DeleteAfterDaysSinceCreationGreaterThanInput"},
			_jsii_.MemberProperty{JsiiProperty: "deleteAfterDaysSinceLastAccessTimeGreaterThan", GoGetter: "DeleteAfterDaysSinceLastAccessTimeGreaterThan"},
			_jsii_.MemberProperty{JsiiProperty: "deleteAfterDaysSinceLastAccessTimeGreaterThanInput", GoGetter: "DeleteAfterDaysSinceLastAccessTimeGreaterThanInput"},
			_jsii_.MemberProperty{JsiiProperty: "deleteAfterDaysSinceModificationGreaterThan", GoGetter: "DeleteAfterDaysSinceModificationGreaterThan"},
			_jsii_.MemberProperty{JsiiProperty: "deleteAfterDaysSinceModificationGreaterThanInput", GoGetter: "DeleteAfterDaysSinceModificationGreaterThanInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetAutoTierToHotFromCoolEnabled", GoMethod: "ResetAutoTierToHotFromCoolEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetDeleteAfterDaysSinceCreationGreaterThan", GoMethod: "ResetDeleteAfterDaysSinceCreationGreaterThan"},
			_jsii_.MemberMethod{JsiiMethod: "resetDeleteAfterDaysSinceLastAccessTimeGreaterThan", GoMethod: "ResetDeleteAfterDaysSinceLastAccessTimeGreaterThan"},
			_jsii_.MemberMethod{JsiiMethod: "resetDeleteAfterDaysSinceModificationGreaterThan", GoMethod: "ResetDeleteAfterDaysSinceModificationGreaterThan"},
			_jsii_.MemberMethod{JsiiMethod: "resetTierToArchiveAfterDaysSinceCreationGreaterThan", GoMethod: "ResetTierToArchiveAfterDaysSinceCreationGreaterThan"},
			_jsii_.MemberMethod{JsiiMethod: "resetTierToArchiveAfterDaysSinceLastAccessTimeGreaterThan", GoMethod: "ResetTierToArchiveAfterDaysSinceLastAccessTimeGreaterThan"},
			_jsii_.MemberMethod{JsiiMethod: "resetTierToArchiveAfterDaysSinceLastTierChangeGreaterThan", GoMethod: "ResetTierToArchiveAfterDaysSinceLastTierChangeGreaterThan"},
			_jsii_.MemberMethod{JsiiMethod: "resetTierToArchiveAfterDaysSinceModificationGreaterThan", GoMethod: "ResetTierToArchiveAfterDaysSinceModificationGreaterThan"},
			_jsii_.MemberMethod{JsiiMethod: "resetTierToCoolAfterDaysSinceCreationGreaterThan", GoMethod: "ResetTierToCoolAfterDaysSinceCreationGreaterThan"},
			_jsii_.MemberMethod{JsiiMethod: "resetTierToCoolAfterDaysSinceLastAccessTimeGreaterThan", GoMethod: "ResetTierToCoolAfterDaysSinceLastAccessTimeGreaterThan"},
			_jsii_.MemberMethod{JsiiMethod: "resetTierToCoolAfterDaysSinceModificationGreaterThan", GoMethod: "ResetTierToCoolAfterDaysSinceModificationGreaterThan"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "tierToArchiveAfterDaysSinceCreationGreaterThan", GoGetter: "TierToArchiveAfterDaysSinceCreationGreaterThan"},
			_jsii_.MemberProperty{JsiiProperty: "tierToArchiveAfterDaysSinceCreationGreaterThanInput", GoGetter: "TierToArchiveAfterDaysSinceCreationGreaterThanInput"},
			_jsii_.MemberProperty{JsiiProperty: "tierToArchiveAfterDaysSinceLastAccessTimeGreaterThan", GoGetter: "TierToArchiveAfterDaysSinceLastAccessTimeGreaterThan"},
			_jsii_.MemberProperty{JsiiProperty: "tierToArchiveAfterDaysSinceLastAccessTimeGreaterThanInput", GoGetter: "TierToArchiveAfterDaysSinceLastAccessTimeGreaterThanInput"},
			_jsii_.MemberProperty{JsiiProperty: "tierToArchiveAfterDaysSinceLastTierChangeGreaterThan", GoGetter: "TierToArchiveAfterDaysSinceLastTierChangeGreaterThan"},
			_jsii_.MemberProperty{JsiiProperty: "tierToArchiveAfterDaysSinceLastTierChangeGreaterThanInput", GoGetter: "TierToArchiveAfterDaysSinceLastTierChangeGreaterThanInput"},
			_jsii_.MemberProperty{JsiiProperty: "tierToArchiveAfterDaysSinceModificationGreaterThan", GoGetter: "TierToArchiveAfterDaysSinceModificationGreaterThan"},
			_jsii_.MemberProperty{JsiiProperty: "tierToArchiveAfterDaysSinceModificationGreaterThanInput", GoGetter: "TierToArchiveAfterDaysSinceModificationGreaterThanInput"},
			_jsii_.MemberProperty{JsiiProperty: "tierToCoolAfterDaysSinceCreationGreaterThan", GoGetter: "TierToCoolAfterDaysSinceCreationGreaterThan"},
			_jsii_.MemberProperty{JsiiProperty: "tierToCoolAfterDaysSinceCreationGreaterThanInput", GoGetter: "TierToCoolAfterDaysSinceCreationGreaterThanInput"},
			_jsii_.MemberProperty{JsiiProperty: "tierToCoolAfterDaysSinceLastAccessTimeGreaterThan", GoGetter: "TierToCoolAfterDaysSinceLastAccessTimeGreaterThan"},
			_jsii_.MemberProperty{JsiiProperty: "tierToCoolAfterDaysSinceLastAccessTimeGreaterThanInput", GoGetter: "TierToCoolAfterDaysSinceLastAccessTimeGreaterThanInput"},
			_jsii_.MemberProperty{JsiiProperty: "tierToCoolAfterDaysSinceModificationGreaterThan", GoGetter: "TierToCoolAfterDaysSinceModificationGreaterThan"},
			_jsii_.MemberProperty{JsiiProperty: "tierToCoolAfterDaysSinceModificationGreaterThanInput", GoGetter: "TierToCoolAfterDaysSinceModificationGreaterThanInput"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_StorageManagementPolicyRuleActionsBaseBlobOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-azurerm.storageManagementPolicy.StorageManagementPolicyRuleActionsOutputReference",
		reflect.TypeOf((*StorageManagementPolicyRuleActionsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "baseBlob", GoGetter: "BaseBlob"},
			_jsii_.MemberProperty{JsiiProperty: "baseBlobInput", GoGetter: "BaseBlobInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "putBaseBlob", GoMethod: "PutBaseBlob"},
			_jsii_.MemberMethod{JsiiMethod: "putSnapshot", GoMethod: "PutSnapshot"},
			_jsii_.MemberMethod{JsiiMethod: "putVersion", GoMethod: "PutVersion"},
			_jsii_.MemberMethod{JsiiMethod: "resetBaseBlob", GoMethod: "ResetBaseBlob"},
			_jsii_.MemberMethod{JsiiMethod: "resetSnapshot", GoMethod: "ResetSnapshot"},
			_jsii_.MemberMethod{JsiiMethod: "resetVersion", GoMethod: "ResetVersion"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "snapshot", GoGetter: "Snapshot"},
			_jsii_.MemberProperty{JsiiProperty: "snapshotInput", GoGetter: "SnapshotInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "version", GoGetter: "Version"},
			_jsii_.MemberProperty{JsiiProperty: "versionInput", GoGetter: "VersionInput"},
		},
		func() interface{} {
			j := jsiiProxy_StorageManagementPolicyRuleActionsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-azurerm.storageManagementPolicy.StorageManagementPolicyRuleActionsSnapshot",
		reflect.TypeOf((*StorageManagementPolicyRuleActionsSnapshot)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-azurerm.storageManagementPolicy.StorageManagementPolicyRuleActionsSnapshotOutputReference",
		reflect.TypeOf((*StorageManagementPolicyRuleActionsSnapshotOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "changeTierToArchiveAfterDaysSinceCreation", GoGetter: "ChangeTierToArchiveAfterDaysSinceCreation"},
			_jsii_.MemberProperty{JsiiProperty: "changeTierToArchiveAfterDaysSinceCreationInput", GoGetter: "ChangeTierToArchiveAfterDaysSinceCreationInput"},
			_jsii_.MemberProperty{JsiiProperty: "changeTierToCoolAfterDaysSinceCreation", GoGetter: "ChangeTierToCoolAfterDaysSinceCreation"},
			_jsii_.MemberProperty{JsiiProperty: "changeTierToCoolAfterDaysSinceCreationInput", GoGetter: "ChangeTierToCoolAfterDaysSinceCreationInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "deleteAfterDaysSinceCreationGreaterThan", GoGetter: "DeleteAfterDaysSinceCreationGreaterThan"},
			_jsii_.MemberProperty{JsiiProperty: "deleteAfterDaysSinceCreationGreaterThanInput", GoGetter: "DeleteAfterDaysSinceCreationGreaterThanInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetChangeTierToArchiveAfterDaysSinceCreation", GoMethod: "ResetChangeTierToArchiveAfterDaysSinceCreation"},
			_jsii_.MemberMethod{JsiiMethod: "resetChangeTierToCoolAfterDaysSinceCreation", GoMethod: "ResetChangeTierToCoolAfterDaysSinceCreation"},
			_jsii_.MemberMethod{JsiiMethod: "resetDeleteAfterDaysSinceCreationGreaterThan", GoMethod: "ResetDeleteAfterDaysSinceCreationGreaterThan"},
			_jsii_.MemberMethod{JsiiMethod: "resetTierToArchiveAfterDaysSinceLastTierChangeGreaterThan", GoMethod: "ResetTierToArchiveAfterDaysSinceLastTierChangeGreaterThan"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "tierToArchiveAfterDaysSinceLastTierChangeGreaterThan", GoGetter: "TierToArchiveAfterDaysSinceLastTierChangeGreaterThan"},
			_jsii_.MemberProperty{JsiiProperty: "tierToArchiveAfterDaysSinceLastTierChangeGreaterThanInput", GoGetter: "TierToArchiveAfterDaysSinceLastTierChangeGreaterThanInput"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_StorageManagementPolicyRuleActionsSnapshotOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-azurerm.storageManagementPolicy.StorageManagementPolicyRuleActionsVersion",
		reflect.TypeOf((*StorageManagementPolicyRuleActionsVersion)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-azurerm.storageManagementPolicy.StorageManagementPolicyRuleActionsVersionOutputReference",
		reflect.TypeOf((*StorageManagementPolicyRuleActionsVersionOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "changeTierToArchiveAfterDaysSinceCreation", GoGetter: "ChangeTierToArchiveAfterDaysSinceCreation"},
			_jsii_.MemberProperty{JsiiProperty: "changeTierToArchiveAfterDaysSinceCreationInput", GoGetter: "ChangeTierToArchiveAfterDaysSinceCreationInput"},
			_jsii_.MemberProperty{JsiiProperty: "changeTierToCoolAfterDaysSinceCreation", GoGetter: "ChangeTierToCoolAfterDaysSinceCreation"},
			_jsii_.MemberProperty{JsiiProperty: "changeTierToCoolAfterDaysSinceCreationInput", GoGetter: "ChangeTierToCoolAfterDaysSinceCreationInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "deleteAfterDaysSinceCreation", GoGetter: "DeleteAfterDaysSinceCreation"},
			_jsii_.MemberProperty{JsiiProperty: "deleteAfterDaysSinceCreationInput", GoGetter: "DeleteAfterDaysSinceCreationInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetChangeTierToArchiveAfterDaysSinceCreation", GoMethod: "ResetChangeTierToArchiveAfterDaysSinceCreation"},
			_jsii_.MemberMethod{JsiiMethod: "resetChangeTierToCoolAfterDaysSinceCreation", GoMethod: "ResetChangeTierToCoolAfterDaysSinceCreation"},
			_jsii_.MemberMethod{JsiiMethod: "resetDeleteAfterDaysSinceCreation", GoMethod: "ResetDeleteAfterDaysSinceCreation"},
			_jsii_.MemberMethod{JsiiMethod: "resetTierToArchiveAfterDaysSinceLastTierChangeGreaterThan", GoMethod: "ResetTierToArchiveAfterDaysSinceLastTierChangeGreaterThan"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "tierToArchiveAfterDaysSinceLastTierChangeGreaterThan", GoGetter: "TierToArchiveAfterDaysSinceLastTierChangeGreaterThan"},
			_jsii_.MemberProperty{JsiiProperty: "tierToArchiveAfterDaysSinceLastTierChangeGreaterThanInput", GoGetter: "TierToArchiveAfterDaysSinceLastTierChangeGreaterThanInput"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_StorageManagementPolicyRuleActionsVersionOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-azurerm.storageManagementPolicy.StorageManagementPolicyRuleFilters",
		reflect.TypeOf((*StorageManagementPolicyRuleFilters)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-azurerm.storageManagementPolicy.StorageManagementPolicyRuleFiltersMatchBlobIndexTag",
		reflect.TypeOf((*StorageManagementPolicyRuleFiltersMatchBlobIndexTag)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-azurerm.storageManagementPolicy.StorageManagementPolicyRuleFiltersMatchBlobIndexTagList",
		reflect.TypeOf((*StorageManagementPolicyRuleFiltersMatchBlobIndexTagList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_StorageManagementPolicyRuleFiltersMatchBlobIndexTagList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-azurerm.storageManagementPolicy.StorageManagementPolicyRuleFiltersMatchBlobIndexTagOutputReference",
		reflect.TypeOf((*StorageManagementPolicyRuleFiltersMatchBlobIndexTagOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "operation", GoGetter: "Operation"},
			_jsii_.MemberProperty{JsiiProperty: "operationInput", GoGetter: "OperationInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetOperation", GoMethod: "ResetOperation"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
			_jsii_.MemberProperty{JsiiProperty: "valueInput", GoGetter: "ValueInput"},
		},
		func() interface{} {
			j := jsiiProxy_StorageManagementPolicyRuleFiltersMatchBlobIndexTagOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-azurerm.storageManagementPolicy.StorageManagementPolicyRuleFiltersOutputReference",
		reflect.TypeOf((*StorageManagementPolicyRuleFiltersOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "blobTypes", GoGetter: "BlobTypes"},
			_jsii_.MemberProperty{JsiiProperty: "blobTypesInput", GoGetter: "BlobTypesInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "matchBlobIndexTag", GoGetter: "MatchBlobIndexTag"},
			_jsii_.MemberProperty{JsiiProperty: "matchBlobIndexTagInput", GoGetter: "MatchBlobIndexTagInput"},
			_jsii_.MemberProperty{JsiiProperty: "prefixMatch", GoGetter: "PrefixMatch"},
			_jsii_.MemberProperty{JsiiProperty: "prefixMatchInput", GoGetter: "PrefixMatchInput"},
			_jsii_.MemberMethod{JsiiMethod: "putMatchBlobIndexTag", GoMethod: "PutMatchBlobIndexTag"},
			_jsii_.MemberMethod{JsiiMethod: "resetMatchBlobIndexTag", GoMethod: "ResetMatchBlobIndexTag"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrefixMatch", GoMethod: "ResetPrefixMatch"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_StorageManagementPolicyRuleFiltersOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-azurerm.storageManagementPolicy.StorageManagementPolicyRuleList",
		reflect.TypeOf((*StorageManagementPolicyRuleList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_StorageManagementPolicyRuleList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-azurerm.storageManagementPolicy.StorageManagementPolicyRuleOutputReference",
		reflect.TypeOf((*StorageManagementPolicyRuleOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "actions", GoGetter: "Actions"},
			_jsii_.MemberProperty{JsiiProperty: "actionsInput", GoGetter: "ActionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "enabledInput", GoGetter: "EnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "filters", GoGetter: "Filters"},
			_jsii_.MemberProperty{JsiiProperty: "filtersInput", GoGetter: "FiltersInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberMethod{JsiiMethod: "putActions", GoMethod: "PutActions"},
			_jsii_.MemberMethod{JsiiMethod: "putFilters", GoMethod: "PutFilters"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_StorageManagementPolicyRuleOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-azurerm.storageManagementPolicy.StorageManagementPolicyTimeouts",
		reflect.TypeOf((*StorageManagementPolicyTimeouts)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-azurerm.storageManagementPolicy.StorageManagementPolicyTimeoutsOutputReference",
		reflect.TypeOf((*StorageManagementPolicyTimeoutsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "create", GoGetter: "Create"},
			_jsii_.MemberProperty{JsiiProperty: "createInput", GoGetter: "CreateInput"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "delete", GoGetter: "Delete"},
			_jsii_.MemberProperty{JsiiProperty: "deleteInput", GoGetter: "DeleteInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "read", GoGetter: "Read"},
			_jsii_.MemberProperty{JsiiProperty: "readInput", GoGetter: "ReadInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetCreate", GoMethod: "ResetCreate"},
			_jsii_.MemberMethod{JsiiMethod: "resetDelete", GoMethod: "ResetDelete"},
			_jsii_.MemberMethod{JsiiMethod: "resetRead", GoMethod: "ResetRead"},
			_jsii_.MemberMethod{JsiiMethod: "resetUpdate", GoMethod: "ResetUpdate"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "update", GoGetter: "Update"},
			_jsii_.MemberProperty{JsiiProperty: "updateInput", GoGetter: "UpdateInput"},
		},
		func() interface{} {
			j := jsiiProxy_StorageManagementPolicyTimeoutsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}
