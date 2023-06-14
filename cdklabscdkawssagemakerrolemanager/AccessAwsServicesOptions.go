package cdklabscdkawssagemakerrolemanager

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecr"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// SageMaker Activity Static Function Options.
// Experimental.
type AccessAwsServicesOptions struct {
	// Experimental.
	SecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// Experimental.
	Subnets *[]awsec2.ISubnet `field:"optional" json:"subnets" yaml:"subnets"`
	// Experimental.
	DataKeys *[]awskms.IKey `field:"optional" json:"dataKeys" yaml:"dataKeys"`
	// Experimental.
	VolumeKeys *[]awskms.IKey `field:"optional" json:"volumeKeys" yaml:"volumeKeys"`
	// Experimental.
	EcrRepositories *[]awsecr.IRepository `field:"required" json:"ecrRepositories" yaml:"ecrRepositories"`
	// Experimental.
	S3Buckets *[]awss3.IBucket `field:"required" json:"s3Buckets" yaml:"s3Buckets"`
}

