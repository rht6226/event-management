
.PHONY run:
	@echo "Running Go application...";
	HOST_NAME="localhost" \
	PORT="5432" \
	DATABASE_NAME="mydatabase" \
	USER="myuser" \
	PASSWORD="mypassword" \
	SSL_MODE="disable" \
	TIMEZONE="UTC" \
	DEPLOYMENT_ENVIRONMENT="dev"\
	go run .

stop:
	docker compose down	

run-db:
	docker compose up -d

