package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Estructura para la solicitud SOAP
type SoapRequest struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    SoapBody `xml:"Body"`
}

type SoapBody struct {
	XMLName   xml.Name     `xml:"Body"`
	HolaMundo HolaMundoReq `xml:"HolaMundo"`
}

type HolaMundoReq struct {
	XMLName xml.Name `xml:"HolaMundo"`
	Nombre  string   `xml:"nombre"`
}

// Estructura para la respuesta SOAP
type SoapResponse struct {
	XMLName xml.Name `xml:"Envelope"`
	Xmlns   string   `xml:"xmlns,attr"`
	Body    SoapBodyResp
}

type SoapBodyResp struct {
	XMLName   xml.Name      `xml:"Body"`
	HolaMundo HolaMundoResp `xml:"HolaMundoResponse"`
}

type HolaMundoResp struct {
	XMLName xml.Name `xml:"HolaMundoResponse"`
	Saludo  string   `xml:"saludo"`
}

func holaMundoHandler(w http.ResponseWriter, r *http.Request) {
	// Leer el cuerpo de la solicitud SOAP
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error al leer la solicitud", http.StatusBadRequest)
		return
	}

	// Parsear el XML de la solicitud
	var request SoapRequest
	err = xml.Unmarshal(body, &request)
	if err != nil {
		http.Error(w, "Error al parsear XML", http.StatusBadRequest)
		return
	}

	// Construir la respuesta SOAP
	response := SoapResponse{
		Xmlns: "http://schemas.xmlsoap.org/soap/envelope/",
		Body: SoapBodyResp{
			HolaMundo: HolaMundoResp{
				Saludo: fmt.Sprintf("¡Hola, %s! Bienvenido al servicio SOAP en Go.", request.Body.HolaMundo.Nombre),
			},
		},
	}

	// Convertir la respuesta a XML
	output, err := xml.MarshalIndent(response, "", "  ")
	if err != nil {
		http.Error(w, "Error al generar XML", http.StatusInternalServerError)
		return
	}

	// Configurar los encabezados y enviar la respuesta
	w.Header().Set("Content-Type", "text/xml")
	w.WriteHeader(http.StatusOK)
	w.Write(output)
}

func main() {
	http.HandleFunc("/soap", holaMundoHandler)
	fmt.Println("Servidor SOAP ejecutándose en http://localhost:8080/soap")
	http.ListenAndServe(":8080", nil)
}
