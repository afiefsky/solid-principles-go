# SOLID Principles in Go

This repository showcases the implementation of the **SOLID** principles using Go programming language. Each principle is demonstrated through clear and simple examples to help you understand their application in real-world scenarios.

## SOLID Principles

The **SOLID** principles consist of five key object-oriented design principles:

- **SRP:** [Single Responsibility Principle](#single-responsibility-principle-srp) – A class should have only one reason to change.
- **OCP:** [Open/Closed Principle](#openclosed-principle-ocp) – Software entities should be open for extension, but closed for modification.
- **LSP:** [Liskov Substitution Principle](#liskov-substitution-principle-lsp) – Objects of a superclass should be replaceable with objects of a subclass without affecting the correctness of the program.
- **ISP:** [Interface Segregation Principle](#interface-segregation-principle-isp) – Clients should not be forced to depend on interfaces they do not use.
- **DIP:** [Dependency Inversion Principle](#dependency-inversion-principle-dip) – High-level modules should not depend on low-level modules. Both should depend on abstractions.

Each principle is explained with an example in its respective folder.

## How to Run

To run each example, use the following commands in your terminal:

```bash
# Single Responsibility Principle (SRP)
go run srp/main.go

# Open/Closed Principle (OCP)
go run ocp/main.go

# Liskov Substitution Principle (LSP)
go run lsp/main.go

# Interface Segregation Principle (ISP)
go run isp/main.go

# Dependency Inversion Principle (DIP)
go run dip/main.go
