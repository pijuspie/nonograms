const sqValue = {
   unknown: 0,
   empty: 1,
   filled: 2,
};

function posibleLines(values: number[], index: number, line: number[], start: number, merge: (line: number[]) => void) {
   const locations = posibleLocations(values[index], line, start);

   locations.forEach(location => {
      const l = [...line];

      if (l[location - 1] !== undefined) l[location - 1] = sqValue.empty;
      if (l[location + values[index]] !== undefined) l[location + values[index]] = sqValue.empty;

      for (let j = 0; j < values[index]; j++) {
         l[location + j] = sqValue.filled;
      }

      if (values.length === index + 1) {
         let total = 0;

         for (let i = 0; i < l.length; i++) {
            if (l[i] === sqValue.unknown) l[i] = sqValue.empty;
            if (l[i] === sqValue.filled) total++;
         }

         if (total === values.reduce((prev, curr) => prev + curr)) {
            merge(l);
         }
      } else {
         posibleLines(values, index + 1, l, location + values[index] + 1, merge);
      }
   });
}

function posibleLocations(square: number, line: number[], start: number): number[] {
   const result: number[] = [];

   for (let i = start; i < line.length; i++) {
      if (line[i - 1] > 1) {
         continue;
      } else if (line[i + square] > 1) {
         continue;
      }

      for (let j = 0; j < square; j++) {
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

function solveLine(values: number[], line: number[]): number[] {
   let result: number[] = [];

   function merge(l: number[]) {
      if (result.length === 0) {
         result = [...l];
         return;
      }

      for (let i = 0; i < l.length; i++) {
         if (result[i] !== l[i]) {
            result[i] = line[i];
         }
      }
   }

   posibleLines(values, 0, line, 0, merge);

   if (result.length === 0) {
      return line;
   } else {
      return result;
   }
}

function createGrid(top: number[][], left: number[][]): number[][] {
   const grid: number[][] = [];

   for (let i = 0; i < left.length; i++) {
      grid[i] = [];
      for (let j = 0; j < top.length; j++) {
         grid[i][j] = sqValue.unknown;
      }
   }

   return grid;
}

function solveTime(top: number[][], left: number[][], grid: number[][]): number[][] {
   const result: number[][] = JSON.parse(JSON.stringify(grid));

   left.forEach((values, i) => {
      result[i] = [...solveLine(values, result[i])];
   });

   top.forEach((values, i) => {
      const column = result.map(row => row[i]);
      const newColumn = solveLine(values, column);

      newColumn.forEach((square, j) => {
         result[j][i] = square;
      });
   });

   return result;
}

function solve(top: number[][], left: number[][]): [number[][], number] {
   let grid = createGrid(top, left);
   let done = false;
   let times = 0;

   while (!done) {
      grid = solveTime(top, left, grid);
      let hasUnknown = false;

      grid.forEach(row => {
         if (row.indexOf(sqValue.unknown) !== -1) {
            hasUnknown = true;
         }
      });

      if (!hasUnknown) {
         done = true;
      }

      times++;
   }

   return [grid, times];
}

function readData(): [number[][], number[][]] {
   const data = `50 35
   16 5
   16 5
   17 5
   15 5
   13 5
   13 5 4
   1 14 9 3
   3 12 2 3 2
   16 2 3 1
   14 1 3
   13 2 3
   12 1 1 3
   9 1 3 2
   6 1 2 2
   6 1 2 2
   6 2 2 3
   5 2 1 2
   4 4 1 2 1
   4 2 1 3 1
   3 3 4 1 4 2
   3 10 3 7 3
   2 6 6 2 5
   3 7 2 2 3 2 1 4
   2 12 5 6 3
   1 9 2 6 5 3
   1 9 1 3 1 6 3
   9 3 2 1 2 1 3
   5 3 3 2 2 1 3
   6 6 2 2 4
   6 6 2 2 5
   5 4 2 2 5
   7 2 1 2 6 4
   9 1 2 9 2
   5 3 3 3 4 1
   6 2 3 3 3 1
   8 1 4 2 3
   8 1 7 2
   7 3 5 1 3
   6 3 1 9 2
   4 1 3 1 3 2
   4 1 1 2
   2 2 2
   2 2
   2 2
   3 1 2 1
   3 3 2
   9 3
   5 4
   5
   5
   2 4 4
   2 4 6
   3 5 7
   4 4 7
   1 1 4 5 12
   2 8 5 15
   17 16
   17 12 2
   17 13
   16 11 2
   15 8 2
   14 11 2
   13 11 10
   13 1 9 2 2 2
   12 1 2 4 6
   11 1 4 4 2 2
   9 2 2 4 3 2
   7 4 3 2
   5 2 3 3 2 2
   4 7 2 3 3 2 7
   3 3 3 6 2 2 5 4
   2 2 10 2 3 2 3
   1 2 3 3 2 3 2 3
   1 3 3 2 2 2 1 2
   2 2 3 2 2 1 2
   2 3 4 2 3 2 2 2
   2 3 1 3 2 4 2
   2 1 1 5 2 2 2
   2 2 4 2 2
   2 2 3 3 3
   5 3 2 1 5 2 3 2 2
   6 3 3 2 4 3 2 3
   7 4 4 12 5 3 4
   8 8 14 9 5
   9 6 18 5 6`;

   const lines = data.split("\n   ");
   const valuesStrings = lines.map(line => line.split(" "));
   const values = valuesStrings.map(line => line.map(value => Number(value)));

   const top = values.slice(1, values[0][0] + 1);
   const left = values.slice(values[0][0] + 1);

   return [top, left];
}

function printResult(grid: number[][], iterations: number) {
   console.log("");

   grid.forEach(row => {
      console.log(row.map(square => (square === sqValue.filled ? "#" : " ")).join(" "));
   });

   console.log("");
   console.log("Iterations count: " + iterations);
}

const data = readData();
const start: any = new Date();
const [result, iterations] = solve(data[0], data[1]);
const end: any = new Date();
printResult(result, iterations);
console.log("Duration: " + Math.round((end - start) / 1000) + "s");
