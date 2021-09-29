// Code generated by goctl. DO NOT EDIT!
// Source: chat.proto

package server

import (
	"context"

	"github.com/acger/chat-svc/internal/logic"
	"github.com/acger/chat-svc/internal/svc"
	"github.com/acger/chat-svc/template"
)

type ChatServer struct {
	svcCtx *svc.ServiceContext
}

func NewChatServer(svcCtx *svc.ServiceContext) *ChatServer {
	return &ChatServer{
		svcCtx: svcCtx,
	}
}

func (s *ChatServer) MessageSave(ctx context.Context, in *template.MsgSaveReq) (*template.Rsp, error) {
	l := logic.NewMessageSaveLogic(ctx, s.svcCtx)
	return l.MessageSave(in)
}

func (s *ChatServer) MessageList(ctx context.Context, in *template.MsgListReq) (*template.MsgListRsp, error) {
	l := logic.NewMessageListLogic(ctx, s.svcCtx)
	return l.MessageList(in)
}

func (s *ChatServer) ChatHistoryList(ctx context.Context, in *template.ChatHistoryReq) (*template.ChatHistoryRsp, error) {
	l := logic.NewChatHistoryListLogic(ctx, s.svcCtx)
	return l.ChatHistoryList(in)
}

func (s *ChatServer) ChatNumber(ctx context.Context, in *template.ChatNumberReq) (*template.ChatNumberRsp, error) {
	l := logic.NewChatNumberLogic(ctx, s.svcCtx)
	return l.ChatNumber(in)
}

func (s *ChatServer) ChatHistorySave(ctx context.Context, in *template.CHSaveReq) (*template.Rsp, error) {
	l := logic.NewChatHistorySaveLogic(ctx, s.svcCtx)
	return l.ChatHistorySave(in)
}
