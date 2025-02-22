package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ ConversationMembersModel = (*customConversationMembersModel)(nil)

type (
	// ConversationMembersModel is an interface to be customized, add more methods here,
	// and implement the added methods in customConversationMembersModel.
	ConversationMembersModel interface {
		conversationMembersModel
	}

	customConversationMembersModel struct {
		*defaultConversationMembersModel
	}
)

// NewConversationMembersModel returns a model for the database table.
func NewConversationMembersModel(conn sqlx.SqlConn) ConversationMembersModel {
	return &customConversationMembersModel{
		defaultConversationMembersModel: newConversationMembersModel(conn),
	}
}
