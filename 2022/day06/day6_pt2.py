with open("input.txt", "r") as file:
    row = file.readline()

    for i in range(len(row) - 14):
        if len(set(row[i : i + 14])) == 14:
            print(i + 14)
            break
