#!/usr/bin/env bash
#wait for the MySQL Server to come up
#sleep 90s

#run the setup script to create the DB and the schema in the DB
mysql -u docker -pdocker test_database < "/docker-entrypoint-initdb.d/create-users-tables.sql"
mysql -u docker -pdocker test_database < "/docker-entrypoint-initdb.d/create-characters-tables.sql"
mysql -u docker -pdocker test_database < "/docker-entrypoint-initdb.d/create-character-user-tables.sql"
mysql -u docker -pdocker test_database < "/docker-entrypoint-initdb.d/create-gacha-contents-tables.sql"
mysql -u docker -pdocker test_database < "/docker-entrypoint-initdb.d/insert-characters.sql"
mysql -u docker -pdocker test_database < "/docker-entrypoint-initdb.d/insert-probability.sql"