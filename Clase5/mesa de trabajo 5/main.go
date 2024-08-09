package main

import "fmt"

/*Repetir el proceso de la ejercitación realizada en clase, pero ahora implementando fmt.Errorf()
para que el mensaje de error reciba por parámetro el valor de salary indicando que no alcanza el
mínimo imponible. El mensaje mostrado por consola deberá decir: “Error: el mínimo imponible es de 150.000 y
el salario ingresado es de: [salary]” (siendo [salary] el valor de tipo int pasado por parámetro).*/

type errSalary struct{
	mensaje string
}
func (e errSalary)Error()string{
	return e.mensaje
}

func main(){
	var salary =300000
	err := impuesto(salary)
	if err != nil {
		fmt.Println(err,salary)
	}else{
		fmt.Println("Debe pagar")
	}
}
func impuesto(salario int)error{
	if salario <150000{
		return errSalary{mensaje: "Error: el mínimo imponible es de 150.000 y el salario ingresado es de: "}
	}
	return nil
}