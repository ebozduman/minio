/*
 * Minio Cloud Storage, (C) 2018 Minio, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package main

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/minio/minio/cmd/logger"
	"github.com/minio/minio/pkg/configuration"
)

func main() {
	var err error
	var keys []string
	var kv map[string]string
	serverConfig := &configuration.ServerConfig{RWMutex: &sync.RWMutex{}}

	// fmt.Println("******************************")
	// fmt.Println("***** Calling SetHandler *****")
	// fmt.Println("******************************")
	// if err = serverConfig.SetHandler("region", "minio-region"); err != nil {
	// 	logger.FatalIf(err, "Failed to load configuration data: ")
	// }
	// fmt.Println("******************************")
	// fmt.Println("***** Calling GetHandler *****")
	// fmt.Println("*****      region        *****")
	// fmt.Println("******************************")
	// keys = []string{"region"}
	// if kv, err = serverConfig.GetHandler(keys); err != nil {
	// 	logger.FatalIf(err, "Failed to get region and version configuration parameters."+err.Error())
	// }
	// fmt.Printf("kv>\n%+v\n\n", kv)

	fmt.Println("******************************")
	fmt.Println("***** Calling GetHandler *****")
	fmt.Println("***** return full config *****")
	fmt.Println("******************************")
	keys = []string{}
	if kv, err = serverConfig.GetHandler(keys); err != nil {
		logger.FatalIf(err, "Failed to get region and version configuration parameters."+err.Error())
	}
	// fmt.Printf("kv>\n%+v\n\n", kv)
	fmt.Println()
	b, err := json.MarshalIndent(kv, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Print(string(b))
	// fmt.Println("******************************")
	// fmt.Println("***** Calling GetHandler *****")
	// fmt.Println("***** region and version *****")
	// fmt.Println("******************************")
	// keys = []string{"region", "version"}
	// if kv, err = serverConfig.GetHandler(keys); err != nil {
	// 	logger.FatalIf(err, "Failed to get region and version configuration parameters."+err.Error())
	// }
	// fmt.Printf("kv>\n%+v\n\n", kv)
}
