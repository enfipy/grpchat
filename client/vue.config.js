const path = require('path')

module.exports = {
  // publicPath: 'grpchat',
  configureWebpack: {
    resolve: {
      alias: {
        schema: path.resolve(__dirname, '../schema/gen/web/')
      }
    }
  }
}
