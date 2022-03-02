package driver

type DriverSql struct {
}

func NewDriverSql() *DriverSql {
	return &DriverSql{}
}

var values map[string]interface{} = make(map[string]interface{})

func (d *DriverSql) Create(key string, value interface{}) bool {

	values[key] = value

	return true
}

func (d *DriverSql) Get(key string) *map[string]interface{} {
	return &values
}

func (d *DriverSql) Update(key string, value interface{}) bool {
	values[key] = value
	return true
}
