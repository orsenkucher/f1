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
    plt.figure(figsize=(16, 6), dpi=80)
    plt.subplot(121)
    scatter(data)
    plt.subplot(122)
    errorbars(data)
    plt.suptitle(file)
    plt.tight_layout()
    # plt.show()
    plt.savefig(file + '.png', dpi=300)
    plt.savefig(os.path.join(plot_dir, file.replace('/','-').replace('\\','-') + '.png'), dpi=300)

def scatter(data):
    for g in data:
        records = g['Records']
        if not records: return
        xs = [float(r['E']) for r in records]
        ys = [float(r['F']) for r in records]
        plt.scatter(xs, ys, s=2)
    # plt.title(file)
    plt.legend([g['Name'] for g in data], loc="upper right")
    plt.xlabel('$E_\gamma, MeV$')
    plt.ylabel('$F, MeV^-3$')
    
def errorbars(data):
    for g in data:
        records = g['Records']
        if not records: return
        xs = [float(r['E']) for r in records]
        ys = [float(r['F']) for r in records]

        lower_err = [float(r['DFMinus']) for r in records]
        upper_err = [float(r['DFPlus']) for r in records]
        obj = {
            'x': xs,
            'y': ys,
            'yerr': np.array(list(zip(lower_err, upper_err))).T
        }
        plt.errorbar(**obj, alpha=.75, fmt=':', capsize=3, capthick=1, elinewidth=1)
        obj = {
            'x': obj['x'],
            'y1': [y - e for y, e in zip(obj['y'], obj['yerr'][0])],
            'y2': [y + e for y, e in zip(obj['y'], obj['yerr'][1])]}
        plt.fill_between(**obj, alpha=.25)
    # plt.title(file)
    plt.legend([g['Name'] for g in data], loc="upper right")
    plt.xlabel('$E_\gamma, MeV$')
    plt.ylabel('$F, MeV^-3$')
    
if __name__ == '__main__': 
    main()
