def get_figure(code1, code2):
    if code1 == "A" and code2 == "Z":
        return "B"
    if code1 == "A" and code2 == "X":
        return "C"
    if code1 == "A" and code2 == "Y":
        return "A"
    if code1 == "B" and code2 == "X":
        return "A"
    if code1 == "B" and code2 == "Y":
        return "B"
    if code1 == "B" and code2 == "Z":
        return "C"
    if code1 == "C" and code2 == "Y":
        return "C"
    if code1 == "C" and code2 == "Z":
        return "A"
    if code1 == "C" and code2 == "X":
        return "B"


code_2_points = {
    "X": 0,
    "Y": 3,
    "Z": 6,
}

figure_2_points = {"A": 1, "B": 2, "C": 3}

with open("input.txt", "r") as file:
    score = 0
    for line in file.readlines():
        code1, code2 = line.strip().split(" ")
        figure_score = figure_2_points[get_figure(code1, code2)]
        outcome_score = code_2_points[code2]
        score += outcome_score + figure_score

print(score)
