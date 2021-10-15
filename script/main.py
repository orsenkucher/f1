import sys
import json

print(sys.argv[1:])
file = sys.argv[1]
with open(file) as f:
    data = json.load(f)
    print(len(data))