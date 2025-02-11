#include "gen_data.h"
#include <random>

namespace Utils {

    std::vector<int> generateRandomData(int k) {
        std::random_device rd;
        std::mt19937 gen(rd());
        std::uniform_int_distribution<> distrib(0, 1); // Генерируем 0 или 1

        std::vector<int> data(k);
        for (int i = 0; i < k; ++i) {
            data[i] = distrib(gen);
        }
        return data;
    }

    double calculateErrorRate(const std::vector<int>& originalData, const std::vector<int>& decodedData) {
        if (originalData.size() != decodedData.size()) {
            return 1.0; // Если размеры не совпадают, считаем, что все неправильно
        }
        int errors = 0;
        for (size_t i = 0; i < originalData.size(); ++i) {
            if (originalData[i] != decodedData[i]) {
                errors++;
            }
        }
        return (double)errors / originalData.size();
    }
}