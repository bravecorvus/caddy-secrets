package secrets

import (
	"errors"
	"github.com/mholt/caddy"
	"github.com/mholt/caddy/caddyhttp/httpserver"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/http"
)

func init() {
	caddy.RegisterPlugin("secrets", caddy.Plugin{
		ServerType: "http",
		Action:     Setup,
	})
}

type SecretsHandler struct {
	Next httpserver.Handler
}

var SecretsMap yaml.MapSlice

func Setup(c *caddy.Controller) error {
	if c.Next() {
		c.Next()

		fileName := c.Val()
		if err := readFile(fileName); err != nil {
			return err
		}

		cfg := httpserver.GetConfig(c)
		mid := func(next httpserver.Handler) httpserver.Handler {
			return SecretsHandler{
				Next: next,
			}
		}
		cfg.AddMiddleware(mid)

		if c.Next() {
			return errors.New("Secrets middleware received more arguments than expected")
		}
	}
	return nil
}

func readFile(fileName string) error {
	m := yaml.MapSlice{}
	if content, err := ioutil.ReadFile(fileName); err != nil {
		return err

	} else {
		if err = yaml.Unmarshal([]byte(content), &m); err != nil {
			return err
		}
		SecretsMap = m
	}
	return nil
}

func (h SecretsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) (int, error) {
	return h.Next.ServeHTTP(w, r)
}
