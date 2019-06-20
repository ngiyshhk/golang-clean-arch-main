package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"regexp"

	"github.com/ngiyshhk/golang-clean-arch-entity/model"
)

// LibraryRepositoryImpl is
type LibraryRepositoryImpl struct{}

// GetAll is
func (*LibraryRepositoryImpl) GetAll() (*model.Libraries, error) {
	url := "http://api.calil.jp/library?format=json&callback=no"
	req, _ := http.NewRequest("GET", url, nil)

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// apiの最後に);がついてくるので、除去
	rep := regexp.MustCompile(`\);$`)
	clearedBody := rep.ReplaceAll(body, []byte{})

	var res model.Libraries
	err = json.Unmarshal(clearedBody, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
