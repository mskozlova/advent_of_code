max_cal = 0

with open("input.txt", "r") as file:
    current_cal = 0
    for line in file.readlines():
        if line.strip() == "":
            max_cal = max(max_cal, current_cal)
            current_cal = 0
        else:
            current_cal += int(line.strip())
    max_cal = max(max_cal, current_cal)

print(max_cal)
