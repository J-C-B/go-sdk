---
title: "lacework generate cloud-account aws"
slug: lacework_generate_cloud-account_aws
hide_title: true
---

## lacework generate cloud-account aws

Generate and/or execute Terraform code for AWS integration

### Synopsis

Use this command to generate Terraform code for deploying Lacework into an AWS environment.

By default, this command interactively prompts for the required information to setup the new cloud account.
In interactive mode, this command will:

* Prompt for the required information to setup the integration
* Generate new Terraform code using the inputs
* Optionally, run the generated Terraform code:
  * If Terraform is already installed, the version is verified as compatible for use
	* If Terraform is not installed, or the version installed is not compatible, a new version will be installed into a temporary location
	* Once Terraform is detected or installed, Terraform plan will be executed
	* The command will prompt with the outcome of the plan and allow to view more details or continue with Terraform apply
	* If confirmed, Terraform apply will be run, completing the setup of the cloud account

This command can also be run in noninteractive mode.
See help output for more details on the parameter value(s) required for Terraform code generation.


```
lacework generate cloud-account aws [flags]
```

### Options

```
      --apply                                 run terraform apply without executing plan or prompting
      --aws_profile string                    specify aws profile
      --aws_region string                     specify aws region
      --aws_subaccount strings                configure an additional aws account; value format must be <aws profile>:<region>
      --bucket_encryption_enabled             enable S3 bucket encryption when creating bucket (default true)
      --bucket_name string                    specify bucket name when creating bucket
      --bucket_sse_key_arn string             specify existing KMS encryption key arn for bucket
      --cloudtrail                            enable cloudtrail integration
      --cloudtrail_name string                specify name of cloudtrail integration
      --config                                enable config integration
      --config_name string                    specify name of config integration
      --consolidated_cloudtrail               use consolidated trail
      --existing_bucket_arn string            specify existing cloudtrail S3 bucket ARN
      --existing_iam_role_arn string          specify existing iam role arn to use
      --existing_iam_role_externalid string   specify existing iam role external_id to use
      --existing_iam_role_name string         specify existing iam role name to use
      --existing_sns_topic_arn string         specify existing SNS topic arn
      --force_destroy_s3                      enable force destroy S3 bucket
  -h, --help                                  help for aws
      --lacework_aws_account_id string        the Lacework AWS root account id
      --output string                         location to write generated content (default is ~/lacework/aws)
      --sns_topic_encryption_enabled          enable encryption on SNS topic when creating one (default true)
      --sns_topic_encryption_key_arn string   specify existing KMS encryption key arn for SNS topic
      --sns_topic_name string                 specify SNS topic name if creating new one
      --sqs_encryption_enabled                enable encryption on SQS queue when creating (default true)
      --sqs_encryption_key_arn string         specify existing KMS encryption key arn for SQS queue
      --sqs_queue_name string                 specify SQS queue name if creating new one
      --use_s3_bucket_notification            enable S3 bucket notifications
```

### Options inherited from parent commands

```
  -a, --account string      account subdomain of URL (i.e. <ACCOUNT>.lacework.net)
  -k, --api_key string      access key id
  -s, --api_secret string   secret access key
      --api_token string    access token (replaces the use of api_key and api_secret)
      --debug               turn on debug logging
      --json                switch commands output from human-readable to json format
      --nocache             turn off caching
      --nocolor             turn off colors
      --noninteractive      turn off interactive mode (disable spinners, prompts, etc.)
      --organization        access organization level data sets (org admins only)
  -p, --profile string      switch between profiles configured at ~/.lacework.toml
      --subaccount string   sub-account name inside your organization (org admins only)
```

### SEE ALSO

* [lacework generate cloud-account](lacework_generate_cloud-account.md)	 - Generate cloud integration IaC

