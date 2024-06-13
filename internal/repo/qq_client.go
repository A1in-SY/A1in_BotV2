package repo

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/go-kratos/kratos/v2/transport/http"
	nethttp "net/http"
	"qqbot/internal/model"

	"time"
)

const (
	QQClientAPI = "http://127.0.0.1:3000/"
)

type QQClient struct {
	client *http.Client
}

func NewQQClient() *QQClient {
	qqClient, _ := http.NewClient(context.Background(), http.WithTimeout(1*time.Second))
	return &QQClient{client: qqClient}
}

func (c *QQClient) Call(ctx context.Context, endpoint string, req interface{}, resp interface{}) (err error) {
	data, err := json.Marshal(req)
	if err != nil {
		return
	}
	hReq, _ := nethttp.NewRequestWithContext(ctx, nethttp.MethodPost, QQClientAPI+endpoint, bytes.NewReader(data))
	hResp, err := c.client.Do(hReq)
	if err != nil {
		return
	}
	defer hResp.Body.Close()
	var buffer bytes.Buffer
	_, err = buffer.ReadFrom(hResp.Body)
	if err != nil {
		return
	}
	qResp := &model.QQCallResp{
		Data: resp,
	}
	err = json.Unmarshal(buffer.Bytes(), qResp)
	if err != nil {
		return
	}
	if qResp.Status != model.QQCallRespStatusOK {
		err = errors.New("qq call resp status not ok")
	}
	return
}
