#include "channel.h"
#include <random>   // ДЛя генерации случайных чисел
#include <cmath>    // sqrt()
#include <algorithm> // std::transform

Channel::Channel(double sigma) : sigma_(sigma) {}

std::vector<double> Channel::awgn(const std::vector<int>& codeWord) {
    std::random_device rd{};
    std::mt19937 gen{rd()};
    std::normal_distribution<> d{0,sigma_};

    std::vector<double> noisyCodeWord(codeWord.size());
    for(size_t i = 0; i < codeWord.size(); ++i) {
       noisyCodeWord[i] = codeWord[i] + d(gen);
    }

    return noisyCodeWord;
}