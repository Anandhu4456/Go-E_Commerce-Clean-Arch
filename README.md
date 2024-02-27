# Yours Store E-Commerce Backend Rest API
***
Yours Store E-Commerce Backend Rest API is a feature-rich backend solution for E-commerce applications developed using Golang with the Gin web framework. This API is designed to efficiently handle routing and HTTP requests while following best practices in code architecture and dependency management.

# Key Features
***
- Clean Code Architecture: The project follows clean code architecture principles, making it maintainable and scalable.
- Dependency Injection: Utilizes the Dependency Injection design pattern for flexible component integration.
- Compile-Time Dependency Injection: Dependencies are managed using Wire for compile-time injection.
- Database: Leverages PostgreSQL for efficient and relational data storage.
- AWS Integration: Integrates with AWS S3 for cloud-based storage solutions.
- E-Commerce Features: Implements a wide range of features commonly found in e-commerce applications, including cart management, wishlist, wallet, offers, and coupon management.
- Code Quality: Implements Continuous Integration(CI) with GitHub Actions.
***
# API Documentation
***
For interactive API documentation, Swagger is implemented. You can explore and test the API endpoints in real-time.
***
# Getting Started
***
To run the project locally, you can follow these steps:

1. Clone the repository.
2. Set up your environment with the required dependencies, including Golang, PostgreSQL, Docker, and Wire.
3. Configure your environment variables (e.g., database credentials, AWS keys, Twilio credentials).
4. Build and run the project.
***
# Environment Variables
***
Before running the project, you need to set the following environment variables with your corresponding values:

# PostgreSQL
***
- `DB_HOST`: Database host
- `DB_NAME`: Database name
- `DB_USER`: Database user
- `DB_PORT`: Database port
- `DB_PASSWORD`: Database password
***
# Twilio
***
- `DB_AUTHTOKEN`: Twilio authentication token
- `DB_ACCOUNTSID`: Twilio account SID
- `DB_SERVICESID`: Twilio services ID
***
# AWS
***
- `AWS_REGION`: AWS region
- `AWS_ACCESS_KEY_ID`: AWS access key ID
- `AWS_SECRET_ACCESS_KEY`: AWS secret access key

Make sure to provide the appropriate values for these environment variables to configure the project correctly.
