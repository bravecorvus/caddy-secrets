package secrets

import (
	//"bytes"
	"fmt"
	"github.com/mholt/caddy"
	"github.com/mholt/caddy/caddyhttp/httpserver"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/http"
	"os/exec"
)

func init() {
	caddy.RegisterPlugin("secrets", caddy.Plugin{
		ServerType: "http",
		Action:     setup,
	})
}

func setup(c *caddy.Controller) error {
	for c.Next() {
		args := c.RemainingArgs()
		fileName := args[0]
		secretsUser := args[1]
		secretsPassword := args[2]

		cmd := exec.Command("/bin/sh", "-c", fmt.Sprintf("echo %s | sudo -S chown `whoami` %s", secretsPassword, fileName))
		cmd.Run()

		// var out bytes.Buffer
		// var stderr bytes.Buffer
		// cmd.Stdout = &out
		// cmd.Stderr = &stderr
		// err := cmd.Run()
		// if err != nil {
		// 	fmt.Printf("err: %s\n", stderr.String())
		// } else {
		// 	fmt.Println(out.String())
		// }

		m := yaml.MapSlice{}
		content, err := ioutil.ReadFile(fileName)
		if err != nil {
			fmt.Println(err.Error())
		} else {

			err := yaml.Unmarshal([]byte(content), &m)
			if err != nil {
				fmt.Printf("error: %v\n", err.Error())
			} else {
				fmt.Printf("yaml: %v\n", m)
			}
			cmd = exec.Command("/bin/sh", "-c", fmt.Sprintf("sudo chown %s %s", secretsUser, fileName))
			cmd.Run()
		}

		cfg := httpserver.GetConfig(c)
		mid := func(next httpserver.Handler) httpserver.Handler {
			return MyHandler{
				SecretsMap: m,
				Next:       next,
			}
		}
		cfg.AddMiddleware(mid)

		// mids := cfg.Middleware()
		// secrets := mids[0](httpserver.EmptyNext)
		// fmt.Printf("\n%T %v\n", secrets, secrets)
	}

	return nil
}

type MyHandler struct {
	Next       httpserver.Handler
	SecretsMap yaml.MapSlice
}

func (h MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) (int, error) {
	return h.Next.ServeHTTP(w, r)
}
