package rutas

import (
	"Go-Proyect/config"
	"Go-Proyect/modelo"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
	paypal "github.com/plutov/paypal/v4"
)

func retornarOrdenPaypal(token string, id int) string {
	errorVariables := godotenv.Load()
	if errorVariables != nil {

		panic(errorVariables)

	}
	time := strings.Split(time.Now().String(), " ")
	orden := string(time[4][6:14])
	urlRes := "https://n1r629ft-8080.usw3.devtunnels.ms/pasarela/buy/paypal/respuesta"
	urlCanc := "https://n1r629ft-8080.usw3.devtunnels.ms/pasarela/buy/paypal/cancelado"
	datos := fmt.Sprintf(`{
		"purchase_units": [
			{
				"amount": {
					"currency_code": "MXN",
					"value": "10.00"
				},
				"reference_id": "orden_2"
			}
		],
		"intent": "CAPTURE",
		"payment_source": {
			"paypal": {
				"experience_context": {
					"payment_method_preference": "IMMEDIATE_PAYMENT_REQUIRED",
					"payment_method_selected": "PAYPAL",
					"brand_name": "Flowers",
					"locale": "es-ES",
					"landing_page": "LOGIN",
					"shipping_preference": "NO_SHIPPING",
					"user_action": "PAY_NOW",
					"return_url": "%s",
					"cancel_url": "%s"
				}
			}
		}
	}`, urlRes, urlCanc)

	byte_arr := []byte(datos)
	client := &http.Client{}
	req, err := http.NewRequest("POST", os.Getenv("PAYPAL_BASE_URI")+"/v2/checkout/orders", bytes.NewBuffer(byte_arr))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("PayPal-Request-Id", "orden_"+orden)
	if err != nil {

	}
	reg, err2 := client.Do(req)
	defer reg.Body.Close()
	if err2 != nil {

	}
	body, err := ioutil.ReadAll(reg.Body)
	paypal := modelo.PaypalOrderResponseModel{}
	errJson := json.Unmarshal(body, &paypal)
	if errJson != nil {

	}
	return paypal.Id
}
func retornarCapturaPaypal(token string, parametro string) string {
	errorVariables := godotenv.Load()
	if errorVariables != nil {

		panic(errorVariables)

	}
	client := &http.Client{}
	req, err := http.NewRequest("POST", os.Getenv("PAYPAL_BASE_URI")+"/v2/checkout/orders/"+parametro+"/capture", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	if err != nil {
		panic(err)
	}
	reg, err2 := client.Do(req)
	defer reg.Body.Close()

	if err2 != nil {

	}
	defer reg.Body.Close()

	if reg.Status == "422 Unprocessable Entity" {

		return "mal"
	} else {

		return "bien"
	}

}
func retornarTokenPaypal() string {
	c, err := paypal.NewClient(os.Getenv("PAYPAL_CLIENT_ID"), os.Getenv("PAYPAL_CLIENT_SECRET"), paypal.APIBaseSandBox)
	if err != nil {
		return "error"
	}
	c.SetLog(os.Stdout)
	accessToken, err := c.GetAccessToken(context.Background())
	if err != nil {
		return "error"
	}
	return accessToken.Token
}

func Pasarelas_paypal(response http.ResponseWriter, request *http.Request) {
	tipo := request.FormValue("tipo")
	_, id := config.ExitUser(request)
	if tipo == "paypal" {
		token := retornarTokenPaypal()
		orden := retornarOrdenPaypal(token, id)
		response.Write([]byte(orden))
		return
	} else {
		response.Write([]byte("Err"))
		return
	}
}

func Pasarelas_paypal_respuesta(response http.ResponseWriter, request *http.Request) {
	template := template.Must(template.ParseFiles("Vista/templates/paypal_respuesta.html", "Vista/Layout/layout-compras.html"))
	exisid, _ := config.ExitUser(request)
	token := request.URL.Query().Get("token")
	if exisid && token != "" {
		estado_paypal := retornarCapturaPaypal(token, request.URL.Query().Get("token"))
		data := map[string]string{
			//"token":         request.URL.Query().Get("token"),
			"estado_paypal": estado_paypal,
		}
		template.Execute(response, data)
	} else {
		Pagina404(response, request)
	}

}
func Pasarelas_paypal_cancelado(response http.ResponseWriter, request *http.Request) {
	token := request.URL.Query().Get("token")
	//fmt.Println(token)
	if token != "" {
		exisid, _ := config.ExitUser(request)
		if exisid {
			config.CrearMensajesFlash(response, request, "Cancel")
			http.Redirect(response, request, "/", http.StatusSeeOther)
		} else {
			Pagina404(response, request)
		}
	} else {
		Pagina404(response, request)
	}

	/*template := template.Must(template.ParseFiles("Vista/templates/paypal_cancelado.html", "Vista/Layout/layout-compras.html"))

	data := map[string]string{

		"token": request.URL.Query().Get("token"),
	}
	template.Execute(response, data)*/

}
