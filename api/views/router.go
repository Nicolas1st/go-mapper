package views

import "net/http"

func NewViewsRouter(pathToTemplates string) *http.ServeMux {
	viewsResource := newViewsResource(pathToTemplates)

	router := http.NewServeMux()

	router.HandleFunc("/order", viewsResource.makeOrderView)
	router.HandleFunc("/signin", viewsResource.SignInView)
	router.HandleFunc("/signup", viewsResource.SignUpView)

	return router
}
