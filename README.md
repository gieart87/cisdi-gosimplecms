### Database Migration
run this command on working dir

`goose -dir db/migrations mysql "root:password@tcp(localhost:33060)/gosimplecms_db" up`