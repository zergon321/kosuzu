k = require("./kosuzu.js");
data = k.serialize(32, {
    name: "Vasya",
    age: 16,
    numbers: new Uint8Array([32, 25, 79])});
console.log(data);