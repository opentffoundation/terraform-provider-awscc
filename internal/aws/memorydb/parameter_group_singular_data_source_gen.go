// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package memorydb

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_memorydb_parameter_group", parameterGroupDataSource)
}

// parameterGroupDataSource returns the Terraform awscc_memorydb_parameter_group data source.
// This Terraform data source corresponds to the CloudFormation AWS::MemoryDB::ParameterGroup resource.
func parameterGroupDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: ARN
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the parameter group.",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the parameter group.",
			Type:        types.StringType,
			Computed:    true,
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "description": "A description of the parameter group.",
			//   "type": "string"
			// }
			Description: "A description of the parameter group.",
			Type:        types.StringType,
			Computed:    true,
		},
		"family": {
			// Property: Family
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the parameter group family that this parameter group is compatible with.",
			//   "type": "string"
			// }
			Description: "The name of the parameter group family that this parameter group is compatible with.",
			Type:        types.StringType,
			Computed:    true,
		},
		"parameter_group_name": {
			// Property: ParameterGroupName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the parameter group.",
			//   "type": "string"
			// }
			Description: "The name of the parameter group.",
			Type:        types.StringType,
			Computed:    true,
		},
		"parameters": {
			// Property: Parameters
			// CloudFormation resource type schema:
			// {
			//   "description": "An map of parameter names and values for the parameter update. You must supply at least one parameter name and value; subsequent arguments are optional.",
			//   "type": "object"
			// }
			Description: "An map of parameter names and values for the parameter update. You must supply at least one parameter name and value; subsequent arguments are optional.",
			Type:        types.MapType{ElemType: types.StringType},
			Computed:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "An array of key-value pairs to apply to this parameter group.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "A key-value pair to associate with a resource.",
			//     "properties": {
			//       "Key": {
			//         "description": "The key for the tag. May not be null.",
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "The tag's value. May be null.",
			//         "maxLength": 256,
			//         "minLength": 1,
			//         "pattern": "",
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Key",
			//       "Value"
			//     ],
			//     "type": "object"
			//   },
			//   "maxItems": 50,
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "An array of key-value pairs to apply to this parameter group.",
			Attributes: tfsdk.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The key for the tag. May not be null.",
						Type:        types.StringType,
						Computed:    true,
					},
					"value": {
						// Property: Value
						Description: "The tag's value. May be null.",
						Type:        types.StringType,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::MemoryDB::ParameterGroup",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::MemoryDB::ParameterGroup").WithTerraformTypeName("awscc_memorydb_parameter_group")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                  "ARN",
		"description":          "Description",
		"family":               "Family",
		"key":                  "Key",
		"parameter_group_name": "ParameterGroupName",
		"parameters":           "Parameters",
		"tags":                 "Tags",
		"value":                "Value",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
