package main

import (
	"bufio"
	"fmt"
	"html/template"
	"image/png"
	"log"
	"net/http"
	"os"

	"strconv"
	"strings"
)



//RootHandler will take files index.html
func RootHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("index.html")

	if err != nil {
		fmt.Println("Index Template Parse Error: ", err)
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		fmt.Println("Index Template Execution Error: ", err)
	}
}

//ASCIIHandler will take the argurment.
func ASCIIHandler(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		fmt.Println("Index Template Parse Error: ", err)
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		fmt.Println("Index Template Execution Error: ", err)
	}

	str := r.FormValue("StringToAscii")

	myFont := r.FormValue("files_name")

	okay := r.FormValue("is_okay")

	text := r.FormValue("choose")

	writeInFile := r.FormValue("NamedFile")

	ASCIIPrint(w, str, myFont, okay, text, writeInFile)

}

//ImageHandler will take the argurment for image
func ImageHandler(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		fmt.Println("Index Template Parse Error: ", err)
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		fmt.Println("Index Template Execution Error: ", err)
	}

	pngImage := r.FormValue("image")

	DrawASCII(w, pngImage)

	

}

func main() {

	fmt.Println("\033[33mLancement du programme...\033[0m")
	fmt.Println("\033[31mCtrl + C\033[0m\033[33m. Pour arrêter le programme.\033[0m")
	style := http.FileServer(http.Dir("css"))
	

	http.HandleFunc("/", RootHandler) // sets router
	http.Handle("/css/", http.StripPrefix("/css/", style))

	http.HandleFunc("/ascii-art", ASCIIHandler)
	http.HandleFunc("/ascii-image", ImageHandler)
	err := http.ListenAndServe(":100", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}

//DrawASCII will draw the image that we choose
func DrawASCII(w http.ResponseWriter, v string) {

	base := "0ND8OZ$7I?+=~:,.."
	f, _ := os.Open(v)
	img, _ := png.Decode(f)
	bounds := img.Bounds()
	ascii := ""
	fmt.Fprintln(w, "<div class=Div2>")
	for y := 0; y < bounds.Dy(); y += 2 {
		for x := 0; x < bounds.Dx(); x++ {
			pixel := img.At(x, y)
			r, g, b, _ := pixel.RGBA()
			r = r & 0xFF
			g = g & 0xFF
			b = b & 0xFF
			gray := 0.299*float64(r) + 0.578*float64(g) + 0.114*float64(b)
			temp := fmt.Sprintf("%.0f", gray*float64(len(base)+1)/255)
			index, _ := strconv.Atoi(temp)
			if index >= len(base) {
				ascii += "."
				ascii += ""

			} else {
				ascii += string(base[index])

			}

		}

		ascii += "<br>\n"

	}
	f.Close()

	fmt.Fprintln(w, ascii)
	fmt.Fprintln(w, "</div>")
}

//ASCIIPrint is our way to print the ASCII-art on the site
func ASCIIPrint(w http.ResponseWriter, str string, myFont string, isOkay string, textorPDF string, writeInFile string) {

	file, _ := os.Open(myFont)
	fileVal := scanFile(file)
	k := ""
	erreur := 0

	insideFile := ""

	narg := strings.Split(str, "\\n")

	fmt.Fprintln(w, "<div class=Div3>")

	if isOkay == "false" {

		for _, strArg := range narg {

			for i := 1; i <= 8; i++ {

				for _, arg := range strArg {

					indexBase := int(rune(arg)-32) * 9
					if indexBase < 0 || indexBase > 856 {

						erreur = 1
						break

					} else if indexBase < 856 {

						k += fileVal[indexBase+i]

						k = strings.Replace(k, " ", "&nbsp;", -1)
					}

				}
				k += "<br>"

			}
		}

	}

	if isOkay == "true" {

		inFile := "css/"
		if textorPDF == "txt" {
			writeInFile += ".txt"
		}

		if writeInFile == "standard.txt" || writeInFile == "shadow.txt" || writeInFile == "thinkertoy.txt" {

			k = "You can't write over this file"
			fmt.Fprintln(w, k)
			//return k
		} else {

			inFile += writeInFile

			file, err := os.Create(inFile)
			if err != nil {

				fmt.Println("Error")
				os.Exit(0)
			}
			defer file.Close()

			if strings.Contains(str, "\\n") {

				narg2 := strings.Split(str, "\\n")

				for _, strPut := range narg2 {

					for i := 1; i <= 8; i++ {

						for _, arg := range strPut {

							indexBase := int(rune(arg)-32) * 9

							if indexBase < 0 || indexBase > 856 {

								erreur = 1
								break

							} else if indexBase < 856 {

								k += fileVal[indexBase+i]

								k = strings.Replace(k, " ", "&nbsp;", -2)

								insideFile += fileVal[indexBase+i]

							}

						}
						k += "<br>"
						insideFile += "\n"

					}
				}
			} else {

				for i := 1; i <= 8; i++ {

					for _, arg := range str {

						indexBase := int(rune(arg)-32) * 9

						if indexBase < 0 || indexBase > 856 {

							erreur = 1
							break

						} else if indexBase < 856 {

							k += fileVal[indexBase+i]

							k = strings.Replace(k, " ", "&nbsp;", -2)

							insideFile += fileVal[indexBase+i]

						}

					}
					k += "<br>"
					insideFile += "\n"

				}
			}

			fmt.Fprintln(file, insideFile)
			k += ""
			k += "<br>"
			k += "<a class=download href='/css/"
			k += writeInFile
			k += "/'download='"
			k += writeInFile
			k += "'><button> Download your "
			k += writeInFile
			k += " here</button></a>"
			insideFile += ""

		}

	}

	if erreur == 0 {

		fmt.Fprint(w, k)
		fmt.Fprintln(w, "</div>")

		if isOkay == "true" {
			fmt.Print("\033[33mUn fichier\033[0m\033[31m ")
			fmt.Print(writeInFile)
			fmt.Println(" \033[33mà été ajouté\033[0m")
		}
	}

	if erreur == 1 {
		fmt.Fprintln(w, "Index of range.<br>If you want to make a backslash n you have to write Word1\\nWord2")
		fmt.Fprintln(w, "<br>You cannot also write something that outside of the original ASCII table.<br>You can only take the decimal 32 to 126")
	}

}

//scanFile will will a create slice of the file
func scanFile(arg *os.File) []string {

	var fileValue []string

	scanner := bufio.NewScanner(arg)

	for scanner.Scan() {

		lines := scanner.Text()

		fileValue = append(fileValue, lines)

	}

	return fileValue
}
