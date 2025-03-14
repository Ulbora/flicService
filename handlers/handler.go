package handlers

/*
 Copyright (C) 2020 Ulbora Labs LLC. (www.ulboralabs.com)
 All rights reserved.
 Copyright (C) 2020 Ken Williamson
 All rights reserved.
 This program is free software: you can redistribute it and/or modify
 it under the terms of the GNU General Public License as published by
 the Free Software Foundation, either version 3 of the License, or
 (at your option) any later version.
 This program is distributed in the hope that it will be useful,
 but WITHOUT ANY WARRANTY; without even the implied warranty of
 MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 GNU General Public License for more details.
 You should have received a copy of the GNU General Public License
 along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

import (
	"net/http"
)

//Handlers Handlers
type Handlers interface {
	FindFlicListByZip(w http.ResponseWriter, r *http.Request)
	FindFlicByKey(w http.ResponseWriter, r *http.Request)
	SetFlicTable(w http.ResponseWriter, r *http.Request)
	SetLogLevel(w http.ResponseWriter, r *http.Request)
}

// //SetupResponse SetupResponse
// func SetupResponse(w *http.ResponseWriter, r *http.Request) {
// 	(*w).Header().Set("Access-Control-Allow-Origin", "*")
// 	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
// 	(*w).Header().Set("Access-Control-Allow-Headers",
// 		"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, X-Requested-With, Authorization, customer-key")
// 	fmt.Println("Access-Control-Allow-Origin: ", (*w).Header().Get("Access-Control-Allow-Origin"))
// 	fmt.Println("Access-Control-Allow-Methods: ", (*w).Header().Get("Access-Control-Allow-Methods"))
// 	//(*w).WriteHeader(http.StatusOK)
// 	if (r).Method == "OPTIONS" {
// 		(*w).WriteHeader(http.StatusOK)
// 		//h.Log.Debug("In preflight: ", w.WriteHeader)

// 	}
// }
