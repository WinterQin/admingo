package itface

type ORMProvider interface {
	Create(model interface{}) error
	Update(model interface{}) error
	Delete(model interface{}) error
	FindByID(model interface{}, id interface{}) error
	FindAll(models interface{}) error
	Paginate(models interface{}, page, pageSize int) (total int64, err error)
}
