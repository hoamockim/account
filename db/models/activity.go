package models

type ActivityType string

const (
	Order ActivityType = "Order"
	View  ActivityType = "View"
)

type Activity struct {
	BaseModel
	CustomerId   int
	ActivityType ActivityType
	Detail       string
	ActivityTime int64
}
