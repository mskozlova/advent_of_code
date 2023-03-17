with open("input.txt", "r") as file:
    row = file.readline()

    for i in range(len(row) - 4):
        if len(set(row[i:i + 4])) == 4:
            print(i + 4)
            break
