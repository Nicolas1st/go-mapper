const path = require('path');

// module.exports = {
//   entry: './web/js/main.js',
//   output: {
//     filename: 'index.js',
//     path: path.resolve(__dirname, 'web/js'),
//   },
//   mode: 'development'
// };

module.exports = {
  entry: {
    main: './web/js/main.js',
    addParkingPlace: './web/js/add-parking-place.js',
    removeParkingPlace: './web/js/remove-parking-place.js',
  },
  output: {
    filename: '[name].js',
    path: path.resolve(__dirname, 'web/dist'),
  },
  mode: 'development'
};