package server

import (
	"net/http"

	"github.com/halorium/httprouter"
	"github.com/jmckee46/deployer/logger"
	"github.com/jmckee46/deployer/request"
	"github.com/jmckee46/deployer/responses"
	"github.com/jmckee46/deployer/serializers"
	"github.com/jmckee46/deployer/stamps"
	"github.com/jmckee46/deployer/uuid"
)

type RouterState struct {
	ID        string              `json:"id"`
	CreatedAt *stamps.Timestamp   `json:"created-at"`
	Nickname  string              `json:"nickname"`
	Request   *request.Request    `json:"-"`
	Response  *responses.Response `json:"-"`
}

// NewRouterState creates a new router state struct
func NewRouterState() *RouterState {
	state := &RouterState{

		// request
		ID:        uuid.NewRandom(),
		CreatedAt: stamps.New(),
	}

	state.Response = responses.NewResponse()
	state.Response.Serializer = serializers.JSONPrettySerializer

	return state
}

func getStatus(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	state := NewRouterState()
	state.Request = request.NewRequest(w, r, &p)
	state.Response.Writer = w

	logger.Debug("status-request", state.Request.Details())

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	logger.Info("status-response", state)
}
