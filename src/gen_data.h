#ifndef GEN_DATA_H
#define GEN_DATA_H

#include <vector>

namespace Utils {
    std::vector<int> generateRandomData(int k);
    double calculateErrorRate(const std::vector<int>& originalData, const std::vector<int>& decodedData);
}

#endif // GEN_DATA_H