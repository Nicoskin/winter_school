#ifndef CHANNEL_H
#define CHANNEL_H

#include <vector>

class Channel {
public:
    Channel(double sigma);
    std::vector<double> awgn(const std::vector<int>& codeWord);

private:
    double sigma_;
};

#endif