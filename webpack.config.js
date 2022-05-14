const path = require('path');

module.exports = {
  entry: {
    order: './web/js/order.js',
    addParkingPlace: './web/js/add-parking-place.js',
    removeParkingPlace: './web/js/remove-parking-place.js',
  },
  output: {
    filename: '[name].js',
    path: path.resolve(__dirname, 'web/dist'),
  },
  mode: 'development'
};