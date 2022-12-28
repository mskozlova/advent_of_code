max_cal = []


def recalc_max_cal(max_cal, new_value):
    if len(max_cal) < 3:
        max_cal.append(new_value)
        return sorted(max_cal, key=lambda x: -x)

    if new_value <= max_cal[-1]:
        return max_cal

    max_cal.pop()
    max_cal.append(new_value)

    return sorted(max_cal, key=lambda x: -x)


with open("input.txt", "r") as file:
    current_cal = 0
    for line in file.readlines():
        if line.strip() == "":
            max_cal = recalc_max_cal(max_cal, current_cal)
            current_cal = 0
        else:
            current_cal += int(line.strip())
    max_cal = recalc_max_cal(max_cal, current_cal)

print(sum(max_cal))
