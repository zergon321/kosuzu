kosuzu = require("./kosuzu.js");

// Serialize without a scheme.
data = kosuzu.serialize(32, {
    name: "Vasya",
    age: 16,
    numbers: new Uint8Array([32, 25, 79])});

console.log("Without a scheme:")
console.log(data);
console.log(kosuzu.packetBytes(data));

// Serialize with a scheme.
scheme = {
    name: "string",
    age: "int32",
    numbers: "[]byte"
};
data = kosuzu.serialize(32, {
    name: "Vasya",
    age: 16,
    numbers: new Uint8Array([32, 25, 79])}, scheme);

console.log("With a scheme:")
console.log(data);
console.log(kosuzu.packetBytes(data));

// Deserialize only works
// if you provide a scheme.
rawPacket = new Uint8Array([0, 0, 0, 32, 0, 0, 0, 0,
                            0, 0, 0, 20, 0, 0, 0, 5, 
                            86, 97, 115, 121, 97, 0, 0, 0,
                            16, 0, 0, 0, 3, 32, 25, 78]);
restoredPacket = kosuzu.packetFromBytes(rawPacket);

console.log("Packet to deserialize:")
console.log(restoredPacket);

scheme = {
    name: "string",
    age: "int32",
    numbers: "[]byte"
};
obj = kosuzu.deserialize(scheme, restoredPacket);

console.log("Deserialized packet:")
console.log(obj);

// Serialize an empty packet.
pk = kosuzu.serialize(13, {});
console.log("Empty packet:");
console.log(pk);
console.log(kosuzu.packetBytes(pk));