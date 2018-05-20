#!/usr/bin/python
import sys
import json
import random
import string

def main():
    num = int(sys.argv[1])
    if num <= 0:
        return
    with open("keyvals.map", 'w') as f:
        keyvals = {}
        gen_count = 0
        while True:
            key = ''.join(random.sample(string.ascii_letters + string.digits, 10))
            val = ''.join(random.sample(string.ascii_letters + string.digits, 10))
            if key in keyvals:
                continue
            keyvals[key] = val
            gen_count += 1
            if gen_count == num:
                break
        s = json.dumps(keyvals)
        f.write(s)

if __name__ == "__main__":
    main()
    

