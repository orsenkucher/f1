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

    plt.subplot(122)
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

    plt.suptitle(file)
    plt.tight_layout()
    # plt.show()
    plt.savefig(file + '.png', dpi=300)
    plt.savefig(os.path.join(plot_dir, file.replace('/','-').replace('\\','-') + '.png'), dpi=300)

if __name__ == '__main__': 
    main()

# import random
# import matplotlib.pyplot as plt

# # create sample data
# N = 8
# data_1 = {
#     'x': list(range(N)),
#     'y': [10. + random.random() for dummy in range(N)],
#     'yerr': [.25 + random.random() for dummy in range(N)]}
# data_2 = {
#     'x': list(range(N)),
#     'y': [10.25 + .5 * random.random() for dummy in range(N)],
#     'yerr': [.5 * random.random() for dummy in range(N)]}

# # plot
# plt.figure()
# # only errorbar
# plt.subplot(211)
# for data in [data_1, data_2]:
#     plt.errorbar(**data, fmt='o')
# # errorbar + fill_between
# plt.subplot(212)
# for data in [data_1, data_2]:
#     plt.errorbar(**data, alpha=.75, fmt=':', capsize=3, capthick=1)
#     data = {
#         'x': data['x'],
#         'y1': [y - e for y, e in zip(data['y'], data['yerr'])],
#         'y2': [y + e for y, e in zip(data['y'], data['yerr'])]}
#     plt.fill_between(**data, alpha=.25)
