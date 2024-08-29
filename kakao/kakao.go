package kakao

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func DoLogin(w http.ResponseWriter, r *http.Request) {

	clientID := "60e0b38a9ee4661ea4323860646e86fe"
	redirectURI := "http://127.0.0.1:5174/kakaoRedirect" // 사바
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

	http.Redirect(w, r, baseURL.String(), http.StatusFound)
}

func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, "Authorization code not found", http.StatusBadRequest)
		panic("Authorization code not found")
	}

	// 토큰 요청
	token, err := getAccessToken(code)
	if err != nil {
		http.Error(w, "Failed to get access token", http.StatusInternalServerError)
		panic("Failed to get access token")
	}

	fmt.Println("AccessToken: ", token.AccessToken)
	// 액세스 토큰과 리프레시 토큰을 쿼리 파라미터로 설정
	redirectURL := fmt.Sprintf("http://localhost:5173/success?access_token=%s&refresh_token=%s",
		url.QueryEscape("myAccessToken"),
		url.QueryEscape("myRefreshToken"))

	http.Redirect(w, r, redirectURL, http.StatusFound)
}

func getAccessToken(code string) (*TokenResponse, error) {
	// 엑세스 토큰 요청을 위한 데이터 준비
	data := url.Values{}
	data.Set("grant_type", "authorization_code")
	data.Set("client_id", "60e0b38a9ee4661ea4323860646e86fe")
	data.Set("redirect_uri", "http://127.0.0.1:5173/kakaoRedirect")
	data.Set("code", code)

	req, err := http.NewRequest("POST", "https://kauth.kakao.com/oauth/token", strings.NewReader(data.Encode()))
	if err != nil {
		panic(err)
		return nil, err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
		return nil, err
	}

	defer resp.Body.Close()

	var tokenResponse TokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&tokenResponse); err != nil {
		panic(err)
		return nil, err
	}

	return &tokenResponse, nil
}

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
	Scope        string `json:"scope"`
}
