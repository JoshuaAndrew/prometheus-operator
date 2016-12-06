// Copyright 2016 The prometheus-operator Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package api

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"
)

type API struct {
}

func New() *API {
	return &API{}
}

var (
	prometheusRoute = regexp.MustCompile("/apis/monitoring.coreos.com/v1alpha1/namespaces/.*/prometheuses/.*/status")
)

func (api *API) Register(mux *http.ServeMux) {
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		if prometheusRoute.MatchString(req.URL.Path) {
			api.prometheusStatus(w, req)
		} else {
			w.WriteHeader(404)
		}
	})
}

func (api *API) prometheusStatus(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	var status = struct{}{}

	b, err := json.Marshal(status)
	if err != nil {
		log.Printf("error: %v", err)
		return
	}
	w.Write(b)
}
