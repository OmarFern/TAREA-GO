
"""
https://appdividend.com/2020/04/16/how-to-join-strings-in-golang-go-string-join-example/
"""


a=["malayalam"]

def validar(a):   
    if ( a ==a[::-1]):
        resultado= True
    else:
        resultado= False
    return resultado
print(validar(a))      

def invertir(a):
    lista=[]
    for i in range(len(a)):
        lista.append(a[i])
    resultado="".join(lista)   
    return resultado  
print(invertir(a))
    
