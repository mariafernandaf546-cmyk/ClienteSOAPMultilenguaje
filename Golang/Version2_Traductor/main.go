package main

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Envelope struct {
	Body Body `xml:"Body"`
}

type Body struct {
	Response NumberToWordsResponse `xml:"NumberToWordsResponse"`
}

type NumberToWordsResponse struct {
	Result string `xml:"NumberToWordsResult"`
}

func convertirNumero(numero int) (string, error) {

	soap := fmt.Sprintf(`<?xml version="1.0" encoding="utf-8"?>
<soap:Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
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

	req.Header.Set("Content-Type", "text/xml; charset=utf-8")
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

	var env Envelope

	xml.Unmarshal(body, &env)

	return env.Body.Response.Result, nil
}

func traducir(texto string) (string, error) {

	urlGoogle :=
		"https://translate.googleapis.com/translate_a/single?client=gtx&sl=en&tl=es&dt=t&q=" +
			url.QueryEscape(texto)

	resp, err := http.Get(urlGoogle)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var datos interface{}

	err = json.Unmarshal(body, &datos)

	if err != nil {
		return "", err
	}

	arr := datos.([]interface{})
	a0 := arr[0].([]interface{})
	a1 := a0[0].([]interface{})

	return a1[0].(string), nil
}

func main() {

	var numero int

	fmt.Print("Ingrese un número: ")

	fmt.Scan(&numero)

	ingles, err := convertirNumero(numero)

	if err != nil {
		panic(err)
	}

	fmt.Println()
	fmt.Println("Número en inglés:")
	fmt.Println(ingles)

	espanol, err := traducir(ingles)

	if err != nil {
		panic(err)
	}

	fmt.Println()
	fmt.Println("Número en español:")
	fmt.Println(espanol)

}