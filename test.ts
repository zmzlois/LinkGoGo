const input1 = "https://mdn.com";
const input2 = "i don't know what this is";

const pattern = /^(https?|ftp):\/\/[^\s/$.?#].[^\s]*$/i;

const result = pattern.test(input1)
const result2 = pattern.test(input2)

console.log("result 1: ", result)
console.log("result 2: ", result2)

const pattern3 = "(https?|ftp):\/\/[^\s/$.?#].[^\s]*";

const result3 = input1.match(pattern3)
const result4 = input2.match(pattern3)

console.log("result 3: ", result3)
console.log("result 4: ", result4)



