const express = require('express');
const mustacheExpress = require('mustache-express');
const app = express();

///
app.engine('mustache', mustacheExpress());
app.set('view engine', 'mustache');

let port = 3000;

let users = [
  {name: "Rover", age: 32},
  {name: "Apple", age: 18},
  {name: "Jane",  age: 27},
]


///
app.get('/home', (req, res) => {
  res.render('home', { animal: 'Alligator' });
})

app.get('/users', (req, res) => {
  res.render('users', { users: users });
})

///
app.listen(port, () => {
  console.log(`Example app listening at http://localhost:${port}`)
})
