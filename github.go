package main

import (
	"errors"
	"fmt"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fastjson"
	"net/http"
)

func getBranchesAndTagsList(token, owner, repo string) (list []string, err error) {
	branches, err := getBranchesList(token, owner, repo, false)
	if err != nil {
		return nil, err
	}
	tags, err := getBranchesList(token, owner, repo, true)
	if err != nil {
		return nil, err
	}
	branches = append(branches, tags...)
	return branches, nil
}

func getBranchesList(token, owner, repo string, tags bool) (branches []string, err error) {
	req := &fasthttp.Request{}
	res := &fasthttp.Response{}
	req.Header.SetMethod(http.MethodGet)
	t := "branches"
	if tags {
		t = "tags"
	}
	req.SetRequestURI(fmt.Sprintf("https://git:%s@api.github.com/repos/%s/%s/%s",
		token, owner, repo, t))
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
