package nexmo_test

import "github.com/ammario/nexmo"
import "os"

//client gets a nexmo client from the environment
func client() *nexmo.Client {
	c := nexmo.New(os.Getenv("API_KEY"), os.Getenv("API_SECRET"))
	if c.APIKey == "" || c.APISecret == "" {
		panic("API_KEY and API_SECRET env must be provided")
	}
	return c
}
