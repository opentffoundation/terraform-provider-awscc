// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package lambda

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_lambda_url", urlDataSource)
}

// urlDataSource returns the Terraform awscc_lambda_url data source.
// This Terraform data source corresponds to the CloudFormation AWS::Lambda::Url resource.
func urlDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AuthType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Can be either AWS_IAM if the requests are authorized via IAM, or NONE if no authorization is configured on the Function URL.",
		//	  "enum": [
		//	    "AWS_IAM",
		//	    "NONE"
		//	  ],
		//	  "type": "string"
		//	}
		"auth_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Can be either AWS_IAM if the requests are authorized via IAM, or NONE if no authorization is configured on the Function URL.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Cors
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "AllowCredentials": {
		//	      "description": "Specifies whether credentials are included in the CORS request.",
		//	      "type": "boolean"
		//	    },
		//	    "AllowHeaders": {
		//	      "description": "Represents a collection of allowed headers.",
		//	      "insertionOrder": true,
		//	      "items": {
		//	        "maxLength": 1024,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "maxItems": 100,
		//	      "minItems": 1,
		//	      "type": "array"
		//	    },
		//	    "AllowMethods": {
		//	      "description": "Represents a collection of allowed HTTP methods.",
		//	      "insertionOrder": true,
		//	      "items": {
		//	        "enum": [
		//	          "GET",
		//	          "PUT",
		//	          "HEAD",
		//	          "POST",
		//	          "PATCH",
		//	          "DELETE",
		//	          "*"
		//	        ],
		//	        "type": "string"
		//	      },
		//	      "maxItems": 6,
		//	      "minItems": 1,
		//	      "type": "array"
		//	    },
		//	    "AllowOrigins": {
		//	      "description": "Represents a collection of allowed origins.",
		//	      "insertionOrder": true,
		//	      "items": {
		//	        "maxLength": 253,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "maxItems": 100,
		//	      "minItems": 1,
		//	      "type": "array"
		//	    },
		//	    "ExposeHeaders": {
		//	      "description": "Represents a collection of exposed headers.",
		//	      "insertionOrder": true,
		//	      "items": {
		//	        "maxLength": 1024,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "maxItems": 100,
		//	      "minItems": 1,
		//	      "type": "array"
		//	    },
		//	    "MaxAge": {
		//	      "maximum": 86400,
		//	      "minimum": 0,
		//	      "type": "integer"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"cors": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: AllowCredentials
				"allow_credentials": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Description: "Specifies whether credentials are included in the CORS request.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: AllowHeaders
				"allow_headers": schema.ListAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Description: "Represents a collection of allowed headers.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: AllowMethods
				"allow_methods": schema.ListAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Description: "Represents a collection of allowed HTTP methods.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: AllowOrigins
				"allow_origins": schema.ListAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Description: "Represents a collection of allowed origins.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: ExposeHeaders
				"expose_headers": schema.ListAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Description: "Represents a collection of exposed headers.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: MaxAge
				"max_age": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: FunctionArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The full Amazon Resource Name (ARN) of the function associated with the Function URL.",
		//	  "pattern": "",
		//	  "type": "string"
		//	}
		"function_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The full Amazon Resource Name (ARN) of the function associated with the Function URL.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: FunctionUrl
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The generated url for this resource.",
		//	  "type": "string"
		//	}
		"function_url": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The generated url for this resource.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: InvokeMode
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The invocation mode for the function's URL. Set to BUFFERED if you want to buffer responses before returning them to the client. Set to RESPONSE_STREAM if you want to stream responses, allowing faster time to first byte and larger response payload sizes. If not set, defaults to BUFFERED.",
		//	  "enum": [
		//	    "BUFFERED",
		//	    "RESPONSE_STREAM"
		//	  ],
		//	  "type": "string"
		//	}
		"invoke_mode": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The invocation mode for the function's URL. Set to BUFFERED if you want to buffer responses before returning them to the client. Set to RESPONSE_STREAM if you want to stream responses, allowing faster time to first byte and larger response payload sizes. If not set, defaults to BUFFERED.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Qualifier
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The alias qualifier for the target function. If TargetFunctionArn is unqualified then Qualifier must be passed.",
		//	  "maxLength": 128,
		//	  "minLength": 1,
		//	  "pattern": "",
		//	  "type": "string"
		//	}
		"qualifier": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The alias qualifier for the target function. If TargetFunctionArn is unqualified then Qualifier must be passed.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: TargetFunctionArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the function associated with the Function URL.",
		//	  "pattern": "",
		//	  "type": "string"
		//	}
		"target_function_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the function associated with the Function URL.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::Lambda::Url",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Lambda::Url").WithTerraformTypeName("awscc_lambda_url")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"allow_credentials":   "AllowCredentials",
		"allow_headers":       "AllowHeaders",
		"allow_methods":       "AllowMethods",
		"allow_origins":       "AllowOrigins",
		"auth_type":           "AuthType",
		"cors":                "Cors",
		"expose_headers":      "ExposeHeaders",
		"function_arn":        "FunctionArn",
		"function_url":        "FunctionUrl",
		"invoke_mode":         "InvokeMode",
		"max_age":             "MaxAge",
		"qualifier":           "Qualifier",
		"target_function_arn": "TargetFunctionArn",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
