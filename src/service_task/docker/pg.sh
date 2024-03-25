docker run -d --rm \
	--name pgdb \
	-e POSTGRES_PASSWORD=admin \
	-e POSTGRES_USER=admin \
	-e POSTGRES_DB=db \
  -p 55000:5432 \
	-e PGDATA=/var/lib/postgresql/data/pgdata \
	-v .:/var/lib/postgresql/data \
	postgres
