package kakao

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func DoLogin() {
	clientID := ""
	redirectURI := ""
	responseType := "code"

	// 요청할 URL 생성
	baseURL, err := url.Parse("https://kauth.kakao.com/oauth/authorize")
	if err != nil {
		panic(err)
		return
	}

	params := url.Values{}
	params.Add("client_id", clientID)
	params.Add("redirect_uri", redirectURI)
	params.Add("response_type", responseType)
	baseURL.RawQuery = params.Encode()

	resp, err := http.Get(baseURL.String())
	if err != nil {
		panic(err)
		return
	}

	defer resp.Body.Close()

	// 응답 본문 읽기
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	// 응답 본문 출력
	fmt.Println("Response Body:")
	fmt.Println(string(body))

}
