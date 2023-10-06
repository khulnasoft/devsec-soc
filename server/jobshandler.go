// Copyright 2019 Jason Ertel (jertel). All rights reserved.
// Copyright 2020-2023 Security Onion Solutions, LLC. All rights reserved.
//
// This program is distributed under the terms of version 2 of the
// GNU General Public License.  See LICENSE for further details.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.

package server

import (
  "context"
  "errors"
  "github.com/khulnasoft/devsec-soc/json"
  "github.com/khulnasoft/devsec-soc/web"
  "net/http"
)

type JobsHandler struct {
  web.BaseHandler
  server *Server
}

func NewJobsHandler(srv *Server) *JobsHandler {
  handler := &JobsHandler{}
  handler.Host = srv.Host
  handler.server = srv
  handler.Impl = handler
  return handler
}

func (jobsHandler *JobsHandler) HandleNow(ctx context.Context, writer http.ResponseWriter, request *http.Request) (int, interface{}, error) {
  switch request.Method {
  case http.MethodGet:
    return jobsHandler.get(ctx, writer, request)
  }
  return http.StatusMethodNotAllowed, nil, errors.New("Method not supported")
}

func (jobsHandler *JobsHandler) get(ctx context.Context, writer http.ResponseWriter, request *http.Request) (int, interface{}, error) {
  kind := request.URL.Query().Get("kind")
  paramsStr := request.URL.Query().Get("parameters")
  var params map[string]interface{}
  if paramsStr != "" {
    err := json.LoadJson([]byte(paramsStr), &params)
    if err != nil {
      return http.StatusBadRequest, nil, err
    }
  }
  return http.StatusOK, jobsHandler.server.Datastore.GetJobs(ctx, kind, params), nil
}
