# SQL migrations for Golang and PostgreSQL [![Build Status](https://travis-ci.org/dskarataev/migrations.svg)](https://travis-ci.org/dskarataev/migrations)

This package allows you to run migrations on your PostgreSQL database using [Golang Postgres client](https://github.com/go-pg/pg).

The difference between this package and original one https://github.com/go-pg/migrations that with this package you can run migrations from your code, not from command line. Also it saves migration's comment to the database for history.
