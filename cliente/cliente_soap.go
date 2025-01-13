package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// URL del servicio SOAP
	url := "http://localhost:8080/soap"

	// Cuerpo de la solicitud SOAP en formato XML
	soapRequest := `
	<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/">
	   <soapenv:Body>
	      <HolaMundo>
	         <nombre>Juan</nombre>
	      </HolaMundo>
	   </soapenv:Body>
	</soapenv:Envelope>`

	// Crear una nueva solicitud HTTP POST
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(soapRequest)))
	if err != nil {
		fmt.Println("Error al crear la solicitud:", err)
		return
	}

	// Configurar los encabezados HTTP necesarios
	req.Header.Set("Content-Type", "text/xml")
	req.Header.Set("SOAPAction", "")

	// Enviar la solicitud al servidor
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error al enviar la solicitud:", err)
		return
	}
	defer resp.Body.Close()

	// Leer la respuesta del servidor
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error al leer la respuesta:", err)
		return
	}

	// Mostrar la respuesta
	fmt.Println("Respuesta del servidor SOAP:")
	fmt.Println(string(body))
}
