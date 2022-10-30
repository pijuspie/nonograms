var __spreadArray = (this && this.__spreadArray) || function (to, from, pack) {
    if (pack || arguments.length === 2) for (var i = 0, l = from.length, ar; i < l; i++) {
        if (ar || !(i in from)) {
            if (!ar) ar = Array.prototype.slice.call(from, 0, i);
            ar[i] = from[i];
        }
    }
    return to.concat(ar || Array.prototype.slice.call(from));
};
var sqValue = {
    unknown: 0,
    empty: 1,
    filled: 2
};
function posibleLines(values, index, line, start, merge) {
    var locations = posibleLocations(values[index], line, start);
    locations.forEach(function (location) {
        var l = __spreadArray([], line, true);
        if (l[location - 1] !== undefined)
            l[location - 1] = sqValue.empty;
        if (l[location + values[index]] !== undefined)
            l[location + values[index]] = sqValue.empty;
        for (var j = 0; j < values[index]; j++) {
            l[location + j] = sqValue.filled;
        }
        if (values.length === index + 1) {
            var total = 0;
            for (var i = 0; i < l.length; i++) {
                if (l[i] === sqValue.unknown)
                    l[i] = sqValue.empty;
                if (l[i] === sqValue.filled)
                    total++;
            }
            if (total === values.reduce(function (prev, curr) { return prev + curr; })) {
                merge(l);
            }
        }
        else {
            posibleLines(values, index + 1, l, location + values[index] + 1, merge);
        }
    });
}
function posibleLocations(square, line, start) {
    var result = [];
    for (var i = start; i < line.length; i++) {
        if (line[i - 1] > 1) {
            continue;
        }
        else if (line[i + square] > 1) {
            continue;
        }
        for (var j = 0; j < square; j++) {
            if (line[i + j] !== sqValue.unknown && line[i + j] !== sqValue.filled) {
                break;
            }
            if (j === square - 1) {
                result.push(i);
            }
        }
    }
    return result;
}
function solveLine(values, line) {
    var result = [];
    function merge(l) {
        if (result.length === 0) {
            result = __spreadArray([], l, true);
            return;
        }
        for (var i = 0; i < l.length; i++) {
            if (result[i] !== l[i]) {
                result[i] = line[i];
            }
        }
    }
    posibleLines(values, 0, line, 0, merge);
    if (result.length === 0) {
        return line;
    }
    else {
        return result;
    }
}
function createGrid(top, left) {
    var grid = [];
    for (var i = 0; i < left.length; i++) {
        grid[i] = [];
        for (var j = 0; j < top.length; j++) {
            grid[i][j] = sqValue.unknown;
        }
    }
    return grid;
}
function solveTime(top, left, grid) {
    var result = JSON.parse(JSON.stringify(grid));
    left.forEach(function (values, i) {
        result[i] = __spreadArray([], solveLine(values, result[i]), true);
    });
    top.forEach(function (values, i) {
        var column = result.map(function (row) { return row[i]; });
        var newColumn = solveLine(values, column);
        newColumn.forEach(function (square, j) {
            result[j][i] = square;
        });
    });
    return result;
}
function solve(top, left) {
    var grid = createGrid(top, left);
    var done = false;
    var times = 0;
    var _loop_1 = function () {
        grid = solveTime(top, left, grid);
        var hasUnknown = false;
        grid.forEach(function (row) {
            if (row.indexOf(sqValue.unknown) !== -1) {
                hasUnknown = true;
            }
        });
        if (!hasUnknown) {
            done = true;
        }
        times++;
    };
    while (!done) {
        _loop_1();
    }
    return [grid, times];
}
function readData() {
    var data = "50 35\n   16 5\n   16 5\n   17 5\n   15 5\n   13 5\n   13 5 4\n   1 14 9 3\n   3 12 2 3 2\n   16 2 3 1\n   14 1 3\n   13 2 3\n   12 1 1 3\n   9 1 3 2\n   6 1 2 2\n   6 1 2 2\n   6 2 2 3\n   5 2 1 2\n   4 4 1 2 1\n   4 2 1 3 1\n   3 3 4 1 4 2\n   3 10 3 7 3\n   2 6 6 2 5\n   3 7 2 2 3 2 1 4\n   2 12 5 6 3\n   1 9 2 6 5 3\n   1 9 1 3 1 6 3\n   9 3 2 1 2 1 3\n   5 3 3 2 2 1 3\n   6 6 2 2 4\n   6 6 2 2 5\n   5 4 2 2 5\n   7 2 1 2 6 4\n   9 1 2 9 2\n   5 3 3 3 4 1\n   6 2 3 3 3 1\n   8 1 4 2 3\n   8 1 7 2\n   7 3 5 1 3\n   6 3 1 9 2\n   4 1 3 1 3 2\n   4 1 1 2\n   2 2 2\n   2 2\n   2 2\n   3 1 2 1\n   3 3 2\n   9 3\n   5 4\n   5\n   5\n   2 4 4\n   2 4 6\n   3 5 7\n   4 4 7\n   1 1 4 5 12\n   2 8 5 15\n   17 16\n   17 12 2\n   17 13\n   16 11 2\n   15 8 2\n   14 11 2\n   13 11 10\n   13 1 9 2 2 2\n   12 1 2 4 6\n   11 1 4 4 2 2\n   9 2 2 4 3 2\n   7 4 3 2\n   5 2 3 3 2 2\n   4 7 2 3 3 2 7\n   3 3 3 6 2 2 5 4\n   2 2 10 2 3 2 3\n   1 2 3 3 2 3 2 3\n   1 3 3 2 2 2 1 2\n   2 2 3 2 2 1 2\n   2 3 4 2 3 2 2 2\n   2 3 1 3 2 4 2\n   2 1 1 5 2 2 2\n   2 2 4 2 2\n   2 2 3 3 3\n   5 3 2 1 5 2 3 2 2\n   6 3 3 2 4 3 2 3\n   7 4 4 12 5 3 4\n   8 8 14 9 5\n   9 6 18 5 6";
    var lines = data.split("\n   ");
    var valuesStrings = lines.map(function (line) { return line.split(" "); });
    var values = valuesStrings.map(function (line) { return line.map(function (value) { return Number(value); }); });
    var top = values.slice(1, values[0][0] + 1);
    var left = values.slice(values[0][0] + 1);
    return [top, left];
}
function printResult(grid, iterations) {
    console.log("");
    grid.forEach(function (row) {
        console.log(row.map(function (square) { return (square === sqValue.filled ? "#" : " "); }).join(" "));
    });
    console.log("");
    console.log("Iterations count: " + iterations);
}
var data = readData();
var start = new Date();
var _a = solve(data[0], data[1]), result = _a[0], iterations = _a[1];
var end = new Date();
printResult(result, iterations);
console.log("Duration: " + Math.round((end - start) / 1000) + "s");
