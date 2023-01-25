target_ticks = [20, 60, 100, 140, 180, 220]
tick_vals = [1]

with open("input.txt", "r") as file:
    for line in file.readlines():
        if line.startswith("noop"):
            tick_vals.append(tick_vals[-1])
        else:
            command, value = line.strip().split(" ")
            tick_vals.extend([tick_vals[-1], tick_vals[-1] + int(value)])

        if len(tick_vals) > 221:
            break

print(sum([tick_vals[tick - 1] * tick for tick in target_ticks]))
