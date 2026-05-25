/**
 * Вариант 17
 * Формула:
 * y = (a + tan^2(bx)) / (b + ctg^2(ax))
 * где ctg(x) = 1 / tan(x)
 */

/**
 * Вычисляет y по формуле варианта.
 * @param a параметр a
 * @param b параметр b
 * @param x аргумент x (в радианах)
 * @returns значение y
 */
function calculateY(a: number, b: number, x: number): number {
  const EPS = 1e-12; 
  const tanBx = Math.tan(b * x);
  const tanAx = Math.tan(a * x);
  if (Math.abs(tanAx) < EPS) {
    throw new Error(
      `Невозможно вычислить ctg(ax): tan(ax) слишком близок к нулю (x = ${x}).`
    );
  }
  const ctgAx = 1 / tanAx;
  const denominator = b + ctgAx * ctgAx;
  if (Math.abs(denominator) < EPS) {
    throw new Error(`Знаменатель формулы равен нулю (x = ${x}).`);
  }
  const numerator = a + tanBx * tanBx;
  return numerator / denominator;
}
function format6(value: number): string {
  return value.toFixed(6);
}
function printRow(x: number, y: number): void {
  console.log(`${format6(x)} | ${format6(y)}`);
}
function main(): void {
  const a = 0.1;
  const b = 0.5;
  const x1 = 0.15;
  const xk = 1.37;
  const dx = 0.25;
  console.log("Задача A");
  console.log("x | y");
  for (let x = x1; x <= xk + 1e-12; x += dx) {
    try {
      const y = calculateY(a, b, x);
      printRow(x, y);
    } catch (error) {
      const message = error instanceof Error ? error.message : String(error);
      console.log(`${format6(x)} | ошибка: ${message}`);
    }
  }
  console.log("");
  const xValues: number[] = [0.2, 0.3, 0.44, 0.6, 0.56];
  console.log("Задача B");
  console.log("x | y");
  for (const x of xValues) {
    try {
      const y = calculateY(a, b, x);
      printRow(x, y);
    } catch (error) {
      const message = error instanceof Error ? error.message : String(error);
      console.log(`${format6(x)} | ошибка: ${message}`);
    }
  }
}
main();
