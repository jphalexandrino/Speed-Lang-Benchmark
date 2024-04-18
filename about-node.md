# About Node.js

Node.js failed to execute the code completely, so it will not be included in the list. However, it may perform more appropriately in other situations.

Since the tests were conducted on a MySQL database, it's possible that Node.js may not be as efficient as other languages for relational database operations. Nevertheless, Node.js remains a great option for other types of utilities.

## To run this file:
- Install `MySQL` in Node.js.
``` bash
npm install mysql
```
- To initialize, use the following command:
```
node --max-old-space-size=9096 index.js
```
This will increase the amount of memory that the Node application can use (Be cautious with the value!).