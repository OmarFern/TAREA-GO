package tp0

// Swap intercambia dos valores enteros.
func Swap(x *int, y *int) {

	    *x,*y =*y,*x
}



// Comparar compara dos arreglos de longitud especificada.
// Devuelve -1 si el primer arreglo es menor que el segundo; 0 si son iguales; o 1 si el primero es el mayor.
// Un arreglo es menor a otro cuando al compararlos elemento a elemento, el primer elemento en el que difieren
// no existe o es menor.
func Comparar(vector1 []int, vector2 []int) int {
	var resultado int
	if (len(vector1)==0) && (len(vector2)==0){
        resultado=0   
	} else if (len(vector1)==len(vector2)){
		contador:=0
		for i:= range vector1 {
			if (vector1[i]==vector2[i]){
				contador+=1
				if (contador==len(vector1)){
					resultado=0 }
			} else if (vector1[i]>vector2[i]){
				    resultado= 1
			} else if (vector1[i]<vector2[i]){
				    resultado= -1}}
		
	} else if (len(vector1)>len(vector2)){
		for i:= range vector2 {
		    if (vector1[i]>vector2[i]){
				resultado=1
		  } else if (vector1[i]<vector2[i]){
			    resultado=-1
		  } else {
			    resultado=1
		  }    }
	} else if (len(vector1)<len(vector2)){
		for i:= range vector1 {
		    if (vector1[i]>vector2[i]){
				 resultado=1
		    } else if (vector1[i]<vector2[i]){
			     resultado=-1
		    } else {
			     resultado=-1
		    }  }
		}
	
	return resultado
}


// Suma devuelve la suma de los elementos de un arreglo. En caso de no tener elementos, debe devolver 0.
// Esta función debe implementarse de forma RECURSIVA. Se puede usar una función auxiliar (que sea
// la recursiva).
func Suma(vector []int) int {
	total:=0
	for _, v := range vector {
		total+=v
	}
	return total
}

