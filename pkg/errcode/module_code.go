package errcode

var (
	ErrorGetTagListFail = NewError(200010001, "获取标签列表失败")
	ErrorCreateTagFail  = NewError(200010002, "创建标签失败")
	ErrorUpdateTagFail  = NewError(200010003, "更新标签失败")
	ErrorDeleteTagFail  = NewError(200010004, "删除标签失败")
	ErrorCountTagFail   = NewError(200010005, "统计标签失败")
	ErrorUploadFileFail   = NewError(20030001, "上传文件失败")
)
