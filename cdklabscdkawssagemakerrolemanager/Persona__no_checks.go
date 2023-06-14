//go:build no_runtime_type_checking

package cdklabscdkawssagemakerrolemanager

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_Persona) validateCreateRoleParameters(scope constructs.Construct, id *string, roleNameSuffix *string) error {
	return nil
}

func (p *jsiiProxy_Persona) validateGrantPermissionsToParameters(identity awsiam.IGrantable) error {
	return nil
}

func validatePersona_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewPersonaParameters(scope constructs.Construct, id *string, props *PersonaProps) error {
	return nil
}

