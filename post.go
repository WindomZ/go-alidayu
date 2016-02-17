package alidayu

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"github.com/WindomZ/go-struct-filler"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strings"
)

func post(arg interface{}) (success bool, response string) {
	if AppKey == "" || AppSecret == "" {
		return false, ERR_NO_APP_KEY.Error()
	}
	body, size := getRequestBody(arg)
	req, _ := http.NewRequest("POST", AppURL, body)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.ContentLength = size
	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		response = err.Error()
		return
	}
	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)
	response = string(data)
	success = strings.Contains(response, "success")
	return
}

func getRequestBody(arg interface{}) (io.Reader, int64) {
	m, err := gsf.StructToStringMap(arg)
	if err != nil {
		return nil, 0
	}
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	v := url.Values{}
	var sign bytes.Buffer
	sign.WriteString(AppSecret)
	for _, k := range keys {
		if len(m[k]) == 0 {
			continue
		}
		v.Set(k, m[k])
		sign.WriteString(k + m[k])
	}
	sign.WriteString(AppSecret)
	v.Set("sign", strings.ToUpper(fmt.Sprintf("%x", md5.Sum(sign.Bytes()))))

	return ioutil.NopCloser(strings.NewReader(v.Encode())), int64(len(v.Encode()))
}
