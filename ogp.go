package main

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/PuerkitoBio/goquery"
)

func init() {
	functions.HTTP("GetOGP", getOGP)
}

func getOGP(w http.ResponseWriter, r *http.Request) {
	// 許可するオリジンのリスト
	allowedOrigins := strings.Split(os.Getenv("ALLOWED_ORIGINS"), " ")

	// CORSの設定
	origin := r.Header.Get("Origin")
	if origin == "" || !contains(allowedOrigins, origin) {
		http.Error(w, "forbidden", http.StatusForbidden)
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", origin)

	url := r.URL.Query().Get("url")
	if url == "" {
		http.Error(w, "url parameter is required", http.StatusBadRequest)
		return
	}

	// 指定されたURLのOGP情報を取得する
	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, "failed to get url", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		http.Error(w, "failed to parse html", http.StatusInternalServerError)
		return
	}

	// OGP情報を格納するための構造体
	type Response struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Image       string `json:"image"`
		Url         string `json:"url"`
	}

	// レスポンス用のJSONを作成
	respJSON := Response{
		Title:       doc.Find("meta[property='og:title']").AttrOr("content", ""),
		Description: doc.Find("meta[property='og:description']").AttrOr("content", ""),
		Image:       doc.Find("meta[property='og:image']").AttrOr("content", ""),
		Url:         doc.Find("meta[property='og:url']").AttrOr("content", ""),
	}

	// JSONをエンコードしてレスポンスを送信
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(respJSON)
}

func contains(list []string, target string) bool {
	for _, v := range list {
		if v == target {
			return true
		}
	}
	return false
}
