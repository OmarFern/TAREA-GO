import doctest




def suma(a):
    """"
    >>> suma([])
    0
    >>> suma([8])
    8
    >>> suma([3, 5, 4, 2, 1])
    15
    >>> suma([4, 8, 15, 16, 23, 42])
    108
    >>> suma([-38, -46, -65, -78])
    -227
    >>> suma([10, 9, -15, 0, 7, -12, 1])
    0
    """
    total=0
    for i in a:
        total+=i
    return total
print(doctest.testmod())