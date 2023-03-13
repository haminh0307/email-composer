# EMAIL COMPOSER

## To run the application

```
go run main.go /path/to/email_template.json /path/to/customers.csv /path/to/output_emails/ /path/to/errors.csv
```

For example,

```
go run main.go assets/email_template.json assets/customers.csv assets/ assets/errors.csv
```

## Design

I follow Clean Architecture, which consists of 4 layers: Entity, Repository, Usecase, and Controller.

Entities are defined in `entity` folder: `Customer`, `EmailTemplate`, and `OutputEmail`. `Customer` has 1 utility function to help write csv file.

The flow of the application is: Controller picks up necessary inputs, i.e. email template and customers list, pass them to Usecase. Usecase will handle it, return output emails, and error customers as well as call Repository to write it out.

Returning output helps Controller can deliver it if we implement RestAPI, and Repository can be changed to write into a database instead. With the help of Clean Architecture, we can apply those changes without much change in the codebase.

## Unit testing

### To run unit test

```
go test ./...
```

Each layer in Clean Architecture can be tested independently of the other. But in this application, repository and controller layers involves file read/write, which is hard to unit test. Therefore, I'll write unit tests for usecase layer only and use gomock to generate mocks for repository layer.

## Area for improvements

- Use fx for dependency injection
- Unit test for repository and controller, which involve file read/write
- Write utility functions to handle file read/write separately.
- The requirement suggests that the error customers should be appended to the `error.csv` file, but the csv writer always writes headers, resulting in headers being written again and again. This problem requires more research.
- Package the application in a Docker image.