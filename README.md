Command to create a new migration logic:

migrate create -seq -ext=.sql -dir=./pkg/db/migrations create_ratings_table
    -seq: to use number for ordering 
    -ext= specify the migration file extension
    -dir= directory to the migrations directory

