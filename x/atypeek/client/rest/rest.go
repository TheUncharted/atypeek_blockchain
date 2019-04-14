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

	r.HandleFunc(fmt.Sprintf("/%s/profiles", storeName), addEndorsementHandler(cdc, cliCtx)).Methods("POST")
	r.HandleFunc(fmt.Sprintf("/%s/profiles/{%s}/profile", storeName, restName), profileHandler(cdc, cliCtx, storeName)).Methods("GET")
}

type addProjectReq struct {
	BaseReq         rest.BaseReq `json:"base_req"`
	Id              string       `json:"id"`
	IdProject       string       `json:"idProject"`
	Contributor     string       `json:"contributor"`
	ContributorName string       `json:"contributorName"`
	Receiver        string       `json:"receiver"`
	ReceiverName    string       `json:"receiverName"`
	Duration        string       `json:"duration"`
	Vote            int          `json:"vote"`
	Comments        string       `json:"comments"`
	Skills          string       `json:"skills"`
}

func addEndorsementHandler(cdc *codec.Codec, cliCtx context.CLIContext) http.HandlerFunc {
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

		c, err := sdk.AccAddressFromBech32(req.Contributor)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		rc, err := sdk.AccAddressFromBech32(req.Receiver)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		msg := atypeek.MsgAddEndorsement{
			IdProject:       req.IdProject,
			IdEndorsement:   req.Id,
			Contributor:     c,
			ContributorName: req.ContributorName,
			Receiver:        rc,
			ReceiverName:    req.ReceiverName,
			Duration:        req.Duration,

			Comments: req.Comments,
			Skills:   req.Skills,
			Vote:     req.Vote,
		}

		fmt.Println("*************")

		err = msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		}

		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		clientrest.CompleteAndBroadcastTxREST(w, cliCtx, baseReq, []sdk.Msg{msg}, cdc)

	}
}

func profileHandler(cdc *codec.Codec, cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		paramType := vars[restName]

		res, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/profile/%s", storeName, paramType), nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}

		rest.PostProcessResponse(w, cdc, res, cliCtx.Indent)
	}
}

//curl -XPOST -s http://localhost:1317/atypeek/resumes --data-binary '{"base_req":{"from":"jack","password":"Xenogears42!","chain_id":"namechain","sequence":"1","account_number":"0"},"id":"3","title":"mon projet3","startDate":"20190101", "endDate": "20190102", "owner":"cosmos10x99urhvuckdrpdnf2jvzct83ufpnk2uud0jmc"}'
