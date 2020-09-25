package vojo

const (
	//the response is normal
	RES_NORMAL = 0
	//request error
	RES_ERROR = 1
	//http status is not 200
	RES_PROTOCAL_ERROR = 2
)
const (
	HTTP_REQ_TYPE = 0
	GRPC_REQ_TYPE = 1
)

type TasksDao struct {
	Task_name  string `db:"task_name" json:"task_name"`
	Task_cron  string `db:"task_cron" json:"task_cron"`
	User_id    int    `db:"user_id" json:"user_id"`
	Url        string `db:"url" json:"url"`
	Id         int    `db:"id" json:"id"`
	ReqType    int    `db:"req_type" json:"req_type"`
	TaskStatus int    `db:"task_status" json:"task_status"`
	Timestamp  string `db:"_timestamp" json:"timestamp"`
}
type TaskDaoListSlice []*TasksDao

func (s TaskDaoListSlice) Len() int      { return len(s) }
func (s TaskDaoListSlice) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s TaskDaoListSlice) Less(i, j int) bool {
	return s[i].Timestamp < s[j].Timestamp
}
