#!/bin/sh

TIMEOUT=5
WAIT=1

while : 
do
    echo waiting for db
    echo "nc -w $TIMEOUT $POSTGRES_HOST $POSTGRES_PORT"
    result=$?
    if [ $result -eq 0 ]; then
        echo db is ready
        break
    else
        echo db is not ready, waiting $TIMEOUT second again
    fi
    sleep $WAIT
done

set -e

CONN_STR="postgres://$POSTGRES_USER:$POSTGRES_PASSWORD@$POSTGRES_HOST:$POSTGRES_PORT/$POSTGRES_DB?sslmode=disable"

echo migrating...

migrate -path /app/db/migrations -database $CONN_STR -verbose up

echo migration completed!

exec "$@"
