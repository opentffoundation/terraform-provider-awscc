// Code generated by generators/plural-data-source/main.go; DO NOT EDIT.

package lex

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_lex_bots", botsDataSource)
}

// botsDataSource returns the Terraform awscc_lex_bots data source.
// This Terraform data source corresponds to the CloudFormation AWS::Lex::Bot resource.
func botsDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]tfsdk.Attribute{
		"id": {
			Description: "Uniquely identifies the data source.",
			Type:        types.StringType,
			Computed:    true,
		},
		"ids": {
			Description: "Set of Resource Identifiers.",
			Type:        types.SetType{ElemType: types.StringType},
			Computed:    true,
		},
	}

	schema := tfsdk.Schema{
		Description: "Plural Data Source schema for AWS::Lex::Bot",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Lex::Bot").WithTerraformTypeName("awscc_lex_bots")
	opts = opts.WithTerraformSchema(schema)

	v, err := NewPluralDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
