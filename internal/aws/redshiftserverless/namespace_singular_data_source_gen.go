// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package redshiftserverless

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	cctypes "github.com/hashicorp/terraform-provider-awscc/internal/types"
)

func init() {
	registry.AddDataSourceFactory("awscc_redshiftserverless_namespace", namespaceDataSource)
}

// namespaceDataSource returns the Terraform awscc_redshiftserverless_namespace data source.
// This Terraform data source corresponds to the CloudFormation AWS::RedshiftServerless::Namespace resource.
func namespaceDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AdminPasswordSecretKmsKeyId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the AWS Key Management Service (KMS) key used to encrypt and store the namespace's admin credentials secret. You can only use this parameter if manageAdminPassword is true.",
		//	  "type": "string"
		//	}
		"admin_password_secret_kms_key_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the AWS Key Management Service (KMS) key used to encrypt and store the namespace's admin credentials secret. You can only use this parameter if manageAdminPassword is true.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: AdminUserPassword
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The password associated with the admin user for the namespace that is being created. Password must be at least 8 characters in length, should be any printable ASCII character. Must contain at least one lowercase letter, one uppercase letter and one decimal digit. You can't use adminUserPassword if manageAdminPassword is true.",
		//	  "maxLength": 64,
		//	  "minLength": 8,
		//	  "pattern": "",
		//	  "type": "string"
		//	}
		"admin_user_password": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The password associated with the admin user for the namespace that is being created. Password must be at least 8 characters in length, should be any printable ASCII character. Must contain at least one lowercase letter, one uppercase letter and one decimal digit. You can't use adminUserPassword if manageAdminPassword is true.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: AdminUsername
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The user name associated with the admin user for the namespace that is being created. Only alphanumeric characters and underscores are allowed. It should start with an alphabet.",
		//	  "pattern": "[a-zA-Z][a-zA-Z_0-9+.@-]*",
		//	  "type": "string"
		//	}
		"admin_username": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The user name associated with the admin user for the namespace that is being created. Only alphanumeric characters and underscores are allowed. It should start with an alphabet.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DbName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The database name associated for the namespace that is being created. Only alphanumeric characters and underscores are allowed. It should start with an alphabet.",
		//	  "maxLength": 127,
		//	  "pattern": "[a-zA-Z][a-zA-Z_0-9+.@-]*",
		//	  "type": "string"
		//	}
		"db_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The database name associated for the namespace that is being created. Only alphanumeric characters and underscores are allowed. It should start with an alphabet.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DefaultIamRoleArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The default IAM role ARN for the namespace that is being created.",
		//	  "type": "string"
		//	}
		"default_iam_role_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The default IAM role ARN for the namespace that is being created.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: FinalSnapshotName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the namespace the source snapshot was created from. Please specify the name if needed before deleting namespace",
		//	  "maxLength": 255,
		//	  "pattern": "[a-z][a-z0-9]*(-[a-z0-9]+)*",
		//	  "type": "string"
		//	}
		"final_snapshot_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the namespace the source snapshot was created from. Please specify the name if needed before deleting namespace",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: FinalSnapshotRetentionPeriod
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The number of days to retain automated snapshot in the destination region after they are copied from the source region. If the value is -1, the manual snapshot is retained indefinitely. The value must be either -1 or an integer between 1 and 3,653.",
		//	  "type": "integer"
		//	}
		"final_snapshot_retention_period": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "The number of days to retain automated snapshot in the destination region after they are copied from the source region. If the value is -1, the manual snapshot is retained indefinitely. The value must be either -1 or an integer between 1 and 3,653.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: IamRoles
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A list of AWS Identity and Access Management (IAM) roles that can be used by the namespace to access other AWS services. You must supply the IAM roles in their Amazon Resource Name (ARN) format. The Default role limit for each request is 10.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "maxLength": 512,
		//	    "minLength": 0,
		//	    "type": "string"
		//	  },
		//	  "type": "array"
		//	}
		"iam_roles": schema.ListAttribute{ /*START ATTRIBUTE*/
			CustomType:  cctypes.NewMultisetTypeOf[types.String](ctx),
			Description: "A list of AWS Identity and Access Management (IAM) roles that can be used by the namespace to access other AWS services. You must supply the IAM roles in their Amazon Resource Name (ARN) format. The Default role limit for each request is 10.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: KmsKeyId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The AWS Key Management Service (KMS) key ID of the encryption key that you want to use to encrypt data in the namespace.",
		//	  "type": "string"
		//	}
		"kms_key_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The AWS Key Management Service (KMS) key ID of the encryption key that you want to use to encrypt data in the namespace.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: LogExports
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The collection of log types to be exported provided by the customer. Should only be one of the three supported log types: userlog, useractivitylog and connectionlog",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "enum": [
		//	      "useractivitylog",
		//	      "userlog",
		//	      "connectionlog"
		//	    ],
		//	    "type": "string"
		//	  },
		//	  "maxItems": 16,
		//	  "minItems": 0,
		//	  "type": "array"
		//	}
		"log_exports": schema.ListAttribute{ /*START ATTRIBUTE*/
			CustomType:  cctypes.NewMultisetTypeOf[types.String](ctx),
			Description: "The collection of log types to be exported provided by the customer. Should only be one of the three supported log types: userlog, useractivitylog and connectionlog",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ManageAdminPassword
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "If true, Amazon Redshift uses AWS Secrets Manager to manage the namespace's admin credentials. You can't use adminUserPassword if manageAdminPassword is true. If manageAdminPassword is false or not set, Amazon Redshift uses adminUserPassword for the admin user account's password.",
		//	  "type": "boolean"
		//	}
		"manage_admin_password": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "If true, Amazon Redshift uses AWS Secrets Manager to manage the namespace's admin credentials. You can't use adminUserPassword if manageAdminPassword is true. If manageAdminPassword is false or not set, Amazon Redshift uses adminUserPassword for the admin user account's password.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Namespace
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Definition of Namespace resource.",
		//	  "properties": {
		//	    "AdminPasswordSecretArn": {
		//	      "type": "string"
		//	    },
		//	    "AdminPasswordSecretKmsKeyId": {
		//	      "type": "string"
		//	    },
		//	    "AdminUsername": {
		//	      "type": "string"
		//	    },
		//	    "CreationDate": {
		//	      "type": "string"
		//	    },
		//	    "DbName": {
		//	      "pattern": "[a-zA-Z][a-zA-Z_0-9+.@-]*",
		//	      "type": "string"
		//	    },
		//	    "DefaultIamRoleArn": {
		//	      "type": "string"
		//	    },
		//	    "IamRoles": {
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "maxLength": 512,
		//	        "minLength": 0,
		//	        "type": "string"
		//	      },
		//	      "type": "array"
		//	    },
		//	    "KmsKeyId": {
		//	      "type": "string"
		//	    },
		//	    "LogExports": {
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "enum": [
		//	          "useractivitylog",
		//	          "userlog",
		//	          "connectionlog"
		//	        ],
		//	        "type": "string"
		//	      },
		//	      "maxItems": 16,
		//	      "minItems": 0,
		//	      "type": "array"
		//	    },
		//	    "NamespaceArn": {
		//	      "type": "string"
		//	    },
		//	    "NamespaceId": {
		//	      "type": "string"
		//	    },
		//	    "NamespaceName": {
		//	      "maxLength": 64,
		//	      "minLength": 3,
		//	      "pattern": "^[a-z0-9-]+$",
		//	      "type": "string"
		//	    },
		//	    "Status": {
		//	      "enum": [
		//	        "AVAILABLE",
		//	        "MODIFYING",
		//	        "DELETING"
		//	      ],
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"namespace": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: AdminPasswordSecretArn
				"admin_password_secret_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: AdminPasswordSecretKmsKeyId
				"admin_password_secret_kms_key_id": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: AdminUsername
				"admin_username": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: CreationDate
				"creation_date": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: DbName
				"db_name": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: DefaultIamRoleArn
				"default_iam_role_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: IamRoles
				"iam_roles": schema.ListAttribute{ /*START ATTRIBUTE*/
					CustomType: cctypes.NewMultisetTypeOf[types.String](ctx),
					Computed:   true,
				}, /*END ATTRIBUTE*/
				// Property: KmsKeyId
				"kms_key_id": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: LogExports
				"log_exports": schema.ListAttribute{ /*START ATTRIBUTE*/
					CustomType: cctypes.NewMultisetTypeOf[types.String](ctx),
					Computed:   true,
				}, /*END ATTRIBUTE*/
				// Property: NamespaceArn
				"namespace_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: NamespaceId
				"namespace_id": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: NamespaceName
				"namespace_name": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: Status
				"status": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Definition of Namespace resource.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: NamespaceName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A unique identifier for the namespace. You use this identifier to refer to the namespace for any subsequent namespace operations such as deleting or modifying. All alphabetical characters must be lower case. Namespace name should be unique for all namespaces within an AWS account.",
		//	  "maxLength": 64,
		//	  "minLength": 3,
		//	  "pattern": "^[a-z0-9-]+$",
		//	  "type": "string"
		//	}
		"namespace_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A unique identifier for the namespace. You use this identifier to refer to the namespace for any subsequent namespace operations such as deleting or modifying. All alphabetical characters must be lower case. Namespace name should be unique for all namespaces within an AWS account.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: NamespaceResourcePolicy
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The resource policy document that will be attached to the namespace.",
		//	  "type": "object"
		//	}
		"namespace_resource_policy": schema.StringAttribute{ /*START ATTRIBUTE*/
			CustomType:  jsontypes.NormalizedType{},
			Description: "The resource policy document that will be attached to the namespace.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: RedshiftIdcApplicationArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ARN for the Redshift application that integrates with IAM Identity Center.",
		//	  "type": "string"
		//	}
		"redshift_idc_application_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ARN for the Redshift application that integrates with IAM Identity Center.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The list of tags for the namespace.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Key": {
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "maxLength": 256,
		//	        "minLength": 0,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key",
		//	      "Value"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "maxItems": 200,
		//	  "minItems": 0,
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
			Description: "The list of tags for the namespace.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::RedshiftServerless::Namespace",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::RedshiftServerless::Namespace").WithTerraformTypeName("awscc_redshiftserverless_namespace")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"admin_password_secret_arn":        "AdminPasswordSecretArn",
		"admin_password_secret_kms_key_id": "AdminPasswordSecretKmsKeyId",
		"admin_user_password":              "AdminUserPassword",
		"admin_username":                   "AdminUsername",
		"creation_date":                    "CreationDate",
		"db_name":                          "DbName",
		"default_iam_role_arn":             "DefaultIamRoleArn",
		"final_snapshot_name":              "FinalSnapshotName",
		"final_snapshot_retention_period":  "FinalSnapshotRetentionPeriod",
		"iam_roles":                        "IamRoles",
		"key":                              "Key",
		"kms_key_id":                       "KmsKeyId",
		"log_exports":                      "LogExports",
		"manage_admin_password":            "ManageAdminPassword",
		"namespace":                        "Namespace",
		"namespace_arn":                    "NamespaceArn",
		"namespace_id":                     "NamespaceId",
		"namespace_name":                   "NamespaceName",
		"namespace_resource_policy":        "NamespaceResourcePolicy",
		"redshift_idc_application_arn":     "RedshiftIdcApplicationArn",
		"status":                           "Status",
		"tags":                             "Tags",
		"value":                            "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
