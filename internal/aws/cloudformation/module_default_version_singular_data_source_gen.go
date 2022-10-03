// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package cloudformation

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_cloudformation_module_default_version", moduleDefaultVersionDataSource)
}

// moduleDefaultVersionDataSource returns the Terraform awscc_cloudformation_module_default_version data source.
// This Terraform data source corresponds to the CloudFormation AWS::CloudFormation::ModuleDefaultVersion resource.
func moduleDefaultVersionDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the module version to set as the default version.",
			//   "pattern": "^arn:aws[A-Za-z0-9-]{0,64}:cloudformation:[A-Za-z0-9-]{1,64}:([0-9]{12})?:type/module/.+/[0-9]{8}$",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the module version to set as the default version.",
			Type:        types.StringType,
			Computed:    true,
		},
		"module_name": {
			// Property: ModuleName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of a module existing in the registry.",
			//   "pattern": "^[A-Za-z0-9]{2,64}::[A-Za-z0-9]{2,64}::[A-Za-z0-9]{2,64}::MODULE",
			//   "type": "string"
			// }
			Description: "The name of a module existing in the registry.",
			Type:        types.StringType,
			Computed:    true,
		},
		"version_id": {
			// Property: VersionId
			// CloudFormation resource type schema:
			// {
			//   "description": "The ID of an existing version of the named module to set as the default.",
			//   "pattern": "^[0-9]{8}$",
			//   "type": "string"
			// }
			Description: "The ID of an existing version of the named module to set as the default.",
			Type:        types.StringType,
			Computed:    true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::CloudFormation::ModuleDefaultVersion",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::CloudFormation::ModuleDefaultVersion").WithTerraformTypeName("awscc_cloudformation_module_default_version")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":         "Arn",
		"module_name": "ModuleName",
		"version_id":  "VersionId",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
