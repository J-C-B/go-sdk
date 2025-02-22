---
title: "lacework agent aws-install ec2ssm"
slug: lacework_agent_aws-install_ec2ssm
hide_title: true
---

## lacework agent aws-install ec2ssm

Use SSM to securely install the Lacework agent on EC2 instances

### Synopsis

This command installs the Lacework agent on all EC2 instances in an AWS account using SSM.

This command will create a role and instance profile with 'SSMManagedInstanceCore'
attached and associate that instance profile with the target instances. If the target
instances already have associated instance profiles, this command will not change
their state. This command will teardown the IAM role and instance profile before exiting.

This command authenticates with AWS credentials from well-known locations on the user's
machine. The principal associated with these credentials should have the
'AmazonEC2FullAccess', 'IAMFullAccess' and 'AmazonSSMFullAccess' policies attached.

Target instances must have the SSM agent installed and running for successful
installation.

To skip IAM role / instance profile creation and instance profile association:

    lacework agent aws-install ec2ssm --skip_iam_role_creation

To provide a preexisting IAM role with the 'SSMManagedInstanceCore' policy

    lacework agent aws-install ec2ssm --iam_role_name IAMRoleName

To filter by one or more regions:

    lacework agent aws-install ec2ssm --include_regions us-west-2,us-east-2

To filter by instance tag:

    lacework agent aws-install ec2ssm --tag TagName,TagValue

To filter by instance tag key:

    lacework agent aws-install ec2ssm --tag_key TagName

To provide an agent access token of your choice, use the command 'lacework agent token list',
select a token and pass it to the '--token' flag. This flag must be selected if the
'--noninteractive' flag is set.

    lacework agent aws-install ec2ssm --token <token>

To explicitly specify the server URL that the agent will connect to:

    lacework agent aws-install ec2ssm --server_url https://your.server.url.lacework.net

To specify an AWS credential profile other than 'default':

    lacework agent aws-install ec2ssm --credential_profile aws-profile-name

AWS credentials are read from the following environment variables:
- AWS_ACCESS_KEY_ID
- AWS_SECRET_ACCESS_KEY
- AWS_SESSION_TOKEN (optional)
- AWS_REGION

```
lacework agent aws-install ec2ssm [flags]
```

### Options

```
      --credential_profile string   AWS credential profile to use (default "default")
  -d, --dry_run                     set this flag to print out the target instances and exit
  -f, --force_reinstall             set this flag to force-reinstall the agent, even if already running on the target instance
  -h, --help                        help for ec2ssm
      --iam_role_name string        IAM role name (not ARN) with SSM policy, if not provided then an ephemeral role will be created
  -r, --include_regions strings     list of regions to filter on
  -n, --max_parallelism int         maximum number of workers executing AWS API calls, set if rate limits are lower or higher than normal (default 50)
      --server_url https://         server URL that agents will talk to, prefixed with https:// (default "https://api.lacework.net")
      --skip_iam_role_creation      set this flag to skip creating an IAM role and instance profile and associating the instance profile. Assumes all instances are already setup for SSM
      --tag strings                 only install agents on infra with this tag
      --tag_key string              only install agents on infra with this tag key set
      --token string                agent access token
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

* [lacework agent aws-install](lacework_agent_aws-install.md)	 - Install the datacollector agent on all remote AWS hosts

