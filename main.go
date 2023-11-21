package main

import (
	controlador "Go-Proyect/Controlador"
	"Go-Proyect/rutas"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {

	mux := mux.NewRouter()

	mux.HandleFunc("/", rutas.Index)
	mux.HandleFunc("/productos", rutas.Prductos)
	mux.HandleFunc("/servicios", rutas.Servicio)
	mux.HandleFunc("/login", rutas.Login)
	mux.HandleFunc("/login/new", rutas.NewRegistro).Methods("POST")
	mux.HandleFunc("/login/validar", rutas.ValidarUsu).Methods("POST")
	mux.HandleFunc("/productos/planta/{id:.*}/{idcat:.*}", rutas.Produc_descripcion)
	mux.HandleFunc("/productos/card", rutas.Card)
	mux.HandleFunc("/operaciones/card/produc", rutas.Operaciones_card).Methods("POST")
	mux.HandleFunc("/logout", rutas.Logout)
	mux.HandleFunc("/productos/card/buy", rutas.Buy)
	mux.HandleFunc("/productos/card/direccion/{id:.*}", rutas.Direccion)
	mux.HandleFunc("/add-upd/direccion/{metodo:.*}", rutas.Add_Direccion).Methods("POST")
	mux.HandleFunc("/pasarela/buy/paypal", rutas.Pasarelas_paypal).Methods("POST")
	mux.HandleFunc("/pasarela/buy/paypal/respuesta", rutas.Pasarelas_paypal_respuesta)
	mux.HandleFunc("/pasarela/buy/paypal/cancelado", rutas.Pasarelas_paypal_cancelado)
	mux.HandleFunc("/send/mail", rutas.Sen_Email).Methods("POST")
	mux.HandleFunc("/productos/card/direccion-cambiar/{id:.*}", rutas.Confirmar_Direccion)
	mux.HandleFunc("/productos/card/direccion-delete/{id:.*}", rutas.Delete_Direccion)

	//archivos estáticos hacia mux
	s := http.StripPrefix("/Vista/", http.FileServer(http.Dir("./Vista/")))
	mux.PathPrefix("/Vista/").Handler(s)

	//error 404
	mux.NotFoundHandler = mux.NewRoute().HandlerFunc(rutas.Pagina404).GetHandler()

	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Println("Desde el celular con esta Ip : http://" + controlador.GetWifiIPv4() + ":8080/")
	fmt.Println("corriendo servidor desde :8080/")
	log.Fatal(server.ListenAndServe())
}
