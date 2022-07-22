package msg

var (
	OK = "ok"
	// ParamsNotValid 参数异常
	ParamsNotValid = "参数异常"
	// DBCreateSuccess 新增成功
	DBCreateSuccess = "新增成功"
	// DBCreateError 新增失败
	DBCreateError = "新增失败"
	// DBGetSuccess 查询成功
	DBGetSuccess = "查询成功"
	// DBGetError 数据查询失败
	DBGetError = "数据查询失败"
	// DBUpdateSuccess 更新成功
	DBUpdateSuccess = "更新成功"
	// DBUpdateError 数据更新失败
	DBUpdateError = "数据更新失败"
	// DBDelSuccess 删除成功
	DBDelSuccess = "删除成功"
	// DBDelError 删除失败
	DBDelError = "删除失败"
	// AuthNotValid 鉴权失败
	AuthNotValid = "鉴权失败"
	// UnAuthorizedError 无权操作
	UnAuthorizedError = "无权操作"
	// ExportExcelNone 没有导出任务
	ExportExcelNone = "no_running"
	// ExportExcelWait Excel正在导出，请稍后
	ExportExcelWait = "running"
	// ExportExcelSuccess 导出Excel成功
	ExportExcelSuccess = "success"
	// ExportExcelError 导出Excel异常
	ExportExcelError = "error"

	LoginSuccess = "登录成功"

	LoginError = "登录失败"

	UploadSuccess = "上传成功"
	UploadError   = "上传失败"
)
