const path = require('path')

module.exports = {
  entry: {
    main: './public/js/src/main.js',
    otherPage: './public/js/src/main.js'
  },
  output: {
    path: path.join(__dirname, '/public/js/dist/'),
    filename: '[name].bundle.js'
  },
  module: {
    rules: [
      
    ]
  }
}
