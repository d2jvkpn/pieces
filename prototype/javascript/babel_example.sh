#! /usr/bin/env bash
set -eu -o pipefail

mkdir project && cd project

cat > main.js << EOF
import fs from 'fs';

const data = fs.readFileSync('example.txt')

console.log(String(data))
EOF

cat > example.txt << EOF
hello, world!
- From the fs pacakage
EOF

npm init -y

npm update -g npm
npm install -g nodemon
# npm install --save-dev nodemon

npm install --save-dev @babel/core @babel/preset-env @babel/cli @babel/node

cat > .babelrc << EOF
{
  "presets": ["@babel/preset-env"]
}
EOF


npx babel-node ./main.js

# or set package.js/scripts["start"] = "nodemon --exec babel-node main.js"
