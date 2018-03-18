import sys
import time

def gcd(a, b):
    if a % b == 0:
        return b
    else:
        return gcd(b, a % b)

print gcd(int(sys.argv[1]), int(sys.argv[2]))
