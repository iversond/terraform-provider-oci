// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

/*
 * This example shows how to use the budget and alert rule resources.
 */

variable "tenancy_ocid" {}

variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "compartment_ocid" {}
variable "region" {}

provider "oci" {
  region           = "${var.region}"
  tenancy_ocid     = "${var.tenancy_ocid}"
  user_ocid        = "${var.user_ocid}"
  fingerprint      = "${var.fingerprint}"
  private_key_path = "${var.private_key_path}"
}

resource "oci_budget_budget" "test_budget" {
  #Required
  amount                = "1"
  compartment_id        = "${var.tenancy_ocid}"
  reset_period          = "MONTHLY"
  target_compartment_id = "${var.compartment_ocid}"

  #Optional
  description  = "budget1 description"
  display_name = "budget1"
}

data "oci_budget_budget" "budget1" {
  budget_id = "${oci_budget_budget.test_budget.id}"
}

output "budget" {
  value = {
    amount                = "${data.oci_budget_budget.budget1.amount}"
    compartment_id        = "${data.oci_budget_budget.budget1.compartment_id}"
    reset_period          = "${data.oci_budget_budget.budget1.reset_period}"
    target_compartment_id = "${data.oci_budget_budget.budget1.target_compartment_id}"
    description           = "${data.oci_budget_budget.budget1.description}"
    display_name          = "${data.oci_budget_budget.budget1.display_name}"
    alert_rule_count      = "${data.oci_budget_budget.budget1.alert_rule_count}"
    state                 = "${data.oci_budget_budget.budget1.state}"
    time_created          = "${data.oci_budget_budget.budget1.time_created}"
    time_updated          = "${data.oci_budget_budget.budget1.time_updated}"
    version               = "${data.oci_budget_budget.budget1.version}"

    # These values are not always present
    //    actual_spend        = "${data.oci_budget_budget.budget1.actual_spend}"
    //    forecasted_spend    = "${data.oci_budget_budget.budget1.forecasted_spend}"
    //    time_spend_computed = "${data.oci_budget_budget.budget1.time_spend_computed}"
  }
}

data "oci_budget_budgets" "test_budgets" {
  #Required
  compartment_id = "${var.tenancy_ocid}"

  #Optional
  //  display_name = "${oci_budget_budget.test_budget.display_name}"
  //  state        = "ACTIVE"
}

output "budgets" {
  value = "${data.oci_budget_budgets.test_budgets.budgets}"
}

resource "oci_budget_alert_rule" "test_alert_rule" {
  #Required
  budget_id      = "${oci_budget_budget.test_budget.id}"
  recipients     = "JohnSmith@example.com"
  threshold      = "100"
  threshold_type = "ABSOLUTE"
  type           = "ACTUAL"

  #Optional
  description  = "alertRuleDescription"
  display_name = "alertRule"
  message      = "possible overspend"
}

data "oci_budget_alert_rules" "test_alert_rules" {
  #Required
  budget_id = "${oci_budget_budget.test_budget.id}"

  #Optional
  //  display_name = "${oci_budget_alert_rule.test_alert_rule.display_name}"
  state = "ACTIVE"
}

output "alert_rule" {
  value = {
    budget_id      = "${data.oci_budget_alert_rule.test_alert_rule.budget_id}"
    recipients     = "${data.oci_budget_alert_rule.test_alert_rule.recipients}"
    description    = "${data.oci_budget_alert_rule.test_alert_rule.description}"
    display_name   = "${data.oci_budget_alert_rule.test_alert_rule.display_name}"
    message        = "${data.oci_budget_alert_rule.test_alert_rule.message}"
    recipients     = "${data.oci_budget_alert_rule.test_alert_rule.recipients}"
    state          = "${data.oci_budget_alert_rule.test_alert_rule.state}"
    threshold      = "${data.oci_budget_alert_rule.test_alert_rule.threshold}"
    threshold_type = "${data.oci_budget_alert_rule.test_alert_rule.threshold_type}"
    time_created   = "${data.oci_budget_alert_rule.test_alert_rule.time_created}"
    time_updated   = "${data.oci_budget_alert_rule.test_alert_rule.time_updated}"
    type           = "${data.oci_budget_alert_rule.test_alert_rule.type}"
    version        = "${data.oci_budget_alert_rule.test_alert_rule.version}"
  }
}

data "oci_budget_alert_rule" "test_alert_rule" {
  #Required
  budget_id     = "${oci_budget_budget.test_budget.id}"
  alert_rule_id = "${oci_budget_alert_rule.test_alert_rule.id}"
}

output "alert_rules" {
  value = "${data.oci_budget_alert_rules.test_alert_rules.alert_rules}"
}
