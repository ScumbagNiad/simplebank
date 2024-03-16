# Setting up the project

## DB

Get the postgres alpine 12 docker image

Make sure you have Docker, Make, SQLC, Go-Migrate.

Start the project by running 

    "make postgres" 

which will start a docker image, after that you can run //If you are using mac you will need to modify the migrateup and down make commands

    "make migrateup"

to setup the DB, after that you can run 

    "make sqlc" 

to make sure all database commands are generated correctly.

## GO

Start project