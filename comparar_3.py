
'''  Comparar compara dos arreglos de longitud especificada.
Devuelve 
          0 si son iguales; 
          1 si el primero es el mayor.
         -1 si el primer arreglo es menor que el segundo;
Un arreglo es menor a otro cuando al compararlos elemento a elemento, 
el primer elemento en el que difieren no existe o es menor. '''


a=[1,2]
b=[1,2,3]
def comparar(a,b):

    if (len(a)==len(b)==0):
        resultado=0     #" A y B son vacios = 0"
    elif(len(a)==len(b)):
        contador=0
        for i in range(len(a)):        
            if a[i]==b[i]:
                contador+=1   
                if(contador==len(a)):
                     resultado= 0 #"A y B son iguales = 0"                     
            elif(a[i]>b[i]):
                resultado= 1 #" A es mayor = 1 "                 
            elif(a[i]<b[i]):
                resultado= -1 #" B es mayor = -1 "   
    elif((len(a)>len(b))):
        for i in range(len(b)):                            
            if(a[i]>b[i]):
                resultado= 1 #" A es mayor = 1 "                 
            elif(a[i]<b[i]):
                resultado=-1 #" B es mayor = -1 "
            else:
                resultado=1 # " A es mayor"
    elif((len(a)<len(b))):
        for i in range(len(a)):                
            if(a[i]>b[i]):
                resultado= 1 #" A es mayor = 1 "                 
            elif(a[i]<b[i]):
                resultado= -1#" B es mayor = -1 "
            else:
                resultado=-1 #" A es mayor"
    return resultado     

print(comparar(a,b))   


