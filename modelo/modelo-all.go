package modelo

//Usuarios
type User struct {
	ID            int
	Nombre        string
	Correo        string
	Edad          string
	Contra        string
	FechaCreacion string
}

//Productos
type Product struct {
	ID          int
	ID_Cat      int
	Nombre      string
	Descripcion string
	Precio      float64
	Cantidad    []int
	Color       string
	Isnew       int
	Img1        string
	Img2        string
	Img3        string
}

//Direccion
type FormData struct {
	ID         int
	FirstName  string
	LastName   string
	Email      string
	AddressOne string
	AddressTwo string
	City       string
	State      string
	Zip        string
	Country    string
	Telefono   string
	Check      int
}

//Carrito
type Produc_Card struct {
	ID_card       int
	Nombre_card   string
	Precio_card   float64
	Cantidad_card int
	Img_card      string
}

//paypal
type PaypalOrderResponseModel struct {
	Id     string
	Status string
	Links  string
}

//categorias
type Categoria struct {
	Flores              string
	Plantas_de_Interior string
	Plantas_de_Sombra   string
	Plantas_de_Exterior string
	Acuaticas           string
}
