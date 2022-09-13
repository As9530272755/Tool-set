package model

import "time"

// 人员信息
type user struct {
	Comment   string `json:"comment"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	Phone     string `json:"phone"`
	Role      string `json:"role"`
	Wechat_id string `json:"wechat_id"`
}

func NewUser() *user {
	return &user{
		Name:      "null",
		Email:     "null",
		Phone:     "null",
		Wechat_id: "null",
		Role:      "null",
		Comment:   "null",
	}
}

// 规则信息
type alarmRules struct {
	Email_receiver   string `json:"email_receiver"`
	Message_receiver string `json:"message_receiver"`
	Restrain_max     int    `json:"restrain_max"`
	Restrain_stage   int    `json:"restrain_stage"`
	Warn_name        string `json:"warn_name"`
	Warn_priority    int    `json:"warn_priority"`
	Wechat_receiver  string `json:"wechat_receiver"`
}

func NewAlarmRules() *alarmRules {

	return &alarmRules{
		Email_receiver:   "",
		Message_receiver: "",
		Restrain_max:     5,
		Restrain_stage:   1,
		Warn_name:        "",
		Warn_priority:    10,
		Wechat_receiver:  "",
	}
}

// 查询规则
type Warnrule struct {
	IndexArray []struct {
		ColName   string `json:"colName"`
		ColType   string `json:"colType"`
		IndexType string `json:"indexType"`
		Relation  string `json:"relation"`
		Value     string `json:"value"`
	} `json:"indexArray"`
	InfoType    string `json:"infoType"`
	IsPage      bool   `json:"isPage"`
	OrderField  string `json:"orderField"`
	OrderOrient string `json:"orderOrient"`
	PageCount   int    `json:"pageCount"`
	PageNum     int    `json:"pageNum"`
}

func NewWarnrule(value string) *Warnrule {
	return &Warnrule{
		PageCount: 10,
		IndexArray: []struct {
			ColName   string `json:"colName"`
			ColType   string `json:"colType"`
			IndexType string `json:"indexType"`
			Relation  string `json:"relation"`
			Value     string `json:"value"`
		}{
			{
				ColName:   "warn_name",
				ColType:   "string",
				IndexType: "7",
				Relation:  "1",
				Value:     value,
			},
		},
	}
}

// 用于实现新增告警字段
type warnRuleDetailUpdate struct {
	WarnField string `json:"warn_field"`
	Logic     int    `json:"logic"`
	WarnValue string `json:"warn_value"`
	RuleId    int    `json:"rule_id"`
}

func NewRuleDetail() *warnRuleDetailUpdate {
	return &warnRuleDetailUpdate{
		WarnField: "",
		Logic:     0,
		WarnValue: "",
		RuleId:    0,
	}
}

// 用于静默时间
type stopTime struct {
	BeginTime    string `json:"begin_time"`
	Comment      string `json:"comment"`
	CreateTime   string `json:"create_time"`
	EndTime      string `json:"end_time"`
	Id           int    `json:"id"`
	Type         int    `json:"type"` // 用于定义静默类型为单次或者每天
	UpdateTime   string `json:"update_time"`
	WarnRuleId   int    `json:"warn_rule_id"`
	WarnRuleName string `json:"warn_rule_name"`
}

func NewStopTime(warnRuleName, beginTime, endTime string, warnRuleId, Type int) *stopTime {
	return &stopTime{
		WarnRuleName: warnRuleName,
		BeginTime:    beginTime,
		EndTime:      endTime,
		WarnRuleId:   warnRuleId,
		Type:         Type,
		CreateTime:   time.Now().Format("2006-01-02 15:04:05"),
		UpdateTime:   time.Now().Format("2006-01-02 15:04:05"),
	}
}
