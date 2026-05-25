/**
 * Лабораторная работа: Вычисление значений функции
 * Вариант 17
 * 
 * Функция: y = (a - lg²(bx)) / (b + arccos²(ax))
 * 
 * Параметры:
 * a = 0.1
 * b = 0.5
 * 
 * Задача А:
 * Значения x: 0.15, 1.37, 0.25, 0.2, 0.3, 0.44, 0.6, 0.56
 * 
 * Задача B:
 * Значения x: 0.15, 0.2, 0.3, 0.44, 0.56, 0.6
 */

// ==========================================
// ФУНКЦИЯ ДЛЯ ВЫЧИСЛЕНИЯ ЗНАЧЕНИЯ Y
// ==========================================

/**
 * Вычисляет значение функции y = (a - lg²(bx)) / (b + arccos²(ax))
 * @param a - параметр a
 * @param b - параметр b
 * @param x - значение аргумента x
 * @returns объект { value, error } - результат вычисления или ошибку
 */
function calculateY(a: number, b: number, x: number): { value: number; error: string | null } {
    const bx: number = b * x;
    const ax: number = a * x;
    if (Math.abs(ax) > 1) {
        return { value: 0, error: `arccos(ax) не определен: |ax| = ${Math.abs(ax).toFixed(2)} > 1` };
    }
    const arccosAx: number = Math.acos(ax);
    const denominator: number = b + arccosAx * arccosAx;
    if (Math.abs(denominator) < 1e-10) {
        return { value: 0, error: `деление на ноль: b + arccos²(ax) = 0` };
    }
    if (bx <= 0) {
        return { value: 0, error: `lg(bx) не определен: bx = ${bx.toFixed(2)} <= 0` };
    }
    const lgBx: number = Math.log10(bx);
    const squaredLg: number = lgBx * lgBx;
    const numerator: number = a - squaredLg;
    const y: number = numerator / denominator;
    return { value: y, error: null };
}
function main(): void {
    const a: number = 0.1;
    const b: number = 0.5;
    console.log("========================================");
    console.log("ЛАБОРАТОРНАЯ РАБОТА");
    console.log("Вычисление значений функции");
    console.log("Вариант 17");
    console.log("========================================");
    console.log(`Функция: y = (a - lg²(bx)) / (b + arccos²(ax))`);
    console.log(`Параметры: a = ${a}, b = ${b}`);
    console.log("========================================\n");
    console.log("=== Задача А (вычисление через массив) ===");
    const xValuesA: number[] = [0.15, 1.37, 0.25, 0.2, 0.3, 0.44, 0.6, 0.56];
    for (let i: number = 0; i < xValuesA.length; i++) {
        const x: number = xValuesA[i];
        const result = calculateY(a, b, x);
        if (result.error !== null) {
            console.log(`При x = ${x.toFixed(2)} \t ОШИБКА: ${result.error}`);
        } else {
            console.log(`При x = ${x.toFixed(2)} \t y = ${result.value.toFixed(4)}`);
        }
    }
    console.log("\n==================================================\n");
    console.log("=== Задача В (вычисление через массив) ===");
    const xValuesB: number[] = [0.15, 0.2, 0.3, 0.44, 0.56, 0.6];
    for (let i: number = 0; i < xValuesB.length; i++) {
        const x: number = xValuesB[i];
        const result = calculateY(a, b, x);
        if (result.error !== null) {
            console.log(`При x = ${x.toFixed(2)} \t ОШИБКА: ${result.error}`);
        } else {
            console.log(`При x = ${x.toFixed(2)} \t y = ${result.value.toFixed(4)}`);
        }
    }
}
main();
export { calculateY };
