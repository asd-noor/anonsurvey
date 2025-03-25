LOCAL_DB_INSTANCE = "anon-mysql"
LOCAL_DB_ROOT_PASS = "toor"
LOCAL_DB_USER_NAME = "anon"
LOCAL_DB_USER_PASS = "anon"
LOCAL_DB_SCHEMA = "anonsurvey"
LOCAL_DB_DATA_DIR ?= "$(pwd)/local-db"

run-db:
	@mkdir -p $(LOCAL_DB_DATA_DIR)
	docker run -d --rm \
		-p 3306:3306 \
		-v $(LOCAL_DB_DATA_DIR):/var/lib/mysql \
		-e MYSQL_ROOT_PASSWORD=$(LOCAL_DB_ROOT_PASS) \
		-e MYSQL_USER=$(LOCAL_DB_USER_NAME) \
		-e MYSQL_PASSWORD=$(LOCAL_DB_USER_PASS) \
		-e MYSQL_DATABASE=$(LOCAL_DB_SCHEMA) \
		--name $(LOCAL_DB_INSTANCE) \
		mysql:8

.SILENT: open-db
open-db:
	docker exec -it $(LOCAL_DB_INSTANCE) mysql -u$(LOCAL_DB_USER_NAME) -p$(LOCAL_DB_USER_PASS) $(LOCAL_DB_SCHEMA)
