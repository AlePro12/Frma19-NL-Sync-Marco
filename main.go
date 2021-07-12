package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/go-ole/go-ole"
	_ "github.com/mattn/go-adodb"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

/**
* Parte F19
* Sincronizador de Farmacia Nuevos Laureles con 19 de Abril.
* @Author Alejandro David Sanchez Marcano (FKGG<3) DiosEnTiConfioC.A.
*
**/

type INV00 struct {
	DESCRIP   string  `json:"DESCRIP"`
	CODP      string  `json:"CODP"`
	DESCORTA  string  `json:"DESCORTA"`
	REF       string  `json:"REF"`
	CATEGORIA string  `json:"CATEGORIA"`
	TIPO      string  `json:"TIPO"`
	MODELO    string  `json:"MODELO"`
	MARCA     string  `json:"MARCA"`
	COSTO_ACT float32 `json:"COSTO_ACT"`
	PRECIO1   float32 `json:"PRECIO1"`
	PRECIOF1  float32 `json:"PRECIOF1"`
	PRECIOD1  float32 `json:"PRECIOD1"`
	UTIL1     float32 `json:"UTIL1"`
	EXIST_ACT float64 `json:"EXIST_ACT"`
	FOTO      string  `json:"FOTO"`
}
type Resultados struct {
	TOKEN   string
	EMPNAME string
	SCAPATH string
}

var EMPNAME, TOKEN, SCAPATH = GetConfigs()

func main() {

	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	}
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/Sync-F19", Sync)

	log.Println("Listing for requests at http://localhost:8000/hello")
	log.Fatal(http.ListenAndServe(":6698", nil))

}
func GetConfigs() (string, string, string) {
	content, err := ioutil.ReadFile("Config.json")
	if err != nil {
		log.Fatal(err)
	}
	var resultados Resultados
	json.Unmarshal([]byte(content), &resultados)
	fmt.Println("Empresa: " + resultados.EMPNAME)
	fmt.Println("Token: " + resultados.TOKEN)
	fmt.Println("SCA Path: " + resultados.SCAPATH)
	return resultados.EMPNAME, resultados.TOKEN, resultados.SCAPATH
}
func Sync(w http.ResponseWriter, req *http.Request) {
	if Getparamets(req, "TOKEN") == TOKEN {

		ole.CoInitialize(0)
		defer ole.CoUninitialize()
		//C:\Users\AP\Documents\SandBox\Scala-Professional\Datos
		db, err := sql.Open("adodb", "Provider=VFPOLEDB.1;Data Source="+SCAPATH+";")
		if err != nil {
			fmt.Println("open", err)
			return
		}
		defer db.Close()
		var SQL string
		SQL = ""
		SQL = "SELECT CODP,DESCRIP,DESCORTA,REF,CATEGORIA,TIPO,MODELO,MARCA,COSTO_ACT,PRECIO1,PRECIOF1,PRECIOD1,UTIL1,EXIST_ACT,FOTO FROM inv00.dbf ORDER BY codp asc"

		rows, err := db.Query(SQL)
		if err != nil {
			fmt.Println("select", err)
			return
		}
		defer rows.Close()
		inv00_part := make(map[int]INV00)
		var ContadorSq int
		ContadorSq = 0
		for rows.Next() {
			var DESCRIP string
			var CODP string
			var DESCORTA string
			var REF string
			var CATEGORIA string
			var TIPO string
			var MODELO string
			var MARCA string
			var COSTO_ACT float32
			var PRECIO1 float32
			var PRECIOF1 float32
			var PRECIOD1 float32
			var UTIL1 float32
			var EXIST_ACT float64
			var FOTO string
			err = rows.Scan(&CODP, &DESCRIP, &DESCORTA, &REF, &CATEGORIA, &TIPO, &MODELO, &MARCA, &COSTO_ACT, &PRECIO1, &PRECIOF1, &PRECIOD1, &UTIL1, &EXIST_ACT, &FOTO)
			//  fmt.Println(DESCRIP)

			checkError(err)

			inv00_part[ContadorSq] = INV00{
				CODP:      strings.TrimSpace(CODP),
				DESCRIP:   strings.TrimSpace(DESCRIP),
				DESCORTA:  strings.TrimSpace(DESCORTA),
				REF:       strings.TrimSpace(REF),
				CATEGORIA: strings.TrimSpace(CATEGORIA),
				TIPO:      strings.TrimSpace(TIPO),
				MODELO:    strings.TrimSpace(MODELO),
				MARCA:     strings.TrimSpace(MARCA),
				FOTO:      strings.TrimSpace(FOTO),
				COSTO_ACT: COSTO_ACT,
				PRECIO1:   PRECIO1,
				PRECIOF1:  PRECIOF1,
				PRECIOD1:  PRECIOD1,
				EXIST_ACT: EXIST_ACT, //GetInv01("01",strings.TrimSpace(CODP))
			}

			ContadorSq = ContadorSq + 1
		}
		jsonString, _ := json.Marshal(inv00_part)
		fmt.Println(ContadorSq)

		fmt.Fprintf(w, string(jsonString))
	} else {
		fmt.Fprintf(w, "TOKEN INVALIDO\n")
	}
	//   fmt.Fprintf(w, "INV Search result print  inv00\n")

}
func Getparamets(r *http.Request, requestKey string) (Result string) {
	keys, ok := r.URL.Query()[requestKey]

	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'key' is missing")
		return
	}

	Result = strings.TrimSpace(keys[0])

	return
}
func checkError(err error) {
	if err != nil {
		log.Println(err)
	}
}
func GetInv01(Nivel string, Codp string) float64 {
	ole.CoInitialize(0)
	defer ole.CoUninitialize()
	db, err := sql.Open("adodb", "Provider=VFPOLEDB.1;Data Source="+SCAPATH+";")
	if err != nil {
		fmt.Println("open", err)
		return 0
	}
	defer db.Close()
	var SQL string
	SQL = "SELECT EXIST FROM INV01.dbf WHERE CODA='" + Nivel + "' AND CODP='" + Codp + "' "
	rows, err := db.Query(SQL)
	if err != nil {
		fmt.Println("select", err)
		return 0
	}
	defer rows.Close()
	var ContadorSq int
	var RETURN float64
	ContadorSq = 0
	for rows.Next() {
		var EXIST float64
		err = rows.Scan(&EXIST)
		RETURN = EXIST
		checkError(err)
		ContadorSq = ContadorSq + 1
	}
	return RETURN
}
