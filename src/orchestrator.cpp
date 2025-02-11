#include "orchestrator.h"
#include "channel.h"
#include "gen_data.h"
#include "coder.h"
#include <iostream>
#include <vector>
#include <fstream>

Orchestrator::Orchestrator(int k, int numSimulations, double startSigma, double endSigma, double sigmaStep) :
    k_(k),
    numSimulations_(numSimulations),
    startSigma_(startSigma),
    endSigma_(endSigma),
    sigmaStep_(sigmaStep) {}

void Orchestrator::runSimulations() {
    std::ofstream outputFile("results.txt"); // Файл для сохранения результатов

    if (!outputFile.is_open()) {
        std::cerr << "Ошибка: не удалось открыть файл для записи результатов!" << std::endl;
        return;
    }

    int k = k_;
    Coder coder(k);

    for (double sigma = startSigma_; sigma <= endSigma_; sigma += sigmaStep_) {
        Channel channel(sigma); // Создаем канал с текущим значением СКО

        int totalErrors = 0;
        for (int i = 0; i < numSimulations_; ++i) {
            std::vector<int> data = Utils::generateRandomData(k);
            std::vector<int> codeWord = coder.encode(data);
            std::vector<double> noisyCodeWord = channel.awgn(codeWord);
            std::vector<int> decodedData = coder.decode(std::vector<int>(noisyCodeWord.begin(), noisyCodeWord.end())); // Преобразуем double в int

            double errorRate = Utils::calculateErrorRate(data, decodedData);
            totalErrors += (errorRate > 0); // Если есть ошибки, считаем как одну ошибку
        }

        double averageErrorRate = (double)totalErrors / numSimulations_;
        std::cout << "Sigma: " << sigma << ", Average Error Rate: " << averageErrorRate << std::endl;
        outputFile << sigma << " " << averageErrorRate << std::endl; // Пишем в файл
    }

    outputFile.close(); // Закрываем файл
}