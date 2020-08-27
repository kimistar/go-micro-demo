package handler

import (
	"context"
	"encoding/json"
	api "github.com/micro/go-micro/v2/api/proto"
	"github.com/micro/go-micro/v2/errors"
	"go-micro-demo/proto/greeter"
	"net/http"
)

type Say struct {
	Client greeter.HelloService
}

func (s *Say) Hello(ctx context.Context, req *api.Request, rsp *api.Response) error {
	name, ok := req.Get["name"]
	if !ok || len(name.Values) == 0 {
		return errors.BadRequest("go.micro.srv.sayhello", "name cannot be empty")
	}

	resp, err := s.Client.Hello(ctx, &greeter.HelloRequest{
		Name: name.Values[0],
	})
	if err != nil {
		return err
	}

	rsp.StatusCode = http.StatusOK
	b, _ := json.Marshal(map[string]interface{}{
		"name": resp.Msg,
	})
	rsp.Body = string(b)
	rsp.Header = map[string]*api.Pair{
		"x-sign": {
			Key:    "x-sign",
			Values: []string{"sign"},
		},
	}
	return nil
}
