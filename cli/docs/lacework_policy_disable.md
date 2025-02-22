---
title: "lacework policy disable"
slug: lacework_policy_disable
hide_title: true
---

## lacework policy disable

Disable policies

### Synopsis

Disable policies by ID or all policies matching a tag.

To disable a single policy by its ID:

	lacework policy disable lacework-policy-id

To disable many policies by ID provide a list of policy ids:

	lacework policy disable lacework-policy-id-one lacework-policy-id-two

To disable all policies for AWS CIS 1.4.0:

	lacework policy disable --tag framework:cis-aws-1-4-0

To disable all policies for GCP CIS 1.3.0:

	lacework policy disable --tag framework:cis-gcp-1-3-0


```
lacework policy disable [policy_id...] [flags]
```

### Options

```
  -h, --help         help for disable
      --tag string   disable all policies with the specified tag
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

* [lacework policy](lacework_policy.md)	 - Manage policies

