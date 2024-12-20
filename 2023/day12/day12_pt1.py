def parse_spring_area(file_name):
    areas = []
    damaged_info = []
    with open(file_name, "r") as file:
        for line in file.readlines():
            area, damaged_groups = line.strip().split(" ")
            areas.append(area)
            damaged_groups = list(map(int, damaged_groups.split(",")))
            damaged_info.append(damaged_groups)
    return areas, damaged_info


def count_options(current_area_pos, current_damaged_group, is_inside_group, inside_group_pos):
    global cache_dict, area_row, damaged_groups
    
    print(current_area_pos, current_damaged_group, is_inside_group, inside_group_pos)
    
    if current_area_pos == 0 and current_damaged_group == 0 and inside_group_pos == 0:
        return 1
    elif current_area_pos == 0:
        return 0
    
    if inside_group_pos < 0:
        return 0
    
    if current_damaged_group < 0:
        return 0
    
    if area_row[current_area_pos - 1] == ".":
        if is_inside_group and inside_group_pos != 0:
            return 0
        
        return count_options(current_area_pos - 1, current_damaged_group, False, 0)
    
    if area_row[current_area_pos - 1] == "#":
        if not is_inside_group:
            return count_options(current_area_pos - 1, current_damaged_group - 1, True, damaged_groups[current_damaged_group - 1])
        else:
            return count_options(current_area_pos - 1, current_damaged_group, True, inside_group_pos - 1)
        
    if area_row[current_area_pos - 1] == "?":
        n_options = 0
        
        # like .
        if (is_inside_group and inside_group_pos == 0) or not is_inside_group:
            n_options += count_options(current_area_pos - 1, current_damaged_group, False, 0)
        
        # like #
        if is_inside_group and inside_group_pos > 0:
            n_options += count_options(current_area_pos - 1, current_damaged_group, True, inside_group_pos - 1)
        
        if not is_inside_group:
            n_options += count_options(current_area_pos - 1, current_damaged_group - 1, True, damaged_groups[current_damaged_group - 1])
            
        return n_options


spring_area, damaged_info = parse_spring_area("input.txt")
total_combos = 0
for area_row, damaged_groups in zip(spring_area, damaged_info):
    print(area_row, damaged_groups)
    print(len(area_row))
    cache_dict = dict()
    row_options = count_options(len(area_row) - 1, len(damaged_groups) - 1, False, 0)
    total_combos += row_options
    print(f"row_options: {row_options}")
print(total_combos)
