include .env
export

create-mc:
	mc alias set $(ALIAS) $(ENDPOINT) $(ACCESS_KEY) $(SECRET_KEY)

create-user:
	mc admin user add $(ALIAS) $(USERNAME) $(SECRET_KEY)

create-policy:
	mc admin policy create $(ALIAS) $(POLICY_NAME) $(POLICY_FILE_PATH)

assign-policy:
	mc admin policy attach $(ALIAS) $(POLICY_NAME) --user=$(USERNAME)

user-info:
	mc admin user info $(ALIAS) $(USERNAME)

create-bucket:
	mc mb $(ALIAS)/$(BUCKET_NAME)

buckets-list:
	mc ls $(ALIAS)

bucket-objects:
	mc ls --recursive --versions $(ALIAS)/$(BUCKET_NAME)