#!/usr/bin/env bash

echo 'Running migrations ...'
/myapp/migrate up > /dev/null 2>&1 &

echo 'Deleting mysql_cloinet ...'
apk del mysql-client

echo 'Start application...'
/myapp/app