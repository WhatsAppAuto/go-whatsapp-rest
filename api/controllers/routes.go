package controllers

import "github.com/Progete-Dev/go-whatsapp-rest/api/middlewares"

func (s *Server) initializeRoutes() {
	// Home Route
	// s.Router.HandleFunc("/home", s.Home).Methods("GET")
	// s.Router.HandleFunc("/login", s.Login).Methods("GET")
	// s.Router.HandleFunc("/dashboard", middlewares.SetMiddlewareAuthentication(s.DB, s.Dashboard)).Methods("GET")
	s.Router.HandleFunc("/file/{id}", middlewares.SetMiddlewareBearerAuthentication(s.DB, s.FileDownload)).Methods("POST")
	// Login Route
	s.Router.HandleFunc("/login", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST")

	//Users routes
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.CreateUser)).Methods("POST")
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareBearerAuthentication(s.DB, s.GetUsers)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(s.GetUser)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareBearerAuthentication(s.DB, s.UpdateUser))).Methods("PUT")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareBearerAuthentication(s.DB, s.DeleteUser)).Methods("DELETE")

	// Whatsapp Routes
	s.Router.HandleFunc("/send/text", middlewares.SetMiddlewareBearerAuthentication(s.DB, s.WhatsAppSendText)).Methods("POST")
	s.Router.HandleFunc("/send/image", middlewares.SetMiddlewareBearerAuthentication(s.DB, s.WhatsAppSendImage)).Methods("POST")
	s.Router.HandleFunc("/send/video", middlewares.SetMiddlewareBearerAuthentication(s.DB, s.WhatsAppSendVideo)).Methods("POST")
	s.Router.HandleFunc("/send/audio", middlewares.SetMiddlewareBearerAuthentication(s.DB, s.WhatsAppSendAudio)).Methods("POST")
	s.Router.HandleFunc("/send/location", middlewares.SetMiddlewareBearerAuthentication(s.DB, s.WhatsAppSendLocation)).Methods("POST")
	s.Router.HandleFunc("/send/document", middlewares.SetMiddlewareBearerAuthentication(s.DB, s.WhatsAppSendDocument)).Methods("POST")
	s.Router.HandleFunc("/wp/login", middlewares.SetMiddlewareBearerAuthentication(s.DB, s.WhatsAppLogin)).Methods("POST")
	s.Router.HandleFunc("/wp/logout", middlewares.SetMiddlewareBearerAuthentication(s.DB, s.WhatsAppLogout)).Methods("POST")
}
