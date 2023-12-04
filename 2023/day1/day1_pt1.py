calibration_sum = 0

with open("input.txt") as file:
    for line in file.readlines():
        digit_symbols = list(filter(lambda x: x.isdigit(), line))
        calibration_sum += int(digit_symbols[0] + digit_symbols[-1])

print(calibration_sum)
