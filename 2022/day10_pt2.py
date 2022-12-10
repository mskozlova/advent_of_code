target_ticks = [20, 60, 100, 140, 180, 220]
tick_vals = [1]

with open("input.txt", "r") as file:
    for line in file.readlines():
        if line.startswith("noop"):
            tick_vals.append(tick_vals[-1])
        else:
            command, value = line.strip().split(" ")
            tick_vals.extend([tick_vals[-1], tick_vals[-1] + int(value)])

        if len(tick_vals) > 6 * 40 + 1:
            break

print(tick_vals)

pixels = []
for i in range(6):
    for j in range(40):
        
        if abs(tick_vals[i * 40 + j] - j) <= 1:
            pixels.append("£")
        else:
            pixels.append(".")

for i in range(6):
    print("".join(pixels[i * 40:(i + 1) * 40]))


