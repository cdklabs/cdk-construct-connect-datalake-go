//go:build no_runtime_type_checking

package cdklabscdkconstructconnectdatalake

// Building without runtime type checking enabled, so all the below just return nil

func validateDataLakeAccess_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewDataLakeAccessParameters(scope constructs.Construct, id *string, props *DataLakeAccessProps) error {
	return nil
}

