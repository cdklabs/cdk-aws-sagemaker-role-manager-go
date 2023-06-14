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
type Persona interface {
	constructs.Construct
	// Experimental.
	Activities() *[]Activity
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Creates role with permissions of persona.
	//
	// Returns: - The role that is created with the permissions of the persona.
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

// The jsii proxy struct for Persona
type jsiiProxy_Persona struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_Persona) Activities() *[]Activity {
	var returns *[]Activity
	_jsii_.Get(
		j,
		"activities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Persona) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Experimental.
func NewPersona(scope constructs.Construct, id *string, props *PersonaProps) Persona {
	_init_.Initialize()

	if err := validateNewPersonaParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_Persona{}

	_jsii_.Create(
		"@cdklabs/cdk-aws-sagemaker-role-manager.Persona",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewPersona_Override(p Persona, scope constructs.Construct, id *string, props *PersonaProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/cdk-aws-sagemaker-role-manager.Persona",
		[]interface{}{scope, id, props},
		p,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func Persona_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePersona_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-aws-sagemaker-role-manager.Persona",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Persona) CreateRole(scope constructs.Construct, id *string, roleNameSuffix *string, roleDescription *string) awsiam.IRole {
	if err := p.validateCreateRoleParameters(scope, id, roleNameSuffix); err != nil {
		panic(err)
	}
	var returns awsiam.IRole

	_jsii_.Invoke(
		p,
		"createRole",
		[]interface{}{scope, id, roleNameSuffix, roleDescription},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Persona) CustomizeKMS(dataKeys *[]awskms.IKey, volumeKeys *[]awskms.IKey) {
	_jsii_.InvokeVoid(
		p,
		"customizeKMS",
		[]interface{}{dataKeys, volumeKeys},
	)
}

func (p *jsiiProxy_Persona) CustomizeVPC(subnets *[]awsec2.ISubnet, securityGroups *[]awsec2.ISecurityGroup) {
	_jsii_.InvokeVoid(
		p,
		"customizeVPC",
		[]interface{}{subnets, securityGroups},
	)
}

func (p *jsiiProxy_Persona) GrantPermissionsTo(identity awsiam.IGrantable) awsiam.Grant {
	if err := p.validateGrantPermissionsToParameters(identity); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		p,
		"grantPermissionsTo",
		[]interface{}{identity},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Persona) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

