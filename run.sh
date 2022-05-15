echo "Running Migrations up";
migrate -path /app/db/migrations/ -database "mysql://$MYSQL_USER:$MYSQL_PASSWORD@tcp($MYSQL_URL:$MYSQL_PORT)/$MYSQL_DATABASE" up

echo "Running FreePad";
cd src && go run .;