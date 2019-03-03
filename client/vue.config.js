const path = require('path')

module.exports = {
  configureWebpack: {
    resolve: {
      alias: {
        schema: path.resolve(__dirname, '../schema/gen/web/')
      }
    }
  }
}
