package handler

import "net/http"

func OrderRegistrationHandler() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		_ = r.Context()
	}
}

func RewardRegistrationHandler() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		_ = r.Context()
	}
}

func AccrualsCalculationHandler() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		_ = r.Context()
	}
}
