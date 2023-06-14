package cdklabscdkawssagemakerrolemanager

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecr"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// Experimental.
type ActivityProps struct {
	// Name of the SageMaker Activity.
	//
	// This name will be used to name the IAM policy that is created from this Activity.
	// Experimental.
	ActivityName *string `field:"required" json:"activityName" yaml:"activityName"`
	// Whether the activity supports customization for kms data keys and volume keys.
	// Experimental.
	IsCustomizationAvailableForKMS *bool `field:"required" json:"isCustomizationAvailableForKMS" yaml:"isCustomizationAvailableForKMS"`
	// Whether the activity supports customization for vpc subnets and vpc security groups.
	// Experimental.
	IsCustomizationAvailableForVPC *bool `field:"required" json:"isCustomizationAvailableForVPC" yaml:"isCustomizationAvailableForVPC"`
	// Names of the Athena workgroups to give query permissions.
	// Experimental.
	AthenaWorkgroupNames *[]*string `field:"optional" json:"athenaWorkgroupNames" yaml:"athenaWorkgroupNames"`
	// ECR Repositories to give image pull permissions.
	// Experimental.
	EcrRepositories *[]awsecr.IRepository `field:"optional" json:"ecrRepositories" yaml:"ecrRepositories"`
	// Names of the Glue Databases to give permissions to search tables.
	// Experimental.
	GlueDatabaseNames *[]*string `field:"optional" json:"glueDatabaseNames" yaml:"glueDatabaseNames"`
	// Roles to allow passing as passed roles to actions.
	// Experimental.
	RolesToPass *[]awsiam.IRole `field:"optional" json:"rolesToPass" yaml:"rolesToPass"`
	// S3 Buckets to give read and write permissions.
	// Experimental.
	S3Buckets *[]awss3.IBucket `field:"optional" json:"s3Buckets" yaml:"s3Buckets"`
	// Version of the SageMaker Activity.
	//
	// This version will be used to fetch the policy template that corresponds to the
	// Activity.
	// Experimental.
	Version *float64 `field:"optional" json:"version" yaml:"version"`
}

