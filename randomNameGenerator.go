package main
import (
 	//"bufio"
  	"fmt"
  	//"log"
	"os"
	"io/ioutil"
	"time"
	"math/rand"
)

var sozluk = map[string]map[string][]string {
	"en": {
		"name": {
			"Apricot",
			"Banana",
			"Course",
			"Dock",
			"Elephant",
			"File",
			"Lion",
			"Leopard",
			"Bird",
			"Bear",
			"Monkey",
		},
		"adjective": {
			"Famous",
			"Poor",
			"Fine",
			"Mountain",
			"Grazzy",
			"Crazy",
			"Nervius",
			"Funny",
			"Helpness",
			"Dangerous",
		},
	},
	"tr": {
		"ad": {
			
			"Balık",
			"Çocuk",
			"mimar",
			"Adam",
			"Ev",
			"Bisiklet",
			"Buyruk",
			"Sabah",
			},
		"sifat": {
			"Latif",
			"Çirkef",
			"Fedakar",
			"Mükemmel",
			"Becerikli",
			"Harika",
			"Güzel",
			"Acayip",
			"Garip",
                        "Kötü",
			"Berbat",
		},
	},
}

func rSU(deger int) int { //random sayi uretici
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(deger)
}

func adSecici(dil string) string { // tr ya da en ad seçimi yapar.
	if dil == "tr" {
		return sozluk["tr"]["ad"][rSU(len(sozluk["tr"]["ad"]))]
	}else {
		return sozluk["en"]["name"][rSU(len(sozluk["en"]["name"]))]
	} 
}

func sifatSecici(dil string) string { 
	if dil == "tr" {
		return sozluk["tr"]["sifat"][rSU(len(sozluk["tr"]["sifat"]))]
	}else {
		return sozluk["en"]["adjective"][rSU(len(sozluk["en"]["adjective"]))]
	} 
}

func dilTamlamaUreteci(dil string) string { //random tamlamalar üretir.
	if dil == "tr" {
		return sifatSecici("tr") + " " + adSecici("tr") 
	}else {
		return sifatSecici("en") + " " + adSecici("en")  
	} 
}

func varMi(ara string, dizin []string) bool { //array içi string arar.
	sonuc := false
	for i := 0; i < len(dizin); i++ {
		if ara == dizin[i] {
			sonuc = true
			break
		}
	}
	return sonuc
}

func createDir(path string, dirName string) {
        os.Mkdir(path + "/" + dirName, 0777)
}

func readDir(path string) []string {
	var list []string
    	files, _ := ioutil.ReadDir(path)
    	for _, f := range files {
		list = append(list, f.Name())
        	
    	}
	return list
}	



func tamlama(dil string, pathfile string, sayi int) []string {
	sayac := 0
	cikti := []string{}
	eklenen := []string{}
	cikti = readDir(pathfile)
	tam := dilTamlamaUreteci(dil)
	for i := 0; i < 1000; i++ {
		if varMi(tam,cikti) == false {
			cikti = append(cikti, tam)
			eklenen = append(eklenen, tam)
			createDir(pathfile, tam)
			tam = dilTamlamaUreteci(dil)
			sayac = sayac + 1
		if sayac == sayi {break}
		}
	}
	return eklenen		
}

func main() {
	fmt.Println(tamlama("tr","../tr_deneme/",5))
	//fmt.Println(tamlama("en","../en_deneme/",5))
}
