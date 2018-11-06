package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/NoahOrberg/nimvle.nvim/nimvle"
	"github.com/neovim/go-client/nvim"
	"github.com/neovim/go-client/nvim/plugin"
)

func main() {
	plugin.Main(func(p *plugin.Plugin) error {
		p.HandleFunction(&plugin.FunctionOptions{Name: "JekyllCurl"}, func(v *nvim.Nvim, args []string) (rs string, rerr error) {
			nimvle := nimvle.New(v, "jekill.nvim")
			// TODO: check args
			//       read current Buffer
			//       curl using post to server (address defines args[0])
			//       show response scratch buffer
			if len(args) != 1 {
				nimvle.Log("argment's length is not 1")
				return
			}
			currentBuf, err := v.CurrentBuffer()
			if err != nil {
				nimvle.Log(err.Error())
				return
			}
			qlQuery, err := nimvle.GetContentFromBuffer(currentBuf)
			if err != nil {
				nimvle.Log(err.Error())
				return
			}
			postBody, err := json.Marshal(struct {
				Q string `json:"query"`
			}{
				Q: qlQuery,
			})
			if err != nil {
				nimvle.Log(err.Error())
				return
			}
			res, err := http.Post(args[0], "", bytes.NewReader(postBody))
			if err != nil {
				nimvle.Log(err.Error())
				return
			}
			b, err := ioutil.ReadAll(res.Body)
			if err != nil {
				nimvle.Log(err.Error())
				return
			}
			// TODO: show in the scratch buffer
			nimvle.Log(string(b))
			return
		})
		return nil
	})
}
