package controllers

import (
	"net/http"
	"log"
	"errors"
)

func NewMux() http.Handler {
	h := http.NewServeMux()
	h.Handle("/api/v1/department", loggerMid(http.HandlerFunc(department)))
	h.Handle("/api/v1/brand", loggerMid(http.HandlerFunc(brand)))
	h.Handle("/api/v1/invoice", loggerMid(http.HandlerFunc(invoice)))
	return h
}

func department(resp http.ResponseWriter, req *http.Request) {
	var err error
	switch req.Method {
	case "GET":
		//Get one or morw deps
	case "POST":
		//create neq deps
	case "UPDATE":
		//update dep
	case "DELETE":
		//delete dep
	default:
		err = errors.New("Method Not Allow")
	}
	if err != nil {
		http.Error(resp, err.Error(), http.StatusInternalServerError)
		return
	}

	return
}

func brand(resp http.ResponseWriter, req *http.Request) {
	var err error
	switch req.Method {
	case "GET":
	//Get one or morw brand
	case "POST":
	//create neq brand
	case "UPDATE":
	//update brand
	case "DELETE":
	//delete brand
	default:
		err = errors.New("Method Not Allow")
	}
	if err != nil {
		http.Error(resp, err.Error(), http.StatusInternalServerError)
		return
	}

	return
}

func invoice(resp http.ResponseWriter, req *http.Request) {
	var err error
	switch req.Method {
	case "GET":
	//Get one or morw invoice
	case "POST":
	//create neq invoice
	case "UPDATE":
	//update invoice
	case "DELETE":
	//delete invoice
	default:
		err = errors.New("Method Not Allow")
	}
	if err != nil {
		http.Error(resp, err.Error(), http.StatusInternalServerError)
		return
	}

	return
}

func loggerMid(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var clIP string
		if r.Header.Get("X-Forwarded-For") == "" {
			clIP = r.RemoteAddr
		} else {
			clIP = r.Header.Get("X-Forwarded-For")
		}

		uAgent := r.Header.Get("User-Agent")
		log.Printf("\"Method\": \"%s\", \"User-Agent\": \"%s\", \"URL\": \"%s\", \"Host\": \"[%s]\", \"Client-IP\": \"%v\"", r.Method, uAgent, r.URL, r.Host, clIP)
		next.ServeHTTP(w, r)
	})
}

func authMid(next http.Handler) http.Handler {
	return http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {

	})
}