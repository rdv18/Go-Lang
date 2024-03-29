package main

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
)

func main() {
	cmd := &cobra.Command{
		Use:          "imgdl",
		Short:        "Image Downloader",
		SilenceUsage: true,
	}
	cmd.AddCommand(initialV(), fetchURL())
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func initialV() *cobra.Command {
	return &cobra.Command{
		Use:     "version",
		Short:   "Get the current version number of imgdl",
		Aliases: []string{"v", "V", "Version", "VERSION"},
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Println("imgdl v0.1")
			return nil
		},
	}
}

func fetchURL() *cobra.Command {
	return &cobra.Command{
		Use:     "get",
		Aliases: []string{"down", "download", "grap"},
		Short:   "Download all images from a specific url from the Internet",
		Args:    cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {

			resp, err := http.Get(args[0])
			if err != nil {
				cmd.Println("No Internet Connection!")
			}
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				cmd.Println("no webpage body!!")
			}
			r, _ := regexp.Compile("[a-zA-Z0-9/_.:-]+.(jpg|png)")

			arr := r.FindAllString(string(body), -1)
			for i := 0; i < len(arr); i++ {
				if strings.HasPrefix(arr[i], "//") {
					arr[i] = "https:" + arr[i]
				}
				if strings.HasPrefix(arr[i], "/") {
					arr[i] = args[0] + arr[i]
				}
				fileName := filepath.Base(arr[i])
				if err := imagedownload(fileName, arr[i]); err != nil {
					panic(err)
				}
				cmd.Println("downloaded image : ", arr[i])
			}
			return nil
		},
	}
}

func imagedownload(filepath string, url string) error {

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create("images/" + filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}
