from utils import get_input_lines

reports = list(map(lambda report: list(map(lambda level: int(level), report)), get_input_lines("inputs/dec2.txt")))

def jump_okay(a: int, b: int, asc: bool):
    if asc and a >= b:
        return False
    elif not asc and a <= b:
        return False
    elif abs(b - a) > 3:
        return False
    return True

def is_safe(report: list[int], dampen: bool = True):
    if len(report) < 2:
        # Single element or empty report are not technically unsafe
        return True
    asc = report[0] < report[1]
    for i in range(len(reports)):
        if i == len(report) - 1:
            break
        if not jump_okay(report[i], report[i+1], asc):
            if not dampen:
                # Allow function to selectively recurse - in this case only one level is allowed
                return False
            for j in range(len(report)):
                # Remove one element at a time and test again
                dampened_report = report[:j] + report[j+1:]
                if is_safe(dampened_report, dampen=False):
                    return True
            return False # TODO: Investigate possible early return - what criteria indicates that this report can't be dampened?
    return True

print(f"Part 1: {len(list(filter(lambda report: is_safe(report, dampen=False), reports)))}") # dampen=False to disable Part 2 logic
print(f"Part 2: {len(list(filter(lambda report: is_safe(report), reports)))}")