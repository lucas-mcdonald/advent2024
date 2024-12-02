from utils import get_input_lines

reports = list(map(lambda report: list(map(lambda level: int(level), report)), get_input_lines("inputs/dec2.txt")))

def jump_okay(a: int, b: int, asc: bool):
    # print(f"Comparing {a}, {b}")
    if asc and a >= b:
        return False
    elif not asc and a <= b:
        return False
    elif abs(b - a) > 3:
        return False
    return True

def is_safe(report: list[int], dampen: bool = True):
    if len(report) < 2:
        return True
    asc = report[0] < report[1]
    for i in range(len(reports)):
        if i == len(report) - 1:
            break
        if not jump_okay(report[i], report[i+1], asc):
            if not dampen:
                return False
            print("UNSAFE, trying to dampen")
            print(f"Report {report} - {'ASC' if asc else 'DESC'}")
            for j in range(len(report)):
                dampened_report = report[:j] + report[j+1:]
                print(f"Subset: {dampened_report}")
                if is_safe(dampened_report, dampen=False):
                    print(f"Dampened! Original - {report}, Subset")
                    return True
            return False
    return True

print(f"Part 1: {len(list(filter(lambda report: is_safe(report, dampen=False), reports)))}")
print(f"Part 2: {len(list(filter(lambda report: is_safe(report), reports)))}")