#include "src/gen_data.h"
#include "src/coder.h"
#include "src/channel.h"
#include "src/orchestrator.h"

#include <iostream>

int main() {

Orchestrator orchestrator(4, 1000, 0.0, 3.0, 0.1);

orchestrator.runSimulations();

return 0;
}