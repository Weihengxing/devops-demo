package store

type Condition map[string]interface{}

type IStore interface {
	Get()
	Del()
	Apply()  //有就更新，没有就添加
}
