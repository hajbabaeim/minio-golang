create-mc:
	mc alias set <ALIAS> <YOUR-MINIO-ENDPOINT> <ACCESS-KEY> <SECRET-KEY>

add-policy:
	mc admin policy add <ALIAS> <POLICY-NAME> ./readwrite-perm.json

create-user:
	mc admin user add <ALIAS> <USERNAME> <ACCESS-KEY> <SECRET-KEY>

assign-policy:
	mc admin policy set <ALIAS> <POLICY-NAME> user=<USERNAME>