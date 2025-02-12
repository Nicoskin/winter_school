#ifndef ORCHESTRATOR_H
#define ORCHESTRATOR_H

class Orchestrator {
public:
    Orchestrator(int k, int numSimulations, double startSigma, double endSigma, double sigmaStep);
    void runSimulations();

private:
    int k_;
    int numSimulations_;
    double startSigma_;
    double endSigma_;
    double sigmaStep_;
};

#endif