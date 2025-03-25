package adapter

// import (
// 	"anonsurvey/internal/usecase"
// 	"anonsurvey/types"
// 	"log"
// 	"net/http"
//
// 	"github.com/mholt/binding"
// )

// type UserController struct {
// 	svc usecase.User
// }
//
// func NewUserController() UserController {
// 	return UserController{}
// }
//
// func (c UserController) Routes() []types.HTTPRoute {
// 	rs := make([]types.HTTPRoute, 0)
// 	return rs
// }
//
// func (c UserController) AddUser(w http.ResponseWriter, r *http.Request) {
// 	t := new(types.ReqLogin)
// 	if errs := binding.Bind(r, t); errs != nil {
// 		// FIXME: remove fatal with response
// 		log.Fatalln("cannot bind request", errs)
// 	}
// }
