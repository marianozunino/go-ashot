/*
Copyright © 2024 Mariano Zunino <marianoz@posteo.net>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	app "github.com/marianozunino/goashot/internal/app/web"
	"github.com/spf13/cobra"
)

var port string
var bindAddress string

// backofficeCmd represents the backoffice command
var backofficeCmd = &cobra.Command{
	Use:   "backoffice",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		app.NewBackofficeApp(
			app.WithPort(port),
			app.WithHost(bindAddress),
		).Run()
	},
}

func init() {
	rootCmd.AddCommand(backofficeCmd)
	backofficeCmd.Flags().StringVarP(&port, "port", "p", "5000", "backoffice port")
	backofficeCmd.Flags().StringVarP(&bindAddress, "address", "a", "localhost", "backoffice bind address")
}