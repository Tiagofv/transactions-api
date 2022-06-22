# Transactions

This is a project made for the pismo team to analyze. The project accepts creating an account, transaction and reading.

# How to run
We are using docker multistage builds to generate a production smallersize image. To run locally you can use the command:
```shell
docker build .
TODO
```
# Concepts
## Clean architecture
This project relies heavily on the concepts presented by uncle in his book "Clean Architecture: a craftsmen guide to software structure and design".
Clean arch allow us to decouple our apps from frameworks, testable design, become independent of external agents.
The most common implementation of Clean architecture is Hexagonal architeture, which is used in this project.

# Entities
Encapsulates enterprises rules. Can be an object with rules or a set of structures or functions.

# Use cases
Application specific business rules. Implements all systems use cases.

# Interface adapters (Presenters)
Converts the data from the use cases to a convenient format for the GUI or client.

