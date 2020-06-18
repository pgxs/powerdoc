docker run --name pg_powerdoc -p 5432:5432 -e POSTGRES_PASSWORD=123456 \
  -e POSTGRES_DB=powerdoc -d postgres:10-alpine