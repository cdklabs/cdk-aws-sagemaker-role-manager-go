// Create roles and policies for ML Activities and ML Personas
package cdklabscdkawssagemakerrolemanager

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-aws-sagemaker-role-manager.AccessAwsServicesOptions",
		reflect.TypeOf((*AccessAwsServicesOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-aws-sagemaker-role-manager.AccessS3AllResourcesOptions",
		reflect.TypeOf((*AccessS3AllResourcesOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-aws-sagemaker-role-manager.AccessS3AllResourcesV2Options",
		reflect.TypeOf((*AccessS3AllResourcesV2Options)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-aws-sagemaker-role-manager.AccessS3BucketsOptions",
		reflect.TypeOf((*AccessS3BucketsOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/cdk-aws-sagemaker-role-manager.Activity",
		reflect.TypeOf((*Activity)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "activityName", GoGetter: "ActivityName"},
			_jsii_.MemberMethod{JsiiMethod: "createPolicy", GoMethod: "CreatePolicy"},
			_jsii_.MemberMethod{JsiiMethod: "createPrincipal", GoMethod: "CreatePrincipal"},
			_jsii_.MemberMethod{JsiiMethod: "createRole", GoMethod: "CreateRole"},
			_jsii_.MemberMethod{JsiiMethod: "customizeKMS", GoMethod: "CustomizeKMS"},
			_jsii_.MemberMethod{JsiiMethod: "customizeVPC", GoMethod: "CustomizeVPC"},
			_jsii_.MemberMethod{JsiiMethod: "grantPermissionsTo", GoMethod: "GrantPermissionsTo"},
			_jsii_.MemberProperty{JsiiProperty: "isKMSCustomized", GoGetter: "IsKMSCustomized"},
			_jsii_.MemberProperty{JsiiProperty: "isVPCCustomized", GoGetter: "IsVPCCustomized"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "version", GoGetter: "Version"},
		},
		func() interface{} {
			j := jsiiProxy_Activity{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-aws-sagemaker-role-manager.ActivityProps",
		reflect.TypeOf((*ActivityProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-aws-sagemaker-role-manager.KMSOptions",
		reflect.TypeOf((*KMSOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-aws-sagemaker-role-manager.ManageEndpointsOptions",
		reflect.TypeOf((*ManageEndpointsOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-aws-sagemaker-role-manager.ManageExperimentsOptions",
		reflect.TypeOf((*ManageExperimentsOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-aws-sagemaker-role-manager.ManageGlueTablesOptions",
		reflect.TypeOf((*ManageGlueTablesOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-aws-sagemaker-role-manager.ManageJobsOptions",
		reflect.TypeOf((*ManageJobsOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-aws-sagemaker-role-manager.ManageModelsOptions",
		reflect.TypeOf((*ManageModelsOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-aws-sagemaker-role-manager.ManagePipelinesOptions",
		reflect.TypeOf((*ManagePipelinesOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-aws-sagemaker-role-manager.MonitorModelsOptions",
		reflect.TypeOf((*MonitorModelsOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/cdk-aws-sagemaker-role-manager.Persona",
		reflect.TypeOf((*Persona)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "activities", GoGetter: "Activities"},
			_jsii_.MemberMethod{JsiiMethod: "createRole", GoMethod: "CreateRole"},
			_jsii_.MemberMethod{JsiiMethod: "customizeKMS", GoMethod: "CustomizeKMS"},
			_jsii_.MemberMethod{JsiiMethod: "customizeVPC", GoMethod: "CustomizeVPC"},
			_jsii_.MemberMethod{JsiiMethod: "grantPermissionsTo", GoMethod: "GrantPermissionsTo"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Persona{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-aws-sagemaker-role-manager.PersonaProps",
		reflect.TypeOf((*PersonaProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-aws-sagemaker-role-manager.QueryAthenaGroupsOptions",
		reflect.TypeOf((*QueryAthenaGroupsOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-aws-sagemaker-role-manager.RunStudioAppsOptions",
		reflect.TypeOf((*RunStudioAppsOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-aws-sagemaker-role-manager.RunStudioAppsV2Options",
		reflect.TypeOf((*RunStudioAppsV2Options)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-aws-sagemaker-role-manager.VPCOptions",
		reflect.TypeOf((*VPCOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-aws-sagemaker-role-manager.VisualizeExperimentsOptions",
		reflect.TypeOf((*VisualizeExperimentsOptions)(nil)).Elem(),
	)
}
