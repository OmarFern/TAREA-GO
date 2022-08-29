""" Seleccion ordena el arreglo recibido mediante el algoritmo de selección."""

#a=[]
import doctest


#a=[4,10,-14,2,5]
def seleccion(a):
    """     
    >>> seleccion([])
    []
    >>> seleccion([8])
    [8]
    >>> seleccion([1, 2, 3, 4, 5]) 
    [1, 2, 3, 4, 5]
    >>> seleccion([4, 8, 15, 16, 23, 42])
    [4, 8, 15, 16, 23, 42]
    >>> seleccion([-78, -65, -46, -38])
    [-78, -65, -46, -38]
    """
    
    if len(a)==0:
        resultado=[]
    elif(len(a)==1): 
        resultado=a
    else:
        n=len(a)
        for i in range(n):
            for j in range(0,n-i-1):
                if a[j]>a[j+1]:
                    a[j],a[j+1]=a[j+1],a[j]
        resultado=a
        return resultado
            
        
    return resultado

print(doctest.testmod())