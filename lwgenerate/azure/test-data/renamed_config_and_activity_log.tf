terraform {
  required_providers {
    lacework = {
      source  = "lacework/lacework"
      version = "~> 1.0"
    }
  }
}

provider "azuread" {
}

provider "azurerm" {
  features {
  }
}

module "az_ad_application" {
  source  = "lacework/ad-application/azure"
  version = "~> 1.0"
}

module "az_config" {
  source                      = "lacework/config/azure"
  version                     = "~> 2.0"
  application_id              = module.az_ad_application.application_id
  application_password        = module.az_ad_application.application_password
  lacework_integration_name   = "Test Config Rename"
  service_principal_id        = module.az_ad_application.service_principal_id
  use_existing_ad_application = true
}

module "az_activity_log" {
  source                      = "lacework/activity-log/azure"
  version                     = "~> 2.0"
  application_id              = module.az_ad_application.application_id
  application_password        = module.az_ad_application.application_password
  lacework_integration_name   = "Test Activity Log Rename"
  service_principal_id        = module.az_ad_application.service_principal_id
  use_existing_ad_application = true
}
