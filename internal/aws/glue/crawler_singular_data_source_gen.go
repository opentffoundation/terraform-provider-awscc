// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package glue

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_glue_crawler", crawlerDataSource)
}

// crawlerDataSource returns the Terraform awscc_glue_crawler data source.
// This Terraform data source corresponds to the CloudFormation AWS::Glue::Crawler resource.
func crawlerDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Classifiers
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A list of UTF-8 strings that specify the names of custom classifiers that are associated with the crawler.",
		//	  "items": {
		//	    "type": "string"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"classifiers": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "A list of UTF-8 strings that specify the names of custom classifiers that are associated with the crawler.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Configuration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Crawler configuration information. This versioned JSON string allows users to specify aspects of a crawler's behavior.",
		//	  "type": "string"
		//	}
		"configuration": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Crawler configuration information. This versioned JSON string allows users to specify aspects of a crawler's behavior.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: CrawlerSecurityConfiguration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the SecurityConfiguration structure to be used by this crawler.",
		//	  "type": "string"
		//	}
		"crawler_security_configuration": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the SecurityConfiguration structure to be used by this crawler.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DatabaseName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the database in which the crawler's output is stored.",
		//	  "type": "string"
		//	}
		"database_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the database in which the crawler's output is stored.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A description of the crawler.",
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A description of the crawler.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: LakeFormationConfiguration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Specifies AWS Lake Formation configuration settings for the crawler",
		//	  "properties": {
		//	    "AccountId": {
		//	      "description": "Required for cross account crawls. For same account crawls as the target data, this can be left as null.",
		//	      "type": "string"
		//	    },
		//	    "UseLakeFormationCredentials": {
		//	      "description": "Specifies whether to use AWS Lake Formation credentials for the crawler instead of the IAM role credentials.",
		//	      "type": "boolean"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"lake_formation_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: AccountId
				"account_id": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Required for cross account crawls. For same account crawls as the target data, this can be left as null.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: UseLakeFormationCredentials
				"use_lake_formation_credentials": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Description: "Specifies whether to use AWS Lake Formation credentials for the crawler instead of the IAM role credentials.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Specifies AWS Lake Formation configuration settings for the crawler",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the crawler.",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the crawler.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: RecrawlPolicy
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "When crawling an Amazon S3 data source after the first crawl is complete, specifies whether to crawl the entire dataset again or to crawl only folders that were added since the last crawler run. For more information, see Incremental Crawls in AWS Glue in the developer guide.",
		//	  "properties": {
		//	    "RecrawlBehavior": {
		//	      "description": "Specifies whether to crawl the entire dataset again or to crawl only folders that were added since the last crawler run. A value of CRAWL_EVERYTHING specifies crawling the entire dataset again. A value of CRAWL_NEW_FOLDERS_ONLY specifies crawling only folders that were added since the last crawler run. A value of CRAWL_EVENT_MODE specifies crawling only the changes identified by Amazon S3 events.",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"recrawl_policy": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: RecrawlBehavior
				"recrawl_behavior": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Specifies whether to crawl the entire dataset again or to crawl only folders that were added since the last crawler run. A value of CRAWL_EVERYTHING specifies crawling the entire dataset again. A value of CRAWL_NEW_FOLDERS_ONLY specifies crawling only folders that were added since the last crawler run. A value of CRAWL_EVENT_MODE specifies crawling only the changes identified by Amazon S3 events.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "When crawling an Amazon S3 data source after the first crawl is complete, specifies whether to crawl the entire dataset again or to crawl only folders that were added since the last crawler run. For more information, see Incremental Crawls in AWS Glue in the developer guide.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Role
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of an IAM role that's used to access customer resources, such as Amazon Simple Storage Service (Amazon S3) data.",
		//	  "type": "string"
		//	}
		"role": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of an IAM role that's used to access customer resources, such as Amazon Simple Storage Service (Amazon S3) data.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Schedule
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "A scheduling object using a cron statement to schedule an event.",
		//	  "properties": {
		//	    "ScheduleExpression": {
		//	      "description": "A cron expression used to specify the schedule. For more information, see Time-Based Schedules for Jobs and Crawlers. For example, to run something every day at 12:15 UTC, specify cron(15 12 * * ? *).",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"schedule": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: ScheduleExpression
				"schedule_expression": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "A cron expression used to specify the schedule. For more information, see Time-Based Schedules for Jobs and Crawlers. For example, to run something every day at 12:15 UTC, specify cron(15 12 * * ? *).",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "A scheduling object using a cron statement to schedule an event.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: SchemaChangePolicy
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The policy that specifies update and delete behaviors for the crawler. The policy tells the crawler what to do in the event that it detects a change in a table that already exists in the customer's database at the time of the crawl. The SchemaChangePolicy does not affect whether or how new tables and partitions are added. New tables and partitions are always created regardless of the SchemaChangePolicy on a crawler. The SchemaChangePolicy consists of two components, UpdateBehavior and DeleteBehavior.",
		//	  "properties": {
		//	    "DeleteBehavior": {
		//	      "description": "The deletion behavior when the crawler finds a deleted object. A value of LOG specifies that if a table or partition is found to no longer exist, do not delete it, only log that it was found to no longer exist. A value of DELETE_FROM_DATABASE specifies that if a table or partition is found to have been removed, delete it from the database. A value of DEPRECATE_IN_DATABASE specifies that if a table has been found to no longer exist, to add a property to the table that says 'DEPRECATED' and includes a timestamp with the time of deprecation.",
		//	      "type": "string"
		//	    },
		//	    "UpdateBehavior": {
		//	      "description": "The update behavior when the crawler finds a changed schema. A value of LOG specifies that if a table or a partition already exists, and a change is detected, do not update it, only log that a change was detected. Add new tables and new partitions (including on existing tables). A value of UPDATE_IN_DATABASE specifies that if a table or partition already exists, and a change is detected, update it. Add new tables and partitions.",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"schema_change_policy": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: DeleteBehavior
				"delete_behavior": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The deletion behavior when the crawler finds a deleted object. A value of LOG specifies that if a table or partition is found to no longer exist, do not delete it, only log that it was found to no longer exist. A value of DELETE_FROM_DATABASE specifies that if a table or partition is found to have been removed, delete it from the database. A value of DEPRECATE_IN_DATABASE specifies that if a table has been found to no longer exist, to add a property to the table that says 'DEPRECATED' and includes a timestamp with the time of deprecation.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: UpdateBehavior
				"update_behavior": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The update behavior when the crawler finds a changed schema. A value of LOG specifies that if a table or a partition already exists, and a change is detected, do not update it, only log that a change was detected. Add new tables and new partitions (including on existing tables). A value of UPDATE_IN_DATABASE specifies that if a table or partition already exists, and a change is detected, update it. Add new tables and partitions.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The policy that specifies update and delete behaviors for the crawler. The policy tells the crawler what to do in the event that it detects a change in a table that already exists in the customer's database at the time of the crawl. The SchemaChangePolicy does not affect whether or how new tables and partitions are added. New tables and partitions are always created regardless of the SchemaChangePolicy on a crawler. The SchemaChangePolicy consists of two components, UpdateBehavior and DeleteBehavior.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: TablePrefix
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The prefix added to the names of tables that are created.",
		//	  "type": "string"
		//	}
		"table_prefix": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The prefix added to the names of tables that are created.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The tags to use with this crawler.",
		//	  "type": "object"
		//	}
		"tags": schema.StringAttribute{ /*START ATTRIBUTE*/
			CustomType:  jsontypes.NormalizedType{},
			Description: "The tags to use with this crawler.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Targets
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Specifies data stores to crawl.",
		//	  "properties": {
		//	    "CatalogTargets": {
		//	      "description": "Specifies AWS Glue Data Catalog targets.",
		//	      "items": {
		//	        "additionalProperties": false,
		//	        "description": "Specifies an AWS Glue Data Catalog target.",
		//	        "properties": {
		//	          "ConnectionName": {
		//	            "description": "The name of the connection for an Amazon S3-backed Data Catalog table to be a target of the crawl when using a Catalog connection type paired with a NETWORK Connection type.",
		//	            "type": "string"
		//	          },
		//	          "DatabaseName": {
		//	            "description": "The name of the database to be synchronized.",
		//	            "type": "string"
		//	          },
		//	          "DlqEventQueueArn": {
		//	            "description": "A valid Amazon dead-letter SQS ARN. For example, arn:aws:sqs:region:account:deadLetterQueue.",
		//	            "type": "string"
		//	          },
		//	          "EventQueueArn": {
		//	            "description": "A valid Amazon SQS ARN. For example, arn:aws:sqs:region:account:sqs.",
		//	            "type": "string"
		//	          },
		//	          "Tables": {
		//	            "description": "A list of the tables to be synchronized.",
		//	            "items": {
		//	              "type": "string"
		//	            },
		//	            "type": "array",
		//	            "uniqueItems": false
		//	          }
		//	        },
		//	        "type": "object"
		//	      },
		//	      "type": "array",
		//	      "uniqueItems": false
		//	    },
		//	    "DeltaTargets": {
		//	      "description": "Specifies an array of Delta data store targets.",
		//	      "items": {
		//	        "additionalProperties": false,
		//	        "description": "Specifies a Delta data store to crawl one or more Delta tables.",
		//	        "properties": {
		//	          "ConnectionName": {
		//	            "description": "The name of the connection to use to connect to the Delta table target.",
		//	            "type": "string"
		//	          },
		//	          "CreateNativeDeltaTable": {
		//	            "description": "Specifies whether the crawler will create native tables, to allow integration with query engines that support querying of the Delta transaction log directly.",
		//	            "type": "boolean"
		//	          },
		//	          "DeltaTables": {
		//	            "description": "",
		//	            "items": {
		//	              "description": "A list of the Amazon S3 paths to the Delta tables.",
		//	              "type": "string"
		//	            },
		//	            "type": "array",
		//	            "uniqueItems": false
		//	          },
		//	          "WriteManifest": {
		//	            "description": "Specifies whether to write the manifest files to the Delta table path.",
		//	            "type": "boolean"
		//	          }
		//	        },
		//	        "type": "object"
		//	      },
		//	      "type": "array",
		//	      "uniqueItems": false
		//	    },
		//	    "DynamoDBTargets": {
		//	      "description": "Specifies Amazon DynamoDB targets.",
		//	      "items": {
		//	        "additionalProperties": false,
		//	        "description": "Specifies an Amazon DynamoDB table to crawl.",
		//	        "properties": {
		//	          "Path": {
		//	            "description": "The name of the DynamoDB table to crawl.",
		//	            "type": "string"
		//	          }
		//	        },
		//	        "type": "object"
		//	      },
		//	      "type": "array",
		//	      "uniqueItems": false
		//	    },
		//	    "IcebergTargets": {
		//	      "description": "Specifies Apache Iceberg data store targets.",
		//	      "items": {
		//	        "additionalProperties": false,
		//	        "description": "Specifies Apache Iceberg data store targets.",
		//	        "properties": {
		//	          "ConnectionName": {
		//	            "description": "The name of the connection to use to connect to the Iceberg target.",
		//	            "type": "string"
		//	          },
		//	          "Exclusions": {
		//	            "description": "A list of global patterns used to exclude from the crawl.",
		//	            "items": {
		//	              "type": "string"
		//	            },
		//	            "type": "array",
		//	            "uniqueItems": false
		//	          },
		//	          "MaximumTraversalDepth": {
		//	            "description": "The maximum depth of Amazon S3 paths that the crawler can traverse to discover the Iceberg metadata folder in your Amazon S3 path. Used to limit the crawler run time.",
		//	            "type": "integer"
		//	          },
		//	          "Paths": {
		//	            "description": "One or more Amazon S3 paths that contains Iceberg metadata folders as s3://bucket/prefix .",
		//	            "items": {
		//	              "type": "string"
		//	            },
		//	            "type": "array",
		//	            "uniqueItems": false
		//	          }
		//	        },
		//	        "type": "object"
		//	      },
		//	      "type": "array",
		//	      "uniqueItems": false
		//	    },
		//	    "JdbcTargets": {
		//	      "description": "Specifies JDBC targets.",
		//	      "items": {
		//	        "additionalProperties": false,
		//	        "description": "Specifies a JDBC data store to crawl.",
		//	        "properties": {
		//	          "ConnectionName": {
		//	            "description": "The name of the connection to use to connect to the JDBC target.",
		//	            "type": "string"
		//	          },
		//	          "EnableAdditionalMetadata": {
		//	            "description": "Specify a value of RAWTYPES or COMMENTS to enable additional metadata in table responses. RAWTYPES provides the native-level datatype. COMMENTS provides comments associated with a column or table in the database.\n\nIf you do not need additional metadata, keep the field empty.",
		//	            "items": {
		//	              "type": "string"
		//	            },
		//	            "type": "array",
		//	            "uniqueItems": false
		//	          },
		//	          "Exclusions": {
		//	            "description": "A list of glob patterns used to exclude from the crawl. For more information, see Catalog Tables with a Crawler.",
		//	            "items": {
		//	              "type": "string"
		//	            },
		//	            "type": "array",
		//	            "uniqueItems": false
		//	          },
		//	          "Path": {
		//	            "description": "The path of the JDBC target.",
		//	            "type": "string"
		//	          }
		//	        },
		//	        "type": "object"
		//	      },
		//	      "type": "array",
		//	      "uniqueItems": false
		//	    },
		//	    "MongoDBTargets": {
		//	      "description": "A list of Mongo DB targets.",
		//	      "items": {
		//	        "additionalProperties": false,
		//	        "description": "Specifies an Amazon DocumentDB or MongoDB data store to crawl.",
		//	        "properties": {
		//	          "ConnectionName": {
		//	            "description": "The name of the connection to use to connect to the Amazon DocumentDB or MongoDB target.",
		//	            "type": "string"
		//	          },
		//	          "Path": {
		//	            "description": "The path of the Amazon DocumentDB or MongoDB target (database/collection).",
		//	            "type": "string"
		//	          }
		//	        },
		//	        "type": "object"
		//	      },
		//	      "type": "array",
		//	      "uniqueItems": false
		//	    },
		//	    "S3Targets": {
		//	      "description": "Specifies Amazon Simple Storage Service (Amazon S3) targets.",
		//	      "items": {
		//	        "additionalProperties": false,
		//	        "description": "Specifies a data store in Amazon Simple Storage Service (Amazon S3).",
		//	        "properties": {
		//	          "ConnectionName": {
		//	            "description": "The name of a connection which allows a job or crawler to access data in Amazon S3 within an Amazon Virtual Private Cloud environment (Amazon VPC).",
		//	            "type": "string"
		//	          },
		//	          "DlqEventQueueArn": {
		//	            "description": "A valid Amazon dead-letter SQS ARN. For example, arn:aws:sqs:region:account:deadLetterQueue.",
		//	            "type": "string"
		//	          },
		//	          "EventQueueArn": {
		//	            "description": "A valid Amazon SQS ARN. For example, arn:aws:sqs:region:account:sqs.",
		//	            "type": "string"
		//	          },
		//	          "Exclusions": {
		//	            "description": "A list of glob patterns used to exclude from the crawl.",
		//	            "items": {
		//	              "type": "string"
		//	            },
		//	            "type": "array",
		//	            "uniqueItems": false
		//	          },
		//	          "Path": {
		//	            "description": "The path to the Amazon S3 target.",
		//	            "type": "string"
		//	          },
		//	          "SampleSize": {
		//	            "description": "Sets the number of files in each leaf folder to be crawled when crawling sample files in a dataset. If not set, all the files are crawled. A valid value is an integer between 1 and 249.",
		//	            "type": "integer"
		//	          }
		//	        },
		//	        "type": "object"
		//	      },
		//	      "type": "array",
		//	      "uniqueItems": false
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"targets": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: CatalogTargets
				"catalog_targets": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
					NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: ConnectionName
							"connection_name": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The name of the connection for an Amazon S3-backed Data Catalog table to be a target of the crawl when using a Catalog connection type paired with a NETWORK Connection type.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: DatabaseName
							"database_name": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The name of the database to be synchronized.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: DlqEventQueueArn
							"dlq_event_queue_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "A valid Amazon dead-letter SQS ARN. For example, arn:aws:sqs:region:account:deadLetterQueue.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: EventQueueArn
							"event_queue_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "A valid Amazon SQS ARN. For example, arn:aws:sqs:region:account:sqs.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: Tables
							"tables": schema.ListAttribute{ /*START ATTRIBUTE*/
								ElementType: types.StringType,
								Description: "A list of the tables to be synchronized.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
					}, /*END NESTED OBJECT*/
					Description: "Specifies AWS Glue Data Catalog targets.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: DeltaTargets
				"delta_targets": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
					NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: ConnectionName
							"connection_name": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The name of the connection to use to connect to the Delta table target.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: CreateNativeDeltaTable
							"create_native_delta_table": schema.BoolAttribute{ /*START ATTRIBUTE*/
								Description: "Specifies whether the crawler will create native tables, to allow integration with query engines that support querying of the Delta transaction log directly.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: DeltaTables
							"delta_tables": schema.ListAttribute{ /*START ATTRIBUTE*/
								ElementType: types.StringType,
								Description: "",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: WriteManifest
							"write_manifest": schema.BoolAttribute{ /*START ATTRIBUTE*/
								Description: "Specifies whether to write the manifest files to the Delta table path.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
					}, /*END NESTED OBJECT*/
					Description: "Specifies an array of Delta data store targets.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: DynamoDBTargets
				"dynamo_db_targets": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
					NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: Path
							"path": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The name of the DynamoDB table to crawl.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
					}, /*END NESTED OBJECT*/
					Description: "Specifies Amazon DynamoDB targets.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: IcebergTargets
				"iceberg_targets": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
					NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: ConnectionName
							"connection_name": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The name of the connection to use to connect to the Iceberg target.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: Exclusions
							"exclusions": schema.ListAttribute{ /*START ATTRIBUTE*/
								ElementType: types.StringType,
								Description: "A list of global patterns used to exclude from the crawl.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: MaximumTraversalDepth
							"maximum_traversal_depth": schema.Int64Attribute{ /*START ATTRIBUTE*/
								Description: "The maximum depth of Amazon S3 paths that the crawler can traverse to discover the Iceberg metadata folder in your Amazon S3 path. Used to limit the crawler run time.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: Paths
							"paths": schema.ListAttribute{ /*START ATTRIBUTE*/
								ElementType: types.StringType,
								Description: "One or more Amazon S3 paths that contains Iceberg metadata folders as s3://bucket/prefix .",
								Computed:    true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
					}, /*END NESTED OBJECT*/
					Description: "Specifies Apache Iceberg data store targets.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: JdbcTargets
				"jdbc_targets": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
					NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: ConnectionName
							"connection_name": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The name of the connection to use to connect to the JDBC target.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: EnableAdditionalMetadata
							"enable_additional_metadata": schema.ListAttribute{ /*START ATTRIBUTE*/
								ElementType: types.StringType,
								Description: "Specify a value of RAWTYPES or COMMENTS to enable additional metadata in table responses. RAWTYPES provides the native-level datatype. COMMENTS provides comments associated with a column or table in the database.\n\nIf you do not need additional metadata, keep the field empty.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: Exclusions
							"exclusions": schema.ListAttribute{ /*START ATTRIBUTE*/
								ElementType: types.StringType,
								Description: "A list of glob patterns used to exclude from the crawl. For more information, see Catalog Tables with a Crawler.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: Path
							"path": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The path of the JDBC target.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
					}, /*END NESTED OBJECT*/
					Description: "Specifies JDBC targets.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: MongoDBTargets
				"mongo_db_targets": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
					NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: ConnectionName
							"connection_name": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The name of the connection to use to connect to the Amazon DocumentDB or MongoDB target.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: Path
							"path": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The path of the Amazon DocumentDB or MongoDB target (database/collection).",
								Computed:    true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
					}, /*END NESTED OBJECT*/
					Description: "A list of Mongo DB targets.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: S3Targets
				"s3_targets": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
					NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: ConnectionName
							"connection_name": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The name of a connection which allows a job or crawler to access data in Amazon S3 within an Amazon Virtual Private Cloud environment (Amazon VPC).",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: DlqEventQueueArn
							"dlq_event_queue_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "A valid Amazon dead-letter SQS ARN. For example, arn:aws:sqs:region:account:deadLetterQueue.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: EventQueueArn
							"event_queue_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "A valid Amazon SQS ARN. For example, arn:aws:sqs:region:account:sqs.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: Exclusions
							"exclusions": schema.ListAttribute{ /*START ATTRIBUTE*/
								ElementType: types.StringType,
								Description: "A list of glob patterns used to exclude from the crawl.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: Path
							"path": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The path to the Amazon S3 target.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: SampleSize
							"sample_size": schema.Int64Attribute{ /*START ATTRIBUTE*/
								Description: "Sets the number of files in each leaf folder to be crawled when crawling sample files in a dataset. If not set, all the files are crawled. A valid value is an integer between 1 and 249.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
					}, /*END NESTED OBJECT*/
					Description: "Specifies Amazon Simple Storage Service (Amazon S3) targets.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Specifies data stores to crawl.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::Glue::Crawler",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Glue::Crawler").WithTerraformTypeName("awscc_glue_crawler")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"account_id":                     "AccountId",
		"catalog_targets":                "CatalogTargets",
		"classifiers":                    "Classifiers",
		"configuration":                  "Configuration",
		"connection_name":                "ConnectionName",
		"crawler_security_configuration": "CrawlerSecurityConfiguration",
		"create_native_delta_table":      "CreateNativeDeltaTable",
		"database_name":                  "DatabaseName",
		"delete_behavior":                "DeleteBehavior",
		"delta_tables":                   "DeltaTables",
		"delta_targets":                  "DeltaTargets",
		"description":                    "Description",
		"dlq_event_queue_arn":            "DlqEventQueueArn",
		"dynamo_db_targets":              "DynamoDBTargets",
		"enable_additional_metadata":     "EnableAdditionalMetadata",
		"event_queue_arn":                "EventQueueArn",
		"exclusions":                     "Exclusions",
		"iceberg_targets":                "IcebergTargets",
		"jdbc_targets":                   "JdbcTargets",
		"lake_formation_configuration":   "LakeFormationConfiguration",
		"maximum_traversal_depth":        "MaximumTraversalDepth",
		"mongo_db_targets":               "MongoDBTargets",
		"name":                           "Name",
		"path":                           "Path",
		"paths":                          "Paths",
		"recrawl_behavior":               "RecrawlBehavior",
		"recrawl_policy":                 "RecrawlPolicy",
		"role":                           "Role",
		"s3_targets":                     "S3Targets",
		"sample_size":                    "SampleSize",
		"schedule":                       "Schedule",
		"schedule_expression":            "ScheduleExpression",
		"schema_change_policy":           "SchemaChangePolicy",
		"table_prefix":                   "TablePrefix",
		"tables":                         "Tables",
		"tags":                           "Tags",
		"targets":                        "Targets",
		"update_behavior":                "UpdateBehavior",
		"use_lake_formation_credentials": "UseLakeFormationCredentials",
		"write_manifest":                 "WriteManifest",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
