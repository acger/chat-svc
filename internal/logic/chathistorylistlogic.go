package logic

import (
	"context"
	"errors"
	"github.com/acger/chat-svc/chat"
	"github.com/acger/chat-svc/internal/svc"
	"github.com/acger/chat-svc/template"
	"github.com/acger/user-svc/userclient"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"

	"github.com/tal-tech/go-zero/core/logx"
)

type ChatHistoryListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewChatHistoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChatHistoryListLogic {
	return &ChatHistoryListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ChatHistoryListLogic) ChatHistoryList(in *template.ChatHistoryReq) (*template.ChatHistoryRsp, error) {
	type ChatHistoryMessage struct {
		Uid    uint64
		Status bool
	}

	var chatHistory []*ChatHistoryMessage

	db := l.svcCtx.DB
	query := db.Table("chat_histories a")
	query = query.Joins("left join chats as b on b.uid = a.to_uid and b.to_uid = ?", in.Id)
	query = query.Where("a.uid = ?", in.Id)
	query = query.Group("a.to_uid")
	query = query.Order("status, max(b.created_at) desc")
	query = query.Select("a.to_uid as uid, IFNULL(MIN(b.status), true) as status")

	result := query.Find(&chatHistory)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return &template.ChatHistoryRsp{}, nil
	}

	var uidList []uint64
	for _, e := range chatHistory {
		uidList = append(uidList, e.Uid)
	}

	//获取用户信息
	userListRsp, _ := l.svcCtx.UserSvc.UserList(l.ctx, &userclient.UserListReq{Id: uidList})
	userMap := make(map[uint64]*userclient.UserInfo)

	for _, u := range userListRsp.User {
		userMap[u.Id] = u
	}

	var chatUserList []*chat.ChatUser

	for _, c := range chatHistory {
		tmp := &chat.ChatUser{
			Id:     c.Uid,
			Status: c.Status,
		}

		copier.Copy(tmp, userMap[c.Uid])
		chatUserList = append(chatUserList, tmp)
	}

	return &template.ChatHistoryRsp{Code: 0, User: chatUserList}, nil
}
