// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package rds

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	cctypes "github.com/hashicorp/terraform-provider-awscc/internal/types"
)

func init() {
	registry.AddDataSourceFactory("awscc_rds_db_proxy", dBProxyDataSource)
}

// dBProxyDataSource returns the Terraform awscc_rds_db_proxy data source.
// This Terraform data source corresponds to the CloudFormation AWS::RDS::DBProxy resource.
func dBProxyDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Auth
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The authorization mechanism that the proxy uses.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "AuthScheme": {
		//	        "description": "The type of authentication that the proxy uses for connections from the proxy to the underlying database. ",
		//	        "enum": [
		//	          "SECRETS"
		//	        ],
		//	        "type": "string"
		//	      },
		//	      "ClientPasswordAuthType": {
		//	        "description": "The type of authentication the proxy uses for connections from clients.",
		//	        "enum": [
		//	          "MYSQL_NATIVE_PASSWORD",
		//	          "POSTGRES_SCRAM_SHA_256",
		//	          "POSTGRES_MD5",
		//	          "SQL_SERVER_AUTHENTICATION"
		//	        ],
		//	        "type": "string"
		//	      },
		//	      "Description": {
		//	        "description": "A user-specified description about the authentication used by a proxy to log in as a specific database user. ",
		//	        "type": "string"
		//	      },
		//	      "IAMAuth": {
		//	        "description": "Whether to require or disallow Amazon Web Services Identity and Access Management (IAM) authentication for connections to the proxy. The ENABLED value is valid only for proxies with RDS for Microsoft SQL Server.",
		//	        "enum": [
		//	          "DISABLED",
		//	          "REQUIRED",
		//	          "ENABLED"
		//	        ],
		//	        "type": "string"
		//	      },
		//	      "SecretArn": {
		//	        "description": "The Amazon Resource Name (ARN) representing the secret that the proxy uses to authenticate to the RDS DB instance or Aurora DB cluster. These secrets are stored within Amazon Secrets Manager. ",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "type": "object"
		//	  },
		//	  "minItems": 1,
		//	  "type": "array"
		//	}
		"auth": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: AuthScheme
					"auth_scheme": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The type of authentication that the proxy uses for connections from the proxy to the underlying database. ",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: ClientPasswordAuthType
					"client_password_auth_type": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The type of authentication the proxy uses for connections from clients.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Description
					"description": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "A user-specified description about the authentication used by a proxy to log in as a specific database user. ",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: IAMAuth
					"iam_auth": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "Whether to require or disallow Amazon Web Services Identity and Access Management (IAM) authentication for connections to the proxy. The ENABLED value is valid only for proxies with RDS for Microsoft SQL Server.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: SecretArn
					"secret_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The Amazon Resource Name (ARN) representing the secret that the proxy uses to authenticate to the RDS DB instance or Aurora DB cluster. These secrets are stored within Amazon Secrets Manager. ",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			CustomType:  cctypes.NewMultisetTypeOf[types.Object](ctx),
			Description: "The authorization mechanism that the proxy uses.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DBProxyArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) for the proxy.",
		//	  "type": "string"
		//	}
		"db_proxy_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) for the proxy.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DBProxyName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The identifier for the proxy. This name must be unique for all proxies owned by your AWS account in the specified AWS Region.",
		//	  "maxLength": 64,
		//	  "pattern": "[0-z]*",
		//	  "type": "string"
		//	}
		"db_proxy_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The identifier for the proxy. This name must be unique for all proxies owned by your AWS account in the specified AWS Region.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DebugLogging
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Whether the proxy includes detailed information about SQL statements in its logs.",
		//	  "type": "boolean"
		//	}
		"debug_logging": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "Whether the proxy includes detailed information about SQL statements in its logs.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Endpoint
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The endpoint that you can use to connect to the proxy. You include the endpoint value in the connection string for a database client application.",
		//	  "type": "string"
		//	}
		"endpoint": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The endpoint that you can use to connect to the proxy. You include the endpoint value in the connection string for a database client application.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: EngineFamily
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The kinds of databases that the proxy can connect to.",
		//	  "enum": [
		//	    "MYSQL",
		//	    "POSTGRESQL",
		//	    "SQLSERVER"
		//	  ],
		//	  "type": "string"
		//	}
		"engine_family": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The kinds of databases that the proxy can connect to.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: IdleClientTimeout
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The number of seconds that a connection to the proxy can be inactive before the proxy disconnects it.",
		//	  "type": "integer"
		//	}
		"idle_client_timeout": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "The number of seconds that a connection to the proxy can be inactive before the proxy disconnects it.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: RequireTLS
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A Boolean parameter that specifies whether Transport Layer Security (TLS) encryption is required for connections to the proxy.",
		//	  "type": "boolean"
		//	}
		"require_tls": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "A Boolean parameter that specifies whether Transport Layer Security (TLS) encryption is required for connections to the proxy.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: RoleArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the IAM role that the proxy uses to access secrets in AWS Secrets Manager.",
		//	  "type": "string"
		//	}
		"role_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the IAM role that the proxy uses to access secrets in AWS Secrets Manager.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An optional set of key-value pairs to associate arbitrary data of your choosing with the proxy.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Key": {
		//	        "maxLength": 128,
		//	        "pattern": "(\\w|\\d|\\s|\\\\|-|\\.:=+-)*",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "maxLength": 128,
		//	        "pattern": "(\\w|\\d|\\s|\\\\|-|\\.:=+-)*",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "type": "object"
		//	  },
		//	  "type": "array"
		//	}
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			CustomType:  cctypes.NewMultisetTypeOf[types.Object](ctx),
			Description: "An optional set of key-value pairs to associate arbitrary data of your choosing with the proxy.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: VpcId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "VPC ID to associate with the new DB proxy.",
		//	  "type": "string"
		//	}
		"vpc_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "VPC ID to associate with the new DB proxy.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: VpcSecurityGroupIds
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "VPC security group IDs to associate with the new proxy.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "type": "string"
		//	  },
		//	  "minItems": 1,
		//	  "type": "array"
		//	}
		"vpc_security_group_ids": schema.ListAttribute{ /*START ATTRIBUTE*/
			CustomType:  cctypes.NewMultisetTypeOf[types.String](ctx),
			Description: "VPC security group IDs to associate with the new proxy.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: VpcSubnetIds
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "VPC subnet IDs to associate with the new proxy.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "type": "string"
		//	  },
		//	  "minItems": 2,
		//	  "type": "array"
		//	}
		"vpc_subnet_ids": schema.ListAttribute{ /*START ATTRIBUTE*/
			CustomType:  cctypes.NewMultisetTypeOf[types.String](ctx),
			Description: "VPC subnet IDs to associate with the new proxy.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::RDS::DBProxy",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::RDS::DBProxy").WithTerraformTypeName("awscc_rds_db_proxy")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"auth":                      "Auth",
		"auth_scheme":               "AuthScheme",
		"client_password_auth_type": "ClientPasswordAuthType",
		"db_proxy_arn":              "DBProxyArn",
		"db_proxy_name":             "DBProxyName",
		"debug_logging":             "DebugLogging",
		"description":               "Description",
		"endpoint":                  "Endpoint",
		"engine_family":             "EngineFamily",
		"iam_auth":                  "IAMAuth",
		"idle_client_timeout":       "IdleClientTimeout",
		"key":                       "Key",
		"require_tls":               "RequireTLS",
		"role_arn":                  "RoleArn",
		"secret_arn":                "SecretArn",
		"tags":                      "Tags",
		"value":                     "Value",
		"vpc_id":                    "VpcId",
		"vpc_security_group_ids":    "VpcSecurityGroupIds",
		"vpc_subnet_ids":            "VpcSubnetIds",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
