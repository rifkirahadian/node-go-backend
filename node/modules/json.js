const fs = require("fs");

exports.readJsonFileSync = (filepath) => {
    filepath = __dirname + '/' + filepath;

    let file = fs.readFileSync(filepath, 'utf8');
    return JSON.parse(file);
}