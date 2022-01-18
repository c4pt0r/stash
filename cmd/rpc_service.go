package cmd

import (
	"net/http"

	"github.com/c4pt0r/log"
	"github.com/c4pt0r/stash/datasource"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
)

type GetArgs struct {
	Key []byte
}

type GetReply struct {
	Result datasource.Item
}

type PutArgs struct {
	Key   []byte
	Value []byte
	Meta  datasource.ItemMeta
}

type PutReply struct{}

type StashRPCService struct{}

func (h *StashRPCService) Get(r *http.Request, args *GetArgs, reply *GetReply) error {
	i, err := Provider().Get(args.Key)
	if err != nil {
		return err
	}
	reply.Result = i
	return nil
}

func (h *StashRPCService) Put(r *http.Request, args *PutArgs, reply *PutReply) error {
	return Provider().Put(args.Key, args.Value, args.Meta)
}

func serve(addr string) error {
	log.I("Start StashRPCService, listening on: ", addr)
	s := rpc.NewServer()
	s.RegisterCodec(json.NewCodec(), "application/json")
	s.RegisterService(&StashRPCService{}, "")
	http.Handle("/rpc", s)
	return http.ListenAndServe(addr, nil)
}
