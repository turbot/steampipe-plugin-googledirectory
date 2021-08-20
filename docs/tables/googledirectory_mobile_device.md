# Table: googledirectory_mobile_device

Query information about mobile devices defined in the Google Workspace directory.

## Examples

### Basic info

```sql
select
  model,
  resource_id,
  device_id,
  brand
from
  googledirectory_mobile_device;
```

### Get Mobile Device by ID

```sql
select
  model,
  resource_id,
  device_id,
  brand
from
  googledirectory_mobile_device
where
  device_id = 'id:03pk8a4z4t34g1w';
```
