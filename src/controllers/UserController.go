package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"stori/src/setup"
	"stori/src/types"
	"stori/src/utils"
)

func UserController(res http.ResponseWriter, req *http.Request) {

	if req.Method != http.MethodPost {
		res.WriteHeader(http.StatusMethodNotAllowed)
		res.Write([]byte("Method Not Allowed"))
	} else {

		var auxUser types.TClient

		json.NewDecoder(req.Body).Decode(&auxUser)

		// READING CSV FILE
		data := setup.CSV(auxUser.Id)

		// START TO MAKE BANK OPERATIONS
		client := utils.Summary(data)

		res.WriteHeader(http.StatusOK)
		res.Header().Set("Content-Type", "text/html; charset=utf-8")
		client.Id = auxUser.Id
		client.Email = auxUser.Email
		client.Name = auxUser.Name

		fmt.Fprint(res, utils.SendMail(client))
	}
}
