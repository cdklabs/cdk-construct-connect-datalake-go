package cdklabscdkconstructconnectdatalake


// Properties for the DataLakeAccess construct.
// Experimental.
type DataLakeAccessProps struct {
	// Array of dataset IDs to associate with the Connect instance.
	//
	// Can include DataType enum values or string dataset IDs.
	// Experimental.
	DatasetIds *[]*string `field:"required" json:"datasetIds" yaml:"datasetIds"`
	// The identifier of the Amazon Connect instance.
	// Experimental.
	InstanceId *string `field:"required" json:"instanceId" yaml:"instanceId"`
	// The identifier of the target account.
	//
	// If not specified, defaults to the current AWS account. For cross-account setups,
	// specify an external account ID to associate datasets and create resources in that account
	//
	// When specified with an external account, `targetAccountRoleArn` is also required.
	// Experimental.
	TargetAccountId *string `field:"optional" json:"targetAccountId" yaml:"targetAccountId"`
	// The IAM role ARN in the target account for cross-account role assumption.
	//
	// Only required when `targetAccountId` is specified with an external account. A template
	// for the required permissions can be found in the [Cross Account Setup](../CROSS_ACCOUNT_SETUP.md) documentation
	// Experimental.
	TargetAccountRoleArn *string `field:"optional" json:"targetAccountRoleArn" yaml:"targetAccountRoleArn"`
}

