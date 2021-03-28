package client

// 由于leveldb是kv性数据库, 所以使用分配好的 PRE + user_id 用于区分业务和用户, 使用前必须分配！！！
const (
	PRE_TEST = "test_"
	PRE_DDL  = "ddl_"
	PRE_TODO = "todo_"
)

// PrefixKey
func PrefixKey(pre, user_id string) string {
	return pre + user_id + "_"
}

// ComposeKey
func ComposeKey(pre, user_id, key string) string {
	return pre + user_id + "_" + key
}
