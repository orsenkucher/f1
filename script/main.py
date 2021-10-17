import os
import sys
import json
import numpy as np
import matplotlib
import matplotlib.pyplot as plt

print(sys.argv[1:])
file = sys.argv[1]
plot_dir = sys.argv[2]

def main():
    with open(file) as f:
        data = json.load(f)
        print(len(data))
        # for g in data:
        #     print (g['Name'])

    for g in data:
        records = g['Records']
        if not records: return
        xs = [float(r['E']) for r in records]
        ys = [float(r['F']) for r in records]
        plt.scatter(xs, ys, s=2)

    plt.title(file)
    plt.legend([g['Name'] for g in data])
    plt.xlabel('$E, MeV$')
    plt.ylabel('$F, MeV^-3$')

    # plt.show()
    plt.savefig(file + '.png', dpi=280)
    plt.savefig(os.path.join(plot_dir, file.replace('/','-').replace('\\','-') + '.png'), dpi=280)

if __name__ == '__main__': 
    main()
