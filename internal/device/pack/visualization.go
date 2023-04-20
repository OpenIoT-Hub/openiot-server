package pack

type (
	TableData struct {
		ID      int         `json:"id"`
		Name    string      `json:"name"`
		Address string      `json:"address"`
		Field1  interface{} `json:"field1"`
		Field2  interface{} `json:"field2"`
	}

	ChartData struct {
		Name  string      `json:"name"`
		Data  interface{} `json:"data"`
		Extra interface{} `json:"extra"`
		Units interface{} `json:"units"`
	}
)
