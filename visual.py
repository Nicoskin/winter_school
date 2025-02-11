import matplotlib.pyplot as plt
import numpy as np

def plot_results(filename="results.txt", k=2):
    """
    Читает данные из файла и строит график.

    Args:
        filename (str, optional): Имя файла с данными. Defaults to "results.txt".
    """
    try:
        # Чтение данных из файла
        data = np.loadtxt(filename)
        sigma_values = data[:, 0]  # Первый столбец: sigma
        error_rates = data[:, 1]   # Второй столбец: error rate

        # Создание графика
        plt.figure(figsize=(10, 6))  # Размер графика
        plt.plot(sigma_values, error_rates, marker='o', linestyle='-', color='blue')  # Линия и маркеры

        # Настройка графика
        plt.title("Зависимость вероятности ошибки от уровня шума | k={}".format(k))
        plt.xlabel("Sigma (СКО шума)")
        plt.ylabel("Вероятность ошибки")
        plt.grid(True)  # Включаем сетку
        plt.xlim(left=0)
        plt.ylim(bottom=0)

        # Отображение графика
        plt.show()

    except FileNotFoundError:
        print(f"Ошибка: Файл '{filename}' не найден.")
    except Exception as e:
        print(f"Произошла ошибка при чтении или обработке файла: {e}")


if __name__ == "__main__":
    plot_results("build/results.txt", 6)  # Вызов функции для построения графика