import os
import sys
import json
import numpy as np
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
    suptitle(data, file)
    plt.tight_layout()
    # plt.show()
    plt.savefig(file + '.png', dpi=128)
    plt.savefig(os.path.join(plot_dir, file.replace(
        '/', '-').replace('\\', '-') + '.png'), dpi=128)


def suptitle(data, file):
    el = get_element(data)
    s = el['Symbol']
    a = el['Mass']
    z = el['Number']
    deform = get_deform(data)
    energy = get_neutron_energy(data)
    if s:
        plt.suptitle(
            f'$^{{{a}}}_{{{z}}}{s},\ \\beta_{{2ef}}: {{{deform}}},\ separation\ energy: {{{energy:.2f}}}, MeV$\n'+file)
    else:
        plt.suptitle(file)


def get_element(data):
    for g in data:
        val = g['Element']
        return val


def get_neutron_energy(data):
    for g in data:
        val = g['NeutronEnergy']
        if not val:
            return None
        return float(val)/1000


def get_deform(data):
    for g in data:
        val = g['Deform']
        if not val:
            return None
        return float(val)


def energy_line(data):
    energy = get_neutron_energy(data)
    if energy:
        plt.axvline(x=energy,
                    c='r', alpha=.7, dashes=(3, 2),
                    )
        y = get_max_y(data)
        plt.text(energy+.1, y,
                 "{:.2f}, MeV".format(energy),
                 rotation=90,
                 verticalalignment='top',
                 alpha=.7, c='r',
                 )


def get_max_y(data):
    m = .0
    for g in data:
        records = g['Records']
        if not records:
            return m
        ys = [float(r['F']) for r in records]
        max_y = max(ys)
        m = max(m, max_y)
    return m


def scatter(data):
    for g in data:
        records = g['Records']
        if not records:
            return
        xs = [float(r['E']) for r in records]
        ys = [float(r['F']) for r in records]
        plt.scatter(xs, ys, s=2)
    energy_line(data)
    # plt.title(file)
    plt.legend([g['Name'] for g in data], loc="upper right")
    plt.xlabel('$E_\gamma, MeV$')
    plt.ylabel('$F, MeV^-3$')


def errorbars(data):
    for g in data:
        records = g['Records']
        if not records:
            return
        xs = [float(r['E']) for r in records]
        ys = [float(r['F']) for r in records]

        lower_err = [float(r['DFMinus']) for r in records]
        upper_err = [float(r['DFPlus']) for r in records]
        obj = {
            'x': xs,
            'y': ys,
            'yerr': np.array(list(zip(lower_err, upper_err))).T
        }
        plt.errorbar(**obj, alpha=.75, fmt=':',
                     capsize=3, capthick=1, elinewidth=1)
        obj = {
            'x': obj['x'],
            'y1': [y - e for y, e in zip(obj['y'], obj['yerr'][0])],
            'y2': [y + e for y, e in zip(obj['y'], obj['yerr'][1])]}
        plt.fill_between(**obj, alpha=.25)
    energy_line(data)
    # plt.title(file)
    plt.legend([g['Name'] for g in data], loc="upper right")
    plt.xlabel('$E_\gamma, MeV$')
    plt.ylabel('$F, MeV^-3$')


if __name__ == '__main__':
    main()
