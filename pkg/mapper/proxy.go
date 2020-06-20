package mapper

import (
	"pgxs.io/powerdoc/pkg/mapper/api"
	"pgxs.io/powerdoc/pkg/mapper/sql"
)

//UserMapper return userMapper(sql/nosql)
func UserMapper() api.UserMapper {
	//check backend db storage is sql or nosql:mongo
	//now only support sql
	return sql.UserSqlMapperInstance()
}
