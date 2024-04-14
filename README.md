# Testing Programming Language Speeds!

Let's see practically which programming language is the best! In this repository, I conducted a test to determine which programming language would be the best. Remember, this repository is not to say which language is better than the other; programming languages are just tools.

## Languages tested:
  <img src="https://raw.githubusercontent.com/devicons/devicon/master/icons/php/php-original.svg" alt="php" width="40" height="40"/><img src="https://raw.githubusercontent.com/devicons/devicon/master/icons/python/python-original.svg" alt="python" width="40" height="40"/><img src="https://raw.githubusercontent.com/devicons/devicon/master/icons/go/go-original.svg" alt="go" width="40" height="40"/>

# Test Method:
The test was performed by downloading a file from the Brazilian government's `open CNPJ data`. The file is in CSV format. In the test, we will send the `CNPJ` which is separated in the lines of the file to the `MySQL` database.

## How it worked?
- Inside the CSV file, each CNPJ will be separated in the lines. For example:
```bash
"99999999";"Name";"9999";"99";"9.99";"99";""
In this example, the line is separated by ";". In this case, we can see this line as follows:
[0] [1] [2] [3] [4] [5] [6] [7]
As we will only take the values
[0], [2] and [5]
```
- After that, we will send the values we got to MySQL. In this case, we are using the `UsbWebServer` because of PHP, which requires an Apache server.

# Dependencies:
Each language has its dependency for using MySQL (except PHP).

## Golang:
```bash
go get github.com/go-sql-driver/mysql
```
This will install the MySQL driver and all its necessary dependencies inside your module.

## Python:
```bash
pip install mysql-connector-python 
```
With this, we will have everything we need to perform the test.

# Code Differences:

### PHP:
PHP had the fewest lines of code, with only 46 lines! (Counting comments)
The code looked like this:
![Code photo](/images/code2.png)

### Python:
Python, on the other hand, had the second-highest number of lines, with 60 lines! Which is 14 more lines. (Counting comments)
The code looked like this:
![Code photo](/images/code3.png)

### Golang:
Go is a bit more verbose, thus it had the most lines. Having 88 lines! Which is almost double the `PHP`.
The code looked like this:
![Code photo](/images/code.png)

# Test Results:
Unlike the number of lines in the codes, the speed score was quite different. The file in question has 4,494,861 lines, so the tests took a long time.

The results of the tests were:

|Language | Time (m)| Errors |
|----------|---------|-------|
|PHP       |  04.75  |   0   |
|Go        |  7.36   |   0   |
|Python    |  11.75  |   0   |

With this, we can determine that in this set of tests performed on `UsbWebServer`, PHP came out as the winner.
However, this result may vary, both from computer to computer and from database to database. But it is important to remember that programming language is just a tool and that there is no "silver bullet" in our field, only languages that we are familiar with and prefer.
