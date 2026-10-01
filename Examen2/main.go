package main

import "fmt"

var nombresVendidos []string
var subtotales []float64

func main() {
	var entrada int

	for {
		fmt.Println("*** MENÚ TIENDA ***")
		fmt.Println("1. Registrar venta")
		fmt.Println("2. Mostrar estadísticas")
		fmt.Println("0. Salir")
		fmt.Print("Ingrese su opción: ")
		fmt.Scanln(&entrada)

		if entrada == 0 {
			fmt.Println("Fin del programa.")
			break
		}

		switch entrada {
		case 1:
			fmt.Println("Productos disponibles:")
			fmt.Println("1. Arroz - $1.25")
			fmt.Println("2. Leche - $0.95")
			fmt.Println("3. Pan   - $0.50")

			var producto int
			fmt.Print("Seleccione el número del producto: ")
			fmt.Scanln(&producto)

			var cantidad int
			fmt.Print("Ingrese la cantidad vendida: ")
			fmt.Scanln(&cantidad)

			switch producto {
			case 1:
				RegistrarVenta("Arroz", 1.25, cantidad)
			case 2:
				RegistrarVenta("Leche", 0.95, cantidad)
			case 3:
				RegistrarVenta("Pan", 0.50, cantidad)
			default:
				fmt.Println("Producto no válido.")
			}

		case 2:
			MostrarEstadisticas()

		default:
			fmt.Println("Opción no válida.")
		}
	}
}

func RegistrarVenta(nombre string, precio float64, cantidad int) {
	subtotal := precio * float64(cantidad)

	nombresVendidos = append(nombresVendidos, nombre)
	subtotales = append(subtotales, subtotal)

	fmt.Println("Venta registrada:", nombre, "- Subtotal: $", subtotal)
}

func MostrarEstadisticas() {
	if len(subtotales) == 0 {
		fmt.Println("No existen ventas registradas todavía.")
		return
	}

	var total float64
	for _, subtotal := range subtotales {
		total += subtotal
	}

	fmt.Println("*** ESTADÍSTICAS ***")
	fmt.Println("Total de ventas realizadas:", len(nombresVendidos))
	fmt.Println("Total recaudado: $", total)
}
