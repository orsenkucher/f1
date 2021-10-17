import sys
import json
import matplotlib
import matplotlib.pyplot as plt

print(sys.argv[1:])
file = sys.argv[1]

def main():
    with open(file) as f:
        data = json.load(f)
        print(len(data))
        for g in data:
            print (g['Name'])

    records = g['Records']
    if not records: return
    for g in data:
        plt.scatter(
            [r['E'] for r in records],
            [r['F'] for r in records],
            s=1,
        )

    plt.legend([g['Name'] for g in data])
    plt.xlabel('E')
    plt.ylabel('F')
    plt.tick_params(
        axis='x',          # changes apply to the x-axis
        which='both',      # both major and minor ticks are affected
        bottom=False,      # ticks along the bottom edge are off
        top=False,         # ticks along the top edge are off
        labelbottom=False
    ) # labels along the bottom edge are offlt.ylabel('F')
    plt.tick_params(
        axis='y',          
        which='both',      
        left=False,      
        right=False,         
        labelleft=False
    )

    # plt.show()
    plt.savefig(file + '.png', dpi=280)

if __name__ == '__main__': 
    main()
