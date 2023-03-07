# Table: oracle_networkfirewall_network_firewall_policies

The composite primary key for this table is (**region**, **compartment_id**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|region (PK)|String|
|compartment_id (PK)|String|
|id (PK)|String|
|display_name|String|
|time_created|Timestamp|
|lifecycle_state|String|
|freeform_tags|JSON|
|defined_tags|JSON|
|time_updated|Timestamp|
|lifecycle_details|String|
|system_tags|JSON|