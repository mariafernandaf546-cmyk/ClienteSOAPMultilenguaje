package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
)

type Envelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    Body     `xml:"Body"`
}

type Body struct {
	Response NumberToWordsResponse `xml:"NumberToWordsResponse"`
}

type NumberToWordsResponse struct {
	Result string `xml:"NumberToWordsResult"`
}

func convertir(numero int) (string, error) {

	soap := fmt.Sprintf(`<?xml version="1.0" encoding="utf-8"?>
<soap:Envelope
xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
xmlns:xsd="http://www.w3.org/2001/XMLSchema"
xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">

<soap:Body>

<NumberToWords xmlns="http://www.dataaccess.com/webservicesserver/">

<ubiNum>%d</ubiNum>

</NumberToWords>

</soap:Body>

</soap:Envelope>`, numero)

	req, err := http.NewRequest(
		"POST",
		"https://www.dataaccess.com/webservicesserver/NumberConversion.wso",
		bytes.NewBuffer([]byte(soap)),
	)

	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "text/xml;charset=utf-8")

	req.Header.Set(
		"SOAPAction",
		"http://www.dataaccess.com/webservicesserver/NumberToWords",
	)

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var envelope Envelope

	xml.Unmarshal(body, &envelope)

	return envelope.Body.Response.Result, nil

}

func main() {

	var numero int

	fmt.Print("Ingrese un número: ")

	fmt.Scan(&numero)

	respuesta, err := convertir(numero)

	if err != nil {

		fmt.Println(err)

		return

	}

	fmt.Println()

	fmt.Println("Número en inglés:")

	fmt.Println(respuesta)

}