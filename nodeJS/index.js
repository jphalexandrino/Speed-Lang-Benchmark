const fs = require('fs');
const readline = require('readline');
const mysql = require('mysql');

// Create MySQL connection
const connection = mysql.createConnection({
    host: 'localhost',
    user: 'root',
    password: 'usbw',
    database: 'test',
    port: 3307
});

// Drop and create table if not exists
connection.connect((err) => {
    if (err) throw err;
    connection.query("DROP TABLE IF EXISTS cnpj", (err, result) => {
        if (err) throw err;
        console.log("Table dropped");
        connection.query("CREATE TABLE cnpj (id INT AUTO_INCREMENT PRIMARY KEY, cnpj VARCHAR(20)) ENGINE=MyISAM", (err, result) => {
            if (err) throw err;
            console.log("Table created");
            readFiles();
        });
    });
});

function readFiles() {
    // Define the path of the folder you want to traverse
    const folder = '../folder/';

    // Register the initial time
    const start = Date.now();

    let allLines = 0;
    let filesProcessed = 0;
    const totalFiles = fs.readdirSync(folder).filter(file => file !== '.DS_Store').length;

    // Read files in the specified folder
    fs.readdir(folder, (err, files) => {
        if (err) throw err;
        files.forEach(file => {
            const filePath = folder + file;
            console.log(filePath);

            // Create read stream
            const fileStream = fs.createReadStream(filePath, { encoding: 'utf8' });

            // Create readline interface
            const rl = readline.createInterface({
                input: fileStream,
                crlfDelay: Infinity
            });

            rl.on('line', (line) => {
                allLines++;
                const parts = line.replace(/"/g, '').split(';');
                const cnpj = parts[0] + parts[2] + parts[5];
                connection.query(`INSERT INTO cnpj (cnpj) VALUES ('${cnpj}')`, (err, result) => {
                    if (err) throw err;
                });
            });

            rl.on('close', () => {
                console.log(`Finished processing file: ${file}`);
                filesProcessed++;
                if (filesProcessed === totalFiles) {
                    connection.end();
                    // Register the final time
                    const end = Date.now();
                    console.log(`Total time for processing files: ${(end - start) / 1000} seconds.`);
                    console.log("Finished!!!");
                }
            });
        });
    });
}
