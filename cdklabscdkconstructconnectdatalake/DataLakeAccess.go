package cdklabscdkconstructconnectdatalake

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-construct-connect-datalake-go/cdklabscdkconstructconnectdatalake/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdklabs/cdk-construct-connect-datalake-go/cdklabscdkconstructconnectdatalake/internal"
)

// Automates Amazon Connect Data Lake integration setup and management.
//
// This construct simplifies the process of associating Amazon Connect analytics datasets with AWS data lake services,
// handling resource sharing, Lake Formation permissions, and data catalog management through a centralized Glue database.
// Experimental.
type DataLakeAccess interface {
	constructs.Construct
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Applies one or more mixins to this construct.
	//
	// Mixins are applied in order. The list of constructs is captured at the
	// start of the call, so constructs added by a mixin will not be visited.
	// Use multiple `with()` calls if subsequent mixins should apply to added
	// constructs.
	//
	// Returns: This construct for chaining.
	// Experimental.
	With(mixins ...constructs.IMixin) constructs.IConstruct
}

// The jsii proxy struct for DataLakeAccess
type jsiiProxy_DataLakeAccess struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_DataLakeAccess) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Experimental.
func NewDataLakeAccess(scope constructs.Construct, id *string, props *DataLakeAccessProps) DataLakeAccess {
	_init_.Initialize()

	if err := validateNewDataLakeAccessParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataLakeAccess{}

	_jsii_.Create(
		"@cdklabs/cdk-construct-connect-datalake.DataLakeAccess",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewDataLakeAccess_Override(d DataLakeAccess, scope constructs.Construct, id *string, props *DataLakeAccessProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/cdk-construct-connect-datalake.DataLakeAccess",
		[]interface{}{scope, id, props},
		d,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Experimental.
func DataLakeAccess_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataLakeAccess_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-construct-connect-datalake.DataLakeAccess",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLakeAccess) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLakeAccess) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		d,
		"with",
		args,
		&returns,
	)

	return returns
}

