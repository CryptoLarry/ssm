package server

// SetupRoutes configures all the routes for this service
func (s *Server) SetupRoutes() {

	// Base Functions

	s.router.Get("/{txid}", s.SSMGet())
	s.router.Post("/add", s.SSMSave())

}
