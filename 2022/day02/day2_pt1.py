def play_game(code1, code2):
    if code1 == "A" and code2 == "Z":
        return 0
    if code1 == "A" and code2 == "X":
        return 3
    if code1 == "A" and code2 == "Y":
        return 6
    if code1 == "B" and code2 == "X":
        return 0
    if code1 == "B" and code2 == "Y":
        return 3
    if code1 == "B" and code2 == "Z":
        return 6
    if code1 == "C" and code2 == "Y":
        return 0
    if code1 == "C" and code2 == "Z":
        return 3
    if code1 == "C" and code2 == "X":
        return 6


code_2_points = {"X": 1, "Y": 2, "Z": 3}

with open("input.txt", "r") as file:
    score = 0
    for line in file.readlines():
        code1, code2 = line.strip().split(" ")
        outcome_score = play_game(code1, code2)
        figure_score = code_2_points[code2]
        score += outcome_score + figure_score

print(score)
