package handlers

import "github.com/jwambugu/mpesa-golang-sdk"

func RegisterC2BURLs(mpesaApp *mpesa.App, shortCode, validationURL, confirmationURL string) error {
	req := mpesa.C2
}