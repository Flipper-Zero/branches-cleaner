package main

import (
	"errors"
	"fmt"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fastjson"
	"net/http"
)

func getBranchesList(token, owner, repo string) (branches []string, err error) {
	req := &fasthttp.Request{}
	res := &fasthttp.Response{}
	req.Header.SetMethod(http.MethodGet)
	req.SetRequestURI(fmt.Sprintf("https://git:%s@api.github.com/repos/%s/%s/branches",
		token, owner, repo))
	err = fasthttp.Do(req, res)
	if err != nil {
		return nil, err
	}
	body := res.Body()

	errorMessage := fastjson.GetString(res.Body(), "message")
	if errorMessage != "" {
		return nil, errors.New(errorMessage)
	}

	val, err := fastjson.ParseBytes(body)
	if err != nil {
		return nil, err
	}
	vals, err := val.Array()
	if err != nil {
		return nil, err
	}

	for _, v := range vals {
		branches = append(branches, string(v.GetStringBytes("name")))
	}

	return
}
