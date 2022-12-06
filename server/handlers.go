package server

import (
	"fmt"
	"math/rand"
	"net/http"
	Platillo "servidor/models"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
	"github.com/shettyh/threadpool"
)

// Template ocupado en el proyecto extension HTML "LOCALIZADO EN TEMPLATE"
var templates = template.Must(template.ParseGlob("template/*"))

// Debug de error
func DebugE(err error) {
	if err != nil {
		panic(err.Error())
	}
}

// Funcion que se encarga de editar un registro
func Editaregistro(w http.ResponseWriter, r *http.Request) {
	//Lectura del Id del registro a editar
	idPlatillo := r.URL.Query().Get("id")
	fmt.Println(idPlatillo)
	//Conexion a la base de datos
	conexionEstablecida := conexionBD()
	//Consulta a la base de datos
	Registro, err := conexionEstablecida.Query("SELECT * FROM platillos WHERE id=?", idPlatillo)
	//Validacion de errores
	DebugE(err)
	//Contruccion del objeto platillo
	platillo := Platillo.Platillo{}

	//Validacion del registro
	for Registro.Next() {
		var id, precio int
		var nombre, descripcion string
		err = Registro.Scan(&id, &nombre, &descripcion, &precio)
		//Validacion de errores
		DebugE(err)
		//Asignacion de valores al objeto platillo
		platillo.Id = id
		platillo.Nombre = nombre
		platillo.Descripcion = descripcion
		platillo.Precio = precio
	}
	//fmt.Println(platillo)
	templates.ExecuteTemplate(w, "editar", platillo)
}

func Borrar(w http.ResponseWriter, r *http.Request) {
	//Lectura del Id del registro a editar
	idPlatillo := r.URL.Query().Get("id")
	fmt.Println(idPlatillo)
	//Conexion a la base de datos
	conexionEstablecida := conexionBD()
	//Consulta a la base de datos
	borrarRegistro, err := conexionEstablecida.Prepare("DELETE FROM platillos WHERE id=?")
	//Validacion de errores
	DebugE(err)
	//Ejecucion de la consulta
	borrarRegistro.Exec(idPlatillo)
	//Redireccionamiento a la pagina principal
	http.Redirect(w, r, "/", 301)

}

// Funcion que se encarga de cargar todos los registros de la base de datos
func index(w http.ResponseWriter, r *http.Request) {
	//Conexion a la base de datos
	conexionEstablecida := conexionBD()
	//Consulta a la base de datos
	Registros, err := conexionEstablecida.Query("SELECT * FROM platillos")
	//Validacion de errores
	DebugE(err)
	//Contruccion del objeto platillo
	platillo := Platillo.Platillo{}
	//Arreglo de objetos platillo
	arregloPlatilos := []Platillo.Platillo{}
	//Validacion de registros
	for Registros.Next() {
		var id, precio int
		var nombre, descripcion string
		//Validacion de datos
		err = Registros.Scan(&id, &nombre, &descripcion, &precio)
		//Validacion de errores
		DebugE(err)
		//Asignacion de valores al objeto platillo
		platillo.Id = id
		platillo.Nombre = nombre
		platillo.Descripcion = descripcion
		platillo.Precio = precio
		//Agregacion de objetos platillo al arreglo
		arregloPlatilos = append(arregloPlatilos, platillo)
	}
	//Renderizado de la pagina
	templates.ExecuteTemplate(w, "inicio", arregloPlatilos)
}

// Funcion que se encarga de cargar el formulario
func form(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "forms", nil)
}

// Funcion que se encarga de insertar un registro
func insertar(w http.ResponseWriter, r *http.Request) {
	//Validacion de metodo
	if r.Method == "POST" {
		//Lectura de los datos del formulario
		nombre := r.FormValue("nombre")
		descripcion := r.FormValue("descripcion")
		precio := r.FormValue("precio")
		//Conexion a la base de datos
		conexionEstablecida := conexionBD()
		//Consulta a la base de datos
		insertarRegistros, err := conexionEstablecida.Prepare("INSERT INTO platillos(nombre,descripcion,precio) VALUES(?,?,?)")
		//Validacion de errores
		DebugE(err)
		//Ejecucion de la consulta
		insertarRegistros.Exec(nombre, descripcion, precio)
		//Redireccionamiento a la pagina principal
		http.Redirect(w, r, "/", 301)
	}
}

// Funcion que se encarga de actualizar un registro
func Actualizar(w http.ResponseWriter, r *http.Request) {
	//Validacion de metodo
	if r.Method == "POST" {
		//Lectura de los datos del formulario
		id := r.FormValue("id")
		nombre := r.FormValue("nombre")
		descripcion := r.FormValue("descripcion")
		precio := r.FormValue("precio")
		//Conexion a la base de datos
		conexionEstablecida := conexionBD()
		//Consulta a la base de datos
		modificarRegistro, err := conexionEstablecida.Prepare("UPDATE platillos SET nombre=?, descripcion=?, precio=? WHERE id=?")
		//Validacion de errores
		DebugE(err)
		//Ejecucion de la consulta
		modificarRegistro.Exec(nombre, descripcion, precio, id)
		//Redireccionamiento a la pagina principal
		http.Redirect(w, r, "/", 301)
	}
}

// Funcion que se encarga de cargar el formulario de concurrencia
func Filtro(w http.ResponseWriter, r *http.Request) {
	//Cargar canal de datos con objetos platillo
	data := make(chan []Platillo.Platillo, 10)
	//Lanzamiento de gorutina
	go func() {
		//Conexion a la base de datos
		conexionEstablecida := conexionBD()
		//Consulta a la base de datos
		Registros, err := conexionEstablecida.Query("SELECT * FROM `platillos`")
		//Validacion de errores
		DebugE(err)
		//Contruccion del objeto platillo
		platillo := Platillo.Platillo{}
		//Arreglo de objetos platillo
		arregloPlatilos := []Platillo.Platillo{}
		//Validacion de registros
		for Registros.Next() {
			//Asignacion de valores al objeto platillo
			var id, precio int
			var nombre, descripcion string
			//Validacion de datos
			err = Registros.Scan(&id, &nombre, &descripcion, &precio)
			//Validacion de errores
			DebugE(err)
			//Asignacion de valores al objeto platillo
			platillo.Id = id
			platillo.Nombre = nombre
			platillo.Descripcion = descripcion
			platillo.Precio = precio
			//Agregacion de objetos platillo al arreglo
			arregloPlatilos = append(arregloPlatilos, platillo)
			//Mandar datos al canal
			data <- arregloPlatilos
		}
	}()
	//Llamada a la funcion que se encarga de leer el canal
	Channel(data)
}

// Funcion que se encarga de leer el canal
func Channel(data chan []Platillo.Platillo) {
	for {
		//Extraer datos del canal
		dato := <-data
		//Impresion de datos
		fmt.Println(dato)
	}
}

type Gorutine struct {
	Cantidad int
}

// ruta de index

func Pool(w http.ResponseWriter, r *http.Request) {

	make := Gorutine{}
	pool := threadpool.NewThreadPool(10, 10)
	//go messj(info) // gorutina
	tarea := &Gorutine{}
	err := pool.Execute(tarea)
	DebugE(err)

	make.Cantidad = tarea.Cantidad
	//Renderizado de la pagina
	templates.ExecuteTemplate(w, "pool", make)

}

func (t *Gorutine) Run() {
	var numero int
	var numero2 int
	var precio int
	arrayNombre := []string{"Papas a la francesa", "Helado", "Filetes", "Manzana"}
	arrayDescripcion := []string{"Papas fritas en aceite", "Crema de nieve de vainilla", "Filete de vaca", "Fruta picada"}
	for i := 1; i <= 20; i++ {
		numero = rand.Intn(4)
		numero2 = rand.Intn(4)
		precio = rand.Intn(1000)
		nombre := arrayNombre[numero]
		descripcion := arrayDescripcion[numero2]
		precio := precio
		//Conexion a la base de datos
		conexionEstablecida := conexionBD()
		//Consulta a la base de datos
		insertarRegistros, err := conexionEstablecida.Prepare("INSERT INTO platillos(nombre,descripcion,precio) VALUES(?,?,?)")
		//Validacion de errores
		DebugE(err)
		//Ejecucion de la consulta
		insertarRegistros.Exec(nombre, descripcion, precio)
		fmt.Println("Ejecutar: ", i, " - ", nombre, " - ", descripcion, " - ", precio)
		t.Cantidad = i
	}
}
