package controllers

import (
	"net/http"
)

// ハンドラー
func top(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, "Hello", "layout", "top")
}
