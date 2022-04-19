package mypackage

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
)

func extraerContenido(archivo string) (string, []byte) {
	datosComoBytes, err := ioutil.ReadFile(archivo)
	if err != nil {
		log.Fatal(err)
	}
	// convertir el arreglo a string
	datosComoString := string(datosComoBytes)
	// imprimir el string
	return datosComoString, datosComoBytes
}

func isDigit(d byte) bool {
	if d == 48 || d == 49 || d == 50 || d == 51 || d == 52 || d == 53 || d == 54 || d == 55 || d == 56 || d == 57 {
		return true
	} else {
		return false
	}
}

func isLetter(l byte) bool {
	if l == 97 || l == 98 || l == 99 || l == 100 || l == 101 || l == 102 || l == 103 || l == 104 || l == 105 || l == 106 || l == 107 || l == 108 || l == 109 || l == 110 || l == 111 || l == 112 || l == 113 || l == 114 || l == 115 || l == 116 || l == 117 || l == 118 || l == 119 || l == 120 || l == 121 || l == 122 {
		return true
	} else {
		return false
	}
}

func Analizador(contenido []byte) (string, [][]string) {
	// contadores de operadores y signos
	var nMas, nMenos, nAsterisco, nBarraI, nPorCiento, nIgual, nDobleIgual, nMenorQue, nMayorQue, nMayorIgual, nMenorIgual, nAbreParentesis, nCierraParentesis, nAbreLlaves, nCierraLlaves, nComillas, nPuntoComa int
	// contadores de palabras reservadas
	var enteroPR, decimalPR, booleanoPR, cadenaPR, siPR, sinoPR, mientrasPR, hacerPR, verdaderoPR, falsoPR int
	// contador de lineas y sensor de cambio de línea
	linea, cambio := 1, 0
	numerosEncontrados, identificadoresEncontrados := [][]string{}, [][]string{}
	agregar := []string{}
	errores := ""

	for i := 0; i < len(contenido); i++ {
		// comprobar si es un número
		if isDigit(contenido[i]) && isDigit(contenido[i+1]) {
			j := i
			var numeros string
			for isDigit(contenido[j]) {
				numeros += string(contenido[j])
				j++
			}
			i = j - 1

			nRepetido := false
			for k := 0; k < len(numerosEncontrados); k++ {
				if numerosEncontrados[k][0] == numeros {
					reemplazo, _ := strconv.Atoi(numerosEncontrados[k][2])
					numerosEncontrados[k][2] = strconv.Itoa(reemplazo + 1)
					nRepetido = true
				}
			}
			if nRepetido == false {
				agregar = []string{numeros, "NUMEROS", "1"}
				numerosEncontrados = append(numerosEncontrados, agregar)
			}
		} else if contenido[i] == 101 && contenido[i+1] == 110 && contenido[i+2] == 116 && contenido[i+3] == 101 && contenido[i+4] == 114 && contenido[i+5] == 111 {
			enteroPR++
			i += 5
			continue
		} else if contenido[i] == 98 && contenido[i+1] == 111 && contenido[i+2] == 111 && contenido[i+3] == 108 && contenido[i+4] == 101 && contenido[i+5] == 97 && contenido[i+6] == 110 && contenido[i+7] == 111 {
			booleanoPR++
			i += 7
			continue
		} else if contenido[i] == 100 && contenido[i+1] == 101 && contenido[i+2] == 99 && contenido[i+3] == 105 && contenido[i+4] == 109 && contenido[i+5] == 97 && contenido[i+6] == 108 {
			decimalPR++
			i += 6
			continue
		} else if contenido[i] == 99 && contenido[i+1] == 97 && contenido[i+2] == 100 && contenido[i+3] == 101 && contenido[i+4] == 110 && contenido[i+5] == 97 {
			cadenaPR++
			i += 5
			continue
		} else if contenido[i] == 115 && contenido[i+1] == 105 {
			if contenido[i+2] == 110 && contenido[i+3] == 111 {
				sinoPR++
				i += 3
			} else {
				siPR++
				i += 1
			}
			continue
		} else if contenido[i] == 109 && contenido[i+1] == 105 && contenido[i+2] == 101 && contenido[i+3] == 110 && contenido[i+4] == 116 && contenido[i+5] == 114 && contenido[i+6] == 97 && contenido[i+7] == 115 {
			mientrasPR++
			i += 7
			continue
		} else if contenido[i] == 104 && contenido[i+1] == 97 && contenido[i+2] == 99 && contenido[i+3] == 101 && contenido[i+4] == 114 {
			hacerPR++
			i += 4
			continue
		} else if contenido[i] == 118 && contenido[i+1] == 101 && contenido[i+2] == 114 && contenido[i+3] == 100 && contenido[i+4] == 97 && contenido[i+5] == 100 && contenido[i+6] == 101 && contenido[i+7] == 114 && contenido[i+8] == 111 {
			verdaderoPR++
			i += 8
			continue
		} else if contenido[i] == 102 && contenido[i+1] == 97 && contenido[i+2] == 108 && contenido[i+3] == 115 && contenido[i+4] == 111 {
			falsoPR++
			i += 4
			continue
		} else {
			switch contenido[i] {
			case 13:
				// ignorar el retorno del carro
				continue
			case 10:
				linea++
				continue
			case 32:
				// ignorar los espacios
				continue
			case 43:
				nMas++
				continue
			case 45:
				nMenos++
				continue
			case 42:
				nAsterisco++
				continue
			case 47:
				nBarraI++
				continue
			case 37:
				nPorCiento++
				continue
			case 61:
				if contenido[i+1] == 61 {
					nDobleIgual++
					i++
				} else {
					nIgual++
				}

				continue
			case 60:
				if contenido[i+1] == 61 {
					nMenorIgual++
					i++
				} else {
					nMenorQue++
				}
				continue
			case 62:
				if contenido[i+1] == 61 {
					nMayorIgual++
				} else {
					nMayorQue++
				}
				continue
			case 40:
				nAbreParentesis++
				continue
			case 41:
				nCierraParentesis++
				continue
			case 123:
				nAbreLlaves++
				continue
			case 125:
				nCierraLlaves++
				continue
			case 34:
				nComillas++
				continue
			case 59:
				nPuntoComa++
				continue
			default:
				// comprobar si es un identificador o un error
				if isLetter(contenido[i]) && isLetter(contenido[i+1]) || isLetter(contenido[i]) && isDigit(contenido[i+1]) {
					j := i
					var identificadores string
					for isLetter(contenido[j]) || isDigit(contenido[j]) {
						identificadores += string(contenido[j])
						j++
					}
					i = j - 1

					iRepetido := false
					for k := 0; k < len(identificadoresEncontrados); k++ {
						if identificadoresEncontrados[k][0] == identificadores {
							reemplazo, _ := strconv.Atoi(identificadoresEncontrados[k][2])
							identificadoresEncontrados[k][2] = strconv.Itoa(reemplazo + 1)
							iRepetido = true
						}
					}
					if iRepetido == false {
						agregar = []string{identificadores, "IDENTIFICADOR", "1"}
						identificadoresEncontrados = append(identificadoresEncontrados, agregar)
					}
				} else {
					if cambio != linea {
						errores += fmt.Sprintf("Error en la línea %d\n", linea)
						cambio = linea
					}
				}
			}
		}
	}

	data := [][]string{[]string{"entero", "Palabra Reservada", strconv.Itoa(enteroPR)},
		{"decimal", "Palabra Reservada", strconv.Itoa(decimalPR)},
		{"booleano", "Palabra Reservada", strconv.Itoa(booleanoPR)},
		{"cadena", "Palabra Reservada", strconv.Itoa(cadenaPR)},
		{"si", "Palabra Reservada", strconv.Itoa(siPR)},
		{"sino", "Palabra Reservada", strconv.Itoa(sinoPR)},
		{"mientras", "Palabra Reservada", strconv.Itoa(mientrasPR)},
		{"hacer", "Palabra Reservada", strconv.Itoa(hacerPR)},
		{"verdadero", "Palabra Reservada", strconv.Itoa(verdaderoPR)},
		{"falso", "Palabra Reservada", strconv.Itoa(falsoPR)},
		{"+", "Operador", strconv.Itoa(nMas)},
		{"-", "Operador", strconv.Itoa(nMenos)},
		{"*", "Operador", strconv.Itoa(nAsterisco)},
		{"/", "Operador", strconv.Itoa(nBarraI)},
		{"%", "Operador", strconv.Itoa(nPorCiento)},
		{"=", "Operador", strconv.Itoa(nIgual)},
		{"==", "Operador", strconv.Itoa(nDobleIgual)},
		{"<", "Operador", strconv.Itoa(nMenorQue)},
		{">", "Operador", strconv.Itoa(nMayorQue)},
		{">=", "Operador", strconv.Itoa(nMayorIgual)},
		{"<=", "Operador", strconv.Itoa(nMenorIgual)},
		{"(", "Operador", strconv.Itoa(nAbreParentesis)},
		{")", "Signos", strconv.Itoa(nCierraParentesis)},
		{"{", "Signos", strconv.Itoa(nAbreLlaves)},
		{"}", "Signos", strconv.Itoa(nCierraLlaves)},
		{"\"", "Signos", strconv.Itoa(nComillas)},
		{";", "Signos", strconv.Itoa(nPuntoComa)}}
	for i := 0; i < len(identificadoresEncontrados); i++ {
		data = append(data, identificadoresEncontrados[i])
	}
	for i := 0; i < len(numerosEncontrados); i++ {
		data = append(data, numerosEncontrados[i])
	}
	return errores, data
}
