package middleware

//func Hmac(log *slog.Logger, store *sessions.FilesystemStore) func(next http.Handler) http.Handler {
//	return func(next http.Handler) http.Handler {
//		allowedRoutes := []string{
//			"/api/v0/user/login",
//		}
//
//		log = log.With(slog.String("component", "middleware/hmac"))
//		log.Info("hmac middleware enabled")
//
//		fn := func(w http.ResponseWriter, r *http.Request) {
//			if slices.Contains(allowedRoutes, r.URL.Path) {
//				next.ServeHTTP(w, r)
//				return
//			}
//
//			session, err := sessionSupport.Get("authSave", r, store)
//			if err != nil {
//				log.Error(err.Error())
//
//				render.Status(r, http.StatusInternalServerError)
//				render.JSON(w, r, response.Error("Fail in server.", response.With{
//					"time": time.Now(),
//				}))
//				return
//			}
//
//			if !HMAC.ValidateHMAC(r, session) {
//				err = session.Save(r, w)
//				if err != nil {
//					log.Error(err.Error())
//
//					render.Status(r, http.StatusInternalServerError)
//					render.JSON(w, r, response.Error("Fail in server.", response.With{
//						"time": time.Now(),
//					}))
//					return
//				}
//
//				render.Status(r, http.StatusBadRequest)
//				render.JSON(w, r, response.Error("Fail to Auth.", response.With{
//					"time": time.Now(),
//				}))
//				log.Error("Request is not protect")
//				return
//			}
//
//			err = session.Save(r, w)
//			if err != nil {
//				log.Error(err.Error())
//
//				render.Status(r, http.StatusInternalServerError)
//				render.JSON(w, r, response.Error("Fail in server.", response.With{
//					"time": time.Now(),
//				}))
//				return
//			}
//
//			next.ServeHTTP(w, r)
//		}
//
//		return http.HandlerFunc(fn)
//	}
//}
