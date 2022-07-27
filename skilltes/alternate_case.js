String.prototype.toAlternatingCase = function () {
    return this.split('').map(Letter => Letter === Letter.toUpperCase() ? 
    Letter.toLocaleLowerCase() : Letter.toUpperCase()).join('');
}

console.log('abc!'.toAlternatingCase());
console.log('ABC!'.toAlternatingCase());
console.log('Hello World!'.toAlternatingCase());

