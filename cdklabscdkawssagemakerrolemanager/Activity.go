package cdklabscdkawssagemakerrolemanager

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-aws-sagemaker-role-manager-go/cdklabscdkawssagemakerrolemanager/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdklabs/cdk-aws-sagemaker-role-manager-go/cdklabscdkawssagemakerrolemanager/internal"
)

// Experimental.
type Activity interface {
	constructs.Construct
	// Experimental.
	ActivityName() *string
	// Experimental.
	IsKMSCustomized() *bool
	// Experimental.
	SetIsKMSCustomized(val *bool)
	// Experimental.
	IsVPCCustomized() *bool
	// Experimental.
	SetIsVPCCustomized(val *bool)
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	Version() *float64
	// Creates policy with permissions of activity.
	//
	// Returns: - The policy that is created with the permissions of the activity.
	// Experimental.
	CreatePolicy(scope constructs.Construct) awsiam.Policy
	// Creates ML Activity service principal using ML Activity trust template.
	//
	// Returns: - The service principal of the ML Activity.
	// Experimental.
	CreatePrincipal() awsiam.ServicePrincipal
	// Creates role with permissions of activity.
	//
	// Returns: - The role that is created with the permissions of the activity.
	// Experimental.
	CreateRole(scope constructs.Construct, id *string, roleNameSuffix *string, roleDescription *string) awsiam.IRole
	// Experimental.
	CustomizeKMS(dataKeys *[]awskms.IKey, volumeKeys *[]awskms.IKey)
	// Experimental.
	CustomizeVPC(subnets *[]awsec2.ISubnet, securityGroups *[]awsec2.ISecurityGroup)
	// Grant permissions of activity to identity.
	//
	// Returns: - The grant with the permissions granted to the identity.
	// Experimental.
	GrantPermissionsTo(identity awsiam.IGrantable) awsiam.Grant
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Activity
type jsiiProxy_Activity struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_Activity) ActivityName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"activityName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Activity) IsKMSCustomized() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isKMSCustomized",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Activity) IsVPCCustomized() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isVPCCustomized",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Activity) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Activity) Version() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}


func (j *jsiiProxy_Activity)SetIsKMSCustomized(val *bool) {
	if err := j.validateSetIsKMSCustomizedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isKMSCustomized",
		val,
	)
}

func (j *jsiiProxy_Activity)SetIsVPCCustomized(val *bool) {
	if err := j.validateSetIsVPCCustomizedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isVPCCustomized",
		val,
	)
}

// Experimental.
func Activity_AccessAwsServices(scope constructs.Construct, id *string, options *AccessAwsServicesOptions) Activity {
	_init_.Initialize()

	if err := validateActivity_AccessAwsServicesParameters(scope, id, options); err != nil {
		panic(err)
	}
	var returns Activity

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-aws-sagemaker-role-manager.Activity",
		"accessAwsServices",
		[]interface{}{scope, id, options},
		&returns,
	)

	return returns
}

// Experimental.
func Activity_AccessS3AllResources(scope constructs.Construct, id *string, options *AccessS3AllResourcesOptions) Activity {
	_init_.Initialize()

	if err := validateActivity_AccessS3AllResourcesParameters(scope, id, options); err != nil {
		panic(err)
	}
	var returns Activity

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-aws-sagemaker-role-manager.Activity",
		"accessS3AllResources",
		[]interface{}{scope, id, options},
		&returns,
	)

	return returns
}

// Experimental.
func Activity_AccessS3AllResourcesV2(scope constructs.Construct, id *string, options *AccessS3AllResourcesV2Options) Activity {
	_init_.Initialize()

	if err := validateActivity_AccessS3AllResourcesV2Parameters(scope, id, options); err != nil {
		panic(err)
	}
	var returns Activity

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-aws-sagemaker-role-manager.Activity",
		"accessS3AllResourcesV2",
		[]interface{}{scope, id, options},
		&returns,
	)

	return returns
}

// Experimental.
func Activity_AccessS3Buckets(scope constructs.Construct, id *string, options *AccessS3BucketsOptions) Activity {
	_init_.Initialize()

	if err := validateActivity_AccessS3BucketsParameters(scope, id, options); err != nil {
		panic(err)
	}
	var returns Activity

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-aws-sagemaker-role-manager.Activity",
		"accessS3Buckets",
		[]interface{}{scope, id, options},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func Activity_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateActivity_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-aws-sagemaker-role-manager.Activity",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Activity_ManageEndpoints(scope constructs.Construct, id *string, options *ManageEndpointsOptions) Activity {
	_init_.Initialize()

	if err := validateActivity_ManageEndpointsParameters(scope, id, options); err != nil {
		panic(err)
	}
	var returns Activity

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-aws-sagemaker-role-manager.Activity",
		"manageEndpoints",
		[]interface{}{scope, id, options},
		&returns,
	)

	return returns
}

// Experimental.
func Activity_ManageExperiments(scope constructs.Construct, id *string, options *ManageExperimentsOptions) Activity {
	_init_.Initialize()

	if err := validateActivity_ManageExperimentsParameters(scope, id, options); err != nil {
		panic(err)
	}
	var returns Activity

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-aws-sagemaker-role-manager.Activity",
		"manageExperiments",
		[]interface{}{scope, id, options},
		&returns,
	)

	return returns
}

// Experimental.
func Activity_ManageGlueTables(scope constructs.Construct, id *string, options *ManageGlueTablesOptions) Activity {
	_init_.Initialize()

	if err := validateActivity_ManageGlueTablesParameters(scope, id, options); err != nil {
		panic(err)
	}
	var returns Activity

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-aws-sagemaker-role-manager.Activity",
		"manageGlueTables",
		[]interface{}{scope, id, options},
		&returns,
	)

	return returns
}

// Experimental.
func Activity_ManageJobs(scope constructs.Construct, id *string, options *ManageJobsOptions) Activity {
	_init_.Initialize()

	if err := validateActivity_ManageJobsParameters(scope, id, options); err != nil {
		panic(err)
	}
	var returns Activity

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-aws-sagemaker-role-manager.Activity",
		"manageJobs",
		[]interface{}{scope, id, options},
		&returns,
	)

	return returns
}

// Experimental.
func Activity_ManageModels(scope constructs.Construct, id *string, options *ManageModelsOptions) Activity {
	_init_.Initialize()

	if err := validateActivity_ManageModelsParameters(scope, id, options); err != nil {
		panic(err)
	}
	var returns Activity

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-aws-sagemaker-role-manager.Activity",
		"manageModels",
		[]interface{}{scope, id, options},
		&returns,
	)

	return returns
}

// Experimental.
func Activity_ManagePipelines(scope constructs.Construct, id *string, options *ManagePipelinesOptions) Activity {
	_init_.Initialize()

	if err := validateActivity_ManagePipelinesParameters(scope, id, options); err != nil {
		panic(err)
	}
	var returns Activity

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-aws-sagemaker-role-manager.Activity",
		"managePipelines",
		[]interface{}{scope, id, options},
		&returns,
	)

	return returns
}

// Experimental.
func Activity_MonitorModels(scope constructs.Construct, id *string, options *MonitorModelsOptions) Activity {
	_init_.Initialize()

	if err := validateActivity_MonitorModelsParameters(scope, id, options); err != nil {
		panic(err)
	}
	var returns Activity

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-aws-sagemaker-role-manager.Activity",
		"monitorModels",
		[]interface{}{scope, id, options},
		&returns,
	)

	return returns
}

// Experimental.
func Activity_QueryAthenaGroups(scope constructs.Construct, id *string, options *QueryAthenaGroupsOptions) Activity {
	_init_.Initialize()

	if err := validateActivity_QueryAthenaGroupsParameters(scope, id, options); err != nil {
		panic(err)
	}
	var returns Activity

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-aws-sagemaker-role-manager.Activity",
		"queryAthenaGroups",
		[]interface{}{scope, id, options},
		&returns,
	)

	return returns
}

// Experimental.
func Activity_RunStudioApps(scope constructs.Construct, id *string, options *RunStudioAppsOptions) Activity {
	_init_.Initialize()

	if err := validateActivity_RunStudioAppsParameters(scope, id, options); err != nil {
		panic(err)
	}
	var returns Activity

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-aws-sagemaker-role-manager.Activity",
		"runStudioApps",
		[]interface{}{scope, id, options},
		&returns,
	)

	return returns
}

// Experimental.
func Activity_RunStudioAppsV2(scope constructs.Construct, id *string, options *RunStudioAppsV2Options) Activity {
	_init_.Initialize()

	if err := validateActivity_RunStudioAppsV2Parameters(scope, id, options); err != nil {
		panic(err)
	}
	var returns Activity

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-aws-sagemaker-role-manager.Activity",
		"runStudioAppsV2",
		[]interface{}{scope, id, options},
		&returns,
	)

	return returns
}

// Experimental.
func Activity_VisualizeExperiments(scope constructs.Construct, id *string, options *VisualizeExperimentsOptions) Activity {
	_init_.Initialize()

	if err := validateActivity_VisualizeExperimentsParameters(scope, id, options); err != nil {
		panic(err)
	}
	var returns Activity

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-aws-sagemaker-role-manager.Activity",
		"visualizeExperiments",
		[]interface{}{scope, id, options},
		&returns,
	)

	return returns
}

func Activity_ACCESS_AWS_SERVICES() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdklabs/cdk-aws-sagemaker-role-manager.Activity",
		"ACCESS_AWS_SERVICES",
		&returns,
	)
	return returns
}

func Activity_ACCESS_S3_ALL_RESOURCES() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdklabs/cdk-aws-sagemaker-role-manager.Activity",
		"ACCESS_S3_ALL_RESOURCES",
		&returns,
	)
	return returns
}

func Activity_ACCESS_S3_BUCKETS() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdklabs/cdk-aws-sagemaker-role-manager.Activity",
		"ACCESS_S3_BUCKETS",
		&returns,
	)
	return returns
}

func Activity_ATHENA_WORKGROUP_NAMES_DEFAULT_VALUE() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@cdklabs/cdk-aws-sagemaker-role-manager.Activity",
		"ATHENA_WORKGROUP_NAMES_DEFAULT_VALUE",
		&returns,
	)
	return returns
}

func Activity_MANAGE_ENDPOINTS_ACTIVITY_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdklabs/cdk-aws-sagemaker-role-manager.Activity",
		"MANAGE_ENDPOINTS_ACTIVITY_NAME",
		&returns,
	)
	return returns
}

func Activity_MANAGE_EXPERIMENTS_ACTIVITY_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdklabs/cdk-aws-sagemaker-role-manager.Activity",
		"MANAGE_EXPERIMENTS_ACTIVITY_NAME",
		&returns,
	)
	return returns
}

func Activity_MANAGE_GLUE_TABLES_ACTIVITY_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdklabs/cdk-aws-sagemaker-role-manager.Activity",
		"MANAGE_GLUE_TABLES_ACTIVITY_NAME",
		&returns,
	)
	return returns
}

func Activity_MANAGE_JOBS_ACTIVITY_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdklabs/cdk-aws-sagemaker-role-manager.Activity",
		"MANAGE_JOBS_ACTIVITY_NAME",
		&returns,
	)
	return returns
}

func Activity_MANAGE_MODELS_ACTIVITY_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdklabs/cdk-aws-sagemaker-role-manager.Activity",
		"MANAGE_MODELS_ACTIVITY_NAME",
		&returns,
	)
	return returns
}

func Activity_MANAGE_PIPELINES_ACTIVITY_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdklabs/cdk-aws-sagemaker-role-manager.Activity",
		"MANAGE_PIPELINES_ACTIVITY_NAME",
		&returns,
	)
	return returns
}

func Activity_MONITOR_MODELS_ACTIVITY_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdklabs/cdk-aws-sagemaker-role-manager.Activity",
		"MONITOR_MODELS_ACTIVITY_NAME",
		&returns,
	)
	return returns
}

func Activity_QUERY_ATHENA_WORKGROUPS() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdklabs/cdk-aws-sagemaker-role-manager.Activity",
		"QUERY_ATHENA_WORKGROUPS",
		&returns,
	)
	return returns
}

func Activity_RUN_STUDIO_APPS() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdklabs/cdk-aws-sagemaker-role-manager.Activity",
		"RUN_STUDIO_APPS",
		&returns,
	)
	return returns
}

func Activity_VISUALIZE_EXPERIMENTS() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdklabs/cdk-aws-sagemaker-role-manager.Activity",
		"VISUALIZE_EXPERIMENTS",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_Activity) CreatePolicy(scope constructs.Construct) awsiam.Policy {
	if err := a.validateCreatePolicyParameters(scope); err != nil {
		panic(err)
	}
	var returns awsiam.Policy

	_jsii_.Invoke(
		a,
		"createPolicy",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Activity) CreatePrincipal() awsiam.ServicePrincipal {
	var returns awsiam.ServicePrincipal

	_jsii_.Invoke(
		a,
		"createPrincipal",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Activity) CreateRole(scope constructs.Construct, id *string, roleNameSuffix *string, roleDescription *string) awsiam.IRole {
	if err := a.validateCreateRoleParameters(scope, id, roleNameSuffix); err != nil {
		panic(err)
	}
	var returns awsiam.IRole

	_jsii_.Invoke(
		a,
		"createRole",
		[]interface{}{scope, id, roleNameSuffix, roleDescription},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Activity) CustomizeKMS(dataKeys *[]awskms.IKey, volumeKeys *[]awskms.IKey) {
	_jsii_.InvokeVoid(
		a,
		"customizeKMS",
		[]interface{}{dataKeys, volumeKeys},
	)
}

func (a *jsiiProxy_Activity) CustomizeVPC(subnets *[]awsec2.ISubnet, securityGroups *[]awsec2.ISecurityGroup) {
	_jsii_.InvokeVoid(
		a,
		"customizeVPC",
		[]interface{}{subnets, securityGroups},
	)
}

func (a *jsiiProxy_Activity) GrantPermissionsTo(identity awsiam.IGrantable) awsiam.Grant {
	if err := a.validateGrantPermissionsToParameters(identity); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		a,
		"grantPermissionsTo",
		[]interface{}{identity},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Activity) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

