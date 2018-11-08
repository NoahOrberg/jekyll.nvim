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

var scratch *nvim.Buffer

func main() {
	plugin.Main(func(p *plugin.Plugin) error {
		p.HandleFunction(&plugin.FunctionOptions{Name: "JekyllCurl"}, func(v *nvim.Nvim, args []string) (rs string, rerr error) {
			// prepare nimvle
			nimvle := nimvle.New(v, "jekill.nvim")
			var err error
			defer func() {
				if err != nil {
					nimvle.Log(err.Error())
				}
			}()

			// check args
			if len(args) != 1 {
				return
			}

			// get query in current buffer
			currentBuf, err := v.CurrentBuffer()
			if err != nil {
				return
			}
			qlQuery, err := nimvle.GetContentFromBuffer(currentBuf)
			if err != nil {
				return
			}

			// request to URL
			postBody, err := json.MarshalIndent(struct {
				Q string `json:"query"`
			}{
				Q: qlQuery,
			}, "", "    ")
			if err != nil {
				return
			}
			res, err := http.Post(args[0], "", bytes.NewReader(postBody))
			if err != nil {
				return
			}

			// prepare scratch buffer
			if scratch == nil {
				scratch, err = nimvle.NewScratchBuffer("*scratch jekyll*", "json")
				if err != nil {
					return
				}
			}

			// set indent
			b, err := ioutil.ReadAll(res.Body)
			if err != nil {
				return
			}
			buf := new(bytes.Buffer)
			err = json.Indent(buf, b, "", "    ")
			if err != nil {
				return
			}

			// write scratch buffer
			err = nimvle.ShowScratchBuffer(*scratch, buf)
			if err != nil {
				return
			}

			return
		})
		return nil
	})
}
