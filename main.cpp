#include "src/gen_data.h"
#include "src/coder.h"
#include "src/channel.h"
#include "src/orchestrator.h"

#include <iostream>

int main() {
int k = 2;

// Coder coder(k);
// Channel channel(1.0); // Примерное значение СКО

Orchestrator orchestrator(4, 1000, 0.0, 3.0, 0.1); // Задаем параметры симуляции

orchestrator.runSimulations();

return 0;
}