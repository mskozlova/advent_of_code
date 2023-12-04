calibration_sum = 0

digits_encoding = {
    "one": 1,
    "two": 2,
    "three": 3,
    "four": 4,
    "five": 5,
    "six": 6,
    "seven": 7,
    "eight": 8,
    "nine": 9,
}


def get_digits(line):
    digits = []

    for i, sym in enumerate(line):
        if sym.isdigit():
            digits.append(int(sym))

        for encoding, digit in digits_encoding.items():
            if line[i:].startswith(encoding):
                digits.append(digit)

    return digits


with open("input.txt") as file:
    for line in file.readlines():
        digits = get_digits(line)
        calibration_sum += digits[0] * 10 + digits[-1]

print(calibration_sum)
