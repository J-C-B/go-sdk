---
title: "lacework agent gcp-install osl"
slug: lacework_agent_gcp-install_osl
hide_title: true
---

## lacework agent gcp-install osl

Use OSLogin to securely connect to GCE instances

### Synopsis

This command installs the agent on all GCE instances in a GCP organization using OSLogin.

The username of the GCP user or service account, in the form `users/<username>`, is a
required argument.

This command will attempt to query the GCE metadata server for the current project. If this
command is not run on a GCE instance, pass the project ID as:

    lacework agent gcp-install osl <gcp_username> --project_id my-project-id

To filter by one or more regions:

    lacework agent gcp-install osl <gcp_username> --include_regions us-west1,europe-west2

To filter by instance metadata:

    lacework agent gcp-install osl <gcp_username> --metadata MetadataKey,MetadataValue

To filter by instance metadata key:

    lacework agent gcp-install osl <gcp_username> --metadata_key MetadataKey

To provide an agent access token of your choice, use the command 'lacework agent token list',
select a token and pass it to the '--token' flag. This flag must be selected if the
'--noninteractive' flag is set.

    lacework agent gcp-install osl <gcp_username> --token <token>

To explicitly specify the server URL that the agent will connect to:

    lacework agent gcp-install osl --server_url https://your.server.url.lacework.net

GCP credentials are read using the following environment variables:
- GOOGLE_APPLICATION_CREDENTIALS

This command will automatically add hosts with successful connections to
'~/.ssh/known_hosts' unless specified with '--trust_host_key=false'.

```
lacework agent gcp-install osl [flags]
```

### Options

```
  -h, --help                      help for osl
  -r, --include_regions strings   list of regions to filter on
  -n, --max_parallelism int       maximum number of workers executing GCP API calls, set if rate limits are lower or higher than normal (default 50)
      --metadata strings          only install agents on infra with this metadata
      --metadata_key string       only install agents on infra with this metadata key set
      --project_id string         ID of the GCP project, set if metadata server does not provide
      --server_url https://       server URL that agents will talk to, prefixed with https:// (default "https://api.lacework.net")
      --token string              agent access token
      --trust_host_key            automatically add host keys to the ~/.ssh/known_hosts file (default true)
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

* [lacework agent gcp-install](lacework_agent_gcp-install.md)	 - Install the datacollector agent on all remote GCE hosts

