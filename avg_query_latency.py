import matplotlib.pyplot as plt
import numpy as np

# 'GET'
frequent_pattern = np.array([
    [2.101, 1.761, 1.763, 1.738, 1.767],
    [1.939,1.522,1.353,1.521,1.558],
    [2.208,1.550,1.429,1.402,1.444],
    [2.098,1.606,1.755,1.534,1.693]
])

# 'DELETE'
infrequent_pattern = np.array([
    [1.746,1.697,1.120,1.353,1.631],
    [1.175,1.133,1.154,1.228,1.315],
    [1.795,1.114,1.200,1.175,1.208],
    [1.267,1.228,1.170,1.306,1.405]
])

# '400'
rare_pattern = np.array([
    [1.142,1.243,1.112,1.244,1.145],
    [2.081,1.165,1.257,1.144,1.173],
    [1.806,1.233,1.133,1.223,1.373],
    [1.231,1.501,1.174,1.254,1.165]
])

fig, (ax0, ax1, ax2) = plt.subplots(nrows=1, ncols=3, sharex=True,sharey=True,
                                    figsize=(12, 6))

ax0.errorbar(range(1, 1+len(frequent_pattern)), np.mean(frequent_pattern, axis=1), np.std(frequent_pattern, axis=1), linestyle='None', marker='^')
ax1.errorbar(range(1, 1+len(infrequent_pattern)), np.mean(infrequent_pattern, axis=1), np.std(infrequent_pattern, axis=1), linestyle='None', marker='^')
ax2.errorbar(range(1, 1+len(rare_pattern)), np.mean(rare_pattern, axis=1), np.std(rare_pattern, axis=1), linestyle='None', marker='^')

ax0.title.set_text('Frequent Pattern')
ax1.title.set_text('Infrequent Pattern')
ax2.title.set_text('Rare Pattern')

fig.supxlabel('Machine Number')
fig.supylabel('Time to serve query (seconds)')

fig.suptitle('Average query latency for various regular expressions')
plt.show()


print(np.mean(frequent_pattern), np.mean(infrequent_pattern), np.mean(rare_pattern))

print(np.mean([np.mean(frequent_pattern), np.mean(infrequent_pattern), np.mean(rare_pattern)]))


