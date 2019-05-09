// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

resource "random_string" "autonomous_data_warehouse_wallet_password" {
  length  = 16
  special = true
}

data "oci_database_autonomous_database_wallet" "autonomous_data_warehouse_wallet" {
  #Required
  autonomous_database_id = "${oci_database_autonomous_database.autonomous_data_warehouse.id}"
  password               = "${random_string.autonomous_data_warehouse_wallet_password.result}"
}

resource "local_file" "autonomous_data_warehouse_wallet_file" {
  content  = "${data.oci_database_autonomous_database_wallet.autonomous_data_warehouse_wallet.content}"
  filename = "${path.module}/autonomous_data_warehouse_wallet.zip"
}

output "autonomous_data_warehouse_wallet_password" {
  value = "${random_string.autonomous_data_warehouse_wallet_password.result}"
}
