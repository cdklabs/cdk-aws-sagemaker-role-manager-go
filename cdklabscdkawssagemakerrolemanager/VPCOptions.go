package cdklabscdkawssagemakerrolemanager

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// Global Condition Customization Options.
// Experimental.
type VPCOptions struct {
	// Experimental.
	SecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// Experimental.
	Subnets *[]awsec2.ISubnet `field:"optional" json:"subnets" yaml:"subnets"`
}

