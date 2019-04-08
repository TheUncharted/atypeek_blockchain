package rest

import (
	"fmt"
	"github.com/theuncharted/atypeek_blockchain/x/atypeek"
	"net/http"

	"github.com/cosmos/cosmos-sdk/client/context"
	clientrest "github.com/cosmos/cosmos-sdk/client/rest"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"
)

const (
	restName = "name"
)

// RegisterRoutes - Central function to define routes that get registered by the main application
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router, cdc *codec.Codec, storeName string) {

	r.HandleFunc(fmt.Sprintf("/%s/resumes", storeName), addProjectHandler(cdc, cliCtx)).Methods("POST")
	r.HandleFunc(fmt.Sprintf("/%s/resumes/{%s}/resume", storeName, restName), resumeHandler(cdc, cliCtx, storeName)).Methods("GET")
}

type addProjectReq struct {
	BaseReq     rest.BaseReq `json:"base_req"`
	Id          string       `json:"id"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	StartDate   string       `json:"startDate"`
	EndDate     string       `json:"endDate"`
	Owner       string       `json:"owner"`
}

func addProjectHandler(cdc *codec.Codec, cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req addProjectReq

		if !rest.ReadRESTReq(w, r, cdc, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		addr, err := sdk.AccAddressFromBech32(req.Owner)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		projectInfo := atypeek.ProjectInfo{
			Id:          req.Id,
			CustomerId:  "",
			Title:       req.Title,
			Description: req.Description,
			StartDate:   req.StartDate,
			EndDate:     req.EndDate,
		}
		msg := atypeek.NewMsgAddProject(projectInfo, addr)
		err = msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		clientrest.CompleteAndBroadcastTxREST(w, cliCtx, baseReq, []sdk.Msg{msg}, cdc)
	}
}

func resumeHandler(cdc *codec.Codec, cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		paramType := vars[restName]

		res, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/resume/%s", storeName, paramType), nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}

		rest.PostProcessResponse(w, cdc, res, cliCtx.Indent)
	}
}

//curl -XPOST -s http://localhost:1317/atypeek/resumes --data-binary '{"base_req":{"from":"jack","password":"Xenogears42!","chain_id":"namechain","sequence":"1","account_number":"0"},"id":"3","title":"mon projet3","startDate":"20190101", "endDate": "20190102", "owner":"cosmos10x99urhvuckdrpdnf2jvzct83ufpnk2uud0jmc"}'
