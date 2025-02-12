#include "coder.h"
#include <iostream>

Coder::Coder(int k) :k_(k) {
    // Copy first k columns from generator_Matrix_A_
    for (size_t i = 0; i < generator_Matrix_A_.size(); i++) {
        std::vector<int> row;
        for (int j = 0; j < k; j++) {
            row.push_back(generator_Matrix_A_[i][j]);
        }
        G.push_back(row);
    }
}

std::vector<int> Coder::encode(const std::vector<int>& data) {
    if (data.size() != k_) {
        std::cerr << "Ошибка: Размер входных данных не соответствует k." << std::endl;
        return {}; // Возвращаем пустой вектор в случае ошибки.
    }

    std::vector<int> codeWord(n_, 0); // Инициализируем кодовое слово нулями.

    // Матричное умножение: codeWord = data * G
    for (int i = 0; i < n_; ++i) {
        for (int j = 0; j < k_; ++j) {
            codeWord[i] ^= data[j] & G[i][j];  // Умножение по модулю 2
        }
    }

    return codeWord;
}

std::vector<int> Coder::decode(const std::vector<double>& receivedCodeWord) {
    if (receivedCodeWord.size() != n_) {
        std::cerr << "Ошибка: Размер принятого кодового слова не соответствует n." << std::endl;
        return {};
    }

    std::vector<int> decodedData(k_, 0);
    int bestIndex = -1;
    int maxCorrelation = -1;

    // Перебираем все возможные сообщения (2^k вариантов).
    for (int i = 0; i < (1 << k_); ++i) {
        std::vector<int> possibleData(k_);
        int temp = i;
        for (int j = k_ - 1; j >= 0; --j) {
            possibleData[j] = temp % 2;
            temp /= 2;
        }

        std::vector<int> possibleCodeWord = encode(possibleData);

        // Вычисляем Log-Likelihood Ratio (LLR)
        double llr = 0.0;
        for (size_t j = 0; j < receivedCodeWord.size(); ++j) {
            llr += (2 * possibleCodeWord[j] - 1) * (2 * receivedCodeWord[j] - 1);
        }

        // Ищем кодовое слово с максимальным LLR
        if (bestIndex == -1 || llr > maxCorrelation) {
            maxCorrelation = llr;
            bestIndex = i;
            decodedData = possibleData;
        }
    }

    if (bestIndex == -1) {
        std::cerr << "Ошибка декодирования." << std::endl;
        return {};
    }

    return decodedData;
}