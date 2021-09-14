var a = {
  i: 1,
  valueOf: function() {
    if (this.i === 1) {
      this.i+=1;
      return 1;
    } else {
      return 12;
    }
  }
};

if (a==1 && a==12) {
  console.log(`a = ${a.i}`);
}
