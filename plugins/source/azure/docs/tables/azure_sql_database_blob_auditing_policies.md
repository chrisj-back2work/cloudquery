# Table: azure_sql_database_blob_auditing_policies


The primary key for this table is **id**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
|subscription_id|String|
|sql_database_id|UUID|
|kind|String|
|state|String|
|storage_endpoint|String|
|storage_account_access_key|String|
|retention_days|Int|
|audit_actions_and_groups|StringArray|
|storage_account_subscription_id|UUID|
|is_storage_secondary_key_in_use|Bool|
|is_azure_monitor_target_enabled|Bool|
|queue_delay_ms|Int|
|id (PK)|String|
|name|String|
|type|String|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|