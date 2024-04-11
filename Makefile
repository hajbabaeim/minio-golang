include .env
export

create-mc:
	mc alias set $(ALIAS) $(ENDPOINT) $(ACCESS_KEY) $(SECRET_KEY)

add-policy:
	mc admin policy add $(ALIAS) $(POLICY_NAME) ./readwrite-perm.json

create-user:
	mc admin user add $(ALIAS) $(USERNAME) $(ACCESS_KEY) $(SECRET_KEY)

assign-policy:
	mc admin policy set $(ALIAS) $(POLICY_NAME) user=$(USERNAME)