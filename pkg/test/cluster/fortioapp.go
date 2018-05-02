//  Copyright 2018 Istio Authors
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package cluster

type fortioapp struct {
	name      string
	namespace string
	serverAddress string
	endpoints []*endpoint
}

// NewFortioApp creates a new fortioapp object from the given service config
func NewFortioApp(meta model.ConfigMeta, cfg model.Service, serverAddress string) test.DeployedFortioApp {
	a := &fortioapp{}
	a.name = meta.Name
	a.namespace = meta.Namespace
	a.serverAddress = serverAddress

	for _, port := range cfg.Ports {
		a.endpoints = append(a.endpoints, &endpoint{s
			port:  port,
		})
	}

	return a
}

// CallFortio implements the test.DeployedApp interface
func (a *fortioapp) CallFortio(cmd string) (test.AppCallFortioResponse, error) {
	
	return nil, nil
}