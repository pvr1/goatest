package hello

import (
	"context"
	"log"

	hellosvc "github.com/pvr1/goatest/gen/hellosvc"
)

// hellosvc service example implementation.
// The example methods log the requests and return zero values.
type hellosvcsrvc struct {
	logger *log.Logger
}

// NewHellosvc returns the hellosvc service implementation.
func NewHellosvc(logger *log.Logger) hellosvc.Service {
	return &hellosvcsrvc{logger}
}

// Greet implements greet.
func (s *hellosvcsrvc) Greet(ctx context.Context, p *hellosvc.GreetPayload) (res string, err error) {
	s.logger.Print("hellosvc.greet")
	return "hello " + p.Name + "!", nil
}
