docker run --name mysql_powerdoc -p 3306:3306  -e MYSQL_ALLOW_EMPTY_PASSWORD=yes -e TZ=Asia/Shanghai -e MYSQL_DATABASE=powerdoc -d mysql:8