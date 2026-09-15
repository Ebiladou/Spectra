                    Spectra CLI
                        │
                        ▼
                 Test Definition
                        │
                        ▼
                  Test Validator
                        │
                        ▼
                  Execution Engine
                        │
             ┌──────────┴──────────┐
             ▼                     ▼
        Load Controller       Virtual Users
                                   │
                                   ▼
                              Scenarios
                                   │
                                   ▼
                                Steps
                                   │
                         ┌─────────┴─────────┐
                         ▼                   ▼
                    HTTP Runner        WebSocket Runner
                         │                   │
                         └─────────┬─────────┘
                                   ▼
                              Target System
                                   │
                                   ▼
                                Metrics
                                   │
                                   ▼
                              Aggregator
                                   │
                         ┌─────────┴─────────┐
                         ▼                   ▼
                    Terminal              JSON
                     Report              Report