# Go Commerce : E-Commerce Platform [ User Service ]

## 📝 Project Overview
Go Commerce is an e-commerce system that provides user account management (buyer and seller roles), product management, and secure payment processing through the payment gateway

## 🎯 Problem Statement & Solution

#### Problem Statement
Many users and sellers face challenges managing their online sales and purchases due to difficulties in tracking products, handling orders, and processing payments securely. Manual processes or fragmented systems can lead to errors, delayed transactions, and poor user experience.

#### Solution
Go Commerce provides a comprehensive e-commerce platform that centralizes user account management, product management, and payment processing. The platform allows users to:

* Create and manage buyer or seller accounts efficiently
* Add, update, and manage product listings
* Process payments securely through integrated payment gateways
* Track orders and transactions seamlessly in one system

By centralizing these processes, Go Commerce improves operational efficiency, reduces errors, and enhances the overall online shopping experience for both buyers and sellers.

## 📚 Documentation
* [Go Commerce with REST API](https://github.com/fildzaanf/go-commerce-api)

## 🚀 Tools and Technologies 
* Go Programming Language
* Echo Go Framework
* GORM for Object Relational Mapping
* MySQL / PostgreSQL for Relational Database
* JSON Web Token (JWT) for Authentication
* Docker for Containerization
* Midtrans Payment Gateway integrated with Webhooks, SMTP, and GoMail for real-time payment notifications
* Amazon Web Services (AWS)
  * Amazon Simple Storage Service (S3)
* GRPC for efficient, low-latency, and strongly-typed API communication

## 🏛️ System Design and Architecture

* Clean Architecture
* Hexagonal Architecture
* Domain-Driven Design (DDD)
* Command Query Responsibility Segregation (CQRS)
* Microservices Architecture
* REST API
* GRPC
* Webhook

## ✨ Features

#### User Management

| Feature                   | Description                                                        |
| ------------------------- | ------------------------------------------------------------------ |
| User Registration & Login | Allows users to register and log in to access the platform         |
| Profile                   | Provides functionality to retrieve user profile information by ID  |

#### Product Management

| Feature           | Description                                                                    |
| ----------------- | ------------------------------------------------------------------------------ |
| Create Product    | Enables adding new products to the platform                                    |
| Update Product    | Allows updating existing product details by product ID                         |
| Delete Product    | Supports removing products from the platform by product ID                     |
| Retrieve Product  | Provides access to a single product by ID or a list of all available products  |

#### Payment Management

| Feature             | Description                                                                          |
| ------------------- | ------------------------------------------------------------------------------------ |
| Create Payment      | Allows users to create new payments for products or services                         |
| Retrieve Payment    | Provides access to all payments or details of a specific payment by ID               |
| Integration Payment | Supports real-time payment updates via Midtrans Webhook integration                  |
| Integration Email   | Integrated Midtrans Payment Gateway using Webhooks with SMTP and Go Mail to automate | 
|                     | event-driven email notifications based on real-time payment status updates           |

## 📡 gRPC Services

#### User

| RPC Method    | RPC Type | Description             |
|--------------|----------|-------------------------|
| RegisterUser | Unary    | Create new user         |
| LoginUser    | Unary    | User authentication     |
| GetUserByID | Unary    | Retrieve user by ID    |


## 📡 API Endpoints

#### Users

| Method | Endpoint        | Description                 |
| ------ | --------------- | --------------------------- |
| POST   | /users/register | Register a new user         |
| POST   | /users/login    | Login user                  |
| GET    | /users/:id      | Retrieve user profile by ID |




