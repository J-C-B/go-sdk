---
title: "lacework compliance google get-report"
slug: lacework_compliance_google_get-report
hide_title: true
---

## lacework compliance google get-report

Get the latest GCP compliance report

### Synopsis

Get the latest compliance assessment report, these reports run on a regular schedule,
typically once a day. The available report formats are human-readable (default), json and pdf.

To list all GCP projects and organizations configured in your account:

    lacework compliance gcp list

To show recommendation details and affected resources for a recommendation id:

    lacework compliance gcp get-report <organization_id> <project_id> [recommendation_id]

To retrieve a specific report by its report name:

    lacework compliance gcp get-report <organization_id> <project_id> --report_name 'GCP Cybersecurity Maturity'


```
lacework compliance google get-report <organization_id> <project_id> [flags]
```

### Options

```
      --category strings     filter report details by category (storage, networking, identity-and-access-management, ...)
      --csv                  output report in CSV format
      --details              increase details about the compliance report
  -h, --help                 help for get-report
      --pdf                  download report in PDF format
      --report_name string   report name to display, run 'lacework report-definitions list' for more information.
      --service strings      filter report details by service (gcp:storage:bucket, gcp:kms:cryptoKey, gcp:project, ...)
      --severity string      filter report details by severity threshold (critical, high, medium, low, info)
      --status string        filter report details by status (non-compliant, requires-manual-assessment, suppressed, compliant, could-not-assess)
      --type string          report type to display, run 'lacework report-definitions list' for more information.
                             valid types:
                             'GCP_CIS','GCP_CIS12','GCP_CIS13','GCP_CIS_1_3_0_NIST_800_171_rev2','GCP_CIS_1_3_0_NIST_800_53_rev5',
                             'GCP_CIS_1_3_0_NIST_CSF','GCP_CMMC_1_02','GCP_HIPAA','GCP_HIPAA_2013','GCP_HIPAA_Rev2',
                             'GCP_ISO_27001','GCP_ISO_27001_2013','GCP_K8S','GCP_NIST_800_171_REV2','GCP_NIST_800_53_REV4',
                             'GCP_NIST_CSF','GCP_PCI','GCP_PCI_DSS_3_2_1','GCP_PCI_Rev2','GCP_SOC',
                             'GCP_SOC_2','GCP_SOC_Rev2', (default "GCP_CIS13")
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

* [lacework compliance google](lacework_compliance_google.md)	 - Compliance for Google Cloud

