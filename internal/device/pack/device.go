package pack

type (
	DeviceData interface {
		GetTableData(TableData, error)
		GetChartData(ChartData, error)
	}

	DataV1 struct {
		ID       int
		Name     string
		Address  string
		version  int
		external map[string]interface{}
	}
)

func (v1 *DataV1) GetTableData() (TableData, error) {
	//  根据version=v1从extranal解析数据
	var (
		data TableData
	)

	return data, nil
}

func (v1 *DataV1) GetChartData() (ChartData, error) {
	// 根据version=v1从extranal解析数据
	var (
		data ChartData
	)

	return data, nil
}
