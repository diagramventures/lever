package lever

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type LeverAPI struct {
	APIKey     string
	HttpClient *http.Client
	BaseURL    string
}

func New(apiKey string) *LeverAPI {
	return &LeverAPI{
		APIKey:     apiKey,
		HttpClient: http.DefaultClient,
		BaseURL:    `https://api.lever.co/v1`,
	}
}

// ListCandidates returns all candidates ever from Lever.
func (api *LeverAPI) ListCandidates() (out []*Candidate, err error) {
	params := P{}
	for {
		var resp *listCandidatesResp
		err = api.call("GET", "candidates", params, nil, &resp)
		if err != nil {
			return
		}

		out = append(out, resp.Data...)
		if !resp.HasNext {
			return
		}
		params["offset"] = resp.Next
	}
}

func (api *LeverAPI) call(method string, endpoint string, params P, body interface{}, out interface{}) error {

	req, err := http.NewRequest(method, fmt.Sprintf("%s/%s?%s", api.BaseURL, endpoint, qsEnc(params)), bodyEnc(body))
	if err != nil {
		return err
	}

	req.SetBasicAuth(api.APIKey, "")

	resp, err := api.HttpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var cnt bytes.Buffer
	_, err = io.Copy(&cnt, resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode > 299 {
		return fmt.Errorf("status code = %d, body=%q", resp.StatusCode, cnt.String())
	}

	if err := json.Unmarshal(cnt.Bytes(), &out); err != nil {
		return err
	}

	return nil
}

type M map[string]interface{}
type P map[string]string

func bodyEnc(v interface{}) io.Reader {
	if v == nil {
		return nil
	}

	cnt, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}

	return bytes.NewReader(cnt)
}

func qsEnc(params P) (out string) {
	if params == nil {
		return
	}
	vals := url.Values{}
	for k, v := range params {
		vals.Set(k, v)
	}
	out = vals.Encode()
	return
}
